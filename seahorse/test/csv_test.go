package test

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"testing"
	"time"
)

type Tick struct {
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
}

func TestCsv(t *testing.T) {
	open, err := os.Open("../Exness_EURUSD_2023.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(open)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		fmt.Println(len(record[0]))
		fmt.Println(record[0], record[2])
	}
}

func TestTime(t *testing.T) {
	// Define the timestamp string
	timestamp := "2023-01-01 22:07:47.707Z"

	// Parse the timestamp string to a Time object
	// The layout string follows the reference time Mon Jan 2 15:04:05 -0700 MST 2006
	layout := "2006-01-02 15:04:05.000Z"
	tx, err := time.Parse(layout, timestamp)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return
	}

	// Convert the Time object to a Unix timestamp
	unixTimestamp := tx.Unix()
	fmt.Println(unixTimestamp)

}

func TestP2(t *testing.T) {
	// Define the string containing the float value
	str := "1.0714299999999999"

	// Convert the string to a float64
	floatVal, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	// Round the float to 5 decimal places
	roundedVal := math.Round(floatVal*100000) / 100000
	fmt.Println(roundedVal)
}
