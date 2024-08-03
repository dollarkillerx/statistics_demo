from sdk import NewSystemSDK

address = "127.0.0.1"
mt5_path = ""
suffix = ""
company_key = "exness"

if __name__ == '__main__':
    sdk = NewSystemSDK(address, mt5_path, suffix, company_key)
    sdk.broadcast()
