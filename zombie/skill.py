import os
import time

import pyautogui
import pyscreeze
from PIL import Image

pyscreeze.USE_IMAGE_NOT_FOUND_EXCEPTION = False
pyscreeze.GRAYSCALE_DEFAULT = False

LOCATE_OPTIONS = {'confidence': 0.98}
LEFT_OFFSET = 1
RIGHT_OFFSET = -1
TOP_OFFSET = 1
BOTTOM_OFFSET = -1

LEFT_BOTTOM_IMG = Image.open('mp/skill-left-bottom.png')
RIGHT_TOP_IMG = Image.open('mp/skill-right-top.png')


def debug_image(im, window, file):
  gim = im.crop((window.left, window.top, window.left + window.width, window.top + window.height))
  gim.save('tmp/' + file + '-' + str(time.time()) + '.png', dpi=(144, 144))
  return gim


def match(file):
  with Image.open(file) as im:
    match_lb = list(pyautogui.locateAll(LEFT_BOTTOM_IMG, im, **LOCATE_OPTIONS))
    match_rt = list(pyautogui.locateAll(RIGHT_TOP_IMG, im, **LOCATE_OPTIONS))
    if len(match_lb) != 3 or len(match_rt) != 3:
      print(f'mismatch: {len(match_lb)} - {len(match_rt)}, file: {file}')
      return False
    else:
      print(f'{file}')
      for i in range(0, 3):
        match_left_bottom = match_lb[i]
        match_right_top = match_rt[i]
        box = pyscreeze.Box(match_left_bottom.left + LEFT_OFFSET, match_right_top.top + TOP_OFFSET,
                            match_right_top.left + match_right_top.width - match_left_bottom.left - LEFT_OFFSET + RIGHT_OFFSET,
                            match_left_bottom.top + match_left_bottom.height - match_right_top.top - TOP_OFFSET + BOTTOM_OFFSET)
        print(f'\t{box}')
        debug_image(im, box, 'skill')
        # lb = debug_image(im, pyscreeze.Box(box.left - 1, box.top + box.height - 50 + 1, 50, 50), 'left-bottom')
        # rt = debug_image(im, pyscreeze.Box(box.left + box.width - 50 + 1, box.top - 1, 50, 75), 'right-top')
        # print(len(list(pyautogui.locateAll(lb, im, **LOCATE_OPTIONS))))
        # print(len(list(pyautogui.locateAll(rt, im, **LOCATE_OPTIONS))))
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
