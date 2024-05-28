from datetime import datetime
import matplotlib.pyplot as plt
import pandas as pd
from pandas.plotting import register_matplotlib_converters
register_matplotlib_converters()
import MetaTrader5 as mt5

# connect to MetaTrader 5
if not mt5.initialize():
    print("initialize() failed")
    mt5.shutdown()

# 连接到指定密码和服务器的交易账户
authorized=mt5.login(76548700, password="@Qq490890221")
if authorized:
    account_info=mt5.account_info()
    if account_info!=None:
        # display trading account data 'as is'
        print(account_info)
        # display trading account data in the form of a dictionary
        print("Show account_info()._asdict():")
        account_info_dict = mt5.account_info()._asdict()
        for prop in account_info_dict:
            print("  {}={}".format(prop, account_info_dict[prop]))
        print()

        # 将词典转换为DataFrame和print
        df=pd.DataFrame(list(account_info_dict.items()),columns=['property','value'])
        print("account_info() as dataframe:")
        print(df)

# request connection status and parameters
print(mt5.terminal_info())
# get data on MetaTrader 5 version
print(mt5.version())

# account_info=mt5.account_info()
# if account_info!=None:
#     # display trading account data 'as is'
#     print(account_info)
#     # display trading account data in the form of a dictionary
#     print("Show account_info()._asdict():")
#     account_info_dict = mt5.account_info()._asdict()
#     for prop in account_info_dict:
#         print("  {}={}".format(prop, account_info_dict[prop]))
#     print()
#
#     # 将词典转换为DataFrame和print
#     df=pd.DataFrame(list(account_info_dict.items()),columns=['property','value'])
#     print("account_info() as dataframe:")
#     print(df)

# 显示有关程序端设置和状态的信息
# terminal_info=mt5.terminal_info()
# if terminal_info!=None:
#     # display the terminal data 'as is'
#     print(terminal_info)
#     # display data in the form of a list
#     print("Show terminal_info()._asdict():")
#     terminal_info_dict = mt5.terminal_info()._asdict()
#     for prop in terminal_info_dict:
#         print("  {}={}".format(prop, terminal_info_dict[prop]))
#     print()
#     # 将词典转换为DataFrame和print
#     df=pd.DataFrame(list(terminal_info_dict.items()),columns=['property','value'])
#     print("terminal_info() as dataframe:")
#     print(df)

# 断开与MetaTrader 5程序端的连接
# 获取交易品种的数量
# symbols=mt5.symbols_total()
# if symbols>0:
#     print("Total symbols =",symbols)

# 获取所有交易品种
# symbols=mt5.symbols_get()
# print('Symbols: ', len(symbols))
# count=0
# # 显示前五个交易品种
# for s in symbols:
#     count+=1
#     print("{}. {}".format(count,s.name))
#     if count==5: break
# print()
#
# # 获取名称中包含RU的交易品种
# ru_symbols=mt5.symbols_get("*RU*")
# print('len(*RU*): ', len(ru_symbols))
# for s in ru_symbols:
#     print(s.name)
# print()

# 获取名称中不包含USD、EUR、JPY和GBP的交易品种
# group_symbols=mt5.symbols_get(group="*,!*USD*,!*EUR*,!*JPY*,!*GBP*")
# print('len(*,!*USD*,!*EUR*,!*JPY*,!*GBP*):', len(group_symbols))
# for s in group_symbols:
#     print(s.name,":",s)

# check the presence of active orders
# 显示有关程序端设置和状态的信息
# terminal_info=mt5.terminal_info()
# if terminal_info!=None:
#     # display the terminal data 'as is'
#     print(terminal_info)
#     # display data in the form of a list
#     print("Show terminal_info()._asdict():")
#     terminal_info_dict = mt5.terminal_info()._asdict()
#     for prop in terminal_info_dict:
#         print("  {}={}".format(prop, terminal_info_dict[prop]))
#     print()
#     # 将词典转换为DataFrame和print
#     df=pd.DataFrame(list(terminal_info_dict.items()),columns=['property','value'])
#     print("terminal_info() as dataframe:")
#     print(df)
#
# gbp_orders=mt5.orders_get(symbol="EURUSDz")
# if gbp_orders is None:
#     print("No orders with group=\"*USD*\", error code={}".format(mt5.last_error()))
# else:
#     print("orders_get(group=\"*USD*\")={}".format(len(gbp_orders)))

print('=============================================')
print(mt5.orders_total())
orders=mt5.orders_get(symbol='USDJPYz')
if orders is None:
    print("No orders on USDJPYz, error code={}".format(mt5.last_error()))
else:
    print("Total orders on USDJPYz:",len(orders))
    # display all active orders
    for order in orders:
        print(order)

mt5.shutdown()
