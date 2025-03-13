from PIL import Image
from pyscreeze import Box

from _game import config
from _image import img
from _locate import _box, locate_all


def _cfg(path: str):
  return config(f'skill.detect.normal.{path}')


class SkillDetectNormal:
  def __init__(self):
    self._LEFT_BOTTOM_IMG = Image.open(img(_cfg('left-bottom-img')))
    self._RIGHT_TOP_IMG = Image.open(img(_cfg('right-top-img')))
    self._SKILL_OFFSET_LEFT = _cfg('skill-offset.left')
    self._SKILL_OFFSET_RIGHT = _cfg('skill-offset.right')
    self._SKILL_OFFSET_TOP = _cfg('skill-offset.top')
    self._SKILL_OFFSET_BOTTOM = _cfg('skill-offset.bottom')
    self._KIND_OFFSET_WIDTH = _cfg('kind-offset.width')
    self._KIND_OFFSET_HEIGHT = _cfg('kind-offset.height')

  def _crop_skill_image(self, screenshot: Image.Image, lb, rt) -> (Box, Image.Image):
    box = _box(lb.left + self._SKILL_OFFSET_LEFT, rt.top + self._SKILL_OFFSET_TOP,
               rt.left + rt.width - lb.left - self._SKILL_OFFSET_LEFT + self._SKILL_OFFSET_RIGHT,
               lb.top + lb.height - rt.top - self._SKILL_OFFSET_TOP + self._SKILL_OFFSET_BOTTOM)
    return box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def crop_kind_image(self, skill_image: Image.Image) -> (Box, Image.Image):
    box = _box(self._KIND_OFFSET_WIDTH, self._KIND_OFFSET_HEIGHT,
               skill_image.width - 2 * self._KIND_OFFSET_WIDTH,
               skill_image.width - 2 * self._KIND_OFFSET_WIDTH)
    return box, skill_image.crop((box.left, box.top, box.left + box.width, box.top + box.height))

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
        yield image_index, skill_rect, skill_image
    else:
      print(f'match_left_bottoms: {len(match_left_bottoms)} - match_right_top: {len(match_right_tops)}')
