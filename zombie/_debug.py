import datetime
import time

from PIL.Image import Image
from pyscreeze import Box

from _image import save_image, img


def now():
  return datetime.datetime.now()


def _tmp_img(filename: str) -> str:
  return f'tmp/{img(filename)}'


def debug_image(im: Image, prefix: str, box: Box = None) -> Image:
  if box is None:
    save_image(im, _tmp_img(f'{prefix}-{time.time()}'))
    return im
  else:
    gim = im.crop((box.left, box.top, box.left + box.width, box.top + box.height))
    save_image(gim, _tmp_img(f'{prefix}-{time.time()}'))
    return gim
