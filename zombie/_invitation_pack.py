import os
import time
from typing import Dict, Optional

from PIL import Image

from _game import config, distribute_file
from _image import img, save_image
from _locate import _box, locate_all, locate


def _cfg(path: str):
  return config(f'team-invitation.invitation.{path}')


class InvitationPack:
  def __init__(self):
    self._LOCATE_IMG = Image.open(img(_cfg('locate-img')))
    self._ITEM_WIDTH, self._ITEM_HEIGHT = _cfg('item.width'), _cfg('item.height')
    self._ITEM_OFFSET_RIGHT, self._ITEM_OFFSET_BOTTOM = _cfg('item.offset-right'), _cfg('item.offset-bottom')
    self._BTN_CANCEL = _box(*tuple(_cfg('btn-cancel').values()))
    self._BTN_CONFIRM = _box(*tuple(_cfg('btn-confirm').values()))
    self._BTN_TITLE = _box(*tuple(_cfg('title').values()))
    self._PERSISTENT_DIR = _cfg('persistent-dir')
    self.invitations: Dict[str, Image.Image] = {}
    self._load()

  def _load(self):
    for file in os.listdir(distribute_file(self._PERSISTENT_DIR)):
      if file.endswith('.png'):
        invitation_name = file[:-4]
        self.invitations[invitation_name] = Image.open(self._invitation_img(invitation_name))
    print(f'加载关卡: {len(self.invitations)}')

  def match_from_screenshot(self, screenshot: Image.Image):
    for match in locate_all(self._LOCATE_IMG, screenshot):
      box = _box(match.left + match.width + self._ITEM_OFFSET_RIGHT - self._ITEM_WIDTH,
                 match.top + match.height + self._ITEM_OFFSET_BOTTOM - self._ITEM_HEIGHT,
                 self._ITEM_WIDTH, self._ITEM_HEIGHT)
      yield box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def _title(self, invitation: Image.Image) -> Image.Image:
    box = self._BTN_TITLE
    return invitation.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def detect(self, invitation: Image.Image) -> Optional[str]:
    title = self._title(invitation)
    for invitation_name, item in self.invitations.items():
      if locate(title, item) is not None:
        return invitation_name

  def _invitation_img(self, invitation_name: str) -> str:
    return img(f'{self._PERSISTENT_DIR}/{invitation_name}')

  def record(self, invitation: Image.Image) -> (str, bool):
    invitation_name = self.detect(invitation)
    if invitation_name is not None:
      return invitation_name, False

    title = self._title(invitation)
    invitation_name = f'invitation-{time.time()}'
    print(f'\trecord invitation: {invitation_name} - {title}')
    self.invitations[invitation_name] = title

    save_image(title, self._invitation_img(invitation_name))

    return invitation_name, True
