import os
import time

from PIL import Image

from _debug import debug_image
from _game import distribute_file, get_distribute, raise_unknown_distribute
from _image import save_image, img
from _locate import locate, Box, locate_all

_SKILL_ROOT_DIR = 'skills'

_SKILL_LEFT_OFFSET = 1
_SKILL_RIGHT_OFFSET = -1
_SKILL_TOP_OFFSET = 1
_SKILL_BOTTOM_OFFSET = -1
_MP_KIND_OFFSET_WIDTH = 38
_MP_KIND_OFFSET_HEIGHT = 76
_IOS_KIND_OFFSET_WIDTH = 39
_IOS_KIND_OFFSET_HEIGHT = 79
_SKILL_CONFIDENCE = 0.98

_SKILL_KINDS = {}
_SKILLS = {}
_LEFT_BOTTOM_IMG = None
_RIGHT_TOP_IMG = None


def _skill_img(kind_name, skill_name):
  return img(f'{_SKILL_ROOT_DIR}/{kind_name}/{skill_name}')


def kind_offset_width():
  distribute = get_distribute()
  if 'mp' == distribute:
    return _MP_KIND_OFFSET_WIDTH
  elif 'ios' == distribute:
    return _IOS_KIND_OFFSET_WIDTH
  else:
    raise_unknown_distribute()


def kind_offset_height():
  distribute = get_distribute()
  if 'mp' == distribute:
    return _MP_KIND_OFFSET_HEIGHT
  elif 'ios' == distribute:
    return _IOS_KIND_OFFSET_HEIGHT
  else:
    raise_unknown_distribute()


def recode_skip():
  distribute = get_distribute()
  if 'mp' == distribute:
    return 1
  elif 'ios' == distribute:
    return 2
  else:
    raise_unknown_distribute()


def _crop_kind_image(skill_image):
  rect = Box(kind_offset_width(), kind_offset_height(), skill_image.width - 2 * kind_offset_width(),
             skill_image.width - 2 * kind_offset_width())
  return rect, skill_image.crop((rect.left, rect.top, rect.left + rect.width, rect.top + rect.height))


def _crop_image(screenshot, match_left_bottom, match_right_top):
  box = Box(match_left_bottom.left + _SKILL_LEFT_OFFSET, match_right_top.top + _SKILL_TOP_OFFSET,
            match_right_top.left + match_right_top.width - match_left_bottom.left - _SKILL_LEFT_OFFSET + _SKILL_RIGHT_OFFSET,
            match_left_bottom.top + match_left_bottom.height - match_right_top.top - _SKILL_TOP_OFFSET + _SKILL_BOTTOM_OFFSET)
  return box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))


def _match_kinds(kinds, kind_image):
  for kind_name, item in kinds.items():
    if locate(kind_image, item, confidence=_SKILL_CONFIDENCE, region=None):
      return kind_name
  return None


def _match_kinds_from_skill(kinds, skill_image):
  _, kind_image = _crop_kind_image(skill_image)
  return _match_kinds(kinds, kind_image), kind_image


def _match_skill(skill_image):
  kind_name, kind_image = _match_kinds_from_skill(_SKILL_KINDS, skill_image)
  if kind_name is None:
    return None, None, kind_image

  for skill_name, item in _SKILLS.get(kind_name).items():
    if locate(skill_image, item, confidence=_SKILL_CONFIDENCE, region=None):
      return kind_name, skill_name, kind_image
  return kind_name, None, kind_image


def _record_kinds(kinds, kind_image):
  kind_name = _match_kinds(kinds, kind_image)
  if kind_name is not None:
    return kind_name, False

  kind_name = f'logo-{time.time()}'
  print(f'\trecord kind: {kind_name} - {kind_image}')
  kinds[kind_name] = kind_image

  if not os.path.exists(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}')):
    os.mkdir(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}'))
  save_image(kind_image, _skill_img(kind_name, kind_name))

  return kind_name, True


def record_skill(image_index, kind_name, kind_image, skill_image):
  if image_index == recode_skip():
    return None, None, False
  else:
    if kind_name is None:
      kind_name, _ = _record_kinds(_SKILL_KINDS, kind_image)

    kind_skills = _SKILLS.get(kind_name)
    if kind_skills is None:
      kind_skills = {}
      _SKILLS[kind_name] = kind_skills

    for skill_name, item in kind_skills.items():
      if locate(skill_image, item, confidence=_SKILL_CONFIDENCE, region=None):
        return kind_name, skill_name, False

    skill_name = f'skill-{time.time()}'
    print(f'\trecord skill: {kind_name} - {skill_name} - {skill_image}')
    kind_skills[skill_name] = skill_image

    save_image(skill_image, _skill_img(kind_name, skill_name))
    return kind_name, skill_name, True


def load_skills():
  global _LEFT_BOTTOM_IMG, _RIGHT_TOP_IMG
  _LEFT_BOTTOM_IMG = Image.open(img('skill-left-bottom'))
  _RIGHT_TOP_IMG = Image.open(img('skill-right-top.png'))
  _SKILL_KINDS.clear()
  _SKILLS.clear()
  for kind_name in os.listdir(distribute_file(_SKILL_ROOT_DIR)):
    if os.path.isdir(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}')):
      _load_kind(_SKILL_KINDS, _SKILLS, kind_name)
  total = 0
  for item in _SKILLS.values():
    total += len(item)
  print(f'load {total} skills of {len(_SKILL_KINDS)} kinds')
  return _SKILL_KINDS, _SKILLS


def _load_kind(kinds, skills, kind_name):
  for skill_name in os.listdir(distribute_file(f'{_SKILL_ROOT_DIR}/{kind_name}')):
    if skill_name.endswith('.png'):
      if skill_name.startswith('logo'):
        _load_logo(kinds, kind_name, skill_name[:len(skill_name) - 4])
      else:
        _load_skill(skills, kind_name, skill_name[:len(skill_name) - 4])


def _load_logo(kinds, kind_name, skill_name):
  kinds[kind_name] = Image.open(_skill_img(kind_name, skill_name))


def _load_skill(skills, kind_name, skill_name):
  kind_skills = skills.get(kind_name)
  if kind_skills is None:
    kind_skills = {}
    skills[kind_name] = kind_skills
  kind_skills[skill_name] = Image.open(_skill_img(kind_name, skill_name))


def _detect_corner(im, box):
  lb = debug_image(im,'skill-left-bottom', Box(box.left - 1, box.top + box.height - 50 + 1, 50, 50))
  rt = debug_image(im,'skill-right-top', Box(box.left + box.width - 50 + 1, box.top - 1, 50, 75))
  print(len(list(locate_all(lb, im))))
  print(len(list(locate_all(rt, im))))


def match_skills_from_screenshot(screenshot):
  match_left_bottoms = list(locate_all(_LEFT_BOTTOM_IMG, screenshot))
  match_right_tops = list(locate_all(_RIGHT_TOP_IMG, screenshot))
  if len(match_left_bottoms) == 3 and len(match_right_tops) == 3:
    for image_index in range(0, 3):
      match_left_bottom = match_left_bottoms[image_index]
      match_right_top = match_right_tops[image_index]
      if match_left_bottom.left >= match_right_top.left or match_left_bottom.top <= match_right_top.top:
        break
      skill_rect, skill_image = _crop_image(screenshot, match_left_bottom, match_right_top)
      # _detect_corner(screenshot, skill_rect)
      kind_name, skill_name, kind_image = _match_skill(skill_image)
      yield image_index, kind_name, skill_name, kind_image, skill_rect, skill_image
