from typing import List

from PIL import Image, ImageDraw
from pyscreeze import Box, Point

from _game import config, get_distribute
from _image import img
from _locate import _box, locate_all
from _skill_detect_base import SkillDetectBase


def _cfg(path: str):
  return config(f'skill.detect.normal.{path}')


def _calculate_kind_corner_points(size: int, core_size: int) -> List[Point]:
  r, pr = (size - 1) / 2, (core_size // 2) ** 2
  points = []
  for x in range(0, size):
    for y in range(0, size):
      if (x - r) ** 2 + (y - r) ** 2 > pr:
        points.append(Point(x, y))
  return points


class SkillDetectNormal(SkillDetectBase):
  def __init__(self):
    super().__init__('normal')
    self._LEFT_BOTTOM_IMG = Image.open(img(_cfg('left-bottom-img')))
    self._RIGHT_TOP_IMG = Image.open(img(_cfg('right-top-img')))

    self._SKILL_OFFSET_LEFT = _cfg('skill-offset.left')
    self._SKILL_OFFSET_RIGHT = _cfg('skill-offset.right')
    self._SKILL_OFFSET_TOP = _cfg('skill-offset.top')
    self._SKILL_OFFSET_BOTTOM = _cfg('skill-offset.bottom')
    self._KIND_LEFT = _cfg('kind.left')
    self._KIND_TOP = _cfg('kind.top')
    self._KIND_SIZE = _cfg('kind.size')
    self._KIND_CORE_SIZE = _cfg('kind.core-size')

    self._kind_corner_points: List[Point] = _calculate_kind_corner_points(self._KIND_SIZE, self._KIND_CORE_SIZE)

  def _crop_skill_image(self, screenshot: Image.Image, lb, rt) -> (Box, Image.Image):
    box = _box(lb.left + self._SKILL_OFFSET_LEFT, rt.top + self._SKILL_OFFSET_TOP,
               rt.left + rt.width - lb.left - self._SKILL_OFFSET_LEFT + self._SKILL_OFFSET_RIGHT,
               lb.top + lb.height - rt.top - self._SKILL_OFFSET_TOP + self._SKILL_OFFSET_BOTTOM)
    return box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def _crop_kind_image(self, image_index: int, skill_image: Image.Image) -> Image.Image:
    box = _box(self._KIND_LEFT - (1 if image_index == 1 and get_distribute() == 'mp' else 0),
               self._KIND_TOP, self._KIND_SIZE, self._KIND_SIZE)
    return self._fill_kind_corner(skill_image.crop((box.left, box.top, box.left + box.width, box.top + box.height)))

  def _fill_kind_corner(self, kind_image: Image.Image) -> Image.Image:
    draw = ImageDraw.Draw(kind_image)
    draw.point(self._kind_corner_points)
    return kind_image

  def match_from_screenshot(self, screenshot: Image.Image):
    match_left_bottoms = list(locate_all(self._LEFT_BOTTOM_IMG, screenshot))
    match_right_tops = list(locate_all(self._RIGHT_TOP_IMG, screenshot))
    if len(match_left_bottoms) == 3 and len(match_right_tops) == 3:
      for image_index in range(0, 3):
        match_left_bottom = match_left_bottoms[image_index]
        match_right_top = match_right_tops[image_index]
        if match_left_bottom.left >= match_right_top.left or match_left_bottom.top <= match_right_top.top:
          break
        skill_rect, skill_image = self._crop_skill_image(screenshot, match_left_bottom, match_right_top)
        kind_name, skill_name, kind_image = self._match_skill(image_index, skill_image)
        yield image_index, kind_name, skill_name, kind_image, skill_rect, skill_image
    else:
      print(f'match_left_bottoms: {len(match_left_bottoms)} - match_right_top: {len(match_right_tops)}')
