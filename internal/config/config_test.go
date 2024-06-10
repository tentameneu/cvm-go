package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "repeating",
			"total":      100,
			"distinct":   50,
			"bufferSize": 10,
		})

		assert.Nil(t, err)
		assert.Equal(t, "repeating", conf.genType)
		assert.Equal(t, 100, conf.total)
		assert.Equal(t, 50, conf.distinct)
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
		assert.EqualError(t, err, "invalid parameter 'genType': must be a string")
	})

	t.Run("InvalidTotal", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "repeating",
			"total":      "100",
			"distinct":   50,
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'total': must be an integer")
	})

	t.Run("InvalidDistinct", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "repeating",
			"total":      100,
			"distinct":   "50",
			"bufferSize": 10,
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'distinct': must be an integer")
	})

	t.Run("InvalidBufferSize", func(t *testing.T) {
		conf, err := NewConfig(map[string]any{
			"genType":    "repeating",
			"total":      100,
			"distinct":   50,
			"bufferSize": "10",
		})

		assert.Nil(t, conf)
		assert.EqualError(t, err, "invalid parameter 'bufferSize': must be an integer")
	})
}

func TestConfigGetters(t *testing.T) {
	conf, err := NewConfig(map[string]any{
		"genType":    "repeating",
		"total":      100,
		"distinct":   50,
		"bufferSize": 10,
	})

	assert.Nil(t, err)
	assert.Equal(t, "repeating", conf.GetGenType())
	assert.Equal(t, 100, conf.GetTotal())
	assert.Equal(t, 50, conf.GetDistinct())
	assert.Equal(t, 10, conf.GetBufferSize())
}
