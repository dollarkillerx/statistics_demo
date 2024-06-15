import json
import time

import utils
import urllib.request

class Order:
    def __init__(self, id: str, type: int, price: float, amount: float, comment: str, symbol: str, created_time: int,magic: int,sl: float, tp: float):
        self.id = str(id)
        self.type = type
        self.price = price
        self.amount = amount
        self.comment = comment
        self.symbol = symbol
        self.created_time = created_time
        self.magic = magic
        self.sl = sl
        self.tp = tp

    def to_dict(self):
        return {
            "id": self.id,
            "type": self.type,
            "price": self.price,
            "amount": self.amount,
            "comment": self.comment,
            "symbol": self.symbol,
            "created_time": self.created_time,
            "magic": self.magic,
            "sl": self.sl,
            "tp": self.tp
        }

class FollowMeSDK:
    address = "http://127.0.0.1:9871"
    token = "FollowMe"
    mt5_path = ""

    def __init__(self,address:str, token: str, mt5_path: str):
        self.address = address
        self.token = token
        self.mt5_path = mt5_path
        self.mt5 = utils.MT5utils(path=mt5_path)

    # 發佈
    def release(self):
        while True:
            orders = []
            positions = self.mt5.positions_get()
            for position in positions:
                orders.append(Order(position.ticket,position.type, position.price_open,
                                    position.volume,position.comment, position.symbol,position.time, position.magic, position.sl, position.tp))
            orders_dict_list = [order.to_dict() for order in orders]
            orders_json = json.dumps({
                "orders": orders_dict_list
            }, indent=4)

            # Define the endpoint URL
            request = urllib.request.Request("{}/release".format(self.address), data=orders_json.encode('utf-8'), headers={'Content-Type': 'application/json','Authorization': self.token})
            try:
                # 发送请求并获取响应
                with urllib.request.urlopen(request) as response:
                    if response.getcode() != 200:
                        # 读取响应内容
                        response_data = response.read()
                        # 打印响应内容
                        # print(response_data.decode('utf-8'))
                    # 打印状态码
                    print('Release Response Code:', response.getcode())
            except urllib.error.HTTPError as e:
                # 打印HTTP错误信息
                print('HTTP Error:', e.code, e.reason)
            except urllib.error.URLError as e:
                # 打印URL错误信息
                print('URL Error:', e.reason)
            time.sleep(100 / 1000)

    # 訂閲
    def subscription(self):
        pass


