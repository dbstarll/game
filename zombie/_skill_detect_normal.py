import os
import time
from typing import Dict, Optional, List

from PIL import Image, ImageDraw
from pyscreeze import Box, Point

from _game import config, distribute_file, get_distribute
from _image import img, save_image
from _locate import _box, locate_all
from _skill_subset import SkillSubset


def _cfg(path: str):
  return config(f'skill.detect.normal.{path}')


class SkillDetectNormal:
  def __init__(self):
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
    self._KIND_CORE_SIZE = _cfg('kind.core-size')
    self._RECORD_KIND = _cfg('record.kind')
    self._RECORD_SKILL = _cfg('record.skill')
    self._RECORD_SKIP = _cfg('record.skip')

    self._series: Dict[str, SkillSubset] = {}
    self._kind_corner_points: List[Point] = self._calculate_kind_corner_points(self._KIND_SIZE, self._KIND_CORE_SIZE)
    self._load()

  def _calculate_kind_corner_points(self, size: int, core_size: int) -> List[Point]:
    r, pr = (size - 1) / 2, (core_size // 2) ** 2
    points = []
    for x in range(0, size):
      for y in range(0, size):
        if (x - r) ** 2 + (y - r) ** 2 > pr:
          points.append(Point(x, y))
    return points

  def _load(self) -> None:
    total = 0
    for kind_name in os.listdir(distribute_file(self._PERSISTENT_DIR)):
      if os.path.isdir(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}')):
        skill_set = self._load_kind(kind_name)
        total += skill_set.size()
        self._series[kind_name] = skill_set
    print(f'加载技能: {total}, 类型: {len(self._series)}')

  def _load_kind(self, kind_name: str) -> SkillSubset:
    skill_set = SkillSubset(kind_name)
    for skill_file_name in os.listdir(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}')):
      if skill_file_name.endswith('.png'):
        skill_name = skill_file_name[:-4]
        if skill_name.startswith('logo'):
          skill_set.set_kind_image(Image.open(self._skill_img(kind_name, skill_name)))
        elif skill_name.startswith('skill-'):
          raise ValueError(f'临时技能文件需要被处理: {self._PERSISTENT_DIR}/{kind_name}/{skill_file_name}')
        else:
          skill_set.add_skill(skill_name, Image.open(self._skill_img(kind_name, skill_name)))
    skill_set.summary()
    return skill_set

  def _skill_img(self, kind_name: str, skill_name: str) -> str:
    return img(f'{self._PERSISTENT_DIR}/{kind_name}/{skill_name}')

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

  def _match_kinds(self, kind_image: Image.Image) -> Optional[str]:
    for kind_name, skill_set in self._series.items():
      if skill_set.match_kind(kind_image) is not None:
        return kind_name

  def _match_skill(self, image_index: int, skill_image: Image.Image) -> (Optional[str], Optional[str], Image.Image):
    kind_image = self._crop_kind_image(image_index, skill_image)
    kind_name = self._match_kinds(kind_image)
    if kind_name is None:
      return None, None, kind_image
    else:
      return kind_name, self._series[kind_name].match_skill(skill_image), kind_image

  def _record_kinds(self, kind_image) -> (str, bool):
    kind_name = self._match_kinds(kind_image)
    if kind_name is not None:
      return kind_name, False

    kind_name = f'logo-{time.time()}'
    print(f'\trecord kind: {kind_name} - {kind_image}')
    self._series[kind_name] = SkillSubset(kind_name).set_kind_image(kind_image)

    if not os.path.exists(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}')):
      os.mkdir(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}'))
    save_image(kind_image, self._skill_img(kind_name, kind_name))

    return kind_name, True

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

  def record(self, image_index: int, skill_image: Image.Image) -> (Optional[str], Optional[str], bool):
    if image_index == self._RECORD_SKIP:
      return None, None, False

    kind_name, skill_name, kind_image = self._match_skill(image_index, skill_image)
    if skill_name is not None:
      return kind_name, skill_name, False

    skill_set: SkillSubset
    if kind_name is not None:
      skill_set = self._series[kind_name]
    elif not self._RECORD_KIND:
      return None, None, False
    else:
      kind_name, _ = self._record_kinds(kind_image)
      skill_set = self._series[kind_name]

    if skill_set.size() > 0:
      skill_name = skill_set.match_skill(skill_image)
      if skill_name is not None:
        return kind_name, skill_name, False

    if not self._RECORD_SKILL:
      return kind_name, None, False

    skill_name = f'skill-{time.time()}'
    print(f'\trecord skill: {kind_name} - {skill_name} - {skill_image}')
    skill_set.add_skill(skill_name, skill_image)
    skill_set.summary()

    save_image(skill_image, self._skill_img(kind_name, skill_name))
    return kind_name, skill_name, True
