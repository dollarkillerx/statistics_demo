import time
import MetaTrader5 as a5
import utils

if __name__ == '__main__':
    mt5 = utils.MT5utils()
    print(mt5.profit())

    mt5.set_currency_suffix("z")
    # mt5.buy("USDJPY",0.1,"1")
    # time.sleep(5)
    # mt5.buy("USDJPY",0.1,"1")
    print(mt5.profit(magic=mt5.magic))

    print(a5.positions_get(ticket=233068760)[0].volume)

