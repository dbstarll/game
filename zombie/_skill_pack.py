import os
import time
from typing import Dict, Optional

from PIL import Image

from _game import config, distribute_file
from _image import img, save_image
from _skill_detect_normal import SkillDetectNormal
from _skill_subset import SkillSubset


def _cfg(path: str):
  return config(f'skill.detect.{path}')


class SkillPack:
  def __init__(self):
    self._RECORD_KIND = _cfg('record.kind')
    self._RECORD_SKILL = _cfg('record.skill')
    self._RECORD_SKIP = _cfg('record.skip')
    self._PERSISTENT_DIR = _cfg('persistent-dir')

    self.series: Dict[str, SkillSubset] = {}
    self.normal = SkillDetectNormal()
    self._load()

  def _load(self) -> None:
    total = 0
    for kind_name in os.listdir(distribute_file(self._PERSISTENT_DIR)):
      if os.path.isdir(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}')):
        skill_set = self._load_kind(kind_name)
        total += skill_set.size()
        self.series[kind_name] = skill_set
    print(f'加载技能: {total}, 类型: {len(self.series)}')

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

  def match_from_screenshot(self, screenshot: Image.Image):
    for image_index, skill_rect, skill_image in self.normal.match_from_screenshot(screenshot):
      kind_name, skill_name, kind_image = self._match_skill(skill_image)
      yield image_index, kind_name, skill_name, kind_image, skill_rect, skill_image

  def _match_kinds(self, kind_image: Image.Image) -> Optional[str]:
    for kind_name, skill_set in self.series.items():
      if skill_set.match_kind(kind_image) is not None:
        return kind_name

  def _match_skill(self, skill_image: Image.Image) -> (Optional[str], Optional[str], Image.Image):
    _, kind_image = self.normal.crop_kind_image(skill_image)
    kind_name = self._match_kinds(kind_image)
    if kind_name is None:
      return None, None, kind_image
    else:
      return kind_name, self.series[kind_name].match_skill(skill_image), kind_image

  def _record_kinds(self, kind_image) -> (str, bool):
    kind_name = self._match_kinds(kind_image)
    if kind_name is not None:
      return kind_name, False

    kind_name = f'logo-{time.time()}'
    print(f'\trecord kind: {kind_name} - {kind_image}')
    self.series[kind_name] = SkillSubset(kind_name).set_kind_image(kind_image)

    if not os.path.exists(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}')):
      os.mkdir(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}'))
    save_image(kind_image, self._skill_img(kind_name, kind_name))

    return kind_name, True

  def record(self, image_index: int, skill_image: Image.Image) -> (Optional[str], Optional[str], bool):
    if image_index == self._RECORD_SKIP:
      return None, None, False

    kind_name, skill_name, kind_image = self._match_skill(skill_image)
    if skill_name is not None:
      return kind_name, skill_name, False

    skill_set: SkillSubset
    if kind_name is not None:
      skill_set = self.series[kind_name]
    elif not self._RECORD_KIND:
      return None, None, False
    else:
      kind_name, _ = self._record_kinds(kind_image)
      skill_set = self.series[kind_name]

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
