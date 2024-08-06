import time

from sdk import NewSystemSDK

address = "http://192.168.40.238:8181"
mt5_path = "C:\\Users\\Administrator\\Desktop\\MT2\\terminal64.exe"
suffix = "c"
company_key = "hfm"

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
