import os
import sys
import time

from PIL import Image

from _game import distribute
from _skill import load_skills, match_skills_from_screenshot, record_skill


def detect_skills_from_file(skills_file):
  matches = 0
  detects = 0
  skill_names = []
  with Image.open(skills_file) as im:
    for image_index, kind_name, skill_name, kind_image, skill_rect, skill_image in match_skills_from_screenshot(im):
      matches += 1
      skill_names.append(skill_name)
      if skill_name is not None:
        detects += 1
      else:
        print(f'\t{image_index}: {kind_name} - {skill_name}, {skill_rect}, {skills_file}')
        record_skill(image_index, kind_name, kind_image, skill_image)

  part = skills_file.split("-")
  if matches == 0:
    if len(part) == 2:
      part.insert(1, 'mismatch')
      os.rename(skills_file, "-".join(part))
    print(f'\tmismatch: {skills_file} -> {"-".join(part)}')
  elif detects == 3:
    if len(part) == 2:
      skill_names.insert(0, part[0])
      skill_names.insert(1, 'full_match')
      skill_names.append(part[1])
      os.rename(skills_file, "-".join(skill_names))
  else:
    print(f'\tpart detected:{skill_names} -  {skills_file}')
  return matches, detects


if __name__ == "__main__":
  dist = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')

  files, full_matches, part_matches, mismatch = 0, 0, 0, 0
  skills = load_skills()
  start = time.time()
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('skills-full_match-') and file.endswith('.png'):
      skills_file = f'tmp/{dist}/{file}'
      if files > 0 and files % 100 == 0:
        print(f'{files} - {time.time() - start}')
      files += 1
      match_skills, detect_skills = detect_skills_from_file(skills_file)
      if detect_skills == 3:
        full_matches += 1
      elif match_skills == 3:
        part_matches += 1
      else:
        mismatch += 1
  print(f'files: {files}, full_matches: {full_matches}, part_matches: {part_matches}, mismatch: {mismatch}')
  print(f'kinds: {len(skills)}')
  for kind, skill_pack in skills.items():
    print(f'skills: {kind} - {skill_pack.size()}')
  print(f'cost - {time.time() - start}')
  # cost - 407.6763346195221  files: 1745, full_matches: 1703, part_matches: 42, mismatch: 0
