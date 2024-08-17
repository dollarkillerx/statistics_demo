import time

from sdk import NewSystemSDK

address = "http://100.109.41.39:5800"
mt5_path = "C:\\Users\\Administrator\\Desktop\\MT1\\terminal64.exe"
suffix = "c"
company_key = "hfm"
subscription_client_id = "exness.76620107"

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
