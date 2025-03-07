import os
import sys

from PIL import Image

from _game import distribute, distribute_file
from _image import img
from _locate import locate
from _rescue import load_rescues, detect_rescues, match_rescues

fight_list_img = None


def match_rescues_from_file(file):
  with Image.open(file) as im:
    if locate(fight_list_img, im) is None:
      raise ValueError('fight-list not found')
    return match_rescues(im)


if __name__ == "__main__":
  dist, _ = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')
  fight_list_img = Image.open(img('fight-list'))

  rescues = load_rescues()
  files = 0
  matches = 0
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('fights-') and file.endswith('.png'):
      files += 1
      for rescue_level, rescue_name, rescue_rect, rescue_image in match_rescues_from_file(
          f'tmp/{distribute_file(file)}'):
        print(f'{rescue_level} - {rescue_name} - {rescue_rect}')
        matches += 1
        if rescue_level == 0:
          _, new_rescue = detect_rescues(rescue_image)
  print(f'rescues: {len(rescues)}')
  print(f'match: {matches} on {files} files')
