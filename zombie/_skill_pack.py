from typing import Optional

from PIL import Image

from _game import config
from _skill_detect_elite import SkillDetectElite
from _skill_detect_normal import SkillDetectNormal


def _cfg(path: str):
  return config(f'skill.detect.{path}')


class SkillPack:
  def __init__(self):
    self.normal = SkillDetectNormal()
    self.elite = SkillDetectElite(self.normal)

  def match_from_screenshot(self, screenshot: Image.Image):
    return self.normal.match_from_screenshot(screenshot)

  def match_elite_from_screenshot(self, screenshot: Image.Image):
    return self.elite.match_from_screenshot(screenshot)

  def record(self, image_index: int, skill_image: Image.Image) -> (Optional[str], Optional[str], bool):
    return self.normal.record(image_index, skill_image)

  def record_elite(self, image_index: int, skill_image: Image.Image) -> (Optional[str], bool):
    return self.elite.record(image_index, skill_image)
