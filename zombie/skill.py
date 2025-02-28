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

SKILL_ROOT_DIR = 'skills'

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

  for kind_name, item in kinds.items():
    if pyautogui.locate(item, kind, **LOCATE_OPTIONS):
      return kind_name, False
  kind_name = f'logo-{time.time()}'
  print(f'\tdetect kind: {kind_name} - {kind}')
  kinds[kind_name] = kind

  if not os.path.exists(f'{SKILL_ROOT_DIR}/{kind_name}'):
    os.mkdir(f'{SKILL_ROOT_DIR}/{kind_name}')
  kind.save(f'{SKILL_ROOT_DIR}/{kind_name}/{kind_name}.png', dpi=(144, 144))

  return kind_name, True


def detect_skills(kinds, skills, skill):
  kind_name, _ = detect_kinds(kinds, skill)
  kind_skills = skills.get(kind_name)
  if kind_skills is None:
    kind_skills = {}
    skills[kind_name] = kind_skills

  for skill_name, item in kind_skills.items():
    if pyautogui.locate(item, skill, **LOCATE_OPTIONS):
      return kind_name, skill_name, False
  skill_name = f'skill-{time.time()}'
  print(f'\tdetect skill: {kind_name} - {skill_name} - {skill}')
  kind_skills[skill_name] = skill

  skill.save(f'{SKILL_ROOT_DIR}/{kind_name}/{skill_name}.png', dpi=(144, 144))
  return kind_name, skill_name, True


def load_skills():
  kinds = {}
  skills = {}
  for kind_name in os.listdir(SKILL_ROOT_DIR):
    if os.path.isdir(f'{SKILL_ROOT_DIR}/{kind_name}'):
      print(f'\tloading kind: {kind_name}')
      for skill_name in os.listdir(f'{SKILL_ROOT_DIR}/{kind_name}'):
        if skill_name.endswith('.png'):
          if skill_name.startswith('logo'):
            print(f'\t\tloading logo: {skill_name}')
            kinds[kind_name] = Image.open(f'{SKILL_ROOT_DIR}/{kind_name}/{skill_name}')
          else:
            print(f'\t\tloading skill: {skill_name}')
            kind_skills = skills.get(kind_name)
            if kind_skills is None:
              kind_skills = {}
              skills[kind_name] = kind_skills
            kind_skills[skill_name] = Image.open(f'{SKILL_ROOT_DIR}/{kind_name}/{skill_name}')

  return kinds, skills


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
