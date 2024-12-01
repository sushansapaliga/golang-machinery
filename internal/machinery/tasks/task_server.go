package tasks

import (
	types "example.com/internal/machinery"

	"github.com/RichardKnop/machinery/v2"
	redisBackend "github.com/RichardKnop/machinery/v2/backends/redis"
	"github.com/RichardKnop/machinery/v2/backends/result"
	redisBroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerLock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/RichardKnop/machinery/v2/tasks"
)

type TaskServer struct {
	server *machinery.Server
}

// Register tasks for worker
func (taskServer TaskServer) RegisterTasks() {
	taskServer.server.RegisterTask(types.TASK_ADDING_MULTIPLE_NUMBER, Add)
}

// Enqueue task
func (taskServer TaskServer) EnqueueTask(task *tasks.Signature) (*result.AsyncResult, error) {
	return taskServer.server.SendTask(task)
}

// NewCustomQueueWorker creates Worker instance with Custom Queue
func (taskServer TaskServer) NewCustomQueueWorker(consumerTag string, concurrency int, queue string) *machinery.Worker {
	return taskServer.server.NewCustomQueueWorker(consumerTag, concurrency, queue)
}

func GetMachineryServer() TaskServer {
	cnf := &config.Config{
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}
	broker := redisBroker.NewGR(cnf, []string{"0.0.0.0:6379"}, 0)
	backend := redisBackend.NewGR(cnf, []string{"0.0.0.0:6379"}, 0)
	lock := eagerLock.New()

	// Initialize the Machinery server
	taskServer := TaskServer{
		server: machinery.NewServer(cnf, broker, backend, lock),
	}

	return taskServer
}
