# 雙向交易
import utils
import tradingview_sdk
from decimal import Decimal, getcontext
getcontext().prec = 28

class Cgs:
    def __init__(self,magic=60606, deviation = 20, currency_suffix = "z", initial_volume = 0.01, increase_multiple = 0.4,symbol = "EURUSD", interval= 6):
        self.magic = magic
        self.deviation = deviation

