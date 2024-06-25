package conf

var Config = config{
	Address: "0.0.0.0:9981",
	Header:  "FollowMe",
}

type config struct {
	Address string `json:"address"`
	Header  string `json:"header"`
}
