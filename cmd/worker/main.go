package main

import (
	"math/rand"
	"time"

	taskServer "example.com/internal/machinery/tasks"
	"example.com/internal/machinery/workers"
)

func main() {
	server := taskServer.GetMachineryServer()

	server.RegisterTasks()

	randLength := 10
	charSet := "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	workerName := StringWithCharset(randLength, charSet)

	workers.StartNormalWorker(server, workerName)

	// workers.StartHighPriorityWorker(server, workerName)
	// workers.StartLowPriorityWorker(server, workerName)
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset)-1)]
	}
	return string(b)
}
