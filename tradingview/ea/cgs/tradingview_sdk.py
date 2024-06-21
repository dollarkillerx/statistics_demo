import json
import urllib.request

class TradingData:
    def __init__(self, ticker, action, contracts, price):
        self.ticker = ticker
        self.action = action
        self.contracts = contracts
        self.price = price

    def __repr__(self):
        return f"TradingData(ticker={self.ticker}, action={self.action}, contracts={self.contracts}, price={self.price})"

class TradingviewSDK:
    def __init__(self, url: str):
        self.url = url

    def get_by_symbol(self,symbol: str):
        url = "{}/by/{}".format(self.url,symbol)
        response = urllib.request.urlopen(url)
        if response.status == 404:
            raise Exception("Error 404: URL not found")
        data = response.read()

        # Parse the JSON data
        json_data = json.loads(data)

        # Create an instance of TradingData class
        return TradingData(
            ticker=json_data['ticker'],
            action=json_data['action'],
            contracts=json_data['contracts'],
            price=json_data['price']
        )