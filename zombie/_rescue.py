import os
import time

from PIL import Image

from _game import distribute_file
from _image import save_image, img
from _locate import locate, _box, locate_all

_RESCUE_ROOT_DIR = 'rescues'
_RESCUES = {}
_RESCUE_IMG = None


def _rescue_img(rescue_name):
  return img(f'{_RESCUE_ROOT_DIR}/{rescue_name}')


def match_rescue(rescue):
  for rescue_name, item in _RESCUES.items():
    if locate(rescue, item, region=None):
      return rescue_name
  return None


def detect_rescues(rescue):
  rescue_name = match_rescue(rescue)
  if rescue_name is not None:
    return rescue_name, False
  rescue_name = f'rescue-{time.time()}'
  print(f'\tdetect rescue: {rescue_name} - {rescue}')
  _RESCUES[rescue_name] = rescue
  save_image(rescue, _rescue_img(rescue_name))
  return rescue_name, True


def _load_rescue(rescues, rescue_name):
  rescues[rescue_name] = Image.open(_rescue_img(rescue_name))


def load_rescues():
  global _RESCUE_IMG
  _RESCUE_IMG = Image.open(img('rescue'))
  for rescue_name in os.listdir(distribute_file(_RESCUE_ROOT_DIR)):
    if rescue_name.endswith('.png'):
      _load_rescue(_RESCUES, rescue_name[:len(rescue_name) - 4])
  print(f'加载寰球救援: {len(_RESCUES)}')
  return _RESCUES


def crop_rescue(im, rect):
  box = _box(rect.left + rect.width * 1.7 - 2, rect.top, rect.width * 0.3 - 2, rect.height)
  return box, im.crop((box.left, box.top, box.left + box.width, box.top + box.height))


def match_rescues(im):
  for match in locate_all(_RESCUE_IMG, im):
    _, rescue_image = crop_rescue(im, match)
    rescue_name = match_rescue(rescue_image)
    rescue_level = int(rescue_name[7:]) if rescue_name is not None else 0
    yield rescue_level, rescue_name, match, rescue_image
