package scheduler

import "go_project/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s QueuedScheduler) ConfigureMasterWorkerChan()  {
	
}