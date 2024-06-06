import utils

if __name__ == '__main__':
    mt5 = utils.MT5utils()
    positions = mt5.positions_get()
    buyNum = 0
    selNum = 0
    for position in positions:
        if position.volume < 0.1:
            continue
        if position.type == 0:
            buyNum += position.profit
        else:
            selNum += position.profit

    print("buy: {} sell: {}".format(buyNum,selNum))
    print(buyNum > selNum)