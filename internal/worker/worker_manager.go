package worker

import (
	"github.com/nmluci/stellar-payment-lite/internal/config"
	"github.com/nmluci/stellar-payment-lite/internal/repository"
	"github.com/rs/zerolog"
)

type WorkerManager struct {
	Repository repository.Repository

	TaskChannel chan ArcDownTask
	DoneChannel chan TaskDone
}

type NewWorkerManagerParams struct {
	Logger     zerolog.Logger
	Config     *config.WorkerConfig
	Repository repository.Repository
}

func NewWorkerManager(params NewWorkerManagerParams) (wm *WorkerManager) {
	taskChan := make(chan ArcDownTask, 20)
	doneChan := make(chan TaskDone, 20)

	manager := &WorkerManager{
		TaskChannel: taskChan,
		DoneChannel: doneChan,
		Repository:  params.Repository,
	}
	return manager
}

func (wm *WorkerManager) StartWorker(workers int) {
	go wm.Orchestrator()

	for i := 0; i < workers; i++ {
	}
}

func (wm *WorkerManager) StopManager() {
}
