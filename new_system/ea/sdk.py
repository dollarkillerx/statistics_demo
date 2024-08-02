import utils
import pandas as pd


class Account:
    def __init__(self, account, leverage, server, company, balance, profit, margin):
        self.account = account
        self.leverage = leverage
        self.server = server
        self.company = company
        self.balance = balance
        self.profit = profit
        self.margin = margin

    def to_dict(self):
        return {
            "account": self.account,
            "leverage": self.leverage,
            "server": self.server,
            "company": self.company,
            "balance": self.balance,
            "profit": self.profit,
            "margin": self.margin,
        }

    @staticmethod
    def from_dict(data):
        return Account(
            account=data["account"],
            leverage=data["leverage"],
            server=data["server"],
            company=data["company"],
            balance=data["balance"],
            profit=data["profit"],
            margin=data["margin"]
        )


class Positions:
    def __init__(self, order_id, direction, symbol, magic, open_price, volume, market, swap, profit, common, opening_time, closing_time):
        self.order_id = order_id
        self.direction = direction
        self.symbol = symbol
        self.magic = magic
        self.open_price = open_price
        self.volume = volume
        self.market = market
        self.swap = swap
        self.profit = profit
        self.common = common
        self.opening_time = opening_time
        self.closing_time = closing_time

    def to_dict(self):
        return {
            "order_id": self.order_id,
            "direction": self.direction.value,
            "symbol": self.symbol,
            "magic": self.magic,
            "open_price": self.open_price,
            "volume": self.volume,
            "market": self.market,
            "swap": self.swap,
            "profit": self.profit,
            "common": self.common,
            "opening_time": self.opening_time,
            "closing_time": self.closing_time,
        }

    @staticmethod
    def from_dict(data):
        return Positions(
            order_id=data["order_id"],
            direction=data["direction"],
            symbol=data["symbol"],
            magic=data["magic"],
            open_price=data["open_price"],
            volume=data["volume"],
            market=data["market"],
            swap=data["swap"],
            profit=data["profit"],
            common=data["common"],
            opening_time=data["opening_time"],
            closing_time=data["closing_time"]
        )

class BroadcastPayload:
    def __init__(self, client_id, account, positions, history):
        self.client_id = client_id
        self.account = account
        self.positions = positions
        self.history = history

    def to_dict(self):
        return {
            "client_id": self.client_id,
            "account": self.account.to_dict(),
            "positions": [position.to_dict() for position in self.positions],
            "history": [history_item.to_dict() for history_item in self.history],
        }

    @staticmethod
    def from_dict(data):
        account = Account.from_dict(data["account"])
        positions = [Positions.from_dict(position) for position in data["positions"]]
        history = [Positions.from_dict(history_item) for history_item in data["history"]]
        return BroadcastPayload(
            client_id=data["client_id"],
            account=account,
            positions=positions,
            history=history
        )


class NewSystemSDK:  # NEW_SYSTEM_SDK_CLASS
    def __init__(self, address: str, mt5_path: str, suffix="", company_key=""):
        self.address = address
        self.mt5_path = mt5_path
        self.suffix = suffix
        self.mt5 = utils.MT5utils(path=mt5_path)
        account = self.mt5.get_mt5().account_info()
        self.client_id = company_key + "." + str(account.login)
        self.Account = Account(account.login, account.leverage, account.server, account.company, account.balance,
                               account.profit, account.margin)
        print(self.Account.to_dict())

