import MetaTrader5 as mt5

if __name__ == '__main__':
    if not mt5.initialize():
        print("initialize() failed")
        mt5.shutdown()
        exit(1)
    sym = "EURUSDz"
    symbol_info = mt5.symbol_info(sym)
    print(symbol_info)
    print("{}".format(symbol_info.point))
