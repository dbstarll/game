import os
import time

import pyautogui
import pyscreeze
from PIL import Image

pyscreeze.USE_IMAGE_NOT_FOUND_EXCEPTION = False
pyscreeze.GRAYSCALE_DEFAULT = False

LOCATE_OPTIONS = {'confidence': 0.98}
LEFT_OFFSET = 1
RIGHT_OFFSET = -1
TOP_OFFSET = 1
BOTTOM_OFFSET = -1
KIND_OFFSET_WIDTH = 38
KIND_OFFSET_HEIGHT = 76

LEFT_BOTTOM_IMG = Image.open('mp/skill-left-bottom.png')
RIGHT_TOP_IMG = Image.open('mp/skill-right-top.png')


def debug_image(im, window, file):
  gim = im.crop((window.left, window.top, window.left + window.width, window.top + window.height))
  gim.save('tmp/' + file + '-' + str(time.time()) + '.png', dpi=(144, 144))
  return gim


def detect_corner(im, box):
  lb = debug_image(im, pyscreeze.Box(box.left - LEFT_OFFSET, box.top + box.height - 50 - BOTTOM_OFFSET, 50, 50),
                   'left-bottom')
  rt = debug_image(im, pyscreeze.Box(box.left + box.width - 50 - RIGHT_OFFSET, box.top - TOP_OFFSET, 50, 75),
                   'right-top')
  print(len(list(pyautogui.locateAll(lb, im, **LOCATE_OPTIONS))))
  print(len(list(pyautogui.locateAll(rt, im, **LOCATE_OPTIONS))))


def match(file):
  with Image.open(file) as im:
    match_lb = list(pyautogui.locateAll(LEFT_BOTTOM_IMG, im, **LOCATE_OPTIONS))
    match_rt = list(pyautogui.locateAll(RIGHT_TOP_IMG, im, **LOCATE_OPTIONS))
    if len(match_lb) != 3 or len(match_rt) != 3:
      print(f'mismatch: {len(match_lb)} - {len(match_rt)}, file: {file}')
    else:
      for i in (0, 2):
        match_left_bottom = match_lb[i]
        match_right_top = match_rt[i]
        box = pyscreeze.Box(match_left_bottom.left + LEFT_OFFSET, match_right_top.top + TOP_OFFSET,
                            match_right_top.left + match_right_top.width - match_left_bottom.left - LEFT_OFFSET + RIGHT_OFFSET,
                            match_left_bottom.top + match_left_bottom.height - match_right_top.top - TOP_OFFSET + BOTTOM_OFFSET)
        # detect_corner(im, box)
        yield im.crop((box.left, box.top, box.left + box.width, box.top + box.height))


def detect_kinds(kinds, skill):
  kind = skill.crop((KIND_OFFSET_WIDTH, KIND_OFFSET_HEIGHT, skill.width - KIND_OFFSET_WIDTH,
                     KIND_OFFSET_HEIGHT + skill.width - 2 * KIND_OFFSET_WIDTH))

  for key, value in kinds.items():
    if pyautogui.locate(value, kind, **LOCATE_OPTIONS):
      return key
  key = f'kind-{len(kinds)}'
  print(f'\tdetect kind: {key} - {kind}')
  kinds[key] = kind

  if not os.path.exists(f'skills/{key}'):
    os.mkdir(f'skills/{key}')
  kind.save(f'skills/{key}/kind-{time.time()}.png', dpi=(144, 144))

  return key


def detect_skills(kinds, skills, skill):
  kind = detect_kinds(kinds, skill)
  kind_skills = skills.get(kind)
  if kind_skills is None:
    kind_skills = []
    skills[kind] = kind_skills

  for each in kind_skills:
    if pyautogui.locate(each, skill, **LOCATE_OPTIONS):
      return False
  print(f'\tdetect skill: {kind} - {skill}')
  kind_skills.append(skill)

  skill.save(f'skills/{kind}/skill-{time.time()}.png', dpi=(144, 144))
  return True


if __name__ == "__main__":
  total = 0
  matches = 0
  kinds = {}
  skills = {}
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
        detect_skills(kinds, skills, box)
  print(f'matches: {matches}/{total}')
  print(f'kinds: {len(kinds)}')
  for kind, kind_skills in skills.items():
    print(f'skills: {kind} - {len(kind_skills)}')
