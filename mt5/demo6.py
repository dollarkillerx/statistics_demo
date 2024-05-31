import utils
import pandas as pd

if __name__ == '__main__':
    mt5 = utils.MT5utils()
    positions = mt5.positions_get()
    for position in positions:
        print(position)
