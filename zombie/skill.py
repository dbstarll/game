import os
import time

import pyautogui
import pyscreeze
from PIL import Image

pyscreeze.USE_IMAGE_NOT_FOUND_EXCEPTION = False
pyscreeze.GRAYSCALE_DEFAULT = False

LOCATE_OPTIONS = {'confidence': 0.98}
LEFT_OFFSET = 3
RIGHT_OFFSET = -1
TOP_OFFSET = -2
BOTTOM_OFFSET = -3

LEFT_TOP_IMG = Image.open('left-top-4.png')
LEFT_TOP_NEW_IMG = Image.open('left-top-new.png')
LEFT_BOTTOM_IMG = Image.open('left-bottom-4.png')
RIGHT_TOP_IMG = Image.open('right-top-4.png')


def debug_image(im, window, file):
  gim = im.crop((window.left, window.top, window.left + window.width, window.top + window.height))
  gim.save('tmp/' + file + '-' + str(time.time()) + '.png', dpi=(144, 144))
  return gim


def match(file):
  with Image.open(file) as im:
    match_lt = list(pyautogui.locateAll(LEFT_TOP_IMG, im, **LOCATE_OPTIONS))
    match_ltn = list(pyautogui.locateAll(LEFT_TOP_NEW_IMG, im, **LOCATE_OPTIONS))
    match_lb = list(pyautogui.locateAll(LEFT_BOTTOM_IMG, im, **LOCATE_OPTIONS))
    match_rt = list(pyautogui.locateAll(RIGHT_TOP_IMG, im, **LOCATE_OPTIONS))
    if len(match_lt) + len(match_ltn) != 3 or len(match_lb) != 3 or len(match_rt) != 3:
      print(f'mismatch: {len(match_lt)} - {len(match_ltn)} - {len(match_lb)} - {len(match_rt)}, file: {file}')
      return False
    else:
      for i in range(0, 3):
        match_left_bottom = match_lb[i]
        match_right_top = match_rt[i]
        box = pyscreeze.Box(match_left_bottom.left + LEFT_OFFSET, match_right_top.top + TOP_OFFSET,
                            match_right_top.left + match_right_top.width - match_left_bottom.left - LEFT_OFFSET + RIGHT_OFFSET,
                            match_left_bottom.top + match_left_bottom.height - match_right_top.top - TOP_OFFSET + BOTTOM_OFFSET)
        print(f'\t{box}')
        debug_image(im, box, 'skill')
        # lb = debug_image(im,pyscreeze.Box(box.left,box.top+box.height-30,50,30),'left-bottom')
        # rt = debug_image(im,pyscreeze.Box(box.left+box.width-50,box.top,50,50),'right-top')
        # print(len(list(pyautogui.locateAll(lb,im,**LOCATE_OPTIONS))))
        # print(len(list(pyautogui.locateAll(rt,im,**LOCATE_OPTIONS))))
      return True


if __name__ == "__main__":
  total = 0
  matches = 0
  for file in os.listdir('tmp'):
    if file.startswith('skills-') and file.endswith('.png'):
      if total >= 100:
        break
      total += 1
      if match('tmp/' + file):
        matches += 1
  print(f'matches: {matches}/{total}')
