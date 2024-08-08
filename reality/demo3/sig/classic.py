import time
import random
import utils
from decimal import Decimal, getcontext
getcontext().prec = 28

class Classic:
    direction = "buy" # 方向
    magic = 0 # 魔术手666666
    deviation = 20 # 滑点
    currency_suffix = "c"  # 货币后缀
    initial_volume = 0.01  # 初始volume
    increase_multiple = 0.4  # 加仓倍数
    symbol = "EURUSD" # 货币
    time_interval = 30  # 时间间隔 default 30分
    interval = 5 # 加仓间隔
    sleepTime = 0 # 休息时间

    highest = 0

    def __init__(self, direction = "buy", magic=60606, deviation = 20, currency_suffix = "z", initial_volume = 0.01, increase_multiple = 0.4,symbol = "EURUSD", interval= 6,time_interval =30):
        self.direction = direction
        self.magic = magic
        self.deviation = deviation
        self.currency_suffix = currency_suffix
        self.initial_volume = initial_volume
        self.increase_multiple = increase_multiple
        self.symbol = symbol
        self.interval = interval
        self.time_interval = time_interval

    def init(self):
        self.mt5 = utils.MT5utils(path="C:\\Users\\Administrator\\Desktop\\MT5\\terminal64.exe")
        self.mt5.set_currency_suffix(self.currency_suffix)
        self.mt5.set_magic(self.magic)
        self.mt5.set_def_deviation(self.deviation)

    # 出场
    def prominence(self):
        profit = self.mt5.profit(magic=self.magic)
        # 移动止损
        if self.highest <= profit:
            self.highest = profit
        if self.highest >= 40:
            if self.highest - profit >= 15:
                self.mt5.close_all(magic=self.magic)
                self.highest = 0
                tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                print(tick)
                self.sleepTime = tick.time + 60*60*3
                return
        if self.highest >= 30:
            if self.highest - profit >= 15:
                self.mt5.close_all(magic=self.magic)
                self.highest = 0
                tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                print(tick)
                self.sleepTime = tick.time + 60*60*3
                return
        if self.highest >= 25:
            if self.highest - profit >= 10:
                self.mt5.close_all(magic=self.magic)
                tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                print(tick)
                self.sleepTime = tick.time + 60*60*3
                self.highest = 0
                return
        if profit < 0:
            if abs(profit) > 1000:
                self.mt5.close_all(magic=self.magic)
                tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                print(tick)
                self.sleepTime = tick.time + 60*60*3
                self.highest = 0
                return
        if self.mt5.positions_total(magic=self.magic) == 1:
            if profit >= 5:
                self.mt5.close_all(magic=self.magic)
                tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                print(tick)
                self.sleepTime = tick.time + 60*60*3
                self.highest = 0
                return

        # if profit >= 5:
        #     self.mt5.close_all(magic=self.magic)
        #     tick = self.mt5.symbol_info_tick2(symbol=self.symbol)
        #     print(tick)
        #     self.sleepTime = tick.time + 60*60*3
        #

    # 随机方向
    def random_direction(self):
        random_number = random.randint(1, 100)
        if random_number > 50:
            return "buy"
        return "sell"

    def run(self):
        while True:
            try:
                symbol_info_tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                # 是否休息
                if self.sleepTime != 0:
                    if self.sleepTime < symbol_info_tick.time:
                        print("continue")
                        continue
                    else:
                        self.sleepTime = 0

                # 出场判断
                self.prominence()

                # 获取当前订单列表
                last_position = self.mt5.last_position()
                if last_position is None:
                    # 随机下单
                    self.direction = self.random_direction()
                    if self.direction == "buy":
                        self.mt5.buy(self.symbol, self.initial_volume, tp=10)
                    else:
                        self.mt5.sell(self.symbol, self.initial_volume, tp=10)
                    print("----------------new-----------------")
                    continue

                # 加仓
                price = 0
                if self.direction == "buy":
                    price = symbol_info_tick.ask
                else:
                    price = symbol_info_tick.bid
                if abs(Decimal(
                        str((Decimal(str(last_position.price_open)) - Decimal(str(price))))) * 10000) > self.interval:
                    if abs(last_position.time_update - symbol_info_tick.time) > self.time_interval * 60:
                        profit = self.mt5.profit()
                        if profit < 0.2:
                            if self.direction == "buy":
                                self.mt5.sell(self.symbol,
                                              round(last_position.volume + self.increase_multiple,
                                                    2))
                            else:
                                self.mt5.buy(self.symbol,
                                             round(last_position.volume + self.increase_multiple,
                                                   2))
            except(Exception) as e:
                print("---------------------------eeeeeeeeeeeeeeeee-------------------------")
                print(e)
                print("---------------------------eeeeeeeeeeeeeeeee-------------------------")
            random_number = random.randint(1, 100)
            print(random_number)
            time.sleep(100 / 1000)







