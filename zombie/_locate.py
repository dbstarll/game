import pyscreeze

pyscreeze.USE_IMAGE_NOT_FOUND_EXCEPTION = False
pyscreeze.GRAYSCALE_DEFAULT = False

_LOCATE_OPTIONS = {'grayscale': False, 'confidence': 0.98}


def locate(needle_image, haystack_image, **kwargs):
  if len(kwargs) == 0:
    return pyscreeze.locate(needle_image, haystack_image, **_LOCATE_OPTIONS)
  else:
    custom_options = _LOCATE_OPTIONS.copy()
    for key, value in kwargs.items():
      custom_options[key] = value
    return pyscreeze.locate(needle_image, haystack_image, **custom_options)


def locate_all(needle_image, haystack_image, **kwargs):
  if len(kwargs) == 0:
    return pyscreeze.locateAll(needle_image, haystack_image, **_LOCATE_OPTIONS)
  else:
    custom_options = _LOCATE_OPTIONS.copy()
    for key, value in kwargs.items():
      custom_options[key] = value
    return pyscreeze.locateAll(needle_image, haystack_image, **custom_options)


def set_game_window(window):
  _LOCATE_OPTIONS['region'] = window


def Box(left, top, width, height) -> pyscreeze.Box:
  return pyscreeze.Box(left, top, width, height)
