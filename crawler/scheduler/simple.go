package scheduler

import "go_project/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request)  {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request)  {
	go func() {
		s.workerChan <- r
	}()
}