package main

import (
	"tradingview/internal"
)

// CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o tradingview cmd/main.go

func main() {
	ser := &internal.Server{}

	if err := ser.Run(); err != nil {
		panic(err)
	}
}
