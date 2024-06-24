import utils
import tradingview_sdk
import time
from decimal import Decimal, getcontext

getcontext().prec = 28

class Encryption:
    buy_highest = 0  # buy 方向的最大盈利
    sell_highest = 0  # sell 方向的最大盈利

    direction = ""  # 方向
    direction_time = 0  # 得到這個方向通知的時間

    def __init__(self, tradingview_sdk_url: str, mt5_path: str, magic=60606, deviation=20, currency_suffix="z",
                 initial_volume=0.01, increase_multiple=0.4, symbol="BTCUSD", interval=100, time_interval=15,
                 max_order=3):
        self.tradingview_sdk_url = tradingview_sdk_url
        self.mt5_path = mt5_path
        self.magic = magic
        self.deviation = deviation
        self.currency_suffix = currency_suffix
        self.initial_volume = initial_volume
        self.increase_multiple = increase_multiple
        self.symbol = symbol
        self.interval = interval
        self.time_interval = time_interval
        self.max_order = max_order

    def init(self):
        self.mt5 = utils.MT5utils(path=self.mt5_path)
        self.mt5.set_currency_suffix(self.currency_suffix)
        self.mt5.set_magic(self.magic)
        self.mt5.set_def_deviation(self.deviation)
        self.tradingview_sdk = tradingview_sdk.TradingviewSDK(url=self.tradingview_sdk_url)

    # 出场
    def prominence(self):
        # buy 方向是否存在订单
        buy_positions = self.mt5.positions_get_by_type(orderType="buy")
        if buy_positions is not None:
            if len(buy_positions) > 0:
                profit = self.mt5.profit_by_order_type(orderType="buy")
                if self.buy_highest <= profit:
                    self.buy_highest = profit

                if len(buy_positions) == 1 and self.direction == "sell":
                    self.mt5.close_all_by_order_type(orderType="buy")
                    self.buy_highest = 0
                    tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                    print(tick)
                    return
                if profit > 5:
                    if self.direction == "sell":
                        self.mt5.close_all_by_order_type(orderType="buy")
                        self.buy_highest = 0
                        tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                        print(tick)
                        return
                if self.buy_highest >= 40:
                    if self.buy_highest - profit >= 15:
                        self.mt5.close_all_by_order_type(orderType="buy")
                        self.buy_highest = 0
                        tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                        print(tick)
                        return
                if self.buy_highest >= 30:
                    if self.buy_highest - profit >= 15:
                        self.mt5.close_all_by_order_type(orderType="buy")
                        self.buy_highest = 0
                        tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                        print(tick)
                        return
                if self.buy_highest >= 25:
                    if self.buy_highest - profit >= 10:
                        self.mt5.close_all_by_order_type(orderType="buy")
                        self.buy_highest = 0
                        tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                        print(tick)
                        return
        # sell 方向是否存在订单
        sell_positions = self.mt5.positions_get_by_type(orderType="sell")
        if sell_positions is not None:
            if len(sell_positions) > 0:
                profit = self.mt5.profit_by_order_type(orderType="buy")
                if len(sell_positions) == 1 and self.direction == "sell":
                    self.mt5.close_all_by_order_type(orderType="sell")
                    self.sell_highest = 0
                    tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                    print(tick)
                    return
                if profit > 5:
                    if self.direction == "buy":
                        self.mt5.close_all_by_order_type(orderType="sell")
                        self.sell_highest = 0
                        tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                        print(tick)
                        return
                if self.sell_highest <= profit:
                    self.sell_highest = profit
                if self.sell_highest >= 40:
                    if self.sell_highest - profit >= 15:
                        self.mt5.close_all_by_order_type(orderType="sell")
                        self.sell_highest = 0
                        tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                        print(tick)
                        return
                if self.sell_highest >= 30:
                    if self.sell_highest - profit >= 15:
                        self.mt5.close_all_by_order_type(orderType="sell")
                        self.sell_highest = 0
                        tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                        print(tick)
                        return
                if self.sell_highest >= 25:
                    if self.sell_highest - profit >= 10:
                        self.mt5.close_all_by_order_type(orderType="sell")
                        self.sell_highest = 0
                        tick = self.mt5.symbol_info_tick(symbol=self.symbol)
                        print(tick)
                        return

    def _current_timestamp(self):
        return int(time.time())

    def run(self):
        while True:
            symbol_info_tick = self.mt5.symbol_info_tick(symbol=self.symbol)

            # 如果buy單是空的 收到了 buy信號 就開單
            buy_positions = self.mt5.positions_get_by_type(symbol=self.symbol, orderType="buy")
            if buy_positions is not None:
                if len(buy_positions) == 0:
                    # 如果 沒有訂單就等待消息
                    if self.direction == "buy":
                        # 收到消息在30s内
                        if abs(self.direction_time - self._current_timestamp()) < 30:
                            self.mt5.buy(self.symbol, self.initial_volume, tp=10)

            # 如果sell單是空的 收到了 sell 信號 就開單
            sell_positions = self.mt5.positions_get_by_type(symbol=self.symbol, orderType="sell")
            if sell_positions is not None:
                if len(sell_positions) == 0:
                    # 如果 沒有訂單就等待消息
                    if self.direction == "sell":
                        # 收到消息在30s内
                        if abs(self.direction_time - self._current_timestamp()) < 30:
                            self.mt5.sell(self.symbol, self.initial_volume, tp=10)

            # 獲取最新的信號
            try:
                resp = self.tradingview_sdk.get_by_symbol(symbol=self.symbol)
                self.direction = resp.action
                self.direction_time = resp.time
            except Exception as e:
                if "404" in "{}".format(e):
                    pass

            # 加倉邏輯
            # buy order 加倉邏輯
            buy_last_order = self.mt5.last_position(orderType="buy")
            if buy_last_order is not None:
                positions = self.mt5.positions_get_by_type(orderType="buy")
                if len(positions) < self.max_order:
                    price = symbol_info_tick.ask
                    if abs(Decimal(
                            str((Decimal(str(buy_last_order.price_open)) - Decimal(
                                str(price))))) * 10000) > self.interval:
                        if abs(buy_last_order.time_update - symbol_info_tick.time) > self.time_interval * 60:
                            self.mt5.buy(self.symbol,
                                         round(buy_last_order.volume + self.increase_multiple,
                                               2))

            # sell order 加倉邏輯
            sell_last_order = self.mt5.last_position(orderType="sell")
            if sell_last_order is not None:
                positions = self.mt5.positions_get_by_type(orderType="sell")
                if len(positions) < self.max_order:
                    price = symbol_info_tick.bid
                    if abs(Decimal(
                            str((Decimal(str(sell_last_order.price_open)) - Decimal(
                                str(price))))) * 10000) > self.interval:
                        if abs(sell_last_order.time_update - symbol_info_tick.time) > self.time_interval * 60:
                            self.mt5.sell(self.symbol,
                                          round(sell_last_order.volume + self.increase_multiple,
                                                2))

            time.sleep(100 / 1000)
