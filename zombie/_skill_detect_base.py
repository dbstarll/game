import os
import time
from typing import Dict, Optional

from PIL import Image

from _game import config, distribute_file
from _image import img, save_image
from _skill_subset import SkillSubset


class SkillDetectBase:
  def __init__(self, style: str):
    self._style = style
    self._PERSISTENT_DIR = config(f'skill.detect.{style}.persistent-dir')
    self._RECORD_KIND = config(f'skill.detect.{style}.record.kind')
    self._RECORD_SKILL = config(f'skill.detect.{style}.record.skill')
    self._RECORD_SKIP = config(f'skill.detect.{style}.record.skip')
    self._series: Dict[str, SkillSubset] = {}
    self.__load()

  def __load(self) -> None:
    total = 0
    for kind_name in os.listdir(distribute_file(self._PERSISTENT_DIR)):
      if os.path.isdir(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}')):
        subset = self.__load_kind(kind_name)
        total += subset.size()
        self._series[kind_name] = subset
    print(f'加载技能[{self._style}]: {total}, 类型: {len(self._series)}')

  def __load_kind(self, kind_name: str) -> SkillSubset:
    subset = SkillSubset(kind_name, self._style)
    for skill_file_name in os.listdir(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}')):
      if skill_file_name.endswith('.png'):
        skill_name = skill_file_name[:-4]
        if skill_name.startswith('logo'):
          subset.set_kind_image(Image.open(self.__skill_img(kind_name, skill_name)))
        elif skill_name.startswith('skill_'):
          raise ValueError(f'临时技能文件需要被处理: {self._PERSISTENT_DIR}/{kind_name}/{skill_file_name}')
        else:
          subset.add_skill(skill_name, Image.open(self.__skill_img(kind_name, skill_name)))
    subset.summary()
    return subset

  def __skill_img(self, kind_name: str, skill_name: str) -> str:
    return img(f'{self._PERSISTENT_DIR}/{kind_name}/{skill_name}')

  def _crop_kind_image(self, image_index: int, skill_image: Image.Image) -> Image.Image:
    return skill_image

  def __match_kinds(self, kind_image: Image.Image) -> Optional[str]:
    for kind_name, subset in self._series.items():
      if subset.match_kind(kind_image) is not None:
        return kind_name

  def _match_skill(self, image_index: int, skill_image: Image.Image) -> (Optional[str], Optional[str], Image.Image):
    for kind_name, subset in self._series.items():
      skill_name = subset.match_skill(skill_image)
      if skill_name is not None:
        return kind_name, skill_name, subset._kind_image
    kind_image = self._crop_kind_image(image_index, skill_image)
    return self.__match_kinds(kind_image), None, kind_image

  def _match_skill2(self, image_index: int, skill_image: Image.Image) -> (Optional[str], Optional[str], Image.Image):
    kind_image = self._crop_kind_image(image_index, skill_image)
    kind_name = self.__match_kinds(kind_image)
    if kind_name is None:
      return None, None, kind_image
    else:
      return kind_name, self._series[kind_name].match_skill(skill_image), kind_image

  def __record_kinds(self, kind_image) -> (str, bool):
    kind_name = self.__match_kinds(kind_image)
    if kind_name is not None:
      return kind_name, False

    kind_name = f'logo_{time.time()}'
    print(f'\trecord kind: {kind_name} - {kind_image}')
    self._series[kind_name] = SkillSubset(kind_name, self._style).set_kind_image(kind_image)

    if not os.path.exists(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}')):
      os.mkdir(distribute_file(f'{self._PERSISTENT_DIR}/{kind_name}'))
    save_image(kind_image, self.__skill_img(kind_name, kind_name))

    return kind_name, True

  def record(self, image_index: int, skill_image: Image.Image) -> (Optional[str], Optional[str], bool):
    if image_index == self._RECORD_SKIP:
      return None, None, False

    kind_name, skill_name, kind_image = self._match_skill(image_index, skill_image)
    if skill_name is not None:
      return kind_name, skill_name, False

    subset: SkillSubset
    if kind_name is not None:
      subset = self._series[kind_name]
    elif not self._RECORD_KIND:
      return None, None, False
    else:
      kind_name, _ = self.__record_kinds(kind_image)
      subset = self._series[kind_name]

    if subset.size() > 0:
      skill_name = subset.match_skill(skill_image)
      if skill_name is not None:
        return kind_name, skill_name, False

    if not self._RECORD_SKILL:
      return kind_name, None, False

    skill_name = f'skill_{time.time()}'
    print(f'\trecord skill: {kind_name} - {skill_name} - {skill_image}')
    subset.add_skill(skill_name, skill_image)
    subset.summary()

    save_image(skill_image, self.__skill_img(kind_name, skill_name))
    return kind_name, skill_name, True
