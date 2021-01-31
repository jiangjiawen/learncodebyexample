import math
import os
import time
from copy import deepcopy

import torch
import torch.backends.cudnn as cudnn
import torch.nn as nn
import torch.nn.functional as F
import torchvision.models as models

def init_seeds(seed=0):
    torch.manual_seed(seed)

    if seed==0:
        cudnn.deterministic = True
        cudnn.benchmark = False
    else:
        cudnn.deterministic = False
        cudnn.benchmark = True

def select_device(device='',apex=False,batch_size=None):
    cpu_request = device.lower() == 'cpu'
    if device and not cpu_request:
        os.environ['CUDA_VISIBLE_DEVICES'] = device
        assert torch.cuda.is_available(), 'cuda unavailable, invalid device %s requested' % device
    
    cuda = False if cpu_request else torch.cuda.is_available()
    if cuda:
        c = 1024**2
        ng = torch.cuda.device_count()
        if ng>1 and batch_size:
            assert batch_size % ng == 0, 'batch-size %g not multiple of GPU count %g' % (batch_size, ng)
        x = [torch.cuda.get_device_properties(i) for i in range(ng)]
        s = 'Using CUDA' + ('Apex' if apex else '')
        for i in range(0, ng):
            if i==1:
                s=' '*len(s)
            print("%sdevice%g _CudaDeviceProperties(name0'%s', total_memory=%dMB)"%(s,i,x[i].name),x[i].total_memory/c)
    
    else:
        print('Using CPU')
    
    print('')
    return torch.device('cuda:0' if cuda else 'cpu')

def time_synchronized():
    torch.cuda.synchronize() if torch.cuda.is_available() else None
    return time.time()

def initialize_weights(model):
    for m in model.modules():
        t = type(m)
        if t is nn.Conv2d:
            pass
            # nn.init.kaiming_normal_(m.weight,mode='fan_out',nonlinearity='relu')
        elif t is nn.BatchNorm2d:
            m.eps = 1e-4
            m.momentum = 0.03
        # relu6 clip to 6 inplace direct change
        elif t in [nn.LeakyReLU, nn.ReLU, nn.ReLU6]:
            m.inplace = True

def find_modules(model, mclass=nn.Conv2d):
    return [i for i,m in enumerate(model.module_list) if isinstance(m, module_list)]

def fuse_conv_and_bn(conv,bn):
    with torch.no_grad():
        fusedconv = torch.nn.Conv2d(conv.in_channels,
                                    conv.out_channels,
                                    kernel_size=conv.kernel_size,
                                    stride=conv.stride,
                                    padding=conv.padding,bias=True)
        
        w_conv = conv.weight.clone().view(conv.out_channels, -1)
        w_bn = torch.diag(bn.weight.div(torch.sqrt(bn.eps+bn.running_var)))
        fusedconv.weight.copy_(torch.mm(w_bn,w_conv).view(fusedconv.weight.size()))

        if conv.bias is not None:
            b_conv=conv.bias
        else:
            b_conv = torch.zeros(conv.weight.size(0),device=conv.weight.device)
        
        b_bn = bn.bias - bn.weight.mul(bn.running_mean).div(torch.sqrt(bn.running_var + bn.eps))
        fusedconv.bias.copy_(torch.mm(w_bn,b_conv.reshape(-1,1)).reshape(-1) + b_bn)

        return fusedconv

def model_info(model, verbose=False):
    n_p=sum(x.numel() for x in model.parameters())
    n_g=sum(x.numel() for x in model.parameters() if x.requires_grad)

    if verbose:
        print('%5s %40s %9s %12s %20s %10s %10s' %('layer','name','gradient','parameters','shape','mu','sigma'))
        for i,(name,p) in enumerate(model.named_parameters()):
            name = name.replace('module_list.','')
            print('%5g %40s %9s %12s %20s %10.3g %10.3g' %(i,name,p.requires_grad,p.numel(),list(p.shape),p.mean(),p.std()))
    
    try:
        from thop import profile
        macs,_=profile(model,inputs=(torch.zeros(1,3,480,640),),verbose=False)
        fs = ', %.1f GFLOPS' %(macs/1e)*2)
    except:
        fs = ''
    
    print('Model Summary:%g layer,%g paramters,%g gradients%s'%(len(list(model.parameters())),n_p,n_g,fs))

def load_classifier(name='resnet101',n=2):
    model=models.__dict__[name](pretrained=True)

    input_size = [3,224,224]
    input_space = 'RGB'
    input_range = [0,1]
    mean = [0.485,0.456,0.406]
    std=[0.229,0.224,0.225]
    for x in [input_size,input_space,input_range,mean,std]:
        print(x+' =', eval(x))
    
    filters = model.fc.weight.shape[1]
    model.fc.bias = torch.nn.Parameter(torch.zeros(n),requires_grad=True)
    model.fc.weight = torch.nn.Parameter(torch.zeros(n, filters), requires_grad=True)
    model.fc.out_features = n
    return model

def scale_img(img,ratio=1.0,same_shape=False):
    h,w=img.shape[2:]
    s=(int(h*ratio),int(w*ratio))
    img = F.iterpolate(img, size=s, mode='bilinear', align_corners=False)
    if not same_shape:
        gs = 32
        h,w=[math.ceil(x.ratio/gs)*gs for x in (h,w)]
    return F.pad(img,[0,w-s[1],0,h-s[0]])

class ModelEMA:

    def __init__(self, model, decay = 0.9999, device=''):
        self.ema = deepcopy(model)
        self.ema.eval()
        self.updates = 0
        self.decay = lambda x:decay * (1-math.exp(-x/2000))
        self.device = device

        if device:
            self.ema.to(device=device)
        for p in self.ema.parameters():
            p.requires_grad_(False)
    
    def update(self, model):
        self.updates += 1
        d = self.decay(self.updates)
        with torch.no_grad():
            if type(mode) in (nn.parallel.DataParallel,nn.parallel.DistributedDataParallel):
                msd,esd=model.module.state_dict(), self.ema.module.state_dict()
            else:
                msd, esd = model.state_dict(), self.ema.state_dict()
            
            for k,v in esd.items():
                if v.dtype.is_floating_point:
                    if v.dtype.is_floating_point:
                        v *= d
                        v += (1.-d)*msd[k].detach()
    
    def update_attr(self, model):
        for k in model.__dict__.keys():
            if not k.startwith('_'):
                setattr(self.ema, k, getattr(model, k))