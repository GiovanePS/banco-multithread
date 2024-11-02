package main

import (
	"github.com/GiovanePS/banco-multithread/server"
)

const numClientRequests = 5

func main() {
	var s *server.Server
	for range 1 {
		s = server.CreateServerThread(3, 3, 3, 3)
		s.StartServerThread()
	}
}
