import classic

#########################################################
# 初始化参数
Direction = "sell"  # buy sell 初始方向
Magic = 666  # 魔术手
Deviation = 30  # 滑点
CurrencySuffix = "c"  # 货币后缀
InitialVolume = 0.02  # 初始volume
IncreaseMultiple = 0.02  # 加仓倍数
Symbol = "EURUSD"  # 基础
Interval = 10  # 加仓间隔
TimeInterval = 30  # 时间间隔 default 30分
#########################################################

if __name__ == '__main__':
    et = classic.Classic(Direction, Magic, Deviation, CurrencySuffix, InitialVolume, IncreaseMultiple,
                         Symbol, Interval, TimeInterval)
    et.init()
    et.run()