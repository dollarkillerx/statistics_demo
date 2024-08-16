import time

from sdk import NewSystemSDK

address = "http://127.0.0.1:5800"
mt5_path = ""
suffix = ""
company_key = "test"

# 发布者
if __name__ == '__main__':
    sdk = NewSystemSDK(address, mt5_path, suffix, company_key)
    while True:
        # 暂停100毫秒
        time.sleep(0.1)
        try:
            sdk.broadcast()
        except Exception as e:
            print(e)
