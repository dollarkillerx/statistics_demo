import encryption

#########################################################
# 初始化参数
Magic = 666  # 魔术手
Deviation = 30  # 滑点
CurrencySuffix = "m"  # 货币后缀
InitialVolume = 0.01  # 初始volume
IncreaseMultiple = 0.01  # 加仓倍数
Symbol = "BTCUSD"  # 基础
Interval = 100  # 加仓间隔
TimeInterval = 15  # 时间间隔 default 30分
tradingview_sdk_url = "https://tradingview.dollarkiller.com"
md5_path="C:\\Users\\Administrator\\Desktop\\MT3\\terminal64.exe"
max_order=3
#########################################################

if __name__ == '__main__':
    cgs = encryption.Encryption(tradingview_sdk_url, md5_path,Magic, Deviation,
                                CurrencySuffix, InitialVolume, IncreaseMultiple,
                         Symbol, Interval, TimeInterval, max_order)
    cgs.init()
    cgs.run()