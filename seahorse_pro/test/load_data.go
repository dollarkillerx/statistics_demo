package test

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"seahorse/internal/conf"
	"seahorse/internal/models"
	"seahorse/internal/storage"
	"strconv"
	"time"
)

// 数据量太大了就不导入了

func main() {
	conf := conf.LoadConf()
	storage := storage.New(conf)

	open, err := os.Open(conf.CsvPath)
	if err != nil {
		panic(err)
	}
	defer open.Close()

	datas := []models.Tick{}

	// Load data from csv file
	reader := csv.NewReader(open)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if record[2] == "Timestamp" {
			continue
		}

		layout := "2006-01-02 15:04:05.000Z"
		tx, err := time.Parse(layout, record[2])
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			return
		}

		bidVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}

		askVal, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}

		// Save data to database
		tick := models.Tick{
			Symbol:    record[1],
			Timestamp: tx.Unix(),
			Ask:       math.Round(bidVal*100000) / 100000,
			Bid:       math.Round(askVal*100000) / 100000,
		}

		if len(datas) < 1000 {
			datas = append(datas, tick)
		} else {
			err = storage.Bb.Model(&models.Tick{}).CreateInBatches(&datas, 300).Error
			if err != nil {
				panic(err)
			}

			datas = []models.Tick{}
		}

	}
}
