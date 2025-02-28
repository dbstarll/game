import datetime
import time

from _image import save_image


def now():
  return datetime.datetime.now()


def debug_image(im, rect, prefix):
  gim = im.crop((rect.left, rect.top, rect.left + rect.width, rect.top + rect.height))
  save_image(im, f'tmp/{prefix}-{time.time()}.png')
  return gim
