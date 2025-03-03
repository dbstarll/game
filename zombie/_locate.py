import pyscreeze

pyscreeze.USE_IMAGE_NOT_FOUND_EXCEPTION = False
pyscreeze.GRAYSCALE_DEFAULT = False

_LOCATE_OPTIONS = {'grayscale': False, 'confidence': 0.98}


def locate(needle_image, haystack_image, **kwargs):
  if len(kwargs) == 0:
    return pyscreeze.locate(needle_image, haystack_image, **_LOCATE_OPTIONS)
  else:
    return pyscreeze.locate(needle_image, haystack_image, **kwargs)


def locate_all(needle_image, haystack_image):
  return pyscreeze.locateAll(needle_image, haystack_image, **_LOCATE_OPTIONS)


def set_game_window(window):
  _LOCATE_OPTIONS['region'] = window


def Box(left, top, width, height):
  return pyscreeze.Box(left, top, width, height)
