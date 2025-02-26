import pyautogui
import pyscreeze
import time
import datetime
import sys

pyscreeze.USE_IMAGE_NOT_FOUND_EXCEPTION = False
LOCATE_OPTIONS = {'grayscale': True, 'confidence': 0.95}
CLICK_INTERVAL = 0.3
DISTRIBUTE = 'mp'

def img(file):
 return DISTRIBUTE + '/' + file

def click(location, offsetX = 0, offsetY = 0):
  center = pyautogui.center(location)
  pyautogui.click(x=center.x//2 + offsetX, y=center.y//2 + offsetY)
  time.sleep(CLICK_INTERVAL)

def debug_image(im,window,file):
  gim = im.crop((window.left, window.top, window.left + window.width, window.top + window.height))
  gim.save(file + '-' + str(time.time()) + '.png')

def get_game_left(screen,locationShop):
  for x in range(locationShop.left,0,-1):
    if screen.getpixel((x, locationShop.top)) == (0,0,0,255):
      return x+1

def get_game_top(screen,locationShop,left):
  for y in range(locationShop.top,0,-1):
    if screen.getpixel((left-1, y)) != (0,0,0,255):
      return y+1

def get_game_bottom(screen,locationShop,left):
  for y in range(locationShop.top+locationShop.height,screen.height):
    if screen.getpixel((left-1, y)) != (0,0,0,255):
      return y-1

def get_game_right(screen,locationShop,bottom):
  for x in range(locationShop.left+locationShop.width,screen.width):
    if screen.getpixel((x, bottom)) == (0,0,0,255):
      return x-1

def get_game_window(screen):
  locationShop = pyautogui.locate(img('rescue-return.png'),screen,**LOCATE_OPTIONS)
  if locationShop:
    left = get_game_left(screen,locationShop)
    if left:
      top = get_game_top(screen,locationShop,left)
      bottom = get_game_bottom(screen,locationShop,left)
      if top and bottom:
        right = get_game_right(screen,locationShop,bottom)
        if right:
          print(f"{datetime.datetime.now()} - 检测到游戏窗口: left: {left}, top: {top}, right: {right}, bottom: {bottom}")
          return pyscreeze.Box(left,top,right-left+1,bottom-top+1)
        else:
          print(f'{datetime.datetime.now()} - not found: right')
      else:
        print(f'{datetime.datetime.now()} - not found: top or bottom')
    else:
      print(f'{datetime.datetime.now()} - not found: left')
  else:
    print(f'{datetime.datetime.now()} - not found: locationShop')

def select_fight(im,window,fights):
  print(len(fights))
  if len(fights) > 1:
    debug_image(im,window,'fights')
  
  for pos in fights:
    print(f"{datetime.datetime.now()} - \t{pos} - ({pos.left//2+160},{pos.top//2-35}) - {pyautogui.position()}")
  for pos in fights:
    return pos
  
def fight_wait(fight,window):
  print(f"{datetime.datetime.now()} - 进入战斗预备, 等待队友开始...: {fight}")
  click(fight, 160, -35)
  start = time.time()
  while True:
    im = pyautogui.screenshot()

    locationInviting = pyautogui.locate(img('inviting.png'),im,**LOCATE_OPTIONS)
    if locationInviting:
      print(f'{datetime.datetime.now()} - 队友已退出: {time.time() - start}')
      break

    locationEnd = pyautogui.locate(img('fight-end.png'),im,**LOCATE_OPTIONS)
    if locationEnd:
      print(f'{datetime.datetime.now()} - 战斗结束: {time.time() - start}')
      click(locationEnd)
      break

    locationOffline = pyautogui.locate(img('offline-confirm.png'),im,**LOCATE_OPTIONS)
    if locationOffline:
      print(f'{datetime.datetime.now()} - 断线重连: {time.time() - start}')
      click(locationOffline)
      continue

    locationSkills = pyautogui.locate(img('select-skill.png'),im,**LOCATE_OPTIONS)
    if locationSkills:
      print(f'{datetime.datetime.now()} - 选择技能: {time.time() - start}')
      debug_image(im,window,'skills')

    time.sleep(5)

def find_fight(window):
  print(f'{datetime.datetime.now()} - 查看副本列表...')
  while True:
    im = pyautogui.screenshot()
    locationFightList = pyautogui.locate(img('fight-list.png'),im,**LOCATE_OPTIONS)
    if locationFightList:
      fight = select_fight(im,window,list(pyautogui.locateAll(img('rescue.png'),im,**LOCATE_OPTIONS)))
      if fight:
        fight_wait(fight,window)
      time.sleep(1)
    else:
      break

def find_invite(window):
  print(f'{datetime.datetime.now()} - 检测副本邀请...')
  while True:
    locationInvite = pyautogui.locateOnScreen(img('rescue-invite.png'),**LOCATE_OPTIONS)
    if locationInvite:
      print(f'{datetime.datetime.now()} - 检测到副本邀请,进入副本列表...')
      click(locationInvite)
      find_fight(window)
      print(f'{datetime.datetime.now()} - 回到主界面')
    time.sleep(1)

if __name__ == "__main__":
  if len(sys.argv) > 1:
    DISTRIBUTE = sys.argv[1]
  print(f'游戏发行版本: {DISTRIBUTE}')

  window = get_game_window(pyautogui.screenshot())
  if window:
    print(f"{datetime.datetime.now()} - 游戏窗口位置: {window}")
    LOCATE_OPTIONS['region'] = window
    find_invite(window)
