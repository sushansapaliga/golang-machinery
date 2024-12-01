package workers

import (
	"log"

	types "example.com/internal/machinery"
	"example.com/internal/machinery/tasks"
)

func StartNormalWorker(taskServer tasks.TaskServer, name string) {
	workerName := "worker-normal-priority-" + name
	worker := taskServer.NewCustomQueueWorker(workerName, 1, types.QUEUE_NORMAL_PRIORITY)

	log.Printf("Launching normal priority worker: " + workerName)
	if err := worker.Launch(); err != nil {
		log.Fatalf("Could not launch worker: %v", err)
	}
}

/**
func StartLowPriorityWorker(taskServer tasks.TaskServer, name string) {
	worker := taskServer.NewCustomQueueWorker("worker-low-priority-" + name, 1, types.QUEUE_LOW_PRIORITY)

	log.Printf("Launching low priority worker...")
	if err := worker.Launch(); err != nil {
		log.Fatalf("Could not launch worker: %v", err)
	}
}

func StartHighPriorityWorker(taskServer tasks.TaskServer, name string) {
	worker := taskServer.NewCustomQueueWorker("worker-high-priority-" + name, 1, types.QUEUE_HIGH_PRIORITY)

	log.Printf("Launching high priority worker...")
	if err := worker.Launch(); err != nil {
		log.Fatalf("Could not launch worker: %v", err)
	}
}
*/
