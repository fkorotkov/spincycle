// Copyright 2017-2018, Square, Inc.

package grapher

import (
	"github.com/square/spincycle/job"
	"github.com/square/spincycle/proto"
)

// noop is the default node spec for sequence fan-out (source) and fan-in (sink) nodes.
var noopSpec = &NodeSpec{
	Name:     "noop",
	Category: "job",
	NodeType: "noop",
}

// noopJob is a no-op job that does nothing and always returns success. It's used
// as the default for sequence start and end.
type noopJob struct {
	id job.Id
}

func (j *noopJob) Create(jobArgs map[string]interface{}) error {
	return nil
}

func (j *noopJob) Serialize() ([]byte, error) {
	return nil, nil
}

func (j *noopJob) Deserialize(bytes []byte) error {
	return nil
}

func (j *noopJob) Run(jobData map[string]interface{}) (job.Return, error) {
	ret := job.Return{
		Exit:   0,
		Error:  nil,
		Stdout: "",
		Stderr: "",
		State:  proto.STATE_COMPLETE,
	}
	return ret, nil
}

func (j *noopJob) Status() string {
	return "nop"
}

func (j *noopJob) Stop() error {
	return nil
}

func (j *noopJob) Id() job.Id {
	return j.id
}
