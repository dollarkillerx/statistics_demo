[earthworm](earthworm.py) 交易策略]

- 进场：
    - 人工tradingview确定一个进场时间和方向
- 加仓：
    - 区间
        - 5point
    - 亏损
        - 逆行情方向 加*1.4
        - 正行情方向 加*1
    - 盈利
        - 亏损方向 加*1.4
        - 盈利方向 加*1
- 退出：
    - 总退出：
        - 利润 > 200 
        - SuperTrend STRATEGY KivancOzbilgic 进场信号
    - 顺市行情退出：
        - 当最新逆行情单出现盈利时
    - 再次进场: SuperTrend STRATEGY KivancOzbilgic 信号
