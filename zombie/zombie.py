import datetime
import os
import subprocess
import sys
import tempfile
import time

import pyautogui
import pyscreeze
from PIL import Image

pyscreeze.USE_IMAGE_NOT_FOUND_EXCEPTION = False
pyscreeze.GRAYSCALE_DEFAULT = False

LOCATE_OPTIONS = {'confidence': 0.98}
GAME_WINDOW_ID = None
GAME_WINDOW_POS = None
CLICK_INTERVAL = 0.3
DISTRIBUTE = 'mp'
ROOM_WAIT_TIMEOUT = 15


def screenshot_osx():
  fh, filepath = tempfile.mkstemp(".png")
  os.close(fh)
  args = ["screencapture"]
  subprocess.call(args + ["-l", GAME_WINDOW_ID, "-o", "-x", filepath])
  im = Image.open(filepath)
  im.load()
  os.unlink(filepath)
  return im


def screenshot():
  return screenshot_osx()


def now():
  return datetime.datetime.now()


def img(file):
  return DISTRIBUTE + '/' + file


def click(location, offset_x=0, offset_y=0):
  center = pyautogui.center(location)
  pyautogui.click(x=GAME_WINDOW_POS[0] + center.x // 2 + offset_x, y=GAME_WINDOW_POS[1] + center.y // 2 + offset_y)
  time.sleep(CLICK_INTERVAL)


def debug_image(im, window, file):
  gim = im.crop((window.left, window.top, window.left + window.width, window.top + window.height))
  gim.save('tmp/' + file + '-' + str(int(time.time())) + '.png', dpi=(144, 144))


def get_game_window_left(screen, location_back):
  for x in range(location_back.left, 0, -1):
    if screen.getpixel((x, location_back.top)) == (0, 0, 0, 255):
      return x + 1


def get_game_window_top(screen, location_back, left):
  for y in range(location_back.top, 0, -1):
    if screen.getpixel((left - 1, y)) != (0, 0, 0, 255):
      return y + 1


def get_game_window_bottom(screen, location_back, left):
  for y in range(location_back.top + location_back.height, screen.height):
    if screen.getpixel((left - 1, y)) != (0, 0, 0, 255):
      return y - 1
  return screen.height - 1


def get_game_window_right(screen, location_back, bottom):
  for x in range(location_back.left + location_back.width, screen.width):
    if screen.getpixel((x, bottom)) == (0, 0, 0, 255):
      return x - 1


def get_game_window(screen):
  print(screen)
  location_back = pyautogui.locate(img('back.png'), screen, **LOCATE_OPTIONS)
  if location_back:
    left = get_game_window_left(screen, location_back)
    if left:
      top = get_game_window_top(screen, location_back, left)
      bottom = get_game_window_bottom(screen, location_back, left)
      if top and bottom:
        right = get_game_window_right(screen, location_back, bottom)
        if right:
          print(f"{now()} - 检测到游戏窗口: left: {left}, top: {top}, right: {right}, bottom: {bottom}")
          return pyscreeze.Box(left, top, right - left + 1, bottom - top + 1)
        else:
          print(f'{now()} - not found: right')
      else:
        print(f'{now()} - not found: top or bottom')
    else:
      print(f'{now()} - not found: left')
  else:
    print(f'{now()} - not found: locationShop')


def check_reconnect(im):
  location_offline = pyautogui.locate(img('offline-confirm.png'), im, **LOCATE_OPTIONS)
  if location_offline:
    print(f'{now()} - 断线重连')
    click(location_offline)
    return True

  location_reconnect = pyautogui.locate(img('reconnect.png'), im, **LOCATE_OPTIONS)
  if location_reconnect:
    print(f'{now()} - 网络断开，重新连接')
    click(location_reconnect)
    return True
  else:
    return False


def select_fight(im, window, fights):
  if len(fights) > 0:
    print(len(fights))
    debug_image(im, window, 'fights')

  for pos in fights:
    print(f"{now()} - \t{pos} - ({pos.left // 2 + 160},{pos.top // 2 - 35}) - {pyautogui.position()}")
  for pos in fights:
    return pos


def fighting(window):
  print(f"{now()} - 开始战斗...")
  start = time.time()
  while True:
    im = screenshot()

    if check_reconnect(im):
      continue

    location_end = pyautogui.locate(img('fight-end.png'), im, **LOCATE_OPTIONS)
    if location_end:
      print(f'{now()} - 战斗结束: {time.time() - start}')
      click(location_end)
      break

    location_skills = pyautogui.locate(img('select-skill.png'), im, **LOCATE_OPTIONS)
    if location_skills:
      match_left_bottoms = list(pyautogui.locateAll(img('skill-left-bottom.png'), im, **LOCATE_OPTIONS))
      match_right_tops = list(pyautogui.locateAll(img('skill-right-top.png'), im, **LOCATE_OPTIONS))
      print(f'{now()} - 选择技能({len(match_left_bottoms)} - {len(match_right_tops)}): {time.time() - start}')
      debug_image(im, window, 'skills')

    location_elite_skills = pyautogui.locate(img('elite-skill-close.png'), im, **LOCATE_OPTIONS)
    if location_elite_skills:
      print(f'{now()} - 精英掉落技能: {time.time() - start}')
      debug_image(im, window, 'elite-skills')

    time.sleep(5)


def fight_prepare(fight, window):
  print(f"{now()} - 进入战斗预备, 等待队友开始...")
  click(fight, 160, -35)
  start = time.time()
  while True:
    im = screenshot()
    location_leave = pyautogui.locate(img('room-leave.png'), im, **LOCATE_OPTIONS)
    if location_leave:
      if time.time() - start > ROOM_WAIT_TIMEOUT:
        print(f"{now()} - 等待超时, 退出战斗: {time.time() - start}")
        click(location_leave)
        break
      else:
        print(f'{now()} - 等待队友开始: {time.time() - start}')
    else:
      location_inviting = pyautogui.locate(img('room-inviting.png'), im, **LOCATE_OPTIONS)
      if location_inviting:
        print(f'{now()} - 队友已退出: {time.time() - start}')
      else:
        fighting(window)
      break
    time.sleep(1)


def find_fight(window):
  print(f'{now()} - 查看副本列表...')
  while True:
    im = screenshot()

    if check_reconnect(im):
      continue

    location_fight_list = pyautogui.locate(img('fight-list.png'), im, **LOCATE_OPTIONS)
    if location_fight_list:
      fight = select_fight(im, window, list(pyautogui.locateAll(img('rescue.png'), im, **LOCATE_OPTIONS)))
      if fight:
        fight_prepare(fight, window)
      time.sleep(1)
    else:
      break


def detect_team_invite(window):
  print(f'{now()} - 检测副本邀请...')
  while True:
    im = screenshot()

    if check_reconnect(im):
      continue

    location_invite = pyautogui.locate(img('team-invite.png'), im, **LOCATE_OPTIONS)
    if location_invite:
      print(f'{now()} - 检测到副本邀请,进入副本列表...')
      click(location_invite)
      find_fight(window)
      print(f'{now()} - 回到主界面')

    time.sleep(1)


def get_window_id_of_game():
  args = ["osascript"]
  return subprocess.run(args + ["get-window-id-of-zombie.scpt"], capture_output=True)


def get_bounds_of_game():
  args = ["osascript"]
  return subprocess.run(args + ["get-bounds-of-zombie.scpt"], capture_output=True)


if __name__ == "__main__":
  if len(sys.argv) > 1:
    DISTRIBUTE = sys.argv[1]
  print(f'游戏发行版本: {DISTRIBUTE}')

  wid = get_window_id_of_game().stdout.decode()
  if wid:
    GAME_WINDOW_ID = wid.rstrip('\n')
    GAME_WINDOW_POS = tuple(map(int, get_bounds_of_game().stdout.decode().rstrip('\n').split(', ')))
    print(f'获得游戏窗口: ID={GAME_WINDOW_ID}, POS={GAME_WINDOW_POS}')
  else:
    print('游戏未启动')
    exit(1)

  window = get_game_window(screenshot())
  if window:
    print(f"{now()} - 游戏窗口位置: {window}")
    LOCATE_OPTIONS['region'] = window
    detect_team_invite(window)
