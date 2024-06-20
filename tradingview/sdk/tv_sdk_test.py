import tradingview_sdk

if __name__ == '__main__':
    tvSDK = tradingview_sdk.TradingviewSDK("https://tradingview.dollarkiller.com")
    try:
        rk = tvSDK.get_by_symbol("EURUSD")
    except Exception as e:
        print(e)

