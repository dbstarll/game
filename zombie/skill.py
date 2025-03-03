import os
import sys

from PIL import Image

from _game import distribute, distribute_file
from _image import img
from _locate import locate_all
from _skill import load_skills, detect_skills, crop_image

LEFT_BOTTOM_IMG = None
RIGHT_TOP_IMG = None


def match_skills(file):
  with Image.open(file) as im:
    match_lb = list(locate_all(LEFT_BOTTOM_IMG, im))
    match_rt = list(locate_all(RIGHT_TOP_IMG, im))
    if len(match_lb) != 3 or len(match_rt) != 3:
      print(f'mismatch: {len(match_lb)} - {len(match_rt)}, file: {file}')
    else:
      for i in (0, 2):
        match_left_bottom = match_lb[i]
        match_right_top = match_rt[i]
        _, skill = crop_image(im, match_left_bottom, match_right_top)
        yield skill


if __name__ == "__main__":
  dist = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')
  LEFT_BOTTOM_IMG = Image.open(img('skill-left-bottom'))
  RIGHT_TOP_IMG = Image.open(img('skill-right-top.png'))

  total = 0
  matches = 0
  kinds, skills = load_skills()
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('skills-') and file.endswith('.png'):
      total += 1
      first = True
      for box in match_skills(f'tmp/{distribute_file(file)}'):
        if first:
          first = False
          matches += 1
        kind_name, skill_name, new_skill = detect_skills(box, f'tmp/{distribute_file(file)}')
  print(f'matches: {matches}/{total}')
  print(f'kinds: {len(kinds)}')
  for kind, kind_skills in skills.items():
    print(f'skills: {kind} - {len(kind_skills)}')
