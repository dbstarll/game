import os
import time

from PIL import Image

from _debug import debug_image
from _game import distribute_file
from _image import save_image, img
from _locate import locate, Box, locate_all

_SKILL_ROOT_DIR = 'skills'

_SKILL_LEFT_OFFSET = 1
_SKILL_RIGHT_OFFSET = -1
_SKILL_TOP_OFFSET = 1
_SKILL_BOTTOM_OFFSET = -1
_KIND_OFFSET_WIDTH = 38
_KIND_OFFSET_HEIGHT = 76

_SKILL_KINDS = {}
_SKILLS = {}


def _skill_img(kind_name, skill_name):
  return img(f'{_SKILL_ROOT_DIR}/{kind_name}/{skill_name}')


def _match_kinds(kinds, skill):
  _, kind = _crop_kind(skill)
  for kind_name, item in kinds.items():
    if locate(kind, item):
      return kind_name, kind
  return None, kind


def _detect_kinds(kinds, skill, file):
  kind_name, kind = _match_kinds(kinds, skill)
  if kind_name is not None:
    return kind_name, True

  kind_name = f'logo-{time.time()}'
  print(f'\tdetect kind: {kind_name} - {file}')
  kinds[kind_name] = kind

  if not os.path.exists(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}')):
    os.mkdir(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}'))
  save_image(kind, _skill_img(kind_name, kind_name))

  return kind_name, True


def match_skills(skill):
  kind_name, _ = _match_kinds(_SKILL_KINDS, skill)
  if kind_name is None:
    return None, None

  for skill_name, item in _SKILLS.get(kind_name).items():
    if locate(skill, item):
      return kind_name, skill_name
  return kind_name, None


def detect_skills(skill, file):
  kind_name, _ = _detect_kinds(_SKILL_KINDS, skill, file)
  kind_skills = _SKILLS.get(kind_name)
  if kind_skills is None:
    kind_skills = {}
    _SKILLS[kind_name] = kind_skills

  for skill_name, item in kind_skills.items():
    if locate(skill, item):
      return kind_name, skill_name, False
  skill_name = f'skill-{time.time()}'
  print(f'\tdetect skill: {kind_name} - {skill_name} - {file}')
  kind_skills[skill_name] = skill

  save_image(skill, _skill_img(kind_name, skill_name))
  return kind_name, skill_name, True


def load_skills():
  _SKILL_KINDS.clear()
  _SKILLS.clear()
  for kind_name in os.listdir(distribute_file(_SKILL_ROOT_DIR)):
    if os.path.isdir(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}')):
      _load_kind(_SKILL_KINDS, _SKILLS, kind_name)
  return _SKILL_KINDS, _SKILLS


def _load_kind(kinds, skills, kind_name):
  print(f'\tloading kind: {kind_name}')
  for skill_name in os.listdir(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}')):
    if skill_name.endswith('.png'):
      if skill_name.startswith('logo'):
        _load_logo(kinds, kind_name, skill_name[:len(skill_name) - 4])
      else:
        _load_skill(skills, kind_name, skill_name[:len(skill_name) - 4])


def _load_logo(kinds, kind_name, skill_name):
  print(f'\t\tloading logo: {skill_name}')
  kinds[kind_name] = Image.open(_skill_img(kind_name, skill_name))


def _load_skill(skills, kind_name, skill_name):
  print(f'\t\tloading skill: {skill_name}')
  kind_skills = skills.get(kind_name)
  if kind_skills is None:
    kind_skills = {}
    skills[kind_name] = kind_skills
  kind_skills[skill_name] = Image.open(_skill_img(kind_name, skill_name))


def _crop_kind(skill):
  box = Box(_KIND_OFFSET_WIDTH, _KIND_OFFSET_HEIGHT, skill.width - 2 * _KIND_OFFSET_WIDTH,
            skill.width - 2 * _KIND_OFFSET_WIDTH)
  return box, skill.crop((box.left, box.top, box.left + box.width, box.top + box.height))


def crop_image(im, match_left_bottom, match_right_top):
  box = Box(match_left_bottom.left + _SKILL_LEFT_OFFSET, match_right_top.top + _SKILL_TOP_OFFSET,
            match_right_top.left + match_right_top.width - match_left_bottom.left - _SKILL_LEFT_OFFSET + _SKILL_RIGHT_OFFSET,
            match_left_bottom.top + match_left_bottom.height - match_right_top.top - _SKILL_TOP_OFFSET + _SKILL_BOTTOM_OFFSET)
  # _detect_corner(im, box)
  return box, im.crop((box.left, box.top, box.left + box.width, box.top + box.height))


def _detect_corner(im, box):
  lb = debug_image(im, Box(box.left - _SKILL_LEFT_OFFSET, box.top + box.height - 50 - _SKILL_BOTTOM_OFFSET, 50, 50),
                   'left-bottom')
  rt = debug_image(im, Box(box.left + box.width - 50 - _SKILL_RIGHT_OFFSET, box.top - _SKILL_TOP_OFFSET, 50, 75),
                   'right-top')
  print(len(list(locate_all(lb, im))))
  print(len(list(locate_all(rt, im))))
