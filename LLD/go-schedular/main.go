package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

// task represents one unit of task
// schedular doesnot care about the, it just calls Execute()
type Task interface {
	// Execute run the task
	// it receive context for cancellation
	// it returns:
	// - task ID
	// - execution duration
	// - error (if any)
	Execute(ctx context.Context) (id string, duration time.Duration, err error)
}

// schedular defines what schedular should do
type Schedular interface {
	AddTasks(task Task) error
	StartExecution(ctx context.Context) error
	GetResults() ([]Result, error)
}

// result store outcome of the execution
type Result struct {
	ID       string
	Duration time.Duration
	Err      error
}

type SimpleSchedular struct {
	tasks      []Task      // all tasks are added before execution
	taskChan   chan Task   // channel used to distribute tasks to workers
	resultChan chan Result // channel used by workers to send back results
	mu         sync.Mutex  // mutext to protect shared data
	results    []Result    // store executed results
	started    bool        // prevent starting execution multiple times
}

// NewSchedular creates and returns schedulater instance
func NewSchedular() Schedular {
	return &SimpleSchedular{}
}

func (s *SimpleSchedular) AddTasks(task Task) error {
	if task == nil {
		return errors.New("task cannot be nil")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.started {
		return errors.New("cannot add task if execution started")
	}

	s.tasks = append(s.tasks, task)
	return nil
}

func (s *SimpleSchedular) StartExecution(ctx context.Context) error {
	s.mu.Lock()
	if s.started {
		s.mu.Unlock()
		return errors.New("execution already started")
	}

	s.started = true
	s.mu.Unlock()

	// create channels
	s.taskChan = make(chan Task)
	s.resultChan = make(chan Result)

	// number of workers can be configured
	workerCount := 3

	// start worker goroutine
	for i := 0; i < workerCount; i++ {
		go worker(ctx, s.taskChan, s.resultChan)
	}

	return nil
}

func worker(ctx context.Context, taskCh <-chan Task, resultCh chan<- Result) {
	for {
		select {
		case task, ok := <-taskCh:
			// if channel is closed no more channel
			if !ok {
				return
			}

			// execute the task
			id, duration, err := task.Execute(ctx)

			// send result to the result Chan
			resultCh <- Result{
				ID:       id,
				Duration: duration,
				Err:      err,
			}
		case <-ctx.Done():
			// stop worker if context is cancelled
			return
		}
	}
}

func (s *SimpleSchedular) GetResults() ([]Result, error) {
	return s.results, nil
}

func main() {

}
