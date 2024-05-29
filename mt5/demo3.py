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
    # print(mt5.profit(magic=mt5.magic))
    #
    # print(a5.positions_get(ticket=233068760)[0].volume)
    # positions = mt5.positions_get()
    # for position in positions:
    #     mt5.close(orderId=position.ticket)
    # print(a5.symbol_info("EURUSDz").point * 100)
    mt5.buy(symbol="EURUSD", volume=0.01, comment="v1", sl=20,tp=20)
    mt5.sell(symbol="EURUSD", volume=0.01, comment="v2",sl=20,tp=20)
