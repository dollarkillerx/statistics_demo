import utils
import time
from decimal import Decimal, getcontext

getcontext().prec = 28


class Earthworm:
    direction = "buy"  # buy sell 初始方向
    magic = 66666  # 魔术手
    deviation = 30  # 滑点
    currency_suffix = "z"  # 货币后缀
    initial_volume = 0.01  # 初始volume
    increase_multiple = 0.4  # 加仓倍数
    base_currency = "EURUSD"  # 基础
    interval = 5  # 加仓间隔
    time_interval = 30  # 时间间隔 default 30分

    def __init__(self, direction, magic, deviation, currency_suffix, initial_volume, increase_multiple, base_currency,
                 interval,
                 time_interval):
        self.mt5 = utils.MT5utils()
        self.direction = direction
        self.magic = magic
        self.deviation = deviation
        self.currency_suffix = currency_suffix
        self.initial_volume = initial_volume
        self.increase_multiple = increase_multiple
        self.base_currency = base_currency
        self.interval = interval
        self.time_interval = time_interval

    def init(self):
        self.mt5.set_currency_suffix(self.currency_suffix)
        self.mt5.set_magic(self.magic)
        self.mt5.set_def_deviation(self.deviation)
        # self.mt5.close_all()
        # exit(0)
        # 是否是再开  （再开恢复）
        positions = self.mt5.positions_get(magic=self.mt5.magic)
        if positions is None:
            raise ValueError("error: {}".format(self.mt5.last_error()))
        if len(positions) == 0:  # 第一次开
            if self.direction == "buy":
                self.mt5.buy(self.base_currency, self.initial_volume, "Genesis")
            else:
                self.mt5.sell(self.base_currency, self.initial_volume, "Genesis")

    # 是否平仓
    def closing_position(self):
        profit = self.mt5.profit(self.magic)
        if profit > 200:
            self.mt5.close_all(magic=self.magic)

    def run(self):
        while True:
            # 1. 查询最近的一个订单
            last_position = self.mt5.last_position()
            if last_position is None:
                print("当前没有一个订单")
                break

            # print(last_position)

            # 获取最新的价格
            symbol_info_tick = self.mt5.symbol_info_tick(self.base_currency)
            if symbol_info_tick is None:
                raise ValueError("啥 这都报错: {}".format(self.mt5.last_error()))
            price = 0
            if self.direction == "buy":
                price = symbol_info_tick.ask
            else:
                price = symbol_info_tick.bid

            # 价格差距是否大于 网格点数
            # print(abs(Decimal(str((Decimal(str(last_position.price_open)) - Decimal(str(price))))) * 10000))
            if abs(Decimal(
                    str((Decimal(str(last_position.price_open)) - Decimal(str(price))))) * 10000) > self.interval:
                # 检查时间是否尚可
                if abs(last_position.time_update - symbol_info_tick.time) > self.time_interval * 60:
                    # 盈利加仓
                    if self.mt5.profit(magic=self.mt5.magic) > 5:
                        if self.direction == "buy":
                            self.mt5.buy(self.base_currency, last_position.volume)
                        else:
                            self.mt5.sell(self.base_currency, last_position.volume)
                    else:  # 亏损加仓
                        if self.direction == "sell":
                            self.mt5.buy(self.base_currency,
                                         round(last_position.volume + self.increase_multiple,
                                               2))
                        else:
                            self.mt5.sell(self.base_currency,
                                          round(last_position.volume + self.increase_multiple,
                                                2))
            self.closing_position()
            time.sleep(100 / 1000)
