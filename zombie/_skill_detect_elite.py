import time
from typing import Optional, Dict

from PIL import Image
from pyscreeze import Box

from _game import config
from _image import img, save_image
from _locate import locate_all, _box, locate


def _cfg(path: str):
  return config(f'skill.detect.elite.{path}')


class SkillDetectElite:
  def __init__(self):
    self._PERSISTENT_DIR = _cfg('persistent-dir')

    self._LEFT_BOTTOM_IMG = Image.open(img(_cfg('left-bottom-img')))
    self._RIGHT_TOP_IMG = Image.open(img(_cfg('right-top-img')))

    self._SKILL_OFFSET_LEFT = _cfg('skill-offset.left')
    self._SKILL_OFFSET_RIGHT = _cfg('skill-offset.right')
    self._SKILL_OFFSET_TOP = _cfg('skill-offset.top')
    self._SKILL_OFFSET_BOTTOM = _cfg('skill-offset.bottom')
    self._KIND_OFFSET_WIDTH = _cfg('kind.offset-width')
    self._KIND_OFFSET_HEIGHT = _cfg('kind.offset-height')
    self._KIND_WIDTH = _cfg('kind.width')
    self._RECORD_KIND = _cfg('record.kind')

    self._series: Dict[str, Image.Image] = {}

  def _kind_img(self, kind_name: str) -> str:
    return img(f'{self._PERSISTENT_DIR}/{kind_name}')

  def _crop_skill_image(self, screenshot: Image.Image, lb, rt) -> (Box, Image.Image):
    box = _box(lb.left + self._SKILL_OFFSET_LEFT, rt.top + self._SKILL_OFFSET_TOP,
               rt.left + rt.width - lb.left - self._SKILL_OFFSET_LEFT + self._SKILL_OFFSET_RIGHT,
               lb.top + lb.height - rt.top - self._SKILL_OFFSET_TOP + self._SKILL_OFFSET_BOTTOM)
    return box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def _crop_kind_image(self, skill_image: Image.Image) -> (Box, Image.Image):
    box = _box(self._KIND_OFFSET_WIDTH, self._KIND_OFFSET_HEIGHT, self._KIND_WIDTH, self._KIND_WIDTH)
    return box, skill_image.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def _match_kinds(self, kind_image: Image.Image) -> Optional[str]:
    for kind_name, item in self._series.items():
      if locate(kind_image, item) is not None:
        return kind_name

  def _match_skill(self, skill_image: Image.Image) -> (Optional[str], Optional[str], Image.Image):
    _, kind_image = self._crop_kind_image(skill_image)
    kind_name = self._match_kinds(kind_image)
    if kind_name is None:
      return None, None, kind_image
    else:
      return kind_name, None, kind_image

  def _record_kinds(self, kind_image) -> (str, bool):
    kind_name = self._match_kinds(kind_image)
    if kind_name is not None:
      return kind_name, False

    kind_name = f'logo-{time.time()}'
    print(f'\trecord kind: {kind_name} - {kind_image}')
    self._series[kind_name] = kind_image

    save_image(kind_image, self._kind_img(kind_name))

    return kind_name, True

  def match_from_screenshot(self, screenshot: Image.Image):
    match_left_bottoms = list(locate_all(self._LEFT_BOTTOM_IMG, screenshot))
    match_right_tops = list(locate_all(self._RIGHT_TOP_IMG, screenshot))
    if len(match_left_bottoms) == len(match_right_tops) and len(match_left_bottoms) > 0:
      for image_index in range(0, len(match_left_bottoms)):
        match_left_bottom = match_left_bottoms[image_index]
        match_right_top = match_right_tops[image_index]
        if match_left_bottom.left >= match_right_top.left or match_left_bottom.top <= match_right_top.top:
          break
        skill_rect, skill_image = self._crop_skill_image(screenshot, match_left_bottom, match_right_top)
        kind_name, skill_name, kind_image = self._match_skill(skill_image)
        yield image_index, kind_name, skill_name, kind_image, skill_rect, skill_image
    else:
      print(f'match_left_bottoms: {len(match_left_bottoms)} - match_right_top: {len(match_right_tops)}')

  def record(self, image_index: int, skill_image: Image.Image) -> (Optional[str], bool):
    kind_name, skill_name, kind_image = self._match_skill(skill_image)
    if skill_name is not None:
      return kind_name, False
    elif kind_name is not None:
      return kind_name, False
    elif not self._RECORD_KIND:
      return None, False
    else:
      return self._record_kinds(kind_image)
