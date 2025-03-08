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
    self._TITLE_END_IMG = Image.open(img(_cfg('title-end-img')))
    self._ITEM_WIDTH, self._ITEM_HEIGHT = _cfg('item.width'), _cfg('item.height')
    self._ITEM_OFFSET_RIGHT, self._ITEM_OFFSET_BOTTOM = _cfg('item.offset-right'), _cfg('item.offset-bottom')
    self._BTN_CANCEL = _box(*tuple(_cfg('btn-cancel').values()))
    self._BTN_CONFIRM = _box(*tuple(_cfg('btn-confirm').values()))
    self._BTN_TITLE = _box(*tuple(_cfg('title').values()))
    self._PERSISTENT_DIR = _cfg('persistent-dir')
    self._TITLE_RESCUE_LEFT_OFFSET = _cfg('title-rescue-left-offset')
    self._TITLE_RESCUE_TOP_OFFSET = _cfg('title-rescue-top-offset')
    self._TITLE_RESCUE_BOTTOM_OFFSET = _cfg('title-rescue-bottom-offset')

    self.invitations: Dict[str, Image.Image] = {}
    self._load()

  def _load(self):
    for file in os.listdir(distribute_file(self._PERSISTENT_DIR)):
      if file.endswith('.png'):
        invitation_name = file[:-4]
        self.invitations[invitation_name] = Image.open(self._invitation_img(invitation_name))
    print(f'加载关卡: {len(self.invitations)}')

  def split_from_screenshot(self, screenshot: Image.Image):
    for match in locate_all(self._LOCATE_IMG, screenshot):
      box = _box(match.left + match.width + self._ITEM_OFFSET_RIGHT - self._ITEM_WIDTH,
                 match.top + match.height + self._ITEM_OFFSET_BOTTOM - self._ITEM_HEIGHT,
                 self._ITEM_WIDTH, self._ITEM_HEIGHT)
      yield box, screenshot.crop((box.left, box.top, box.left + box.width, box.top + box.height))

  def _title(self, invitation: Image.Image) -> (Optional[Image.Image], bool):
    title_end = locate(self._TITLE_END_IMG, invitation)
    if title_end is not None:
      width = title_end.left + title_end.width - self._BTN_TITLE.left
      if width > self._TITLE_RESCUE_LEFT_OFFSET:
        box = _box(self._BTN_TITLE.left + self._TITLE_RESCUE_LEFT_OFFSET - 1,
                   title_end.top + self._TITLE_RESCUE_TOP_OFFSET, width - self._TITLE_RESCUE_LEFT_OFFSET + 1,
                   self._BTN_TITLE.height - self._TITLE_RESCUE_TOP_OFFSET - self._TITLE_RESCUE_BOTTOM_OFFSET)
        return invitation.crop((box.left, box.top, box.left + box.width, box.top + box.height)), True
      else:
        box = _box(self._BTN_TITLE.left, title_end.top, width, self._BTN_TITLE.height)
        return invitation.crop((box.left, box.top, box.left + box.width, box.top + box.height)), False
    return None, False

  def detect(self, invitation: Image.Image) -> (Optional[str], bool):
    title, is_rescue = self._title(invitation)
    if title is not None:
      for invitation_name, item in self.invitations.items():
        if title.width <= item.width and locate(title, item) is not None:
          return invitation_name, is_rescue, title
    return None, is_rescue, title

  def _invitation_img(self, invitation_name: str) -> str:
    return img(f'{self._PERSISTENT_DIR}/{invitation_name}')

  def record(self, invitation: Image.Image) -> (str, bool):
    invitation_name, is_rescue, title = self.detect(invitation)
    if invitation_name is not None:
      return invitation_name, is_rescue, False

    if title is None:
      return None, is_rescue, False

    invitation_name = f'invitation-{time.time()}'
    print(f'record invitation: {invitation_name} - {is_rescue} - {title}')
    self.invitations[invitation_name] = title

    save_image(title, self._invitation_img(invitation_name))

    return invitation_name, is_rescue, True

  def match_from_screenshot(self, screenshot: Image.Image):
    for box, invitation in self.split_from_screenshot(screenshot):
      invitation_name, is_rescue, title = self.detect(invitation)
      yield invitation_name, is_rescue, box, invitation, title
