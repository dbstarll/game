import os
import sys

from PIL import Image

from _game import distribute, distribute_file, config
from _invitation_pack import InvitationPack


def _cfg(path: str):
  return config(f'team-invitation.invitation.{path}')


def detect_invitations_from_file(fights_file: str) -> (int, int, int):
  matches, detects, records = 0, 0, 0
  invitation_names = []
  with Image.open(fights_file) as im:
    for invitation_name, is_rescue, box, invitation, title in pack.match_from_screenshot(im):
      matches += 1
      if invitation_name is not None:
        detects += 1
        invitation_names.append(''.join(invitation_name.split('-')[:2 if is_rescue else 1]))
      else:
        invitation_names.append('none')
        if is_rescue:
          _, _, create = pack.record(invitation)
          if create:
            records += 1

  part = fights_file.split("-")
  if matches == 0:
    if len(part) == 2:
      part.insert(1, 'mismatch')
      os.rename(fights_file, "-".join(part))
    print(f'\tmismatch: {fights_file} -> {"-".join(part)}')
  elif matches == detects:
    if len(part) == 2:
      invitation_names.insert(0, part[0])
      invitation_names.insert(1, 'full_match')
      invitation_names.append(part[1])
      os.rename(fights_file, "-".join(invitation_names))
  else:
    if len(part) == 2:
      invitation_names.insert(0, part[0])
      invitation_names.insert(1, 'part_match')
      invitation_names.append(part[1])
      os.rename(fights_file, "-".join(invitation_names))

  return matches, detects, records


if __name__ == "__main__":
  dist, _ = distribute(sys.argv, "ios")
  print(f'游戏发行版本: {dist}')

  pack = InvitationPack()

  files = 0
  matches = 0
  detects = 0
  records = 0
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('fights-17') and file.endswith('.png'):
      full_path = f'tmp/{distribute_file(file)}'
      files += 1
      m, d, r = detect_invitations_from_file(full_path)
      matches += m
      detects += d
      records += r

  print(f'match: {matches} on {files} files, detected: {detects}, records: {records}')
