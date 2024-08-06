import random
import requests
import utils
from datetime import datetime, timedelta
import json


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
    def __init__(self, order_id, direction, symbol, magic, open_price, volume, market, swap, profit, common,
                 opening_time, closing_time):
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
            "direction": self.direction,
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
        history = [History.from_dict(history_item) for history_item in data["history"]]
        return BroadcastPayload(
            client_id=data["client_id"],
            account=account,
            positions=positions,
            history=history
        )


class History:
    def __init__(self, ticket, time_setup, type_, magic, position_id, volume_initial, price_current, symbol, comment):
        self.ticket = ticket
        self.time_setup = time_setup
        self.type = type_
        self.magic = magic
        self.position_id = position_id
        self.volume_initial = volume_initial
        self.price_current = price_current
        self.symbol = symbol
        self.comment = comment

    @classmethod
    def from_dict(cls, data):
        return cls(
            ticket=data.get("ticket"),
            time_setup=data.get("time_setup"),
            type_=data.get("type"),
            magic=data.get("magic"),
            position_id=data.get("position_id"),
            volume_initial=data.get("volume_initial"),
            price_current=data.get("price_current"),
            symbol=data.get("symbol"),
            comment=data.get("comment")
        )

    def to_dict(self):
        return {
            "ticket": self.ticket,
            "time_setup": self.time_setup,
            "type": self.type,
            "magic": self.magic,
            "position_id": self.position_id,
            "volume_initial": self.volume_initial,
            "price_current": self.price_current,
            "symbol": self.symbol,
            "comment": self.comment
        }


from pydantic import BaseModel
from typing import List, Optional


class Position(BaseModel):
    order_id: int
    direction: str
    symbol: str
    magic: int
    open_price: float
    volume: float
    market: float
    swap: float
    profit: float
    common: str
    opening_time: int
    closing_time: int
    common_internal: str
    opening_time_system: int
    closing_time_system: int


class Data(BaseModel):
    client_id: str
    subscription_client_id: str
    open_positions: Optional[List[Position]] = None  # 将 open_positions 设置为可选
    close_position: Optional[List[Position]] = None  # 将 close_positions 设置为可选


class ResponseBody(BaseModel):
    code: int
    msg: str
    data: Data


