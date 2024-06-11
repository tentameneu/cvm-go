package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("ValidIncremental", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "incremental",
			"total":      100,
			"distinct":   50,
			"bufferSize": 10,
		})

		assert.Nil(t, err)
		assert.Equal(t, "incremental", conf.genType)
		assert.Equal(t, 100, conf.total)
		assert.Equal(t, 50, conf.distinct)
		assert.Equal(t, 10, conf.bufferSize)
	})

	t.Run("ValidRandom", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  0,
			"randomMax":  1_000_000,
			"bufferSize": 10,
		})

		assert.Nil(t, err)
		assert.Equal(t, "random", conf.genType)
		assert.Equal(t, 100, conf.total)
		assert.Equal(t, 50, conf.distinct)
		assert.Equal(t, 0, conf.randomMin)
		assert.Equal(t, 1_000_000, conf.randomMax)
		assert.Equal(t, 10, conf.bufferSize)
	})

	t.Run("InvalidGenType", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    12345,
			"total":      100,
			"distinct":   50,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'gen-type': must be a string")
	})

	t.Run("InvalidTotal", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "incremental",
			"total":      "100",
			"distinct":   50,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'total': must be an integer")
	})

	t.Run("InvalidDistinct", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "incremental",
			"total":      100,
			"distinct":   "50",
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'distinct': must be an integer")
	})

	t.Run("InvalidBufferSize", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "incremental",
			"total":      100,
			"distinct":   50,
			"bufferSize": "10",
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'buffer-size': must be an integer")
	})

	t.Run("NegativeTotal", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "incremental",
			"total":      -100,
			"distinct":   50,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'total': must be a positive integer")
	})

	t.Run("NegativeDistinct", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "incremental",
			"total":      100,
			"distinct":   -50,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'distinct': must be a positive integer")
	})

	t.Run("NegativeBufferSize", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "incremental",
			"total":      100,
			"distinct":   50,
			"bufferSize": -10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'buffer-size': must be a positive integer")
	})

	t.Run("Total<Distinct", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "incremental",
			"total":      100,
			"distinct":   500,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'total < distinct': total number of elements can't be smaller than distinct number of elements")
	})

	t.Run("InvalidRandomMin", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  "100",
			"randomMax":  1_000_000,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'random-min': must be an integer")
	})

	t.Run("InvalidRandomMax", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  100,
			"randomMax":  "1_000_000",
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'random-max': must be an integer")
	})

	t.Run("NegativeRandomMin", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  -100,
			"randomMax":  1_000_000,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'random-min': must be a positive integer or 0")
	})

	t.Run("NegativeRandomMax", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  100,
			"randomMax":  -1_000_000,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'random-max': must be a positive integer")
	})

	t.Run("RandomMax<RandomMin", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  1_000_000,
			"randomMax":  100,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'random-max < random-min': random-max can't be smaller than random-min")
	})

	t.Run("RandomMax+RandomMin<Distinct", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "random",
			"total":      1_000,
			"distinct":   500,
			"randomMin":  10,
			"randomMax":  100,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'random-max - random-min < distinct': (random-max - random-min) can't be smaller than distinct number of elements")
	})
}

func TestConfigGetters(t *testing.T) {
	conf, err := NewConfig(map[string]any{
		"genType":    "incremental",
		"total":      100,
		"distinct":   50,
		"bufferSize": 10,
	})

	assert.Nil(t, err)
	assert.Equal(t, "incremental", conf.GetGenType())
	assert.Equal(t, 100, conf.GetTotal())
	assert.Equal(t, 50, conf.GetDistinct())
	assert.Equal(t, 10, conf.GetBufferSize())
}
