package cvm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tentameneu/cvm-go/internal/config"
	"github.com/tentameneu/cvm-go/internal/stream"
)

func newTestStreamRunner(conf *config.Config) *CVMRunner {
	streamGenerator, _ := stream.NewStreamGenerator(conf)
	return NewCVMRunner(streamGenerator.Generate(), conf.GetBufferSize())
}

func TestRun(t *testing.T) {
	t.Run("SmallerBuffer", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"streamType": "incremental",
			"total":      1_000_000,
			"distinct":   10_000,
			"bufferSize": 1_000,
		})
		runner := newTestStreamRunner(conf)
		n := runner.Run()
		assert.InDelta(t, 10_000, n, 1_000)
	})

	t.Run("ExactBuffer", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"streamType": "incremental",
			"total":      1_000_000,
			"distinct":   10_000,
			"bufferSize": 10_000,
		})
		runner := newTestStreamRunner(conf)
		n := runner.Run()
		assert.Exactly(t, 10_000, n)
	})
}
