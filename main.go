package main

import (
	"github.com/GiovanePS/banco-multithread/server"
)

func main() {
	var s *server.Server
	for range 1 {
		s = server.CreateServerThread(10)
		s.StartServerThread()
	}
}
