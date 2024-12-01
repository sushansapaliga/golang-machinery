package tasks

import (
	"log"
	"time"
)

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	time.Sleep(90 * time.Second)
	log.Printf("Calculated the sum: %v", sum)
	return sum, nil
}
