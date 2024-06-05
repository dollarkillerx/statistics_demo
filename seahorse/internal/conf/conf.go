package conf

import (
	"encoding/json"
	"os"
)

type Conf struct {
	Address string  `json:"address"`
	CsvPath string  `json:"csv_path"`
	Symbol  string  `json:"symbol"`
	DB      DB      `json:"db"`
	Account Account `json:"account"`
}

type DB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type Account struct {
	Account string  `json:"account"`
	Balance float64 `json:"balance"` // 资金
	Lever   int     `json:"lever"`   // 杠杆
}

func LoadConf() *Conf {
	file, err := os.ReadFile("config/conf.json")
	if err != nil {
		panic(err)
	}
	var conf Conf
	err = json.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
