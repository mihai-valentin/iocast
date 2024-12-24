package iocast

import (
	"errors"
)

type pipeline[T any] struct {
	id         string
	head       *task[T]
	resultChan chan Result[T]
}

// NewPipeline links tasks together to execute them in order, returns a pipeline instance.
func NewPipeline[T any](id string, tasks ...*task[T]) (*pipeline[T], error) {
	if len(tasks) < 2 {
		return nil, errors.New("at least two tasks must be linked to create a pipeline")
	}
	head := tasks[0]
	for i, t := range tasks {
		if i < len(tasks)-1 {
			t.link(tasks[i+1])
		}
	}
	return &pipeline[T]{
		id:         id,
		head:       head,
		resultChan: head.resultChan,
	}, nil
}

// Wait awaits for the final result of the pipeline (last task in the order).
func (p *pipeline[T]) Wait() <-chan Result[T] {
	return p.head.resultChan
}

// Exec executes the linked tasks of the pipeline.
func (p *pipeline[T]) Exec() {
	p.head.Exec()
}

func (p *pipeline[T]) Write() error {
	return p.head.Write()
}

// Id is an ID geter.
func (p *pipeline[T]) Id() string {
	return p.id
}
