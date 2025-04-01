import os
import sys
import time
from typing import List

from PIL import Image

from _game import distribute, distribute_file
from _skill_pack import SkillPack


def rename(old_name: str, new_name: str):
  if rename_file:
    os.rename(old_name, new_name)


def detect_full_match(skills_file: str, detect_names: List[str]) -> int:
  part = skills_file.split("-")
  detects = 0
  if len(part) - 4 == len(detect_names):
    for i in range(0, len(detect_names)):
      detect = detect_names[i]
      expect = part[i + 3]
      if detect == expect:
        detects += 1
      else:
        part[i + 3] = detect
        print(f'expect: {expect}, detect: {detect} at {skills_file}')
    if detects != len(detect_names):
      print(f'\tre_full_match: {skills_file} -> {"-".join(part)}')
      rename(skills_file, "-".join(part))
  else:
    new_part = part[:3]
    for i in range(0, len(detect_names)):
      new_part.append(detect_names[i])
    new_part.append(part[len(part) - 1])
    print(f'\tre_full_match: {skills_file} -> {"-".join(new_part)}')
    rename(skills_file, "-".join(new_part))
  return detects


def detect_skills_from_file(skills_file: str) -> (int, int, int):
  matches = detects = records = 0
  skill_names = []
  with Image.open(skills_file) as im:
    for image_index, kind_name, skill_name, _, skill_rect, skill_image in skill_pack.match_elite_from_screenshot(
        im):
      matches += 1
      skill_names.append(skill_name)
      if skill_name is not None:
        detects += 1
      else:
        print(f'\t{image_index}: {kind_name} - {skill_name}, {skill_rect}, {skills_file}')
        _, _, new_skill = skill_pack.record_elite(image_index, skill_image)
        if new_skill:
          records += 1

  part = skills_file.split("-")
  if matches > 0 and matches == detects and 'full_match' == part[2]:
    return matches, detect_full_match(skills_file, skill_names), records

  if matches == 0:
    if len(part) == 3:
      part.insert(2, 'mismatch')
      rename(skills_file, "-".join(part))
      print(f'\tmismatch: {skills_file} -> {"-".join(part)}')
  elif detects == matches:
    if len(part) == 3:
      skill_names.insert(0, part[0])
      skill_names.insert(1, part[1])
      skill_names.insert(2, 'full_match')
      skill_names.append(part[2])
      rename(skills_file, "-".join(skill_names))
      print(f'\tfull_match: {skills_file} -> {"-".join(skill_names)}')
  else:
    print(f'\tpart detected:{skill_names} - {skills_file}')

  return matches, detects, records


if __name__ == "__main__":
  dist, _ = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')
  skill_pack = SkillPack()
  rename_file, reset = False, False

  files = full_matches = part_matches = mismatch = records = 0
  start = time.time()
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('elite-skills-17') and file.endswith('.png'):
      skills_file = f'tmp/{distribute_file(file)}'
      if reset:
        part = skills_file.split("-")
        if len(part) > 3:
          new_part = part[:2]
          new_part.append(part[len(part) - 1])
          print(f'\t{skills_file}: {new_part}')
          os.rename(skills_file, "-".join(new_part))
        continue

      if files > 0 and files % 100 == 0:
        print(f'{files} - {time.time() - start}')
      files += 1
      match_skills, detect_skills, record_skills = detect_skills_from_file(skills_file)
      records += record_skills
      if detect_skills > 0 and detect_skills == match_skills:
        full_matches += 1
      elif match_skills > 0:
        part_matches += 1
      else:
        mismatch += 1
  print(
    f'files: {files}, full_matches: {full_matches}, part_matches: {part_matches}, mismatch: {mismatch}, records: {records}')
  print(f'cost - {time.time() - start}')
