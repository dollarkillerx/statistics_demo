# Earthworm 蚯蚓   贪婪的网格马丁
# @adapawang@gmail.com
import earthworm

#########################################################
# 初始化参数
Direction = "buy"  # buy sell 初始方向
Magic = 66666  # 魔术手
Deviation = 30  # 滑点
CurrencySuffix = "z"  # 货币后缀
InitialVolume = 0.01  # 初始volume
IncreaseMultiple = 0.4  # 加仓倍数
BaseCurrency = "EURUSD"  # 基础
Interval = 5  # 加仓间隔
TimeInterval = 30  # 时间间隔 default 30分
#########################################################


if __name__ == '__main__':
    ehw = earthworm.Earthworm(Direction, Magic, Deviation, CurrencySuffix, InitialVolume, IncreaseMultiple,
                              BaseCurrency, Interval, TimeInterval)
    ehw.init()
    ehw.run()
