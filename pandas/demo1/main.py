import MetaTrader5 as mt5
import pandas as pd


if __name__ == '__main__':
    mt5.initialize()
    print(mt5.terminal_info())
    pd.set_option('display.max_columns', 500)  # 显示的最大列数
    pd.set_option('display.width', 1500)  # 控制显示输出的最大宽度
    # 从当日获取10个GBPUSD D1柱形图
    rates = mt5.copy_rates_from_pos("GBPUSDm", mt5.TIMEFRAME_D1, 0, 10)
    # 从所获得的数据创建DataFrame
    rates_frame = pd.DataFrame(rates)
    # 将时间（以秒为单位）转换为日期时间格式
    rates_frame['time'] = pd.to_datetime(rates_frame['time'], unit='s')
    # 显示数据
    print("\nDisplay dataframe with data")
    print(rates_frame)

