import followme_sdk

if __name__ == '__main__':
    sdk = followme_sdk.FollowMeSDK("http://127.0.0.1:9871","FollowMe","C:\\Users\\Administrator\\Desktop\\mt1\\terminal64.exe")
    sdk.release()