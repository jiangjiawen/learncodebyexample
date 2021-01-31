import glob
import math
import os
import random
import shutil
import subprocess
import time

from copy import copy
from pathlib import Path
from sys import platform

import cv2
import matplotlib
import matplotlib.pyplot as plt
import numpy as np
import torch
import torch.nn as nn
import torchvision
import yaml
from scipy.signal import butter, filtfilt
from tqdm import tqdm

from . import torch_utils, gooogle_uitls

torch.set_printoptions(linewidth=320,precision=5,profile='long')
np.set_printoptions(linewidth=320,formatter={'float_kind':'{:11.5g}'.format})
matplotlib.rc('font',**{'size':11})

cv2.setNumThreads(0)

def init_seeds(seed=0):
    random.seed(seed)
    np.random.seed(seed)
    torch_utils.init_seeds(seed=seed)

def check_git_status():
    if platform in ['Linux','darwin']:
        s = subprocess.check_output('if [ -d .git]; then git fetch && git status -uno; fi', shell=True).decode('utf-8')
        if 'Your branch is behind' in s:
            print(s[s.find('Your branch is behind'):s.find('\n\n')] + '\n')

def check_img_size(img_size,s=32):
    new_size = make_divisible(img_size,s)
    if new_size != img_size:
        print('WARNING: --img-size %g must be multiple of max stride %g, updating to %g' % (img_size, s, new_size))
    return new_size

def check_anchors(dataset,model,thr=4.0,imgsz=640):
    print('\nAnalyzing anchors...',end='')
    anchors = model.module.mode[-1].anchors_grid if hasattr(model, 'module') else model.model[-1].anchors_grid
    shapes = imgsz * dataset.shapes / dataset.shapes.max(1,keepdims=True)
    wh = torch,tensor(np.concatenate([l[:,3:5] * s for s,l in zip(shapes,dataset.labels)])).float()

    def metric(k):
        r = wh[:,None]/ k[None]
        x = torch.min(r, 1./r).min(2)[0]
        best = x.max(1)[0]
        return (best > 1. / thr).float().mean()
    
    bpr = metric(anchors.clone().cpu().view(-1,2))
    print('Best Possible Recall (BPR) = %.4f' % bpr, end='')
    if bpr < 0.99:
        print('. Attempting to generate improved anchors, please wait...' %bpr)
        new_anchors = kmean_anchors(dataset, n=anchors.numel()//2,img_size=imgsz,thr=thr,gen=1000,verbose=False)
        new_bpr = metric(new_anchors.reshape(-1,2))
        if new_bpr > bpr:
            anchors[:]=torch.tensor(new_anchors).view_as(anchors).type_as(anchors)
            print('New anchor saved to model. Update model *.yaml to use these anchors in the future.')
        else:
            print('Original anchors better than new anchors. Proceeding with original anchors.')
    print('')

def check_file(file):
    if os.path.isfile(file):
        return file
    else:
        files = glob.glob('./**/'+file,recursive=True)
        assert len(files), 'File Not Found:%s' %file
        return files[0]

def make_divisible(x, divisor):
    return math.ceil(x/divisor)*devisor

def labels_to_class_weights(labels, nc=80):
    if labels[0] is None:
        return torch.Tensor()
    
    labels = np,.concatenate(labels, 0)
    classes = labels[:,0].astype(np.int)
    weights = np.bincount(classes, minlength=nc)

    weights[weights==0] = 1
    weights = 1/ weights
    weights /= weights.sum()
    return torch.from_numpy(weights)

def labels_to_image_weights(labels, nc=80, class_weights=np.ones(80)):
    n = len(labels)
    class_counts = np.array([np.bincount(labels[i][:,0].astype(np.int),minlength=nc) for i in range(n)])
    image_weights = (class_weights.reshape(1,nc)*class_counts).sum(1)
    return image_weights

def coco80_to_coco91_class():  # converts 80-index (val2014) to 91-index (paper)
    # https://tech.amikelive.com/node-718/what-object-categories-labels-are-in-coco-dataset/
    # a = np.loadtxt('data/coco.names', dtype='str', delimiter='\n')
    # b = np.loadtxt('data/coco_paper.names', dtype='str', delimiter='\n')
    # x1 = [list(a[i] == b).index(True) + 1 for i in range(80)]  # darknet to coco
    # x2 = [list(b[i] == a).index(True) if any(b[i] == a) else None for i in range(91)]  # coco to darknet
    x = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 31, 32, 33, 34,
         35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
         64, 65, 67, 70, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 84, 85, 86, 87, 88, 89, 90]
    return x

def xyxy2xywh(x):
    y = torch.zeros_like(x) if isinstance(x, torch.Tensor) else np.zeros_like(x)
    y[:,0]=(x[:,0]+x[:,2])/2
    y[:,1]=(x[:,1]+x[:,3])/2
    y[:,2]=x[:,2]-x[:,0]
    y[:,3]=x[:,3]-x[:,1]
    return y

