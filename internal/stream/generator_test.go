package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tentameneu/cvm-go/internal/config"
)

func TestNewStreamGenerator(t *testing.T) {
	t.Run("Repeating", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"genType":    "repeating",
			"total":      100,
			"distinct":   10,
			"bufferSize": 100,
		})
		generator, err := NewStreamGenerator(conf)

		assert.NotNil(t, generator)
		assert.Nil(t, err)
	})

	t.Run("Unknown", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"genType":    "unknown",
			"total":      100,
			"distinct":   10,
			"bufferSize": 100,
		})
		generator, err := NewStreamGenerator(conf)

		assert.Nil(t, generator)
		assert.Error(t, err, ("unknown generator type 'unknown'"))
	})
}
