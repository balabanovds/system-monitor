package cmd

import (
	"context"
)

type Result struct {
	Data []byte
	Err  error
}

type (
	InCh  = <-chan Result
	OutCh = InCh
	BiCh  = chan Result
)

type stage func(inCh InCh) (outCh OutCh)

type pipeline struct {
	stages     []stage
	dataStream InCh
}

func newPipeline(inCh InCh, stages ...stage) *pipeline {
	return &pipeline{
		stages:     stages,
		dataStream: inCh,
	}
}

// Exec run stages one by one.
func (p *pipeline) exec(ctx context.Context) OutCh {
	for _, stage := range p.stages {
		stageStream := make(BiCh)

		go func(stageStream BiCh, inStream InCh) {
			defer close(stageStream)

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-inStream:
					if !ok {
						return
					}
					if data.Err != nil {
						return
					}
					if data.Data != nil {
						select {
						case <-ctx.Done():
							return
						case stageStream <- data:
						}
					}
				}
			}
		}(stageStream, p.dataStream)

		p.dataStream = stage(stageStream)
	}

	return p.dataStream
}

func execPipeline(ctx context.Context, in InCh, stages ...stage) OutCh {
	return newPipeline(in, stages...).exec(ctx)
}
