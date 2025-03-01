import os
import sys

from PIL import Image

from _game import distribute
from _image import img
from _locate import locate, locate_all, Box
from _rescue import load_rescues, detect_rescues

rescue_img = None
fight_list_img = None


def match_rescues(file):
  with Image.open(file) as im:
    if locate(fight_list_img, im) is None:
      raise ValueError('fight-list not found')
    for match in locate_all(rescue_img, im):
      box = Box(match.left + match.width * 1.7 - 2, match.top, match.width * 0.3 - 2, match.height)
      yield im.crop((box.left, box.top, box.left + box.width, box.top + box.height))


if __name__ == "__main__":
  dist = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')
  rescue_img = Image.open(img('rescue'))
  fight_list_img = Image.open(img('fight-list'))

  rescues = load_rescues()
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('fights-') and file.endswith('.png'):
      for img in match_rescues(f'tmp/{dist}/{file}'):
        rescue_name, new_rescue = detect_rescues(rescues, img)
  print(f'rescues: {len(rescues)}')
