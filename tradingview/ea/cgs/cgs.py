# 雙向交易
import utils
import tradingview_sdk
from decimal import Decimal, getcontext
getcontext().prec = 28

class Cgs:
    buy_highest = 0 # buy 方向的最大盈利
    sell_highest = 0 # sell 方向的最大盈利

    def __init__(self,tradingview_sdk_url: str,mt5_path: str,magic=60606, deviation = 20, currency_suffix = "z", initial_volume = 0.01, increase_multiple = 0.4,symbol = "EURUSD", interval= 6):
        self.tradingview_sdk_url = tradingview_sdk_url
        self.mt5_path = mt5_path
        self.magic = magic
        self.deviation = deviation
        self.currency_suffix = currency_suffix
        self.initial_volume = initial_volume
        self.increase_multiple = increase_multiple
        self.symbol = symbol
        self.interval = interval

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
                profit = self.mt5.profit_by_order_type(orderType="sell")
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

    def run(self):
        while True:
            symbol_info_tick = self.mt5.symbol_info_tick(symbol=self.symbol)