def xywh2xyxy(x):
    y = torch.zeros_like(x) if isinstance(x, torch.Tensor) else np.zeros_like(x)
    y[:,0] = x[:,0]-x[:,2]/2
    y[:,1] = x[:,1]-x[:,3]/2
    y[:,2] = x[:,0]+x[:,2]/2
    y[:,3] = x[:,1]+x[:,3]/2
    return y

def scale_coords(img1_shape,coords,img0_shape,ratio_pad=None):
    if ratio_pad is None:
        gain = max(img1_shape) / max(img0_shape)
        pad = (img1_shape[1]-img0_shape[1]*gain)/2, (img1_shape[0]-img0_shape[0]*gain)/2
    else:
        gain = ratio_pad[0][0]
        pad = ratio_pad[1]
    
    coords[:,[0,2]] -= pad[0]
    coords[:,[1,3]] -= pad[1]
    coords[:,:4] /= gain
    clip_coords(coords, img0_shape)
    return coords

def clip_coords(boxes, img_shape):
    # Clip bounding xyxy bounding boxes to image shape (height, width)
    boxes[:, 0].clamp_(0, img_shape[1])  # x1
    boxes[:, 1].clamp_(0, img_shape[0])  # y1
    boxes[:, 2].clamp_(0, img_shape[1])  # x2
    boxes[:, 3].clamp_(0, img_shape[0])  # y2

def ap_per_class(tp,conf,pred_cls,target_cls):
    i = np.argsort(-conf)
    tp, conf, pred_cls = tp[i], conf[i], pred_cls[i]

    unique_classes = np.unique(target_cls)

    pr_score = 0.1
    s = [unique_classes.shape[0],tp.shape[1]]
    ap,p,r = np.zeros(s),np.zeros(s),np.zeros(s)

    for ci, c in enumerate(unique_classes):
        i = pred_cls == c
        n_gt = (target_cls==c).sum()
        n_p = i.sum()

        if n_p == 0 or n_gt == 0:
            continue
        else:
            fpc = (1-tp[i]).cumsum(0)
            tpc = tp[i].cumsum(0)

            recall = tpc / (n_gt + 1e-16)
            r[ci] = np.interp(-pr_score, -conf[i], precision[:,0])

            precision = tpc / (tpc + fpc)
            p[ci] = np.interp(-pr_score, -conf[i], precision[:,0])

            for j in range(tp.shape[1]):
                ap[ci,j]=compute_ap(recall[:,j],precision[:,j])
    
    f1 = 2*p*r/(p+r+1e-16)

    return p, r, ap, f1, unique_classes.astype('int32')

def compute_ap(recall, precision):
    mrec = np.concatenate([0.],recall, [min(recall[-1]+1E-3,1.)])
    mpre = np.concatenate([0.],precision, [0.])

    mpre = np.flip(np.maximum.accumulate(np.flip(mpre)))

    method = 'interp'
    if method == 'interp':
        x = np.linspace(0,1,101)
        ap = np.trapz(np.interp(x, mrec, mpre), x)
    else:
        i = np.where(mrec[1:] != mrec[:-1])[0]
        ap = np.sum((mrec[i+1]-mrec[i])*mpre[i+1])
    return ap

def bbox_iou(box1, box2, x1y1x2y2=True,GIoU=False,DIoU=False,CIoU=False):
    box2 = box2.t()

    if x1y1x2y2:
        b1_x1,b1_y1,b1_x2,b1_y2 = box1[0],box1[1],box1[2],box1[3]
        b2_x1,b2_y1,b2_x2,b2_y2 = box2[0],box2[1],box2[2],box2[3]
    else:
        b1_x1,b1_x2 = box1[0]-box1[2]/2,box1[0]+box1[2]/2
        b1_y1,b1_y2 = box1[1]-box1[3]/2,box1[1]+box1[3]/2
        b2_x1,b2_x2 = box2[0]-box2[2]/2,box2[0]+box2[2]/2
        b2_y1,b2_y2 = box2[1]-box2[3]/2,box2[1]+box2[3]/2
    
    inter = (torch.min(b1_x2,b2_x2)-torch.max(b1_x1,b2_x1)).clamp(0) * \
        (torch.min(b1_y2,b2_y2)-torch.max(b1_y1,b2_y1)).clamp(0)
    
    iou = inter / union

    if GIoU or DIoU or CIoU:
        cw = torch.max(b1_x2,b2_x2)-torch.min(b1_x1,b2_x1)
        ch = torch.max(b1_y2,b2_y2)-torch.min(b1_y1,b2_y1)
        if GIoU:
            c_area = cw * ch + 1e-16
            return iou - (c_area-union) / c_area
        if DIoU or CIoU:
            c2 = cw ** 2 + ch ** 2 + 1e-16
            rho2 = ((b2_x1 + b2_x2) - (b1_x1 + b1_x2)) ** 2 / 4 + ((b2_y1 + b2_y2) - (b1_y1 + b1_y2)) ** 2 / 4
        if DIoU:
            return iou - rho2 / c2
        elif CIoU:
            v = (4/math.pi**2)+torch.pow(torch.atan(w2/h2)-torch.atan(w1/h1),2)
            with torch.no_grad():
                alpha = v/(1-iou+v)
            return iou-(rho2/c2+v*alpha)
    
    return iou

