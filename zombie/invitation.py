import os
import sys

from PIL import Image

from _game import distribute, distribute_file, config
from _invitation_pack import InvitationPack


def _cfg(path: str):
  return config(f'team-invitation.invitation.{path}')


if __name__ == "__main__":
  dist, _ = distribute(sys.argv, "mp")
  print(f'游戏发行版本: {dist}')

  pack = InvitationPack()

  files = 0
  matches = 0
  detects = 0
  records = 0
  for file in os.listdir(f'tmp/{dist}'):
    if file.startswith('fights-') and file.endswith('.png'):
      full_path = f'tmp/{distribute_file(file)}'
      files += 1
      with Image.open(full_path) as im:
        for invitation_name, is_rescue, box, invitation, title in pack.match_from_screenshot(im):
          matches += 1
          if invitation_name is not None:
            detects += 1
            print(f'\t{invitation_name} - {is_rescue} - {full_path}')
          elif is_rescue:
            _, _, create = pack.record(invitation)
            if create:
              records += 1

  print(f'match: {matches} on {files} files, detected: {detects}, records: {records}')
