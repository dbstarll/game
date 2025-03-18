from typing import List

from PIL import Image, ImageDraw
from pyscreeze import Box, Point

from _game import config
from _image import img
from _locate import locate_all, _box
from _skill_detect_base import SkillDetectBase


def _cfg(path: str, default_none: bool = False):
  return config(f'skill.detect.elite.{path}', default_none)


def _calculate_kind_corner_points(size: int) -> List[Point]:
  r, pr = (size - 1) / 2, (size // 2) ** 2
  points = []
  for x in range(0, size):
    for y in range(0, size):
      if (x - r) ** 2 + (y - r) ** 2 > pr:
        points.append(Point(x, y))
  return points


class SkillDetectElite(SkillDetectBase):
  def __init__(self):
    super().__init__('elite')
    self._LEFT_BOTTOM_IMG = Image.open(img(_cfg('left-bottom-img')))
    self._RIGHT_TOP_IMG = Image.open(img(_cfg('right-top-img')))
    cfg_left_of_kind_img = _cfg('left-of-kind-img', True)
    self._LEFT_OF_KIND_IMG = None if cfg_left_of_kind_img is None else Image.open(img(cfg_left_of_kind_img))

    self._SKILL_OFFSET_LEFT = _cfg('skill-offset.left')
    self._SKILL_OFFSET_RIGHT = _cfg('skill-offset.right')
    self._SKILL_OFFSET_TOP = _cfg('skill-offset.top')
    self._SKILL_OFFSET_BOTTOM = _cfg('skill-offset.bottom')
    self._KIND_LEFT = _cfg('kind.left')
    self._KIND_TOP = _cfg('kind.top')
    self._KIND_SIZE = _cfg('kind.size')

    self._kind_corner_points: List[Point] = _calculate_kind_corner_points(self._KIND_SIZE)

  def _crop_skill_image(self, screenshot: Image.Image, lb, rt) -> (Box, Image.Image):
    box = _box(lb.left + self._SKILL_OFFSET_LEFT, rt.top + self._SKILL_OFFSET_TOP,
               rt.left + rt.width - lb.left - self._SKILL_OFFSET_LEFT + self._SKILL_OFFSET_RIGHT,
               lb.top + lb.height - rt.top - self._SKILL_OFFSET_TOP + self._SKILL_OFFSET_BOTTOM)
    return box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def _crop_skill_image_by_left_of_kind(self, screenshot: Image.Image, left_of_kind) -> (Box, Image.Image):
    box = _box(left_of_kind.left + self._SKILL_OFFSET_LEFT, left_of_kind.top + self._SKILL_OFFSET_TOP,
               left_of_kind.width + self._SKILL_OFFSET_RIGHT - self._SKILL_OFFSET_LEFT,
               left_of_kind.height + self._SKILL_OFFSET_BOTTOM - self._SKILL_OFFSET_TOP)
    return box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def _crop_kind_image(self, image_index: int, skill_image: Image.Image) -> (Box, Image.Image):
    box = _box(self._KIND_LEFT, self._KIND_TOP, self._KIND_SIZE, self._KIND_SIZE)
    return self._fill_kind_corner(skill_image.crop((box.left, box.top, box.left + box.width, box.top + box.height)))

  def _fill_kind_corner(self, kind_image: Image.Image) -> Image.Image:
    draw = ImageDraw.Draw(kind_image)
    draw.point(self._kind_corner_points)
    return kind_image

  def match_from_screenshot(self, screenshot: Image.Image):
    if self._LEFT_OF_KIND_IMG is None:
      return self.__match_by_left_bottom_and_right_top(screenshot)
    else:
      return self.__match_by_left_of_kind(screenshot)

  def __match_by_left_of_kind(self, screenshot: Image.Image):
    image_index: int = 0
    matches: List[Box] = []
    for match_left_of_kind in locate_all(self._LEFT_OF_KIND_IMG, screenshot):
      if len(matches) > 0 and match_left_of_kind.top - matches[len(matches) - 1].top > 5:
        yield self.__match_one_by_left_of_kinds(screenshot, image_index, matches)
        image_index += 1
        matches.clear()
      matches.append(match_left_of_kind)
    if len(matches) > 0:
      yield self.__match_one_by_left_of_kinds(screenshot, image_index, matches)

  def __match_one_by_left_of_kinds(self, screenshot: Image.Image, image_index: int, matches: List[Box]):
    skill_rect, skill_image = self._crop_skill_image_by_left_of_kind(screenshot, matches[len(matches) - 1])
    kind_name, skill_name, kind_image = self._match_skill(image_index, skill_image)
    return image_index, kind_name, skill_name, kind_image, skill_rect, skill_image

  def __match_by_left_bottom_and_right_top(self, screenshot: Image.Image):
    match_left_bottoms = list(locate_all(self._LEFT_BOTTOM_IMG, screenshot))
    match_right_tops = list(locate_all(self._RIGHT_TOP_IMG, screenshot))
    if len(match_left_bottoms) == len(match_right_tops) and len(match_left_bottoms) > 0:
      for image_index in range(0, len(match_left_bottoms)):
        match_left_bottom = match_left_bottoms[image_index]
        match_right_top = match_right_tops[image_index]
        if match_left_bottom.left >= match_right_top.left or match_left_bottom.top <= match_right_top.top:
          print(f'match_left_bottom: {match_left_bottom}, match_right_top: {match_right_top}')
          break
        skill_rect, skill_image = self._crop_skill_image(screenshot, match_left_bottom, match_right_top)
        kind_name, skill_name, kind_image = self._match_skill(image_index, skill_image)
        yield image_index, kind_name, skill_name, kind_image, skill_rect, skill_image
    else:
      print(f'match_left_bottoms: {len(match_left_bottoms)} - match_right_top: {len(match_right_tops)}')
