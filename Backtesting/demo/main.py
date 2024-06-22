import yfinance as yf
from backtesting import Backtest, Strategy
from backtesting.lib import crossover
from backtesting.test import SMA

# 下载 EUR/USD 的 5 分钟 K 线数据
data = yf.download('EURUSD=X', interval='5m', start='2023-01-01', end='2023-02-20')

# 将数据调整为 backtesting 库需要的格式
data.columns = ['Open', 'High', 'Low', 'Close', 'Adj Close', 'Volume']
data = data.drop(columns=['Adj Close', 'Volume'])

class SmaCross(Strategy):
    n1 = 10  # 第一个移动平均线的周期
    n2 = 20  # 第二个移动平均线的周期

    def init(self):
        # 初始化移动平均线指标
        self.sma1 = self.I(SMA, self.data.Close, self.n1)  # 计算第一个移动平均线
        self.sma2 = self.I(SMA, self.data.Close, self.n2)  # 计算第二个移动平均线

    def next(self):
        # 在每个新的数据点上调用，检查移动平均线是否交叉
        if crossover(self.sma1, self.sma2):  # 如果第一个移动平均线向上穿过第二个
            self.buy()  # 买入
        elif crossover(self.sma2, self.sma1):  # 如果第二个移动平均线向上穿过第一个
            self.sell()  # 卖出

if __name__ == '__main__':
    # 创建 Backtest 对象，传入历史数据，策略类，初始现金金额和交易佣金
    bt = Backtest(data, SmaCross, cash=10000, commission=.002)
    stats = bt.run()  # 运行回测并获得结果统计
    bt.plot()  # 绘制回测结果
