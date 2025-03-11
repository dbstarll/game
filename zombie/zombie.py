import sys
import time
from typing import List

from PIL import Image
from past.builtins import cmp
from pyscreeze import Box

from _debug import now, debug_image
from _game import distribute, init_game_window, screenshot, click, get_distribute, config
from _image import img
from _invitation_pack import InvitationPack
from _locate import locate
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


def select_fight(im, invitations):
  _max = -2
  _max_pos = None
  for invitation_name, is_rescue, box, _, _ in invitations:
    print(f"{now()} - \t{invitation_name} - {is_rescue} - {box}")
    if is_rescue:
      rescue_level = 0 if invitation_name is None else int(invitation_name.split('-')[1])
    else:
      rescue_level = -1
    if rescue_level > _max:
      _max = rescue_level
      _max_pos = box

  if _max_pos is not None:
    debug_image(im, 'fights')

  return _max, _max_pos


def fighting():
  print(f"{now()} - 开始战斗...")
  start = time.time()
  last_skills = []
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
      last_skills = select_skill(im, last_skills)

    location_elite_skills = locate(img('elite-skill-close'), im)
    if location_elite_skills:
      print(f'{now()} - 精英掉落技能: {time.time() - start}')
      debug_image(im, 'elite-skills')

    time.sleep(5)


def select_skill(im: Image.Image, last_skills: List[str]) -> List[str]:
  min_idx, min_idx_rect, min_idx_name = 100, None, None
  skills, exist_unknown = [], False
  for image_index, kind_name, skill_name, _, skill_rect, _ in match_skills_from_screenshot(im):
    kind_and_skill = f'{kind_name}:{skill_name}'
    skills.append(kind_and_skill)
    if skill_name is None:
      exist_unknown = True
    if PREFER_SKILLS.count(kind_and_skill) == 1:
      idx = PREFER_SKILLS.index(kind_and_skill)
      if idx < min_idx:
        min_idx = idx
        min_idx_rect = skill_rect
        min_idx_name = kind_and_skill
  if len(skills) > 0 and cmp(last_skills, skills) != 0:
    print(f'{now()} - \t可选技能: {skills}')
    if exist_unknown:
      debug_image(im, 'skills')
    if min_idx_rect is not None:
      print(f'{now()} - 选择技能: {min_idx_name}')
      if config('skill.auto-select'):
        click(min_idx_rect)
    return skills
  else:
    return last_skills


def fight_prepare(fight):
  print(f"{now()} - 进入战斗预备, 等待队友开始...")
  click(invitation_pack.confirm(fight))
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
      level, invitation = select_fight(im, invitation_pack.match_from_screenshot(im))
      if invitation is not None:
        if level >= mini_level or level == 0:
          fight_prepare(invitation)
        elif level < 0:
          print(f"{now()} - 非寰球救援, 拒绝战斗")
          click(invitation_pack.cancel(invitation))
          continue
        else:
          print(f"{now()} - 寰球等级[{level}]太低, 拒绝战斗")
          click(invitation_pack.cancel(invitation))
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
  load_skills()
  invitation_pack = InvitationPack()

  detect_team_invite()
