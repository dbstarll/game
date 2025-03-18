import os
import subprocess
import tempfile
import time
from typing import Optional, Dict, Any

import pyautogui
import yaml
from PIL import Image
from pyscreeze import Box

from _ios import obtain_game_ui_ios
from _locate import _box
from _mp import obtain_game_ui_mp

_DISTRIBUTE: Optional[str] = None
_GAME_WINDOW_ID: Optional[str] = None
_GAME_WINDOW_RECT: Optional[Box] = None
_GAME_UI_RECT: Optional[Box] = None
_CLICK_INTERVAL = 0.2
_CONFIG: Optional[Dict[str, Any]] = None


def _error_distribute_not_set() -> ValueError:
  return ValueError('distribute not set')


def error_unknown_distribute() -> ValueError:
  return ValueError(f'unknown distribute: {_DISTRIBUTE}')


def _load_game_config(config_file: str):
  rf = open(file=distribute_file(config_file), mode='r', encoding='utf-8')
  crf = rf.read()
  rf.close()
  return yaml.load(stream=crf, Loader=yaml.FullLoader)


def distribute(args, default=Optional[str]) -> (str, Dict[str, Any]):
  global _DISTRIBUTE, _CONFIG
  if len(args) > 1:
    _DISTRIBUTE = args[1]
  elif default is not None:
    _DISTRIBUTE = default
  else:
    raise _error_distribute_not_set()
  _CONFIG = _load_game_config('config.yaml')
  return _DISTRIBUTE, _CONFIG


def distribute_file(filename: str) -> str:
  if _DISTRIBUTE is None:
    raise _error_distribute_not_set()
  return f'{_DISTRIBUTE}/{filename}'


def get_distribute() -> str:
  if _DISTRIBUTE is None:
    raise _error_distribute_not_set()
  return _DISTRIBUTE


def _screenshot_wid(window_id: str) -> Image.Image:
  fh, filepath = tempfile.mkstemp(".png")
  os.close(fh)
  subprocess.call(["screencapture", "-l", window_id, "-o", "-x", filepath])
  im = Image.open(filepath)
  im.load()
  os.unlink(filepath)
  return im


def _screenshot_rect(rect: Box) -> Image.Image:
  fh, filepath = tempfile.mkstemp(".png")
  os.close(fh)
  region = f"{rect.left},{rect.top},{rect.width},{rect.height}"
  subprocess.call(["screencapture", "-R", region, "-x", filepath])
  im = Image.open(filepath)
  im.load()
  os.unlink(filepath)
  return im


def screenshot() -> Image.Image:
  ss = _screenshot_wid(_GAME_WINDOW_ID) if _GAME_WINDOW_ID is not None else _screenshot_rect(_GAME_WINDOW_RECT)
  return ss if _GAME_UI_RECT is None else ss.crop((_GAME_UI_RECT.left, _GAME_UI_RECT.top,
                                                   _GAME_UI_RECT.left + _GAME_UI_RECT.width,
                                                   _GAME_UI_RECT.top + _GAME_UI_RECT.height))


def _osascript(script_file: str) -> str:
  return subprocess.run(["osascript", script_file], capture_output=True).stdout.decode().rstrip('\n')


def _obtain_window() -> Optional[Box]:
  position_result = _osascript(f'get-position-of-{_DISTRIBUTE}.scpt')
  if position_result:
    position = tuple(map(int, position_result.split(', ')))
    size_result = _osascript(f'get-size-of-{_DISTRIBUTE}.scpt')
    if size_result:
      size = tuple(map(int, size_result.split(', ')))
      return _box(position[0], position[1], size[0], size[1])


def _init_game_window_mp() -> (Box, Box):
  global _GAME_WINDOW_ID, _GAME_WINDOW_RECT, _GAME_UI_RECT
  wid = _osascript('get-window-id-of-zombie.scpt')
  if wid:
    _GAME_WINDOW_ID = wid
    window = _obtain_window()
    if window:
      _GAME_WINDOW_RECT = window
      print(f'获得游戏窗口: id={_GAME_WINDOW_ID}, window={window}')
      _GAME_UI_RECT = obtain_game_ui_mp(screenshot())
      return window, _GAME_UI_RECT
  print('游戏未启动')
  exit(1)


def _init_game_window_ios() -> (Box, Box):
  global _GAME_WINDOW_RECT, _GAME_UI_RECT
  window = _obtain_window()
  if window:
    _GAME_WINDOW_RECT = window
    print(f'获得游戏窗口: window={window}')
    _GAME_UI_RECT = obtain_game_ui_ios(screenshot())
    return window, _GAME_UI_RECT
  print('游戏未启动')
  exit(1)


def init_game_window() -> (Box, Box):
  if _DISTRIBUTE is None:
    raise _error_distribute_not_set()
  if 'mp' == _DISTRIBUTE:
    print('initializing mp...')
    return _init_game_window_mp()
  elif 'ios' == _DISTRIBUTE:
    print('initializing ios...')
    return _init_game_window_ios()
  else:
    raise error_unknown_distribute()


def click(location):
  _activate()
  center = pyautogui.center(location)
  pyautogui.click(x=_GAME_WINDOW_RECT.left + (_GAME_UI_RECT.left + center.x) // 2,
                  y=_GAME_WINDOW_RECT.top + (_GAME_UI_RECT.top + center.y) // 2)
  time.sleep(_CLICK_INTERVAL)


def _activate() -> None:
  _osascript(f'activate-{_DISTRIBUTE}.scpt')


def config(path: str, default_none: bool = False) -> Any:
  data = _CONFIG
  for part in path.split('.'):
    data = data[part] if not default_none else data.get(part)
  return data
