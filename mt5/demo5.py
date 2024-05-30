from datetime import datetime
import MetaTrader5 as mt5
# 显示有关MetaTrader 5程序包的数据
print("MetaTrader5 package author: ",mt5.__author__)
print("MetaTrader5 package version: ",mt5.__version__)

# 导入'pandas'模块，用于以表格形式显示获得的数据
import pandas as pd
pd.set_option('display.max_columns', 500) # number of columns to be displayed
pd.set_option('display.width', 1500)      # max table width to display

# 建立与MetaTrader 5程序端的连接
if not mt5.initialize():
    print("initialize() failed, error code =",mt5.last_error())
    quit()

# 从当日获取10个GBPUSD D1柱形图
rates = mt5.copy_rates_from_pos("EURUSDz", mt5.TIMEFRAME_M1, 0, 10)

# 断开与MetaTrader 5程序端的连接
mt5.shutdown()
# 在新行显示所获得数据的每个元素
# print("Display obtained data 'as is'")
# for rate in rates:
#     print(rate)

# 从所获得的数据创建DataFrame
rates_frame = pd.DataFrame(rates)
# 将时间（以秒为单位）转换为日期时间格式
rates_frame['time']=pd.to_datetime(rates_frame['time'], unit='s')

# 显示数据
print("\nDisplay dataframe with data")
print(rates_frame)