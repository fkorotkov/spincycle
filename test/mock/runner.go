// Copyright 2017, Square, Inc.

package mock

import (
	"errors"
	"sync"

	"github.com/square/spincycle/job-runner/runner"
	"github.com/square/spincycle/proto"
)

var (
	ErrRunner = errors.New("forced error in runner")
)

type RunnerFactory struct {
	RunnersToReturn map[string]*Runner // Keyed on job name.
	MakeErr         error
}

func (f *RunnerFactory) Make(job proto.Job, requestId string) (runner.Runner, error) {
	return f.RunnersToReturn[job.Id], f.MakeErr
}

type Runner struct {
	RunReturn    byte
	RunErr       error
	RunFunc      func(jobData map[string]interface{}) byte // can use this instead of RunErr and RunFunc for more involved mocks
	AddedJobData map[string]interface{}                    // Data to add to jobData.
	RunWg        *sync.WaitGroup                           // WaitGroup that gets released from when a runner starts running.
	RunBlock     chan struct{}                             // Channel that runner.Run() will block on, if defined.
	StopErr      error
	StatusResp   string

	stopped bool // if Stop was called
}

func (r *Runner) Run(jobData map[string]interface{}) byte {
	// If RunFunc is defined, use that.
	if r.RunFunc != nil {
		return r.RunFunc(jobData)
	}

	if r.RunWg != nil {
		r.RunWg.Done()
	}
	if r.RunBlock != nil {
		<-r.RunBlock
		if r.stopped {
			return proto.STATE_FAIL
		}
	}
	// Add job data.
	for k, v := range r.AddedJobData {
		jobData[k] = v
	}

	return r.RunReturn
}

func (r *Runner) Stop() error {
	r.stopped = true
	if r.RunBlock != nil {
		close(r.RunBlock)
	}
	return nil
}

func (r *Runner) Status() string {
	if r.RunBlock != nil {
		close(r.RunBlock)
	}
	return r.StatusResp
}