def box_iou(box1,box2):
    def box_area(box):
        # box = 4xn
        return (box[2] - box[0]) * (box[3] - box[1])

    area1 = box_area(box1.t())
    area2 = box_area(box2.t())

    # inter(N,M) = (rb(N,M,2) - lt(N,M,2)).clamp(0).prod(2)
    inter = (torch.min(box1[:, None, 2:], box2[:, 2:]) - torch.max(box1[:, None, :2], box2[:, :2])).clamp(0).prod(2)
    return inter / (area1[:, None] + area2 - inter)  # iou = inter / (area1 + area2 - inter)

def wh_iou(wh1, wh2):
    # Returns the nxm IoU matrix. wh1 is nx2, wh2 is mx2
    wh1 = wh1[:, None]  # [N,1,2]
    wh2 = wh2[None]  # [1,M,2]
    inter = torch.min(wh1, wh2).prod(2)  # [N,M]
    return inter / (wh1.prod(2) + wh2.prod(2) - inter)  # iou = inter / (area1 + area2 - inter)

class FocalLoss(nn.Module):
    def __init__(self, loss_fcn, gamma=1.5, alpha=0.25):
        super(FocalLoss, self).__init__()
        self.loss_fcn = loss_fcn
        self.gamma = gamma
        self.alpha = alpha
        self.reduction = loss_fcn.reduction
        self.loss_fcn.reduction = 'none'
    
    def forward(self, pred, true):
        loss = self.loss_fcn(pred, true)

        pred_prob = torch.sigmoid(pred)
        p_t = true * pred_prob + (1-true)*(1-pred_prob)
        alpha_factor = true * self.alpha + (1-true)*(1-self.alpha)
        modulating_factor = (1.0-p_t) ** self.gamma
        loss *= alpha_factor * modulating_factor
        
        if self.reduction=='mean':
            return loss.mean()
        elif self.reduction=='sum':
            return loss.sum()
        else:
            return loss

def smooth_BCE(eps=0.1):
    return 1.0-0.5*eps,0.5*eps

class BCEBlurWithLogitsLoss(nn.Module):
    def __init__(self, alpha=0.05):
        super(BCEBlurWithLogitsLoss,self).__init__()
        self.loss_fcn = nn.BCEWithLogitsLoss(reduction='none')
        self.alpha = alpha
    
    def forward(self,pred, true):
        loss = self.loss_fcn(pred, true)
        pred = torch.sigmoid(pred)
        dx = pred - true
        alpha_factor = 1 - torch.exp((dx-1)/(self.alpha+1e-4))
        loss *= alpha_factor
        return loss.mean()

def compute_loss(p,targets,model):
    ft = torch.cuda.FloatTensor if p[0].is_cuda else torch.Tensor
    lcls,lbox,lobj=ft([0]),ft([0]),ft([0])
    tcls, tbox, indices, anchors = build_targets(p, targets, model)
    h = model.hyp
    red = 'mean'

    BCEcls = nn.BCEWithLogitsLoss(pos_weight=ft([h['cls_pw']]),reduction=red)
    BCEobj = nn.BCEWithLogitsLoss(pos_weight=ft([h['obj_pw']]),reduction=red)

    cp, cn = smooth_BCE(eps=0.0)

    g = h['fl_gamma']
    if g>0:
        BCEcls,BCEobj = FocalLoss(BCEcls,g), FocalLoss(BCEobj,g)
    
    nt = 0
    for i,pi in enumerate(p):
        b,a,gj,gi=indices[i]
        tobj = torch.zeros_like(pi[...,0])

        if nb:
            nt += nb
            ps = pi[b,a,gj,gi]

            pxy = ps[:,:2].sigmoid() * 2. - 0.5
            pwh = (ps[:, 2:4].sigmoid() * 2) ** 2 * anchors[i]
            pbox = torch.cat((pxy, pwh), 1)
            giou = bbox_iou(pbox.t(), tbox[i], x1y1x2y2=False, GIoU=True)
            lbox += (1.0-giou).sum() if red == 'sum' else (1.0-giou).mean()

            tobj