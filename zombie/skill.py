import os

from PIL import Image

from _debug import debug_image
from _locate import locate_all, Box
from _skills import load_skills, detect_skills

LEFT_OFFSET = 1
RIGHT_OFFSET = -1
TOP_OFFSET = 1
BOTTOM_OFFSET = -1

LEFT_BOTTOM_IMG = Image.open('mp/skill-left-bottom.png')
RIGHT_TOP_IMG = Image.open('mp/skill-right-top.png')


def detect_corner(im, box):
  lb = debug_image(im, Box(box.left - LEFT_OFFSET, box.top + box.height - 50 - BOTTOM_OFFSET, 50, 50),
                   'left-bottom')
  rt = debug_image(im, Box(box.left + box.width - 50 - RIGHT_OFFSET, box.top - TOP_OFFSET, 50, 75),
                   'right-top')
  print(len(list(locate_all(lb, im))))
  print(len(list(locate_all(rt, im, ))))


def match(file):
  with Image.open(file) as im:
    match_lb = list(locate_all(LEFT_BOTTOM_IMG, im))
    match_rt = list(locate_all(RIGHT_TOP_IMG, im))
    if len(match_lb) != 3 or len(match_rt) != 3:
      print(f'mismatch: {len(match_lb)} - {len(match_rt)}, file: {file}')
    else:
      for i in (0, 2):
        match_left_bottom = match_lb[i]
        match_right_top = match_rt[i]
        box = Box(match_left_bottom.left + LEFT_OFFSET, match_right_top.top + TOP_OFFSET,
                  match_right_top.left + match_right_top.width - match_left_bottom.left - LEFT_OFFSET + RIGHT_OFFSET,
                  match_left_bottom.top + match_left_bottom.height - match_right_top.top - TOP_OFFSET + BOTTOM_OFFSET)
        # detect_corner(im, box)
        yield im.crop((box.left, box.top, box.left + box.width, box.top + box.height))


if __name__ == "__main__":
  total = 0
  matches = 0
  kinds, skills = load_skills()
  for file in os.listdir('tmp'):
    if file.startswith('skills-') and file.endswith('.png'):
      if total >= 1000:
        break
      total += 1
      first = True
      for box in match('tmp/' + file):
        if first:
          first = False
          matches += 1
        kind_name, skill_name, new_skill = detect_skills(kinds, skills, box)
  print(f'matches: {matches}/{total}')
  print(f'kinds: {len(kinds)}')
  for kind, kind_skills in skills.items():
    print(f'skills: {kind} - {len(kind_skills)}')
