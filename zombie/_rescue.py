import os
import time

from PIL import Image

from _game import distribute_file
from _image import save_image, img
from _locate import locate, Box

_RESCUE_ROOT_DIR = 'rescues'
_RESCUES = {}


def _rescue_img(rescue_name):
  return img(f'{_RESCUE_ROOT_DIR}/{rescue_name}')


def match_rescues(rescue):
  for rescue_name, item in _RESCUES.items():
    if locate(rescue, item, True):
      return rescue_name
  return None


def detect_rescues(rescue):
  rescue_name = match_rescues(rescue)
  if rescue_name is not None:
    return rescue_name, False
  rescue_name = f'rescue-{time.time()}'
  print(f'\tdetect rescue: {rescue_name} - {rescue}')
  _RESCUES[rescue_name] = rescue
  save_image(rescue, _rescue_img(rescue_name))
  return rescue_name, True


def _load_rescue(rescues, rescue_name):
  print(f'\tloading rescue: {rescue_name}')
  rescues[rescue_name] = Image.open(_rescue_img(rescue_name))


def load_rescues():
  global _RESCUES
  for rescue_name in os.listdir(distribute_file(_RESCUE_ROOT_DIR)):
    if rescue_name.endswith('.png'):
      _load_rescue(_RESCUES, rescue_name[:len(rescue_name) - 4])
  return _RESCUES


def crop_rescue(im, rect):
  box = Box(rect.left + rect.width * 1.7 - 2, rect.top, rect.width * 0.3 - 2, rect.height)
  return box, im.crop((box.left, box.top, box.left + box.width, box.top + box.height))
