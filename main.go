package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/GiovanePS/banco-multithread/server"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: go run [numWorkers] [numClients] [maxRequests] [serviceTimeInSeconds]")
		os.Exit(1)
	}
	numWorkers, _ := strconv.Atoi(os.Args[1])
	numClients, _ := strconv.Atoi(os.Args[2])
	maxRequests, _ := strconv.Atoi(os.Args[3])
	serviceTime, _ := strconv.Atoi(os.Args[4])
	s := server.CreateServerThread(numWorkers, numClients, maxRequests, serviceTime)
	s.StartServerThread()
}
