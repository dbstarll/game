import os
import sys

from PIL import Image

from _game import distribute
from _image import img
from _locate import locate, locate_all
from _rescue import load_rescues, detect_rescues, crop_rescue

rescue_img = None
fight_list_img = None


def match_rescues(file):
  with Image.open(file) as im:
    if locate(fight_list_img, im) is None:
      raise ValueError('fight-list not found')
    for match in locate_all(rescue_img, im):
      _, img = crop_rescue(im, match)
      yield img


if __name__ == "__main__":
  dist = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')
  rescue_img = Image.open(img('rescue'))
  fight_list_img = Image.open(img('fight-list'))

  rescues = load_rescues()
  files = 0
  matches = 0
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('fights-') and file.endswith('.png'):
      files += 1
      for img in match_rescues(f'tmp/{dist}/{file}'):
        matches += 1
        rescue_name, new_rescue = detect_rescues(rescues, img)
  print(f'rescues: {len(rescues)}')
  print(f'match: {matches} on {files} files')
