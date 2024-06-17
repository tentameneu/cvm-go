package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("ValidIncremental", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   50,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.Nil(t, err)
		assert.Equal(t, "incremental", Conf.streamType)
		assert.Equal(t, 100, Conf.total)
		assert.Equal(t, 50, Conf.distinct)
		assert.Equal(t, 10, Conf.bufferSize)
		assert.Equal(t, "info", Conf.logLevel)
	})

	t.Run("ValidRandom", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  0,
			"randomMax":  1_000_000,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.Nil(t, err)
		assert.Equal(t, "random", Conf.streamType)
		assert.Equal(t, 100, Conf.total)
		assert.Equal(t, 50, Conf.distinct)
		assert.Equal(t, 0, Conf.randomMin)
		assert.Equal(t, 1_000_000, Conf.randomMax)
		assert.Equal(t, 10, Conf.bufferSize)
	})

	t.Run("InvalidStreamType", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": 12345,
			"total":      100,
			"distinct":   50,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'stream-type': must be a string")
	})

	t.Run("InvalidTotal", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      "100",
			"distinct":   50,
			"bufferSize": 10,
			"logLevel":   "info",
		})

		assert.EqualError(t, err, "invalid parameter 'total': must be an integer")
	})

	t.Run("InvalidDistinct", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   "50",
			"bufferSize": 10,
			"logLevel":   "info",
		})

		assert.EqualError(t, err, "invalid parameter 'distinct': must be an integer")
	})

	t.Run("InvalidBufferSize", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   50,
			"bufferSize": "10",
			"logLevel":   "info",
		})

		assert.EqualError(t, err, "invalid parameter 'buffer-size': must be an integer")
	})

	t.Run("NegativeTotal", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      -100,
			"distinct":   50,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'total': must be a positive integer")
	})

	t.Run("NegativeDistinct", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   -50,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'distinct': must be a positive integer")
	})

	t.Run("NegativeBufferSize", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   50,
			"bufferSize": -10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'buffer-size': must be a positive integer")
	})

	t.Run("Total<Distinct", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   500,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'total < distinct': total number of elements can't be smaller than distinct number of elements")
	})

	t.Run("InvalidRandomMin", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  "100",
			"randomMax":  1_000_000,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'random-min': must be an integer")
	})

	t.Run("InvalidRandomMax", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  100,
			"randomMax":  "1_000_000",
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'random-max': must be an integer")
	})

	t.Run("NegativeRandomMin", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  -100,
			"randomMax":  1_000_000,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'random-min': must be a positive integer or 0")
	})

	t.Run("NegativeRandomMax", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  100,
			"randomMax":  -1_000_000,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'random-max': must be a positive integer")
	})

	t.Run("RandomMax<RandomMin", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "random",
			"total":      100,
			"distinct":   50,
			"randomMin":  1_000_000,
			"randomMax":  100,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'random-max < random-min': random-max can't be smaller than random-min")
	})

	t.Run("RandomMax+RandomMin<Distinct", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "random",
			"total":      1_000,
			"distinct":   500,
			"randomMin":  10,
			"randomMax":  100,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'random-max - random-min < distinct': (random-max - random-min) can't be smaller than distinct number of elements")
	})

	t.Run("InvalidLogLevel", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   50,
			"bufferSize": 10,
			"logLevel":   12345,
			"filePath":   "",
		})

		assert.EqualError(t, err, "invalid parameter 'log-level': must be a string")
	})

	t.Run("InvalidFilePath", func(t *testing.T) {
		err := SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   50,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   123,
		})

		assert.EqualError(t, err, "invalid parameter 'file-path': must be a string")
	})
}

func TestConfigGetters(t *testing.T) {
	err := SetConfig(map[string]any{
		"streamType": "random",
		"total":      100,
		"distinct":   50,
		"randomMin":  100,
		"randomMax":  1_000_000,
		"bufferSize": 10,
		"logLevel":   "info",
		"filePath":   "./stream",
	})

	assert.Nil(t, err)
	assert.Equal(t, "random", StreamType())
	assert.Equal(t, 100, Total())
	assert.Equal(t, 50, Distinct())
	assert.Equal(t, 100, RandomMin())
	assert.Equal(t, 1_000_000, RandomMax())
	assert.Equal(t, 10, BufferSize())
	assert.Equal(t, "info", LogLevel())
	assert.Equal(t, "./stream", FilePath())
}
