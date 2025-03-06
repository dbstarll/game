from PIL import Image

from _game import distribute_file


def save_image(im: Image, filename: str) -> None:
  return im.save(filename, dpi=(144, 144))


def img(filename: str, ext: str = 'png') -> str:
  if filename.lower().endswith(ext):
    return distribute_file(f'{filename}')
  else:
    return distribute_file(f'{filename}.{ext}')
