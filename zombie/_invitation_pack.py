import collections
import os
import time
from typing import Dict, Optional

from PIL import Image
from pyscreeze import Box

from _game import config, distribute_file
from _image import img, save_image
from _locate import _box, locate_all, locate

Title = collections.namedtuple('Title',
                               'left_in_invitation height top_offset width_offset padding_top padding_bottom rescue_left')


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
    self.__TITLE: Title = Title(*tuple(_cfg('title').values()))
    self._PERSISTENT_DIR = _cfg('persistent-dir')

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

  def _title(self, invitation: Image.Image) -> (Optional[Image.Image], Optional[Image.Image], bool):
    title_end = locate(self._TITLE_END_IMG, invitation)
    if title_end is not None:
      record = _box(self.__TITLE.left_in_invitation, title_end.top + self.__TITLE.top_offset,
                    title_end.left + title_end.width - self.__TITLE.left_in_invitation - self.__TITLE.width_offset,
                    self.__TITLE.height)
      is_rescue = record.width > self.__TITLE.rescue_left
      if is_rescue:
        record = _box(record.left + self.__TITLE.rescue_left - 1, record.top,
                      record.width - self.__TITLE.rescue_left + 1, record.height)
      match = _box(record.left, record.top + self.__TITLE.padding_top,
                   record.width, record.height - self.__TITLE.padding_top - self.__TITLE.padding_bottom)
      return (invitation.crop((record.left, record.top, record.left + record.width, record.top + record.height)),
              invitation.crop((match.left, match.top, match.left + match.width, match.top + match.height)),
              is_rescue)
    return None, None, False

  def detect(self, invitation: Image.Image) -> (Optional[str], bool, Optional[Image.Image]):
    title_record, title_match, is_rescue = self._title(invitation)
    if title_record is not None:
      for invitation_name, item in self.invitations.items():
        if title_match.width <= item.width and locate(title_match, item) is not None:
          return invitation_name, is_rescue, title_record
    return None, is_rescue, title_record

  def _invitation_img(self, invitation_name: str) -> str:
    return img(f'{self._PERSISTENT_DIR}/{invitation_name}')

  def record(self, invitation: Image.Image) -> (Optional[str], bool, bool):
    invitation_name, is_rescue, title = self.detect(invitation)
    if invitation_name is not None:
      return invitation_name, is_rescue, False

    if title is None:
      return None, is_rescue, False

    invitation_name = f'invitation_{time.time()}'
    print(f'record invitation: {invitation_name} - {is_rescue} - {title}')
    self.invitations[invitation_name] = title

    save_image(title, self._invitation_img(invitation_name))

    return invitation_name, is_rescue, True

  def match_from_screenshot(self, screenshot: Image.Image):
    for box, invitation in self.split_from_screenshot(screenshot):
      invitation_name, is_rescue, title = self.detect(invitation)
      yield invitation_name, is_rescue, box, invitation, title

  def cancel(self, invitation: Box) -> Box:
    return _box(invitation.left + self._BTN_CANCEL.left, invitation.top + self._BTN_CANCEL.top,
                self._BTN_CANCEL.width, self._BTN_CANCEL.height)

  def confirm(self, invitation: Box) -> Box:
    return _box(invitation.left + self._BTN_CONFIRM.left, invitation.top + self._BTN_CONFIRM.top,
                self._BTN_CONFIRM.width, self._BTN_CONFIRM.height)