class NewSystemSDK:  # NEW_SYSTEM_SDK_CLASS
    def __init__(self, address: str, mt5_path: str, suffix="", company_key="",multiple=1,hardTakeProfit=0):
        self.address = address
        self.mt5_path = mt5_path
        self.suffix = suffix
        self.mt5 = utils.MT5utils(path=mt5_path)
        self.mt5.set_currency_suffix(suffix)
        self.multiple = multiple
        self.hardTakeProfit = hardTakeProfit
        account = self.mt5.get_mt5().account_info()
        self.client_id = company_key + "." + str(account.login)
        self.Account = Account(account.login, account.leverage, account.server, account.company, account.balance,
                               account.profit, account.margin)
        print(self.Account.to_dict())

    def broadcast(self):
        resPos = []
        # 获取当前持仓
        positions = self.mt5.positions_get()
        for position in positions:
            resPos.append(Positions.from_dict({
                "order_id": position.ticket,
                "direction": "SELL" if position.type == 1 else "BUY",
                "symbol": position.symbol,
                "magic": position.magic,
                "open_price": position.price_open,
                "volume": position.volume,
                "market": position.price_current,
                "swap": position.swap,
                "profit": position.profit,
                "common": position.comment,
                "opening_time": position.time,
                "closing_time": 0
            }))
        # 获取历史持仓
        # 获取当前时间
        current_time = datetime.now()
        # 设置历史数据的开始和结束时间
        start_time = current_time - timedelta(days=30)  # 假设从一年前开始获取数据
        end_time = current_time

        # 获取历史交易记录
        deals = self.mt5.get_mt5().history_orders_get(start_time, end_time)

        cm = {}
        history = []
        for deal in deals:
            newDeals = self.mt5.get_mt5().history_orders_get(position=deal.position_id)
            for item in newDeals:
                deal_dict = History.from_dict({
                    "ticket": item.ticket,
                    "time_setup": item.time_setup,
                    "type": "SELL" if item.type == 1 else "BUY",
                    "magic": item.magic,
                    "position_id": item.position_id,
                    "volume_initial": item.volume_initial,
                    "price_current": item.price_current,
                    "symbol": item.symbol,
                    "comment": item.comment,
                })
                if cm.get(item.ticket) is None:
                    history.append(deal_dict)
                    cm[item.ticket] = deal_dict

        account = self.mt5.get_mt5().account_info()
        account = Account(account.login, account.leverage, account.server, account.company, account.balance,
                          account.profit, account.margin)

        r = BroadcastPayload(self.client_id, account, resPos, history)
        response = requests.post(self.address + "/ea/broadcast", data=json.dumps(r.to_dict()),
                                 headers={"Content-Type": "application/json"})

        print(response.status_code, "   ", random.Random().randint(0, 10))

    def subscription(self, subscription_client_id, strategy_code="Reverse"):
        resPos = []
        # 获取当前持仓
        positions = self.mt5.positions_get()
        for position in positions:
            resPos.append(Positions.from_dict({
                "order_id": position.ticket,
                "direction": "SELL" if position.type == 1 else "BUY",
                "symbol": position.symbol,
                "magic": position.magic,
                "open_price": position.price_open,
                "volume": position.volume,
                "market": position.price_current,
                "swap": position.swap,
                "profit": position.profit,
                "common": position.comment,
                "opening_time": position.time,
                "closing_time": 0
            }))
        # 获取历史持仓
        # 获取当前时间
        current_time = datetime.now()
        # 设置历史数据的开始和结束时间
        start_time = current_time - timedelta(days=30)  # 假设从一年前开始获取数据
        end_time = current_time

        # 获取历史交易记录
        deals = self.mt5.get_mt5().history_orders_get(start_time, end_time)

        cm = {}
        history = []
        for deal in deals:
            newDeals = self.mt5.get_mt5().history_orders_get(position=deal.position_id)
            for item in newDeals:
                deal_dict = History.from_dict({
                    "ticket": item.ticket,
                    "time_setup": item.time_setup,
                    "type": "SELL" if item.type == 1 else "BUY",
                    "magic": item.magic,
                    "position_id": item.position_id,
                    "volume_initial": item.volume_initial,
                    "price_current": item.price_current,
                    "symbol": item.symbol,
                    "comment": item.comment,
                })
                if cm.get(item.ticket) is None:
                    history.append(deal_dict)
                    cm[item.ticket] = deal_dict

        account = self.mt5.get_mt5().account_info()
        account = Account(account.login, account.leverage, account.server, account.company, account.balance,
                          account.profit, account.margin)

        r = BroadcastPayload(self.client_id, account, resPos, history)
        rj = r.to_dict()
        rj['subscription_client_id'] = subscription_client_id
        rj['strategy_code'] = strategy_code
        response = requests.post(self.address + "/ea/subscription", data=json.dumps(rj),
                                 headers={"Content-Type": "application/json"})
        print(response.status_code, "   ", random.Random().randint(0, 10))

        response_body = ResponseBody.parse_raw(response.text)

        positions = self.mt5.positions_get()

        # 关闭订单
        if response_body.data.close_position:
            for closePos in response_body.data.close_position:
                for position in positions:
                    if position.comment == str(closePos.order_id):
                        self.mt5.close(position.ticket)

        if account.profit > self.hardTakeProfit:
            self.mt5.close_all()
            exit(0)

        # 开新订单
        if response_body.data.open_positions:
            for pos in response_body.data.open_positions:
                ex = False
                for position in positions:
                    if position.comment == str(pos.order_id):
                        ex = True
                # 执行卖单
                if ex == False:
                    current_timestamp = datetime.now().timestamp()
                    symbol = pos.symbol[0:len(pos.symbol)-1]
                    # print(pos.opening_time)
                    print(pos.opening_time,"   ", pos.symbol,current_timestamp - pos.opening_time > 600)
                    if current_timestamp - pos.opening_time > 600:
                        continue

                    symbol_info = self.mt5.symbol_info(symbol)
                    if symbol_info is None:
                        continue
                    if symbol_info.visible == False:
                        continue
                    # buy
                    if pos.direction == "BUY":
                        self.mt5.buy(symbol=symbol, volume=round(pos.volume * self.multiple,2), comment=str(pos.order_id))
                    # sell
                    if pos.direction == "SELL":
                        self.mt5.sell(symbol=symbol, volume=round(pos.volume * self.multiple,2), comment=str(pos.order_id))
