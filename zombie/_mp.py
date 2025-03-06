from typing import Optional

from PIL.Image import Image
from pyscreeze import Box

from _locate import _box


def _get_game_window_left(screen: Image, box: Box) -> Optional[int]:
  for x in range(box.left, 0, -1):
    if screen.getpixel((x, box.top)) == (0, 0, 0, 255):
      return x + 1


def _get_game_window_right(screen: Image, box: Box) -> Optional[int]:
  for x in range(box.left + box.width, screen.width):
    if screen.getpixel((x, box.top)) == (0, 0, 0, 255):
      return x - 1


def _get_game_window_top(screen: Image, location_back: Box, left: int) -> Optional[int]:
  for y in range(location_back.top, 0, -1):
    if screen.getpixel((left - 1, y)) != (0, 0, 0, 255):
      return y + 1


def _get_game_window_bottom(screen: Image, location_back: Box, left: int) -> int:
  for y in range(location_back.top + location_back.height, screen.height):
    if screen.getpixel((left - 1, y)) != (0, 0, 0, 255):
      return y - 1
  return screen.height - 1


def obtain_game_ui_mp(screen: Image) -> Box:
  print(f"屏幕: {screen}")
  screen_width, screen_height = screen.width, screen.height
  box = _box(screen_width // 3, screen_height // 3, screen_width // 3, screen_height // 3)
  left, right = _get_game_window_left(screen, box), _get_game_window_right(screen, box)
  if left is not None and right is not None:
    top, bottom = _get_game_window_top(screen, box, left), _get_game_window_bottom(screen, box, left)
    if top is not None:
      print(f"获得游戏UI: left: {left}, top: {top}, right: {right}, bottom: {bottom}")
      return _box(left, top, right - left + 1, bottom - top + 1)
    else:
      raise ValueError('game window not found: top')
  else:
    raise ValueError('game window not found: left or right')
