import time
import random
import seahorse as utils
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
        self.mt5 = utils.Seahorse("127.0.0.1:8475", "my_test",5000,2000)
        self.mt5.init_account()
        self.mt5.set_currency_suffix(self.currency_suffix)
        self.mt5.set_magic(self.magic)
        self.mt5.set_def_deviation(self.deviation)

    # 出场
    def prominence(self):
        profit = self.mt5.profit(magic=self.magic)
        if profit > 3:
            self.mt5.close_all()


    # 随机方向
    def random_direction(self):
        random_number = random.randint(1, 100)
        if random_number > 50:
            return "buy"
        return "sell"

    def next_direction(self):
        positions = self.mt5.positions_get(symbol=self.symbol, magic=self.magic)
        buyFloat = 0.0
        selFloat = 0.0
        for position in positions:
            if position.type == 0:
                buyFloat += position.profit
            else:
                selFloat += position.profit
        if buyFloat > selFloat:
            return "buy"
        return "sell"

    def ok(self, ask: float):
        positions = self.mt5.positions_get(symbol=self.symbol, magic=self.magic)
        for position in positions:
            if abs(position.price_open -  ask) < 5:
                return False
        return True

    def run(self):
        while True:
            self.mt5.next() # 获取tick
            symbol_info_tick = self.mt5.symbol_info_tick(symbol=self.symbol)

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
                # time.sleep(100 / 1000)
                continue

            # 加仓
            price = symbol_info_tick.ask
            if abs(Decimal(
                    str((Decimal(str(last_position.price_open)) - Decimal(str(price))))) * 10000) > self.interval:
                        if abs(last_position.time_update - symbol_info_tick.time) > self.time_interval * 60:
                                profit = self.mt5.profit()
                                if profit < 5:
                                    if self.ok(price):
                                        if self.next_direction() == "buy":
                                            self.mt5.buy(self.symbol,
                                                         round(last_position.volume * 2,
                                                               2))
                                        else:
                                            self.mt5.buy(self.symbol,
                                                         round(last_position.volume * 2,
                                                               2))



            time.sleep(100 / 1000)



