// @Time : 2018/6/1 20:10
// @Author : minigeek
package scheduler

import "crawler/engine"

type SimpleSchduler struct {
	workerChan chan engine.Request
}

func (s *SimpleSchduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleSchduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleSchduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleSchduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}
