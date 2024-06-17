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
    suffix = ""

    def __init__(self,address:str, token: str, mt5_path: str, suffix = ""):
        self.address = address
        self.token = token
        self.mt5_path = mt5_path
        self.mt5 = utils.MT5utils(path=mt5_path)
        self.suffix = suffix

    # 發佈
    def release(self):
        print("Release")
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
        print("Subscribing...")
        while True:
            time.sleep(100 / 1000)
            # 创建请求对象并设置header
            request = urllib.request.Request("{}/subscription".format(self.address), headers={'Authorization': self.token})
            try:
                # 发送请求并获取响应
                with urllib.request.urlopen(request) as response:
                    print('Subscription Response Code:', response.getcode())
                    # 读取响应内容
                    response_data = response.read()
                    # 解析JSON数据
                    json_data = json.loads(response_data)
                    # 定义Order类
                    # 解析orders数据
                    orders = [Order(**order) for order in json_data['orders']]
                    # 打印Order对象列表
                    dictMap = {}
                    removeOrderMap = {}
                    print("orders: {}".format(len(orders)))

                    for index, order in enumerate(orders):
                        removeOrderMap[order.id] = 0

                    positions = self.mt5.positions_get(magic=self.mt5.magic)
                    for position in positions:
                        if position.comment not in removeOrderMap:
                            # close
                            self.mt5.close(position.ticket)

                    if len(orders) == 0:
                        self.mt5.close_all(magic=self.mt5.magic)

                    if len(orders) == 1:
                        continue

                    for position in positions:
                        dictMap[position.comment] = 0

                    for index, order in enumerate(orders):
                        if index == 0:
                            continue
                        if order.id not in dictMap:
                            # 獲取當前時間
                            symbol = ""
                            if len(order.symbol) != 6:
                                symbol = order.symbol[:len(order.symbol)-1] + self.suffix

                            tick = self.mt5.symbol_info_tick(symbol=symbol)
                            if tick is None:
                                print("Error1: 找不到貨幣: {}".format(symbol))
                                exit(1)

                            if abs(tick.time - order.created_time) < 6*60*60: # 6h
                                # 下單
                                if order.type == 0: # 反向
                                    self.mt5.sell(symbol, order.amount, order.id)
                                else:
                                    self.mt5.buy(symbol, order.amount, order.id)



            except urllib.error.HTTPError as e:
                # 打印HTTP错误信息
                print('HTTP Error:', e.code, e.reason)
            except urllib.error.URLError as e:
                # 打印URL错误信息
                print('URL Error:', e.reason)


