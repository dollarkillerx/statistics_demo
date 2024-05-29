import MetaTrader5 as mt5
import pandas as pd


class MT5utils:
    currency_suffix = ''
    magic = 15770
    deviation = 30

    def __init__(self, user='', password=''):
        if not mt5.initialize():
            print("initialize() failed")
            mt5.shutdown()
            exit(1)
        if user != '' and password != '':
            authorized = mt5.login(user, password=password)
            if authorized:
                account_info = mt5.account_info()
                if account_info != None:
                    account_info_dict = mt5.account_info()._asdict()
                    df = pd.DataFrame(list(account_info_dict.items()), columns=['property', 'value'])
                    print("account_info() as dataframe:")
                    print(df)
        # print(mt5.terminal_info())
        # get data on MetaTrader 5 version
        # print(mt5.version())
        print("MetaTrader5 package author: ", mt5.__author__)
        print("MetaTrader5 package version: ", mt5.__version__)

    def shutdown(self):
        mt5.shutdown()

    def set_magic(self, magic):
        self.magic = magic

    def set_def_deviation(self, deviation):
        self.deviation = deviation

    # 设置货币后缀
    def set_currency_suffix(self, currency_suffix):
        self.currency_suffix = currency_suffix
        # try
        if mt5.symbol_info('USDJPY{}'.format(self.currency_suffix)) is None:
            mt5.shutdown()
            print('Wrong currency suffix USDJPY{}'.format(self.currency_suffix))
            exit(1)

    # [内部] 获取货币名称
    def _get_currency_name(self, symbol):
        return '{}{}'.format(symbol, self.currency_suffix)

    # [内部] symbol_info
    def _get_symbol_info(self, symbol):
        symbol_info = mt5.symbol_info(self._get_currency_name(symbol))
        if symbol_info is None:
            raise ValueError("symbol invalid")
        if not symbol_info.visible:
            raise ValueError("symbol invalid")
        return symbol_info

    # buy 市价
    def buy(self, symbol: str, volume: float, comment='', sl=0, tp=0, deviation=0):
        try:
            if deviation == 0:
                deviation = self.deviation
            point = self._get_symbol_info(symbol).point
            price = mt5.symbol_info_tick(self._get_currency_name(symbol)).ask
            request = {
                "action": mt5.TRADE_ACTION_DEAL,
                "symbol": self._get_currency_name(symbol),
                "volume": volume,
                "type": mt5.ORDER_TYPE_BUY,
                "price": price,
                "deviation": deviation,
                "magic": self.magic,
                "comment": comment,
                "type_time": mt5.ORDER_TIME_GTC,  # 订单将一直保留在队列中，直到手动取消
                "type_filling": mt5.ORDER_FILLING_FOK,  # 不满足条件不执行
            }
            if sl != 0:
                request['sl'] = price - sl * point * 10
            if tp != 0:
                request['tp'] = price + tp * point * 10
            result = mt5.order_send(request)
            print("[buy] {} {} lots at {} with deviation={} points".format(symbol, volume, price, deviation))
            if result.retcode != mt5.TRADE_RETCODE_DONE:
                print("order_send failed, retcode={}".format(result.retcode))
                # 请求词典结果并逐个元素显示
                result_dict = result._asdict()
                for field in result_dict.keys():
                    print("   {}={}".format(field, result_dict[field]))
                    # if this is a trading request structure, display it element by element as well
                    if field == "request":
                        traderequest_dict = result_dict[field]._asdict()
                        for tradereq_filed in traderequest_dict:
                            print(
                                "       traderequest: {}={}".format(tradereq_filed, traderequest_dict[tradereq_filed]))
                raise ValueError("order_send failed, retcode={}".format(result.retcode))
        except Exception as e:
            print(f'buy exception: {e}')
            exit(1)

    # sell 市价
    def sell(self, symbol: str, volume: float, comment='', sl=0, tp=0, deviation=0):
        try:
            if deviation == 0:
                deviation = self.deviation
            point = self._get_symbol_info(symbol).point
            price = mt5.symbol_info_tick(self._get_currency_name(symbol)).bid
            request = {
                "action": mt5.TRADE_ACTION_DEAL,
                "symbol": self._get_currency_name(symbol),
                "volume": volume,
                "type": mt5.ORDER_TYPE_SELL,
                "price": price,
                "deviation": deviation,
                "magic": self.magic,
                "comment": comment,
                "type_time": mt5.ORDER_TIME_GTC,  # 订单将一直保留在队列中，直到手动取消
                "type_filling": mt5.ORDER_FILLING_FOK,  # 不满足条件不执行
            }
            if sl != 0:
                request['sl'] = price + sl * point * 10
            if tp != 0:
                request['tp'] = price - tp * point * 10
            result = mt5.order_send(request)
            print("[sell] {} {} lots at {} with deviation={} points".format(symbol, volume, price, deviation))
            if result.retcode != mt5.TRADE_RETCODE_DONE:
                print("order_send failed, retcode={}".format(result.retcode))
                # 请求词典结果并逐个元素显示
                result_dict = result._asdict()
                for field in result_dict.keys():
                    print("   {}={}".format(field, result_dict[field]))
                    # if this is a trading request structure, display it element by element as well
                    if field == "request":
                        traderequest_dict = result_dict[field]._asdict()
                        for tradereq_filed in traderequest_dict:
                            print(
                                "       traderequest: {}={}".format(tradereq_filed, traderequest_dict[tradereq_filed]))
                raise ValueError("order_send failed, retcode={}".format(result.retcode))
        except Exception as e:
            print(f'sell exception: {e}')
            exit(1)

    # 持仓订单
    def positions_total(self, magic=0):
        if magic == 0:
            return mt5.positions_total()
        return len(self.positions_get(magic=magic))

    # 获取当前持仓  获取当ea的持仓 mt5.positions_get(magic=mt5.magic)
    def positions_get(self, symbol="", magic=0):
        # ticket: id
        # time: unix
        # type: 0buy 1sell
        # magic:
        # volume:
        # sl
        # tp
        # price_open 开户价格
        # price_current 当前价格
        # swap 库存费
        # comment
        positions = ()
        if symbol != "":
            symbol = self._get_currency_name(symbol)
            positions = mt5.positions_get(symbol=symbol)
            if positions is None:
                raise ValueError("error: {} {}".format(mt5.last_error(), symbol))
        else:
            positions = mt5.positions_get()
            if positions is None:
                raise ValueError("error: {} {}".format(mt5.last_error(), symbol))

        if magic == 0:
            return positions
        filtered_positions = [pos for pos in positions if pos.magic == magic]

        return filtered_positions

    # 获取利润
    def profit(self, magic=0):
        if magic == 0:
            account_info = mt5.account_info()
            if account_info == None:
                raise ValueError("account_info error: {}".format(mt5.last_error()))
            return account_info.profit
        profit = 0
        positions = self.positions_get(magic=self.magic)
        for position in positions:
            profit += position.profit
        return profit

    # 关闭订单
    def close(self, orderId, deviation=0):
        if deviation == 0:
            deviation = self.deviation
        positions = mt5.positions_get(ticket=orderId)
        if positions is None:
            raise ValueError("positions not found {}".format(orderId))
        if len(positions) != 1:
            raise ValueError("positions not found {}".format(orderId))
        position = positions[0]

        point = mt5.symbol_info(position.symbol).point
        symbol_info_tick = mt5.symbol_info_tick(position.symbol)

        bid_price = symbol_info_tick.bid
        ask_price = symbol_info_tick.ask
        price = 0
        type = mt5.ORDER_TYPE_SELL

        if position.type == 0:
            # buy
            price = bid_price
            type = mt5.ORDER_TYPE_SELL
        else:
            # sell
            price = ask_price
            type = mt5.ORDER_TYPE_BUY

        request = {
            "action": mt5.TRADE_ACTION_DEAL,
            "symbol": position.symbol,
            "volume": position.volume,
            "type": type,
            "position": position.ticket,  # 订单号
            "price": price,
            "deviation": deviation,
            "magic": self.magic,
            "comment": position.comment,
            "type_time": mt5.ORDER_TIME_GTC,  # 订单将一直保留在队列中，直到手动取消
            "type_filling": mt5.ORDER_FILLING_FOK,  # 不满足条件不执行
        }
        result = mt5.order_send(request)
        print("[close] {} {} {} lots at {} with deviation={} points".format(position.ticket, position.symbol,
                                                                            position.volume, price, deviation))
        if result.retcode != mt5.TRADE_RETCODE_DONE:
            print("order_send failed, retcode={}".format(result.retcode))
            # 请求词典结果并逐个元素显示
            result_dict = result._asdict()
            for field in result_dict.keys():
                print("   {}={}".format(field, result_dict[field]))
                # if this is a trading request structure, display it element by element as well
                if field == "request":
                    traderequest_dict = result_dict[field]._asdict()
                    for tradereq_filed in traderequest_dict:
                        print(
                            "       traderequest: {}={}".format(tradereq_filed, traderequest_dict[tradereq_filed]))
            raise ValueError("order_send failed, retcode={}".format(result.retcode))
