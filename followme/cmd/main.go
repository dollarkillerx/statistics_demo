package main

import "followme/internal/server"

func main() {
	s := server.New()
	if err := s.Run(); err != nil {
		panic(err)
	}
}
