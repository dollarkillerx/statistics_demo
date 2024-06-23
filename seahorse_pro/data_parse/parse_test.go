package data_parse

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

// https://www.histdata.com/
// 数据下载

func TestDataParse(t *testing.T) {
	input := "EURUSDm_M1_202403180821_202405312058.csv"
	output := "eurusd_m1.csv"

	open, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer open.Close()

	datas := make([]Data, 0)

	reader := csv.NewReader(open)
	reader.Comma = '\t'
	for {
		record, err := reader.Read()
		if err == nil {
			item := Data{
				TimeStr: record[0] + " " + record[1],
				Time:    timeStringToTimestamp(record[0] + " " + record[1]),
				Open:    ParseFloat(record[2]),
				High:    ParseFloat(record[3]),
				Low:     ParseFloat(record[4]),
				Close:   ParseFloat(record[5]),
				Volume:  ParseFloat(record[6]),
				Spread:  ParseInt(record[8]),
			}
			datas = append(datas, item)
		} else {
			break
		}
	}

	create, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(create)
	for _, data := range datas {
		writer.Write([]string{
			data.TimeStr,
			data.Time,
		})
	}

}

type Data struct {
	TimeStr string  `json:"time_str"`
	Time    int64   `json:"time"`
	Open    float64 `json:"open"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Close   float64 `json:"close"`
	Volume  float64 `json:"volume"`
	Spread  int64   `json:"spread"`
}

func (d *Data) Print() {
	indent, err := json.MarshalIndent(d, "", "  ")
	if err == nil {
		fmt.Println(string(indent))
	}
}

func ParseFloat(str string) float64 {
	floatNum, _ := strconv.ParseFloat(str, 64)
	return floatNum
}

func ParseInt(str string) int64 {
	floatNum, _ := strconv.ParseInt(str, 10, 64)
	return floatNum
}

// 将时间字符串转换为 Unix 时间戳的函数
func timeStringToTimestamp(timeStr string) int64 {
	// 定义时间格式
	timeFormat := "2006.01.02 15:04:05"

	// 解析时间字符串
	t, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		return 0
	}

	// 返回 Unix 时间戳
	return t.Unix()
}
