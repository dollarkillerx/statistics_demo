import MetaTrader5 as mt5
import pandas as pd

if __name__ == '__main__':
    if not mt5.initialize():
        print("initialize() failed")
        mt5.shutdown()
        exit(1)
    account_info = mt5.account_info()
    account_info_dict = mt5.account_info()._asdict()
    df = pd.DataFrame(list(account_info_dict.items()), columns=['property', 'value'])
    print("account_info() as dataframe:")
    print(df)
    positions = mt5.positions_get()
    if positions is not None:
        for pos in positions:
            print(pos)