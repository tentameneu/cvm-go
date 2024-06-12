package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tentameneu/cvm-go/internal/config"
)

func TestNewStreamGenerator(t *testing.T) {
	t.Run("incremental", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   10,
			"bufferSize": 100,
			"logLevel":   "info",
		})
		generator, err := NewStreamGenerator(conf)

		assert.NotNil(t, generator)
		assert.IsType(t, &incrementalStreamGenerator{}, generator)
		assert.Nil(t, err)
	})

	t.Run("Random", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"streamType": "random",
			"total":      100,
			"distinct":   10,
			"randomMin":  100,
			"randomMax":  1000000,
			"bufferSize": 100,
			"logLevel":   "info",
		})
		generator, err := NewStreamGenerator(conf)

		assert.NotNil(t, generator)
		assert.IsType(t, &randomStreamGenerator{}, generator)
		assert.Nil(t, err)
	})

	t.Run("Unknown", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"streamType": "unknown",
			"total":      100,
			"distinct":   10,
			"bufferSize": 100,
			"logLevel":   "info",
		})
		generator, err := NewStreamGenerator(conf)

		assert.Nil(t, generator)
		assert.Error(t, err, ("unknown generator type 'unknown'"))
	})
}

func TestStreamGenerate(t *testing.T) {
	t.Run("incremental", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"streamType": "incremental",
			"total":      10,
			"distinct":   5,
			"bufferSize": 10,
			"logLevel":   "info",
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

	t.Run("Random", func(t *testing.T) {
		conf, _ := config.NewConfig(map[string]any{
			"streamType": "random",
			"total":      10,
			"distinct":   5,
			"randomMin":  10,
			"randomMax":  25,
			"bufferSize": 10,
			"logLevel":   "info",
		})
		generator, err := NewStreamGenerator(conf)
		assert.NotNil(t, generator)
		assert.Nil(t, err)

		stream := generator.Generate()
		assert.Equal(t, 10, len(stream))
		assert.Condition(t, func() (success bool) { return conf.GetRandomMin() <= stream[0] && stream[0] <= conf.GetRandomMax() })
		assert.Condition(t, func() (success bool) { return conf.GetRandomMin() <= stream[1] && stream[1] <= conf.GetRandomMax() })
		assert.Condition(t, func() (success bool) { return conf.GetRandomMin() <= stream[2] && stream[2] <= conf.GetRandomMax() })
		assert.Condition(t, func() (success bool) { return conf.GetRandomMin() <= stream[3] && stream[3] <= conf.GetRandomMax() })
		assert.Condition(t, func() (success bool) { return conf.GetRandomMin() <= stream[4] && stream[4] <= conf.GetRandomMax() })
		assert.Equal(t, stream[0], stream[5])
		assert.Equal(t, stream[1], stream[6])
		assert.Equal(t, stream[2], stream[7])
		assert.Equal(t, stream[3], stream[8])
		assert.Equal(t, stream[4], stream[9])
	})
}
