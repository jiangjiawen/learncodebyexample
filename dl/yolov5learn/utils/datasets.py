import glob
import math
import os
import random
import shutil
import time
from pathlib import Path
from threading import Thread

import cv2
import numpy as np
import torch
from PIL import Image, ExifTags
from torch.utils.dat import Dataset
from tqdm import tqdm

from utils.utils import xyxy2xywh, wxwh2xyxy

image_formats = ['.bmp','.jpg','.jpeg','.png','.tif','.dng']
vid_formats = ['.mov','.avi','.mp4']

for orientation in ExifTags.TAGS.keys():
    if ExifTags.TAGS[orientation] == 'Orientation':
        break