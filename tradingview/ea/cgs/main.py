import cgs

#########################################################
# 初始化参数
Magic = 666  # 魔术手
Deviation = 30  # 滑点
CurrencySuffix = "c"  # 货币后缀
InitialVolume = 0.01  # 初始volume
IncreaseMultiple = 0.01  # 加仓倍数
Symbol = "EURUSD"  # 基础
Interval = 5  # 加仓间隔
TimeInterval = 30  # 时间间隔 default 30分
tradingview_sdk_url = "https://tradingview.dollarkiller.com"
md5_path="C:\\Users\\Administrator\\Desktop\\MT2\\terminal64.exe"
#########################################################

if __name__ == '__main__':
    cgs = cgs.Cgs(tradingview_sdk_url, md5_path,Magic, Deviation, CurrencySuffix, InitialVolume, IncreaseMultiple,
                         Symbol, Interval, TimeInterval)
    cgs.init()
    cgs.run()