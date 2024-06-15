package conf

var Config = config{
	Address: "localhost:9871",
	Header:  "FollowMe",
}

type config struct {
	Address string `json:"address"`
	Header  string `json:"header"`
}
