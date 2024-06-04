import utils

if __name__ == '__main__':
    mt5 = utils.MT5utils()
    print(mt5.positions_total())
    mt5.set_magic(66666)
    last_position = mt5.last_position()
    if last_position is None:
        print("????")
        exit(0)
    # print(last_position.volume)
    # print(last_position.volume * 0.4)
    # print(round(last_position.volume * 0.4, 4))
    # print(last_position.volume)
    print("{}".format(0.01*0.4))
    print("{}".format(round(0.01*0.4,2)))
    print("{}".format(round(0.01 + round(0.01 * 0.4, 2),
                            2)))
