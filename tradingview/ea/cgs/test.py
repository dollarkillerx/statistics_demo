import utils

if __name__ == '__main__':
    mt5 = utils.MT5utils()
    mt5.set_magic(0)
    positions = mt5.positions_get_by_type(orderType="sell")
    if positions is not None:
        for pos in positions:
            print(pos)

    print("=-=-=-=-=-=-=-=-=-=-=-=")
    positions = mt5.positions_get_by_type(orderType="buy")
    if positions is not None:
        for pos in positions:
            print(pos)

    print(mt5.profit_by_order_type(orderType="sell"))
    print(mt5.profit_by_order_type(orderType="buy"))

    mt5.close_all_by_order_type(orderType="sell")
    print(mt5.profit_by_order_type(orderType="sell"))
    print(mt5.profit_by_order_type(orderType="buy"))
