import datetime
import time

from _image import save_image, img


def now():
  return datetime.datetime.now()


def _tmp_img(filename):
  return f'tmp/{img(filename)}'


def debug_image(im, rect, prefix):
  gim = im.crop((rect.left, rect.top, rect.left + rect.width, rect.top + rect.height))
  save_image(gim, _tmp_img(f'{prefix}-{time.time()}'))
  return gim
