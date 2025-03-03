import os
import sys

from PIL import Image

from _game import distribute
from _skill import load_skills, match_skills_from_screenshot, record_skill


def detect_skills(skills_file):
  detects = 0
  with Image.open(skills_file) as im:
    for image_index, kind_name, skill_name, kind_image, skill_rect, skill_image in match_skills_from_screenshot(im):
      if skill_name is not None:
        detects += 1
      else:
        print(f'\t{image_index}: {kind_name} - {skill_name}, {skill_rect}, {skills_file}')
        record_skill(image_index, kind_name, kind_image, skill_image)

  return detects == 3


if __name__ == "__main__":
  dist = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')

  total = 0
  matches = 0
  kinds, skills = load_skills()
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('skills-') and file.endswith('.png'):
      skills_file = f'tmp/{dist}/{file}'
      total += 1
      if detect_skills(skills_file):
        matches += 1
  print(f'matches: {matches}/{total}')
  print(f'kinds: {len(kinds)}')
  for kind, kind_skills in skills.items():
    print(f'skills: {kind} - {len(kind_skills)}')
