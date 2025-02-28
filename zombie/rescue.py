import os

from PIL import Image

from _locate import locate_all, Box
from _rescue import load_rescues, detect_rescues

rescue_img = Image.open('mp/rescue.png')


def match_rescues(file):
  with Image.open(file) as im:
    for match in locate_all(rescue_img, im):
      box = Box(match.left, match.top, match.width * 1.9, match.height)
      yield im.crop((box.left, box.top, box.left + box.width, box.top + box.height))


if __name__ == "__main__":
  rescues = load_rescues()
  for file in os.listdir('tmp'):
    if file.startswith('fights-') and file.endswith('.png'):
      for box in match_rescues(f'tmp/{file}'):
        rescue_name, new_rescue = detect_rescues(rescues, box)
  print(f'rescues: {len(rescues)}')
