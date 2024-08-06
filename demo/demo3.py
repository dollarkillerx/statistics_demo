from datetime import datetime, timedelta
import yfinance as yf

if __name__ == '__main__':
    # 获取今年8月5日之前90天的日元对美元汇率
    end_date = datetime(2024, 8, 5)
    start_date = end_date - timedelta(days=90)

    data_90_days = yf.download('JPY=X', start=start_date, end=end_date, interval='1d')

    # 计算中心值、平均值、最频值
    median_rate_90_days = data_90_days['Close'].median()
    mean_rate_90_days = data_90_days['Close'].mean()
    mode_rate_90_days = data_90_days['Close'].mode().iloc[0]

    print("计算中心值: ",median_rate_90_days, "平均值: ",mean_rate_90_days, "最频值: ",mode_rate_90_days)
