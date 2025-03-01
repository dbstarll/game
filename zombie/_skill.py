import os
import time

from PIL import Image

from _game import distribute_file
from _image import save_image, img
from _locate import locate

_KIND_OFFSET_WIDTH = 38
_KIND_OFFSET_HEIGHT = 76
_SKILL_ROOT_DIR = 'skills'


def _skill_img(kind_name, skill_name):
  return img(f'{_SKILL_ROOT_DIR}/{kind_name}/{skill_name}')


def _match_kinds(kinds, skill):
  kind = skill.crop((_KIND_OFFSET_WIDTH, _KIND_OFFSET_HEIGHT, skill.width - _KIND_OFFSET_WIDTH,
                     _KIND_OFFSET_HEIGHT + skill.width - 2 * _KIND_OFFSET_WIDTH))
  for kind_name, item in kinds.items():
    if locate(kind, item):
      return kind_name, kind
  return None, kind


def _detect_kinds(kinds, skill):
  kind_name, kind = _match_kinds(kinds, skill)
  if kind_name is not None:
    return kind_name, True

  kind_name = f'logo-{time.time()}'
  print(f'\tdetect kind: {kind_name} - {kind}')
  kinds[kind_name] = kind

  if not os.path.exists(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}')):
    os.mkdir(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}'))
  save_image(kind, _skill_img(kind_name, kind_name))

  return kind_name, True


def match_skills(kinds, skills, skill):
  kind_name, _ = _match_kinds(kinds, skill)
  if kind_name is None:
    return None, None

  for skill_name, item in skills.get(kind_name).items():
    if locate(skill, item):
      return kind_name, skill_name
  return kind_name, None


def detect_skills(kinds, skills, skill):
  kind_name, _ = _detect_kinds(kinds, skill)
  kind_skills = skills.get(kind_name)
  if kind_skills is None:
    kind_skills = {}
    skills[kind_name] = kind_skills

  for skill_name, item in kind_skills.items():
    if locate(skill, item):
      return kind_name, skill_name, False
  skill_name = f'skill-{time.time()}'
  print(f'\tdetect skill: {kind_name} - {skill_name} - {skill}')
  kind_skills[skill_name] = skill

  save_image(skill, _skill_img(kind_name, skill_name))
  return kind_name, skill_name, True


def load_skills():
  kinds = {}
  skills = {}
  for kind_name in os.listdir(distribute_file(_SKILL_ROOT_DIR)):
    if os.path.isdir(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}')):
      _load_kind(kinds, skills, kind_name)
  return kinds, skills


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
