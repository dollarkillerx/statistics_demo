import time

from sdk import NewSystemSDK

address = "http://100.109.41.39:5800"
mt5_path = "C:\\Users\\Administrator\\Desktop\\MT3\\terminal64.exe"
suffix = "m"
company_key = "exness"

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
