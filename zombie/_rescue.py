import os
import time

from PIL import Image

from _game import distribute_file
from _image import save_image, img
from _locate import locate, Box

RESCUE_ROOT_DIR = 'rescues'


def _rescue_img(rescue_name):
  return img(f'{RESCUE_ROOT_DIR}/{rescue_name}')


def match_rescues(rescues, rescue):
  for rescue_name, item in rescues.items():
    if locate(rescue, item, True):
      return rescue_name
  return None


def detect_rescues(rescues, rescue):
  rescue_name = match_rescues(rescues, rescue)
  if rescue_name is not None:
    return rescue_name, False
  rescue_name = f'rescue-{time.time()}'
  print(f'\tdetect rescue: {rescue_name} - {rescue}')
  rescues[rescue_name] = rescue
  save_image(rescue, _rescue_img(rescue_name))
  return rescue_name, True


def _load_rescue(rescues, rescue_name):
  print(f'\tloading rescue: {rescue_name}')
  rescues[rescue_name] = Image.open(_rescue_img(rescue_name))


def load_rescues():
  rescues = {}
  for rescue_name in os.listdir(distribute_file(RESCUE_ROOT_DIR)):
    if rescue_name.endswith('.png'):
      _load_rescue(rescues, rescue_name[:len(rescue_name) - 4])
  return rescues


def crop_rescue(im, rect):
  box = Box(rect.left + rect.width * 1.7 - 2, rect.top, rect.width * 0.3 - 2, rect.height)
  return box, im.crop((box.left, box.top, box.left + box.width, box.top + box.height))
