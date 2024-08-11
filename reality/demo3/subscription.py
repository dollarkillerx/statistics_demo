import time

from sdk import NewSystemSDK

address = "http://192.168.40.198:8182"
mt5_path = "C:\\Users\\Administrator\\Desktop\\MT4\\terminal64.exe"
suffix = "m"
company_key = "exness"
subscription_client_id = "exness.76667118"
multiple=2.5
hardTakeProfit = 100

# 订阅者
if __name__ == '__main__':
    sdk = NewSystemSDK(address, mt5_path, suffix, company_key,multiple,hardTakeProfit)
    while True:
        # 暂停100毫秒
        time.sleep(0.1)
        try:
            sdk.subscription(subscription_client_id)
        except Exception as e:
            print(e)
