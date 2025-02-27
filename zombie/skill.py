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
TOP_OFFSET = -2
BOTTOM_OFFSET = 1


def debug_image(im, window, file):
  gim = im.crop((window.left, window.top, window.left + window.width, window.top + window.height))
  gim.save('tmp/' + file + '-' + str(time.time()) + '.png', dpi=(144, 144))
  return gim


def match(file):
  with Image.open(file) as im:
    print(f'{file}: {im}')
    match_left_bottoms = list(pyautogui.locateAll('left-bottom.png', im, **LOCATE_OPTIONS))
    match_right_tops = list(pyautogui.locateAll('right-top.png', im, **LOCATE_OPTIONS))
    if len(match_left_bottoms) != 3 or len(match_right_tops) != 3:
      print(f'mismatch: {len(match_left_bottoms)} - {len(match_right_tops)}, file: {file}')
    else:
      for i in range(0, 3):
        match_left_bottom = match_left_bottoms[i]
        match_right_top = match_right_tops[i]
        box = pyscreeze.Box(match_left_bottom.left + LEFT_OFFSET, match_right_top.top + TOP_OFFSET,
                            match_right_top.left + match_right_top.width - match_left_bottom.left - LEFT_OFFSET + RIGHT_OFFSET,
                            match_left_bottom.top + match_left_bottom.height - match_right_top.top - TOP_OFFSET + BOTTOM_OFFSET)
        print(f'\t{box}')
        debug_image(im, box, 'skill')
        # lb = debug_image(im,pyscreeze.Box(box.left,box.top+box.height-30,50,30),'left-bottom')
        # rt = debug_image(im,pyscreeze.Box(box.left+box.width-50,box.top,50,50),'right-top')
        # print(len(list(pyautogui.locateAll(lb,im,**LOCATE_OPTIONS))))
        # print(len(list(pyautogui.locateAll(rt,im,**LOCATE_OPTIONS))))


if __name__ == "__main__":
  for file in os.listdir('tmp'):
    if file.startswith('skills-') and file.endswith('.png'):
      match('tmp/' + file)
