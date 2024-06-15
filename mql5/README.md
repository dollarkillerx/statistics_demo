# MQL5 參考

### 基礎函數
``` 
int OnInit()

void OnDeinit(const int reason) 程序退出時執行

void OnTick() // tick 跳動時執行

void OnTimer() // OnInit中設置  EventSetTimer(60// 秒数);  更具設置執行 
 
void OnTrade() // 交易發生時執行

void OnTradeTransaction(const MqlTradeTransaction& trans,
                        const MqlTradeRequest& request,
                        const MqlTradeResult& result)


```