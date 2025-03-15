import os
import time
from typing import Optional, Dict, List

from PIL import Image, ImageDraw
from pyscreeze import Box, Point

from _game import config, distribute_file
from _image import img, save_image
from _locate import locate_all, _box, locate
from _skill_detect_normal import SkillDetectNormal


def _cfg(path: str):
  return config(f'skill.detect.elite.{path}')


class SkillDetectElite:
  def __init__(self, normal: SkillDetectNormal):
    self._PERSISTENT_DIR = _cfg('persistent-dir')

    self._LEFT_BOTTOM_IMG = Image.open(img(_cfg('left-bottom-img')))
    self._RIGHT_TOP_IMG = Image.open(img(_cfg('right-top-img')))

    self._SKILL_OFFSET_LEFT = _cfg('skill-offset.left')
    self._SKILL_OFFSET_RIGHT = _cfg('skill-offset.right')
    self._SKILL_OFFSET_TOP = _cfg('skill-offset.top')
    self._SKILL_OFFSET_BOTTOM = _cfg('skill-offset.bottom')
    self._KIND_LEFT = _cfg('kind.left')
    self._KIND_TOP = _cfg('kind.top')
    self._KIND_SIZE = _cfg('kind.size')
    self._RECORD_KIND = _cfg('record.kind')

    self._normal = normal
    self._series: Dict[str, Image.Image] = {}
    self._kind_corner_points: List[Point] = self._calculate_kind_corner_points(self._KIND_SIZE)
    self._load()

  def _calculate_kind_corner_points(self, size: int) -> List[Point]:
    r, pr = (size - 1) / 2, (size // 2) ** 2
    points = []
    for x in range(0, size):
      for y in range(0, size):
        if (x - r) ** 2 + (y - r) ** 2 > pr:
          points.append(Point(x, y))
    return points

  def _load(self) -> None:
    for kind_name in os.listdir(distribute_file(self._PERSISTENT_DIR)):
      if kind_name.endswith('.png'):
        kind_name = kind_name[:-4]
        self._series[kind_name] = Image.open(self._kind_img(kind_name))
    print(f'加载技能类型: {len(self._series)}')

  def _kind_img(self, kind_name: str) -> str:
    return img(f'{self._PERSISTENT_DIR}/{kind_name}')

  def _crop_skill_image(self, screenshot: Image.Image, lb, rt) -> (Box, Image.Image):
    box = _box(lb.left + self._SKILL_OFFSET_LEFT, rt.top + self._SKILL_OFFSET_TOP,
               rt.left + rt.width - lb.left - self._SKILL_OFFSET_LEFT + self._SKILL_OFFSET_RIGHT,
               lb.top + lb.height - rt.top - self._SKILL_OFFSET_TOP + self._SKILL_OFFSET_BOTTOM)
    return box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def _crop_kind_image(self, skill_image: Image.Image) -> (Box, Image.Image):
    box = _box(self._KIND_LEFT, self._KIND_TOP, self._KIND_SIZE, self._KIND_SIZE)
    return self._fill_kind_corner(skill_image.crop((box.left, box.top, box.left + box.width, box.top + box.height)))

  def _fill_kind_corner(self, kind_image: Image.Image) -> Image.Image:
    draw = ImageDraw.Draw(kind_image)
    draw.point(self._kind_corner_points)
    return kind_image

  def _match_kinds(self, kind_image: Image.Image) -> Optional[str]:
    for kind_name, item in self._series.items():
      if locate(kind_image, item) is not None:
        return kind_name

  def _match_skill(self, skill_image: Image.Image) -> (Optional[str], Optional[str], Image.Image):
    kind_image = self._crop_kind_image(skill_image)
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
