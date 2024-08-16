import time

from sdk import NewSystemSDK

address = "http://127.0.0.1:5800"
mt5_path = "C:\\MT2\\terminal64.exe"
suffix = ""
company_key = "test"
subscription_client_id = "test.85468370"

# 订阅者
if __name__ == '__main__':
    sdk = NewSystemSDK(address, mt5_path, suffix, company_key)
    while True:
        # 暂停100毫秒
        time.sleep(0.1)
        try:
            sdk.subscription(subscription_client_id)
        except Exception as e:
            print(e)
