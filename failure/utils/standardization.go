package utils

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

func CleanData(path string) []ExchangeRate {
	open, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var items []ExchangeRate
	reader := csv.NewReader(bufio.NewReader(open))
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		timestamp, err := DateStringToTimestamp(record[0])

		f, err := strconv.ParseFloat(record[1], 64)
		if err == nil {
			items = append(items, ExchangeRate{
				Time:     timestamp,
				Timeunix: timestamp.Unix(),
				Price:    f,
			})
		}
	}

	return items
}

func DateStringToTimestamp(dateString string) (time.Time, error) {
	// 使用指定的布局解析日期字符串
	layout := "2006-01-02" // 因为 Go 的时间格式必须是 2006-01-02 15:04:05 这样的形式
	t, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
