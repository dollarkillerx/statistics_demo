import MetaTrader5 as mt5
import time
from utils import MT5utils

class MartingaleGridTrading:
    def __init__(self, user, password, symbol, grid_points=5, initial_volume=0.01, max_risk=0.1, target_profit=0.02):
        self.mt5 = MT5utils(user, password)
        self.symbol = symbol
        self.grid_points = grid_points
        self.initial_volume = initial_volume
        self.max_risk = max_risk
        self.target_profit = target_profit
        self.initial_balance = self.mt5.balance()
        self.positions = []

    def start(self):
        # Initialize positions
        self.buy_order(self.initial_volume)
        self.sell_order(self.initial_volume)

        while True:
            try:
                self.check_positions()
                self.check_risk()
                self.check_target_profit()
                time.sleep(60)  # Check every minute
            except Exception as e:
                print(f"Error: {e}")
                self.mt5.shutdown()
                break

    def buy_order(self, volume):
        result = self.mt5.buy(self.symbol, volume, take_profit=self.grid_points)
        if result:
            self.positions.append(result.order)
        print(f"Placed buy order: volume={volume}")

    def sell_order(self, volume):
        result = self.mt5.sell(self.symbol, volume, take_profit=self.grid_points)
        if result:
            self.positions.append(result.order)
        print(f"Placed sell order: volume={volume}")

    def check_positions(self):
        for position in self.mt5.positions_get():
            if position.symbol == self.symbol:
                if position.profit < 0:  # Losing position
                    if position.type == mt5.ORDER_TYPE_BUY:
                        new_volume = position.volume * 1.4
                        self.sell_order(new_volume)
                    else:
                        new_volume = position.volume * 1.4
                        self.buy_order(new_volume)

    def check_risk(self):
        current_equity = self.mt5.equity()
        drawdown = self.initial_balance - current_equity
        if drawdown / self.initial_balance > self.max_risk:
            print(f"Max risk exceeded: {drawdown / self.initial_balance:.2%}")
            self.close_all_positions()
            raise Exception("Max risk exceeded")

    def check_target_profit(self):
        current_equity = self.mt5.equity()
        profit = current_equity - self.initial_balance
        if profit / self.initial_balance >= self.target_profit:
            print(f"Target profit reached: {profit / self.initial_balance:.2%}")
            self.close_all_positions()
            raise Exception("Target profit reached")

    def close_all_positions(self):
        self.mt5.close_all()
        self.positions = []
        print("All positions closed")

# Example usage
if __name__ == "__main__":
    user = "your_username"
    password = "your_password"
    symbol = "EURUSD"

    trader = MartingaleGridTrading(user, password, symbol)
    trader.start()
