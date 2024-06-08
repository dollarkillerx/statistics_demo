import json
import urllib.request
from types import SimpleNamespace

class Seahorse:
    def __init__(self, address, account: str, balance: float, lever: float):
        self.address = address
        self.account = account
        self.balance = balance
        self.lever = lever

    # 初始化账户
    def init_account(self):
        url = "http://{}/api/v1/init".format(self.address)
        data = {
            "account": self.account,
            "balance": self.balance,
            "lever": self.lever
        }

        # 将数据编码为 JSON
        json_data = json.dumps(data).encode('utf-8')

        # 创建请求对象
        req = urllib.request.Request(url, data=json_data, headers={'Content-Type': 'application/json'})
        try:
            with urllib.request.urlopen(req) as response:
                bytes = response.read()
                # 如果需要，将字节数据解码为字符串
                response_data = bytes.decode('utf-8')
                print(response_data)
        except urllib.error.URLError as e:
            print(e)

    def shutdown(self):
        pass

    def set_magic(self, magic):
        self.magic = magic

    def set_def_deviation(self, deviation):
        self.deviation = deviation

    # 设置货币后缀
    def set_currency_suffix(self, currency_suffix):
        self.currency_suffix = currency_suffix

    # 获取利润
    def profit(self, magic=0):
        if magic == 0:
            account_info = self.account_info()
            if account_info == None:
                raise ValueError("account_info error: {}".format(self.last_error()))
            return account_info.profit
        profit = 0
        positions = self.positions_get(magic=self.magic)
        for position in positions:
            profit += position.profit
        return profit

    def symbol_info_tick(self, symbol):
        url = "http://{}/api/v1/symbol_info_tick".format(self.address)
        data = {
            "symbol": symbol,
        }

        # 将数据编码为 JSON
        json_data = json.dumps(data).encode('utf-8')

        # 创建请求对象
        req = urllib.request.Request(url, data=json_data, headers={'Content-Type': 'application/json'})
        try:
            with urllib.request.urlopen(req) as response:
                bytes = response.read()
                # 将字节数据解码为字符串
                response_data = bytes.decode('utf-8')
                # 将响应字符串转换为 JSON
                response_json = json.loads(response_data)
                return SimpleNamespace(**response_json)
        except urllib.error.URLError as e:
            return {
                "error": e
            }

    def symbol_info_tick2(self, symbol):
        url = "http://{}/api/v1/symbol_info_tick2".format(self.address)
        data = {
            "symbol": symbol,
        }

        # 将数据编码为 JSON
        json_data = json.dumps(data).encode('utf-8')

        # 创建请求对象
        req = urllib.request.Request(url, data=json_data, headers={'Content-Type': 'application/json'})
        try:
            with urllib.request.urlopen(req) as response:
                bytes = response.read()
                # 将字节数据解码为字符串
                response_data = bytes.decode('utf-8')
                # 将响应字符串转换为 JSON
                response_json = json.loads(response_data)
                return SimpleNamespace(**response_json)
        except urllib.error.URLError as e:
            return {
                "error": e
            }

        # buy 市价
    def buy(self, symbol: str, volume: float, comment='', sl=0, tp=0, deviation=0):
        try:
            if symbol == '':
                print("what fuck")
                exit(0)
            # if deviation == 0:
            #     deviation = self.deviation
            point = 0.00001
            print(self.symbol_info_tick2(symbol))
            price = self.symbol_info_tick2(symbol).ask
            request = {
                "symbol": symbol,
                "volume": volume,
                "type": 0,
                "price": price,
                "account": self.account,
            }
            if sl != 0:
                request['sl'] = price - sl * point * 10
            if tp != 0:
                request['tp'] = price + tp * point * 10
            self._orderSend(0, symbol, volume, price, 0)
        except Exception as e:
            print(f'buy exception: {e}')
            exit(1)

        # sell 市价
    def sell(self, symbol: str, volume: float, comment='', sl=0, tp=0, deviation=0):
        try:
            point = 0.00001
            price = self.symbol_info_tick2(symbol).bid
            request = {
                "symbol": symbol,
                "volume": volume,
                "type": 1,
                "price": price,
                "deviation": deviation,
                "comment": comment,
                "account": self.account,
            }
            if sl != 0:
                request['sl'] = price + sl * point * 10
            if tp != 0:
                request['tp'] = price - tp * point * 10
            self._orderSend(0, symbol, volume, price, 1)
        except Exception as e:
            print(f'sell exception: {e}')
            exit(1)

    def positions_total(self, magic=0):
        url = "http://{}/api/v1/positions_total".format(self.address)
        data = {
            "account": self.account,
        }

        # 将数据编码为 JSON
        json_data = json.dumps(data).encode('utf-8')

        # 创建请求对象
        req = urllib.request.Request(url, data=json_data, headers={'Content-Type': 'application/json'})
        try:
            with urllib.request.urlopen(req) as response:
                bytes = response.read()
                # 将字节数据解码为字符串
                response_data = bytes.decode('utf-8')
                # 将响应字符串转换为 JSON
                response_json = json.loads(response_data)
                return SimpleNamespace(**response_json)

        except urllib.error.URLError as e:
            raise ValueError(e)

    def account_info(self):
        url = "http://{}/api/v1/account_info".format(self.address)
        data = {
            "account": self.account,
        }

        # 将数据编码为 JSON
        json_data = json.dumps(data).encode('utf-8')

        # 创建请求对象
        req = urllib.request.Request(url, data=json_data, headers={'Content-Type': 'application/json'})
        try:
            with urllib.request.urlopen(req) as response:
                bytes = response.read()
                # 将字节数据解码为字符串
                response_data = bytes.decode('utf-8')
                # 将响应字符串转换为 JSON
                response_json = json.loads(response_data)
                if response_json['balance'] + response_json['profit'] - response_json['margin'] < -10:
                    exit(0)
                    # self.close_all()
                return SimpleNamespace(**response_json)
        except urllib.error.URLError as e:
            raise ValueError(e)

    def _positions_get(self, symbol = "", ticket = 0):
        url = "http://{}/api/v1/positions_get".format(self.address)
        data = {
            "account": self.account,
            "symbol": symbol,
            "ticket": ticket,
        }

        # 将数据编码为 JSON
        json_data = json.dumps(data).encode('utf-8')

        # 创建请求对象
        req = urllib.request.Request(url, data=json_data, headers={'Content-Type': 'application/json'})
        try:
            with urllib.request.urlopen(req) as response:
                bytes = response.read()
                # 将字节数据解码为字符串
                response_data = bytes.decode('utf-8')
                # 将响应字符串转换为 JSON
                response_json = json.loads(response_data)
                resp = []
                for item in response_json['items']:
                    item['time_update'] = item['time']
                    resp.append(SimpleNamespace(**item))
                return resp
        except urllib.error.URLError as e:
            raise ValueError(e)

    # 获取当前持仓  获取当ea的持仓 mt5.positions_get(magic=mt5.magic)
    def positions_get(self, symbol="", magic=0):
        # ticket: id
        # time: unix
        # type: 0buy 1sell
        # magic:
        # volume:
        # sl
        # tp
        # price_open 开户价格
        # price_current 当前价格
        # swap 库存费
        # comment
        positions = ()
        if symbol != "":
            symbol = symbol
            positions = self._positions_get(symbol=symbol)
            if positions is None:
                raise ValueError("error: {} {}".format("error", symbol))
        else:
            positions = self._positions_get()
            if positions is None:
                raise ValueError("error: {} {}".format("error", symbol))

        if magic == 0:
            return positions
        filtered_positions = [pos for pos in positions if pos.magic == magic]

        return filtered_positions

    def _orderSend(self, position: int, symbol: str, volume: float, price: float, typ: int):
        url = "http://{}/api/v1/order_send".format(self.address)
        data = {
            "position": position,
            "symbol":  symbol,
            "volume":  volume,
            "type":    typ,
            "price":   price,
            "account": self.account,
        }

        # 将数据编码为 JSON
        json_data = json.dumps(data).encode('utf-8')

        # 创建请求对象
        req = urllib.request.Request(url, data=json_data, headers={'Content-Type': 'application/json'})
        try:
            with urllib.request.urlopen(req) as response:
                bytes = response.read()
                # 将字节数据解码为字符串
                response_data = bytes.decode('utf-8')
                # 将响应字符串转换为 JSON
                response_json = json.loads(response_data)
                return response_json
        except urllib.error.URLError as e:
            raise ValueError(e)

    # 关闭订单
    def close(self, orderId, deviation=0):
        try:
            # if deviation == 0:
            #     deviation = self.deviation
            positions = self._positions_get(ticket=orderId)
            print('------------------------------------')
            print(positions)
            if positions is None:
                raise ValueError("positions not found {}".format(orderId))
            if len(positions) != 1:
                raise ValueError("positions not found {}".format(orderId))
            position = positions[0]

            point = 0.00001
            symbol_info_tick = self.symbol_info_tick2(position.symbol)

            bid_price = symbol_info_tick.bid
            ask_price = symbol_info_tick.ask
            price = 0
            type = 1

            if position.type == 0:
                # buy
                price = bid_price
                type = 1
            else:
                # sell
                price = ask_price
                type = 0

            self._orderSend(position.ticket, position.symbol, position.volume, price, type)
        except Exception as e:
            print(f'close exception: {e}')
            exit(1)

    # 关闭所有订单
    def close_all(self, magic=0,comment = ''):
        url = "http://{}/api/v1/close_all".format(self.address)
        data = {
            "account": self.account,
            "comment":  comment,
        }

        # 将数据编码为 JSON
        json_data = json.dumps(data).encode('utf-8')

        # 创建请求对象
        req = urllib.request.Request(url, data=json_data, headers={'Content-Type': 'application/json'})
        try:
            with urllib.request.urlopen(req) as response:
                bytes = response.read()
                # 将字节数据解码为字符串
                response_data = bytes.decode('utf-8')
                # 将响应字符串转换为 JSON
                response_json = json.loads(response_data)
                return response_json
        except urllib.error.URLError as e:
            raise ValueError(e)

    def last_error(self):
        return ""

    # 获取最新的一个订单  buy, sell
    def last_position(self, orderType=""):
        positions = self.positions_get()
        if positions is None:
            raise ValueError("last_positions {}".format(self.last_error()))
        if orderType == "buy":
            rs = [pos for pos in positions if pos.type == 0]
            if len(rs) > 0:
                return rs[len(rs) - 1]
            return None
        elif orderType == "sell":
            rs = [pos for pos in positions if pos.type == 1]
            if len(rs) > 0:
                return rs[len(rs) - 1]
            return None
        else:
            if len(positions) > 0:
                return positions[len(positions) - 1]
            return None

