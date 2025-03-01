from _game import distribute_file


def save_image(im, filename):
  im.save(filename, dpi=(144, 144))


def img(filename, ext: str = 'png'):
  if filename.lower().endswith(ext):
    return distribute_file(f'{filename}')
  else:
    return distribute_file(f'{filename}.{ext}')
