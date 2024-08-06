import yfinance as yf
import pandas as pd

if __name__ == '__main__':
    # 获取1985年到2024年的日元对美元汇率
    data = yf.download('JPY=X', start='1985-01-01', end='2024-12-31', interval='1d')

    # 只保留每年的最后一个交易日
    data['Year'] = data.index.year
    data_yearly = data.groupby('Year').tail(1)

    # 计算中心值、平均值、最频值
    median_rate = data_yearly['Close'].median()
    mean_rate = data_yearly['Close'].mean()
    mode_rate = data_yearly['Close'].mode().iloc[0]

    print("计算中心值: ",median_rate, "平均值: ",mean_rate, "最频值: ",mode_rate)
