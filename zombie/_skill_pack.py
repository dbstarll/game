from typing import Optional, Dict, List

from PIL import Image
from pyscreeze import Box

from _locate import locate, locate_all


class SkillPack:
  def __init__(self, kind_name: str):
    self.kind_name: str = kind_name
    self.skills: Dict[str, Image] = {}
    self.width: int = 0
    self.height: int = 0

  def set_kind_image(self, kind_image: Image):
    self.kind_image = kind_image
    return self

  def add_skill(self, skill_name: str, skill_image: Image):
    self.skills[skill_name] = skill_image
    self.width = max(self.width, skill_image.width)
    self.height = max(self.height, skill_image.height)

  def summary(self) -> None:
    self.image_all_in_one: Image = Image.new('RGBA', (self.width * self.size(), self.height))
    self.names: List[str] = []
    for skill_name, skill_image in self.skills.items():
      self.image_all_in_one.paste(skill_image, (self.width * len(self.names), 0))
      self.names.append(skill_name)
    print(f'summary: {self.kind_name} - {len(self.names)} - {self.image_all_in_one}')

  def size(self) -> int:
    return len(self.skills)

  def match_kind(self, kind_image: Image) -> Optional[Box]:
    return locate(kind_image, self.kind_image)

  def match_skill(self, skill_image: str) -> Optional[str]:
    return self._match_skill_from_all_in_one(skill_image)

  def _match_skill_from_skills(self, skill_image: str) -> Optional[str]:
    for skill_name, item in self.skills.items():
      if locate(skill_image, item, grayscale=True) is not None:
        return skill_name

  def _match_skill_from_all_in_one(self, skill_image: str) -> Optional[str]:
    match: List[int] = []
    for box in locate_all(skill_image, self.image_all_in_one, grayscale=True):
      index = int(round(box.left / self.width))
      if match.count(index) == 0:
        match.append(index)
        if len(match) > 1:
          print(f'match more then one: {index} - {self.names[index]}')
    return self.names[match[0]] if len(match) >= 1 else None
