import os
import subprocess
import tempfile
import time

import pyautogui
import pyscreeze
from PIL import Image

from _locate import Box

_DISTRIBUTE = None
_GAME_WINDOW_ID = None
_GAME_WINDOW_RECT: pyscreeze.Box
_CLICK_INTERVAL = 0.2


def _raise_distribute_not_set():
  raise ValueError('distribute not set')


def raise_unknown_distribute():
  raise ValueError(f'unknown distribute: {_DISTRIBUTE}')


def distribute(args, default=None) -> str:
  global _DISTRIBUTE
  if len(args) > 1:
    _DISTRIBUTE = args[1]
  elif default is not None:
    _DISTRIBUTE = default
  else:
    _raise_distribute_not_set()
  return _DISTRIBUTE


def distribute_file(filename: str) -> str:
  if _DISTRIBUTE is None:
    _raise_distribute_not_set()
  return f'{_DISTRIBUTE}/{filename}'


def get_distribute():
  if _DISTRIBUTE is None:
    _raise_distribute_not_set()
  return _DISTRIBUTE


def _screenshot_wid(window_id):
  fh, filepath = tempfile.mkstemp(".png")
  os.close(fh)
  subprocess.call(["screencapture", "-l", window_id, "-o", "-x", filepath])
  im = Image.open(filepath)
  im.load()
  os.unlink(filepath)
  return im


def _screenshot_rect(rect):
  fh, filepath = tempfile.mkstemp(".png")
  os.close(fh)
  subprocess.call(["screencapture", "-R", f"{rect.left},{rect.top},{rect.width},{rect.height}", "-x", filepath])
  im = Image.open(filepath)
  im.load()
  os.unlink(filepath)
  return im


def screenshot():
  if _GAME_WINDOW_ID is not None:
    return _screenshot_wid(_GAME_WINDOW_ID)
  else:
    return _screenshot_rect(_GAME_WINDOW_RECT)


def _osascript(script_file):
  return subprocess.run(["osascript", script_file], capture_output=True).stdout.decode().rstrip('\n')


def _init_window():
  global _GAME_WINDOW_RECT
  position_result = _osascript(f'get-position-of-{_DISTRIBUTE}.scpt')
  if not position_result:
    print('游戏未启动')
    exit(1)
  position = tuple(map(int, position_result.split(', ')))
  size = tuple(map(int, _osascript(f'get-size-of-{_DISTRIBUTE}.scpt').split(', ')))
  _GAME_WINDOW_RECT = Box(position[0], position[1], size[0], size[1])
  print(f'获得游戏窗口: ID={_GAME_WINDOW_ID}, POS={_GAME_WINDOW_RECT}')
  return _GAME_WINDOW_RECT


def _init_mp():
  global _GAME_WINDOW_ID
  wid = _osascript('get-window-id-of-zombie.scpt')
  if wid:
    _GAME_WINDOW_ID = wid
    return _init_window()
  else:
    print('游戏未启动')
    exit(1)


def _init_ios():
  return _init_window()


def init_game():
  if _DISTRIBUTE is None:
    _raise_distribute_not_set()
  if 'mp' == _DISTRIBUTE:
    print('initializing mp...')
    return _init_mp()
  elif 'ios' == _DISTRIBUTE:
    print('initializing ios...')
    return _init_ios()
  else:
    raise_unknown_distribute()


def click(location, offset_x=0, offset_y=0, once=False):
  center = pyautogui.center(location)
  pyautogui.click(x=_GAME_WINDOW_RECT.left + center.x // 2 + offset_x,
                  y=_GAME_WINDOW_RECT.top + center.y // 2 + offset_y)
  time.sleep(_CLICK_INTERVAL)
  if not once:
    pyautogui.click(x=_GAME_WINDOW_RECT.left + center.x // 2 + offset_x,
                    y=_GAME_WINDOW_RECT.top + center.y // 2 + offset_y)
    time.sleep(_CLICK_INTERVAL)
