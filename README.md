# statistics_demo
[kcgi] ITのための統計学

[DEMO1 交易策略]
- 进场：  
  - 随机进场
- 加仓：
  - 区间
    - 15point
  - 亏损
    - 逆行情方向 加*0.4
    - 正行情方向 加*0.2
  - 盈利
    - 亏损方向 加*0.4
    - 盈利方向 加*0.2
- 退出：
  - 总退出： 
    - 利润 > 1000 ?
    - 亏损 > 1000 ?
  - 顺市行情退出：
    - 当最新逆行情单出现盈利时
  - 再次进场: sleep 4h

### MetaQuotes Ltd 公司是傻逼这文档写的真好 太好了
https://www.mql5.com/zh/docs/python_metatrader5

