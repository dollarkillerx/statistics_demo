# EyeTruth 真实之眼
# @adapawang@gmail.com
import eye_truth

#########################################################
# 初始化参数
Direction = "sell"  # buy sell 初始方向
Magic = 66600  # 魔术手
Deviation = 30  # 滑点
CurrencySuffix = "m"  # 货币后缀
InitialVolume = 0.01  # 初始volume
IncreaseMultiple = 0.01  # 加仓倍数
BaseCurrency = "EURUSD"  # 基础
Interval = 5  # 加仓间隔
TimeInterval = 30  # 时间间隔 default 30分
#########################################################

if __name__ == '__main__':
    et = eye_truth.EyeTruth(Direction, Magic, Deviation, CurrencySuffix, InitialVolume, IncreaseMultiple,
                             BaseCurrency, Interval, TimeInterval)
    et.init()
    et.run()
