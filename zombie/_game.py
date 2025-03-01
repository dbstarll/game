DISTRIBUTE = None


def distribute(args, default=None):
  global DISTRIBUTE
  if len(args) > 1:
    DISTRIBUTE = args[1]
  elif default is not None:
    DISTRIBUTE = default
  else:
    raise ValueError('distribute not set')
  return DISTRIBUTE


def distribute_file(filename):
  if DISTRIBUTE is None:
    raise ValueError('distribute not set')
  return f'{DISTRIBUTE}/{filename}'
