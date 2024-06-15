import MetaTrader5 as mt5
import pandas as pd

if __name__ == '__main__':
    if not mt5.initialize(path="d:\\mt5_1\\terminal64.exe", portable=True):
        print("initialize() failed")
        mt5.shutdown()
        exit(1)
    sym = "EURUSDm"
    symbol_info = mt5.symbol_info(sym)
    print(symbol_info)
    print("{}".format(symbol_info.point))
    account_info = mt5.account_info()
    if account_info != None:
        account_info_dict = mt5.account_info()._asdict()
        df = pd.DataFrame(list(account_info_dict.items()), columns=['property', 'value'])
        print("account_info() as dataframe:")
        print(df)