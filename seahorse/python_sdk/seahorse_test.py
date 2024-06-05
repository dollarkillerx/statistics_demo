import seahorse

if __name__ == '__main__':
    # sh = seahorse.Seahorse("127.0.0.1:8475", "test_account1", 1000, 2000)
    sh = seahorse.Seahorse("127.0.0.1:8475", "my_test", 1000, 2000)
    # sh.init_account()
    symbol = "EURUSD"
    print(sh.symbol_info_tick(symbol))
    print(sh.positions_total())
    print(sh.account_info())
    print(len(sh.positions_get()))
    print(sh.positions_get())
    #
    # sh.buy(symbol, 0.01)
    # sh.sell(symbol, 0.02)

    # sh.close(17)
    print(sh.last_position())
    print(sh.close_all())