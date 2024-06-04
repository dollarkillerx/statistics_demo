package demo1

import (
	"fmt"
	"github.com/dollarkillerx/statistics_demo/utils"
	"testing"
)

func TestDemo1(t *testing.T) {
	// 1985-2024 jpy/usd 汇率 最高值 最低 平均 中心 最平指
	// 提取数据并结构化
	data := utils.CleanData("USD_JPY.csv")

	// 上面数据为月度数据，我们需要将其转换为年度数据
	var yearMap = map[int][]float64{}
	for _, v := range data {
		year := v.Time.Year()
		if year < 1000 {
			continue
		}
		yearMap[year] = append(yearMap[year], v.Price)
	}
	fmt.Println("ok")

}
