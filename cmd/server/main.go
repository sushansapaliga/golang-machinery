package main

import (
	"log"
	"net/http"

	taskServer "example.com/internal/machinery/tasks"
	taskTypes "example.com/internal/machinery"
	"github.com/RichardKnop/machinery/v2/tasks"
)

func main() {
	server := taskServer.GetMachineryServer()

	http.HandleFunc("/schedule-task-high-priority", func(w http.ResponseWriter, r *http.Request) {
		task := &tasks.Signature{
			Name: taskTypes.TASK_ADDING_MULTIPLE_NUMBER,
			Args: []tasks.Arg{
				{Type: "int64", Value: 5},
				{Type: "int64", Value: 6},
			},
			RoutingKey: taskTypes.TASK_PRIORITY_HIGH, // Send the task to high-priority queue
		}

		_, err := server.EnqueueTask(task)
		if err != nil {
			log.Printf("Error sending task: %v", err)
		}
	})

	http.HandleFunc("/schedule-task-low-priority", func(w http.ResponseWriter, r *http.Request) {
		task := &tasks.Signature{
			Name: taskTypes.TASK_ADDING_MULTIPLE_NUMBER,
			Args: []tasks.Arg{
				{Type: "int64", Value: 5},
				{Type: "int64", Value: 6},
			},
			RoutingKey: taskTypes.TASK_PRIORITY_LOW, // Send the task to low-priority queue
		}

		_, err := server.EnqueueTask(task)
		if err != nil {
			log.Printf("Error sending task: %v", err)
		}
	})

	http.HandleFunc("/schedule-task", func(w http.ResponseWriter, r *http.Request) {
		task := &tasks.Signature{
			Name: taskTypes.TASK_ADDING_MULTIPLE_NUMBER,
			Args: []tasks.Arg{
				{Type: "int64", Value: 5},
				{Type: "int64", Value: 6},
			},
			RoutingKey: taskTypes.TASK_PRIORITY_NORMAL, // Send the task to normal-priority queue
		}

		_, err := server.EnqueueTask(task)
		if err != nil {
			log.Printf("Error sending task: %v", err)
		}
	})

	log.Printf("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
