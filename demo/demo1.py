import scipy.stats as stats

# 给定数据
sample_mean = 40.7
population_mean = 41
sample_std = 0.9
sample_size = 25
alpha = 0.01

# 计算t统计量
t_statistic = (sample_mean - population_mean) / (sample_std / (sample_size ** 0.5))

# 计算临界值
df = sample_size - 1
critical_value = stats.t.ppf(alpha, df)

# 显示结果
print(f"t统计量: {t_statistic}")
print(f"临界值: {critical_value}")

# 进行决策
if t_statistic < critical_value:
    print("拒绝原假设: 劳动时间显著缩短")
else:
    print("不拒绝原假设: 没有足够证据表明劳动时间缩短")
