import subprocess
import sys
import time

import pyautogui
import pyscreeze

from _debug import now, debug_image
from _game import distribute
from _image import img
from _locate import locate, locate_all, set_game_window

CLICK_INTERVAL = 0.2
ROOM_WAIT_TIMEOUT = 15


def screenshot():
  return pyscreeze.screenshot()


def click(location, offset_x=0, offset_y=0):
  center = pyautogui.center(location)
  pyautogui.click(x=center.x // 2 + offset_x, y=center.y // 2 + offset_y)
  time.sleep(CLICK_INTERVAL)
  pyautogui.click(x=center.x // 2 + offset_x, y=center.y // 2 + offset_y)
  time.sleep(CLICK_INTERVAL)


def get_game_window_left(location_back):
  return location_back.left - 33


def get_game_window_top(location_back):
  return location_back.top - 1650


def get_game_window_bottom(location_back):
  return location_back.top + location_back.height + 37


def get_game_window_right(location_back):
  return location_back.left + location_back.width + 704


def get_game_window(screen):
  print(screen)
  location_back = locate(img('back.png'), screen)
  if location_back:
    left = get_game_window_left(location_back)
    if left:
      top = get_game_window_top(location_back)
      bottom = get_game_window_bottom(location_back)
      if top and bottom:
        right = get_game_window_right(location_back)
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
  location_offline = locate(img('offline-confirm.png'), im)
  if location_offline:
    print(f'{now()} - 断线重连')
    click(location_offline)
    return True

  location_reconnect = locate(img('reconnect.png'), im)
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

    location_end = locate(img('fight-end.png'), im)
    if location_end:
      print(f'{now()} - 战斗结束: {time.time() - start}')
      click(location_end)
      break

    location_skills = locate(img('select-skill.png'), im)
    if location_skills:
      match_left_bottoms = list(locate_all(img('skill-left-bottom.png'), im))
      match_right_tops = list(locate_all(img('skill-right-top.png'), im, ))
      print(f'{now()} - 选择技能({len(match_left_bottoms)} - {len(match_right_tops)}): {time.time() - start}')
      debug_image(im, window, 'skills')

    location_elite_skills = locate(img('elite-skill-close.png'), im)
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
    debug_image(im, window, 'prepare')
    location_leave = locate(img('room-leave.png'), im)
    if location_leave:
      if time.time() - start > ROOM_WAIT_TIMEOUT:
        print(f"{now()} - 等待超时, 退出战斗: {time.time() - start}")
        click(location_leave)
        break
      else:
        print(f'{now()} - 等待队友开始: {time.time() - start}')
    else:
      location_inviting = locate(img('room-inviting.png'), im)
      if location_inviting:
        print(f'{now()} - 队友已退出: {time.time() - start}')
      else:
        fighting(window)
      break
    time.sleep(0.7)


def find_fight(window):
  print(f'{now()} - 查看副本列表...')
  while True:
    im = screenshot()

    if check_reconnect(im):
      continue

    location_fight_list = locate(img('fight-list.png'), im)
    if location_fight_list:
      fight = select_fight(im, window, list(locate_all(img('rescue.png'), im)))
      if fight:
        fight_prepare(fight, window)
      time.sleep(0.7)
    else:
      break


def detect_team_invite(window):
  print(f'{now()} - 检测副本邀请...')
  while True:
    im = screenshot()

    if check_reconnect(im):
      continue

    location_invite = locate(img('team-invite.png'), im)
    if location_invite:
      print(f'{now()} - 检测到副本邀请,进入副本列表...')
      click(location_invite)
      find_fight(window)
      print(f'{now()} - 回到主界面')

    time.sleep(0.7)


def get_window_id_of_game():
  args = ["osascript"]
  return subprocess.run(args + ["get-window-id-of-zombie.scpt"], capture_output=True)


def get_bounds_of_game():
  args = ["osascript"]
  return subprocess.run(args + ["get-bounds-of-zombie.scpt"], capture_output=True)


if __name__ == "__main__":
  print(f'游戏发行版本: {distribute(sys.argv, "mp")}')

  window = get_game_window(screenshot())
  if window:
    print(f"{now()} - 游戏窗口位置: {window}")
    set_game_window(window)
    detect_team_invite(window)
