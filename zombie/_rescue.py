import os
import time

from PIL import Image

from _image import save_image
from _locate import locate

RESCUE_ROOT_DIR = 'rescues'


def detect_rescues(rescues, rescue):
  for rescue_name, item in rescues.items():
    if locate(item, rescue):
      return rescue_name, False
  rescue_name = f'rescue-{time.time()}'
  print(f'\tdetect rescue: {rescue_name} - {rescue}')
  rescues[rescue_name] = rescue
  save_image(rescue, f'{RESCUE_ROOT_DIR}/{rescue_name}.png')
  return rescue_name, True


def _load_rescue(rescues, rescue_name):
  print(f'\tloading rescue: {rescue_name}')
  rescues[rescue_name] = Image.open(f'{RESCUE_ROOT_DIR}/{rescue_name}.png')


def load_rescues():
  rescues = {}
  for rescue_name in os.listdir(RESCUE_ROOT_DIR):
    if rescue_name.endswith('.png'):
      _load_rescue(rescues, rescue_name[:len(rescue_name) - 4])
  return rescues
