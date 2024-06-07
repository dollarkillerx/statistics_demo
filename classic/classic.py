import time
import random
import seahorse
from decimal import Decimal, getcontext
getcontext().prec = 28

class Classic:
    direction = "buy" # 方向
    magic = 60606 # 魔术手
    deviation = 20 # 滑点
    currency_suffix = "z"  # 货币后缀
    initial_volume = 0.01  # 初始volume
    increase_multiple = 0.4  # 加仓倍数
    symbol = "EURUSD" # 货币
    interval = 5 # 加仓间隔
    sleepTime = 0 # 休息时间


    def __init__(self, direction = "buy", magic=60606, deviation = 20, currency_suffix = "z", initial_volume = 0.01, increase_multiple = 0.4,symbol = "EURUSD", interval= 5):
        self.direction = direction
        self.magic = magic
        self.deviation = deviation
        self.currency_suffix = currency_suffix
        self.initial_volume = initial_volume
        self.increase_multiple = increase_multiple
        self.symbol = symbol
        self.interval = interval

    def init(self):
        self.mt5 = seahorse.Seahorse("127.0.0.1:8475", "my_test", 1000, 2000)
        self.mt5.set_currency_suffix(self.currency_suffix)
        self.mt5.set_magic(self.magic)
        self.mt5.set_def_deviation(self.deviation)

    # 出场
    def prominence(self):
        profit = self.mt5.profit(magic=self.magic)
        if profit >= 5:
            self.mt5.close_all(magic=self.magic)
            tick = self.mt5.symbol_info_tick2(symbol=self.magic)
            self.sleepTime = tick.time + 60*60*3

    # 随机方向
    def random_direction(self):
        random_number = random.randint(1, 100)
        if random_number > 50:
            return "buy"
        return "sell"

    def run(self):
        while True:
            tick = self.mt5.symbol_info_tick(symbol=self.symbol)
            # 是否休息
            if self.sleepTime != 0:
                if self.sleepTime < tick.time:
                    continue
                else:
                    self.sleepTime = 0

            # 获取当前订单列表
            positions = self.mt5.positions_get()
            if len(positions) == 0:
                # 随机下单
                direction = self.random_direction()
                if direction == "buy":
                    self.mt5.buy(self.symbol, self.initial_volume, tp=5)
                else:
                    self.mt5.sell(self.symbol, self.initial_volume, tp=5)
                time.sleep(100 / 1000)
                continue

            # 出场判断
            self.prominence()
            # 加仓


