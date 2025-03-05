from PIL import Image
from pyscreeze import Box

from _locate import _box

NORMAL_LEFT_OFFSET = 17
NORMAL_RIGHT_OFFSET = 536
NORMAL_TOP_OFFSET = 1279
NORMAL_BOTTOM_OFFSET = 20

BIG_LEFT_OFFSET = 11
BIG_RIGHT_OFFSET = 12
BIG_TOP_OFFSET = 180
BIG_BOTTOM_OFFSET = 12


def obtain_game_ui_ios(screen: Image) -> Box:
  print(f"屏幕: {screen}")
  screen_width, screen_height = screen.width, screen.height
  left = BIG_LEFT_OFFSET
  top = BIG_TOP_OFFSET
  bottom = screen_height - BIG_BOTTOM_OFFSET
  right = screen_width - BIG_RIGHT_OFFSET
  print(f"获得游戏UI: left: {left}, top: {top}, right: {right}, bottom: {bottom}")
  return _box(left, top, right - left + 1, bottom - top + 1)
