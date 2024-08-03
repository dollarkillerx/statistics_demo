import MetaTrader5 as mt5
import pandas as pd
from datetime import datetime, timedelta

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


    # 获取当前时间
    current_time = datetime.now()

    # 设置历史数据的开始和结束时间
    start_time = current_time - timedelta(days=1)  # 假设从一年前开始获取数据
    end_time = current_time

    # 获取历史交易记录
    deals = mt5.history_orders_get(start_time, end_time)

    if deals is None:
        print("没有找到交易记录")
    else:
        print(f"共找到 {len(deals)} 条交易记录")
        # 获取最新的30个交易记录
        latest_deals = deals[-30:]
        for deal in latest_deals:
            print(deal)
