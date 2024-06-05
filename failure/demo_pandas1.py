import pandas as pd

if __name__ == '__main__':
    data = {
        'Name': ['Alice', 'Bob', 'Charlie'],
        'Age': [25, 30, 35],
        'City': ['New York', 'Los Angeles', 'Chicago']
    }
    # Pandas的主要数据结构是DataFrame（二维表格数据）和Series（一维数据）
    df = pd.DataFrame(data)
    print(df)

    # 查看前几行数据
    print("------------head-----------")
    print(df.head())

    # 查看数据概况
    print("------------info-----------")
    print(df.info())

    # 描述性统计
    print("------------describe-----------")
    print(df.describe())

    # 选择单列
    print("------------选择单列-----------")
    print(df['Name'])

    # 选择多列
    print("------------选择多列-----------")
    print(df[['Name', 'Age']])

    # 通过行索引选择
    print("------------通过行索引选择-----------")
    print(df.loc[0])

    # 通过行和列索引选择
    print("------------通过行和列索引选择-----------")
    print(df.loc[0, 'Name'])

    # 条件选择
    print("------------条件选择-----------")
    print(df[df['Age'] > 30])

    # 添加列
    print("------------添加列-----------")
    df['Salary'] = [50000, 60000, 70000]
    print(df)

    # 删除列
    print("------------删除列-----------")
    df = df.drop('City', axis=1)
    print(df)

    # 修改数据
    print("------------修改数据-----------")
    df.loc[0, 'Age'] = 26
    print(df)