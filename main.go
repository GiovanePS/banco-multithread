package main

import (
	"github.com/GiovanePS/banco-multithread/server"
)

func main() {
	s := server.CreateServerThread(10)
	s.StartServerThread()
}
