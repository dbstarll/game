import sys
import time

from PIL import Image
from pyscreeze import Box

from _debug import now, debug_image
from _game import distribute, init_game_window, screenshot, click, get_distribute
from _image import img
from _locate import locate, locate_all
from _rescue import match_rescues, load_rescues
from _skill import load_skills, match_skills_from_screenshot

ROOM_WAIT_TIMEOUT = 15

PREFER_SKILLS = ['枪械:分裂冰片', '枪械:连发+', '枪械:齐射+', '枪械:急冻子弹+', '枪械:子弹爆炸',
                 '枪械:伤害增幅', '枪械:分裂子弹', '枪械:分裂子弹四射', '枪械:分裂子弹爆炸', '枪械:全子弹增幅',
                 '装甲车:装甲车', '装甲车:焦土策略']


def check_unexpected(im: Image.Image) -> bool:
  if get_distribute() == 'ios' and check_reconnect_ios(im):
    return True

  return check_reconnect(im)


def check_reconnect_ios(im: Image.Image) -> bool:
  btn_reconnect = locate(img('reconnect-ios'), im)
  if btn_reconnect is not None:
    print(f'{now()} - 重新连接iPhone镜像')
    click(btn_reconnect)
    return True
  else:
    return False


def check_reconnect(im: Image.Image) -> bool:
  location_offline = locate(img('offline-confirm'), im)
  if location_offline:
    print(f'{now()} - 断线重连')
    click(location_offline)
    return True

  location_reconnect = locate(img('reconnect'), im)
  if location_reconnect:
    print(f'{now()} - 网络断开，重新连接')
    click(location_reconnect)
    return True
  else:
    return False


def select_fight(im, fights):
  if len(fights) > 0:
    debug_image(im, 'fights')

  _max = -1
  _max_pos = None
  for rescue_level, rescue_name, rescue_rect, rescue_image in fights:
    print(f"{now()} - \t{rescue_level} - {rescue_name} - {rescue_rect}")
    if rescue_level > _max:
      _max = rescue_level
      _max_pos = rescue_rect
  return _max, _max_pos


def fighting():
  print(f"{now()} - 开始战斗...")
  start = time.time()
  while True:
    im = screenshot()

    if check_unexpected(im):
      continue

    location_end = locate(img('fight-end'), im)
    if location_end:
      print(f'{now()} - 战斗结束: {time.time() - start}')
      click(location_end)
      break

    location_skills = locate(img('select-skill'), im)
    if location_skills:
      match_left_bottoms = list(locate_all(img('skill-left-bottom'), im))
      match_right_tops = list(locate_all(img('skill-right-top'), im))
      print(f'{now()} - 可选技能({len(match_left_bottoms)} - {len(match_right_tops)}): {time.time() - start}')
      min_idx = 100
      min_idx_rect = None
      min_idx_name = None
      for image_index, kind_name, skill_name, _, skill_rect, _ in match_skills_from_screenshot(im):
        print(f'{now()} - \t{image_index} - 技能[{kind_name} - {skill_name}]: {time.time() - start}')
        kind_and_skill = f'{kind_name}:{skill_name}'
        if PREFER_SKILLS.count(kind_and_skill) == 1:
          idx = PREFER_SKILLS.index(kind_and_skill)
          if idx < min_idx:
            min_idx = idx
            min_idx_rect = skill_rect
            min_idx_name = kind_and_skill
      if min_idx_rect is not None:
        print(f'{now()} - 选择技能: {min_idx_name}: {time.time() - start}')
        # click(min_idx_rect)

      debug_image(im, 'skills')

    location_elite_skills = locate(img('elite-skill-close'), im)
    if location_elite_skills:
      print(f'{now()} - 精英掉落技能: {time.time() - start}')
      debug_image(im, 'elite-skills')

    time.sleep(5)


def fight_prepare(fight):
  print(f"{now()} - 进入战斗预备, 等待队友开始...")
  click(fight, 160, -35)
  start = time.time()
  while True:
    im = screenshot()
    location_leave = locate(img('room-leave'), im)
    if location_leave:
      if time.time() - start > ROOM_WAIT_TIMEOUT:
        print(f"{now()} - 等待超时, 退出战斗: {time.time() - start}")
        click(location_leave)
        break
      else:
        print(f'{now()} - 等待队友开始: {time.time() - start}')
    else:
      location_inviting = locate(img('room-inviting'), im)
      if location_inviting:
        print(f'{now()} - 队友已退出: {time.time() - start}')
      else:
        debug_image(im, 'fighting')
        location_at_home = locate(img('at-home'), im)
        if location_at_home:
          print(f'{now()} - 房间满，回到首页: {time.time() - start}')
        else:
          fighting()
      break
    time.sleep(0.7)


def find_fight():
  print(f'{now()} - 查看副本列表...')
  while True:
    im = screenshot()

    if check_unexpected(im):
      continue

    location_fight_list = locate(img('fight-list'), im)
    if location_fight_list:
      level, fight = select_fight(im, list(match_rescues(im)))
      if fight is not None:
        if level >= mini_level or level == 0:
          fight_prepare(fight)
        else:
          print(f"{now()} - 寰球等级[{level}]太低, 拒绝战斗")
          click(fight, 160, 0, True)
          continue
      time.sleep(0.7)
    else:
      break


def to_team_invite_list(invite: Box):
  print(f'{now()} - 检测到副本邀请,进入副本列表...')
  click(invite)
  find_fight()
  print(f'{now()} - 回到主界面')


def detect_team_invite():
  print(f'{now()} - 检测副本邀请...')
  while True:
    im = screenshot()

    if check_unexpected(im):
      continue

    location_invite_at_fight = locate(img('team-invite-at-fight'), im)
    if location_invite_at_fight is not None:
      to_team_invite_list(location_invite_at_fight)
    else:
      location_invite = locate(img('team-invite'), im)
      if location_invite is not None:
        to_team_invite_list(location_invite)

    time.sleep(0.7)


if __name__ == "__main__":
  dist, _ = distribute(sys.argv, "mp")
  mini_level = 5 if 'mp' == dist else 1
  print(f'游戏发行版本: {dist}')
  init_game_window()
  load_rescues()
  load_skills()

  detect_team_invite()
