class FollowMeSDK:
    address = "http://127.0.0.1:9871"
    token = "FollowMe"
    mt5_path = ""

    def __init__(self,address:str, token: str, mt5_path: str):
        self.address = address
        self.token = token
        self.mt5_path = mt5_path

    # 發佈
    def release(self):
        pass

    # 訂閲
    def subscription(self):
        pass