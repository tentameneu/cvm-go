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

func TestStreamGenerate(t *testing.T) {
	t.Run("Repeating", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"genType":    "repeating",
			"total":      10,
			"distinct":   5,
			"bufferSize": 10,
		})
		generator, err := NewStreamGenerator(conf)
		assert.NotNil(t, generator)
		assert.Nil(t, err)

		stream := generator.Generate()
		assert.Equal(t, 10, len(stream))
		assert.Equal(t, 0, stream[0])
		assert.Equal(t, 1, stream[1])
		assert.Equal(t, 2, stream[2])
		assert.Equal(t, 3, stream[3])
		assert.Equal(t, 4, stream[4])
		assert.Equal(t, 0, stream[5])
		assert.Equal(t, 1, stream[6])
		assert.Equal(t, 2, stream[7])
		assert.Equal(t, 3, stream[8])
		assert.Equal(t, 4, stream[9])
	})
}