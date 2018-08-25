# Hello

# Hello params
nb: public(int128)

@public
def __init__(_startval: int128):
    self.nb = _startval

@public
def increment():
    self.nb = self.nb + 1
