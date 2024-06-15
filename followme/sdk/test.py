import urllib.request
import json


r = "USDJPYz"
print(r[:len(r)-1])

exit(0)
# 定义目标URL和token
url = 'http://127.0.0.1:9871/subscription'  # 请替换为实际的URL
token = 'FollowMe'  # 请替换为实际的token

# 创建请求对象并设置header
request = urllib.request.Request(url, headers={'Authorization': token})

try:
    # 发送请求并获取响应
    with urllib.request.urlopen(request) as response:
        # 读取响应内容
        response_data = response.read()
        # 打印响应内容
        print(response_data.decode('utf-8'))
        # 解析JSON数据
        json_data = json.loads(response_data)

        # 定义Order类
        class Order:
            def __init__(self, id: str, type: int, price: float, amount: float, comment: str, symbol: str, magic: int, sl: float, tp: float, created_time: int):
                self.id = id
                self.type = type
                self.price = price
                self.amount = amount
                self.comment = comment
                self.symbol = symbol
                self.magic = magic
                self.sl = sl
                self.tp = tp
                self.created_time = created_time

            def __repr__(self):
                return f"Order(id={self.id}, type={self.type}, price={self.price}, amount={self.amount}, comment='{self.comment}', symbol='{self.symbol}', magic={self.magic}, sl={self.sl}, tp={self.tp}, created_time={self.created_time})"

        # 解析orders数据
        orders = [Order(**order) for order in json_data['orders']]

        # 打印Order对象列表
        for order in orders:
            print(order)

except urllib.error.HTTPError as e:
    # 打印HTTP错误信息
    print('HTTP Error:', e.code, e.reason)
except urllib.error.URLError as e:
    # 打印URL错误信息
    print('URL Error:', e.reason)