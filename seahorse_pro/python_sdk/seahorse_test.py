import seahorse

if __name__ == '__main__':
    sh = seahorse.Seahorse("127.0.0.1:8475", "my_test", 1000, 2000)
    symbol = "EURUSD"
    sh.next()
    sh.init_account()
    r = sh.symbol_info_tick(symbol=symbol)
    print(r)