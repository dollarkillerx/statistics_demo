import utils
import MetaTrader5 as a5
import pandas as pd

if __name__ == '__main__':
    mt5 = utils.MT5utils()

    symbol = "EURUSD"

    # buy
    mt5.set_currency_suffix("z")
    # mt5.buy(symbol, 0.01, "buy_1")
    # mt5.sell(symbol, 0.02, "sell_1")
    # print(mt5.positions_total())
    # USDCHFのポジションを取得する
    positions=a5.positions_get(symbol="EURUSDz")
    if positions==None:
        print("No positions on USDCHF, error code={}".format(a5.last_error()))
    elif len(positions)>0:
        print("Total positions on USDCHF =",len(positions))
        # すべてのポジションを表示する
        for position in positions:
            print(position)

    # 名前に「*USD*」が含まれる銘柄のポジションリストを取得する
    usd_positions=a5.positions_get(group="*USD*")
    if usd_positions==None:
        print("No positions with group=\"*USD*\", error code={}".format(a5.last_error()))
    elif len(usd_positions)>0:
        print("positions_get(group=\"*USD*\")={}".format(len(usd_positions)))
        # pandas.DataFrameを使用してこれらのポジションを表として表示する
        df=pd.DataFrame(list(usd_positions),columns=usd_positions[0]._asdict().keys())
        df['time'] = pd.to_datetime(df['time'], unit='s')
        df.drop(['time_update', 'time_msc', 'time_update_msc', 'external_id'], axis=1, inplace=True)
    print(df)


mt5.shutdown()
