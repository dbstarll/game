import os
import sys
import time

from PIL import Image

from _game import distribute, distribute_file, config
from _invitation_pack import InvitationPack


def _cfg(path: str):
  return config(f'team-invitation.invitation.{path}')


def rename(old_name: str, new_name: str):
  if rename_file:
    os.rename(old_name, new_name)


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
      rename(fights_file, "-".join(part))
      print(f'\tmismatch: {fights_file} -> {"-".join(part)}')
  elif matches == detects:
    if len(part) == 2:
      invitation_names.insert(0, part[0])
      invitation_names.insert(1, 'full_match')
      invitation_names.append(part[1])
      rename(fights_file, "-".join(invitation_names))
      print(f'\tfull_match: {fights_file} -> {"-".join(invitation_names)}')
  else:
    print(f'\tpart detected:{invitation_names} - {fights_file}')

  return matches, detects, records


if __name__ == "__main__":
  dist, _ = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')
  pack = InvitationPack()
  rename_file, reset = False, False

  files = matches = detects = records = 0
  start = time.time()
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('fights') and file.endswith('.png'):  # -1740769503.3104131
      fights_file = f'tmp/{distribute_file(file)}'
      if reset:
        part = fights_file.split("-")
        if len(part) > 2:
          new_part = part[:1]
          new_part.append(part[len(part) - 1])
          print(f'\t{fights_file}: {"-".join(new_part)}')
          os.rename(fights_file, "-".join(new_part))
        continue

      if files > 0 and files % 100 == 0:
        print(f'{files} - {time.time() - start}')
      files += 1
      m, d, r = detect_invitations_from_file(fights_file)
      matches += m
      detects += d
      records += r

  print(f'match: {matches} on {files} files, detected: {detects}, records: {records}')
