import os
import time

from PIL import Image

from _image import save_image
from _locate import locate

KIND_OFFSET_WIDTH = 38
KIND_OFFSET_HEIGHT = 76
SKILL_ROOT_DIR = 'skills'


def _detect_kinds(kinds, skill):
  kind = skill.crop((KIND_OFFSET_WIDTH, KIND_OFFSET_HEIGHT, skill.width - KIND_OFFSET_WIDTH,
                     KIND_OFFSET_HEIGHT + skill.width - 2 * KIND_OFFSET_WIDTH))

  for kind_name, item in kinds.items():
    if locate(item, kind):
      return kind_name, False
  kind_name = f'logo-{time.time()}'
  print(f'\tdetect kind: {kind_name} - {kind}')
  kinds[kind_name] = kind

  if not os.path.exists(f'{SKILL_ROOT_DIR}/{kind_name}'):
    os.mkdir(f'{SKILL_ROOT_DIR}/{kind_name}')
  save_image(kind, f'{SKILL_ROOT_DIR}/{kind_name}/{kind_name}.png')

  return kind_name, True


def detect_skills(kinds, skills, skill):
  kind_name, _ = _detect_kinds(kinds, skill)
  kind_skills = skills.get(kind_name)
  if kind_skills is None:
    kind_skills = {}
    skills[kind_name] = kind_skills

  for skill_name, item in kind_skills.items():
    if locate(item, skill):
      return kind_name, skill_name, False
  skill_name = f'skill-{time.time()}'
  print(f'\tdetect skill: {kind_name} - {skill_name} - {skill}')
  kind_skills[skill_name] = skill

  save_image(skill, f'{SKILL_ROOT_DIR}/{kind_name}/{skill_name}.png')
  return kind_name, skill_name, True


def load_skills():
  kinds = {}
  skills = {}
  for kind_name in os.listdir(SKILL_ROOT_DIR):
    if os.path.isdir(f'{SKILL_ROOT_DIR}/{kind_name}'):
      _load_kind(kinds, skills, kind_name)
  return kinds, skills


def _load_kind(kinds, skills, kind_name):
  print(f'\tloading kind: {kind_name}')
  for skill_name in os.listdir(f'{SKILL_ROOT_DIR}/{kind_name}'):
    if skill_name.endswith('.png'):
      if skill_name.startswith('logo'):
        _load_logo(kinds, kind_name, skill_name[:len(skill_name) - 4])
      else:
        _load_skill(skills, kind_name, skill_name[:len(skill_name) - 4])


def _load_logo(kinds, kind_name, skill_name):
  print(f'\t\tloading logo: {skill_name}')
  kinds[kind_name] = Image.open(f'{SKILL_ROOT_DIR}/{kind_name}/{skill_name}.png')


def _load_skill(skills, kind_name, skill_name):
  print(f'\t\tloading skill: {skill_name}')
  kind_skills = skills.get(kind_name)
  if kind_skills is None:
    kind_skills = {}
    skills[kind_name] = kind_skills
  kind_skills[skill_name] = Image.open(f'{SKILL_ROOT_DIR}/{kind_name}/{skill_name}.png')
