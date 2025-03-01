_DISTRIBUTE = None


def distribute(args, default=None):
  global _DISTRIBUTE
  if len(args) > 1:
    _DISTRIBUTE = args[1]
  elif default is not None:
    _DISTRIBUTE = default
  else:
    raise ValueError('distribute not set')
  return _DISTRIBUTE


def distribute_file(filename):
  if _DISTRIBUTE is None:
    raise ValueError('distribute not set')
  return f'{_DISTRIBUTE}/{filename}'
