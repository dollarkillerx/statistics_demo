import time

from sdk import NewSystemSDK

address = "http://127.0.0.1:8181"
mt5_path = ""
suffix = ""
company_key = "exness"
subscription_client_id = "exness.76620107"

# 订阅者
if __name__ == '__main__':
    sdk = NewSystemSDK(address, mt5_path, suffix, company_key)
    sdk.subscription(subscription_client_id)
    # while True:
    #     # 暂停100毫秒
    #     time.sleep(0.1)
    #     try:
    #         sdk.subscription(subscription_client_id)
    #     except Exception as e:
    #         print(e)
