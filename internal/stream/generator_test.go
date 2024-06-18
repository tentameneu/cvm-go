package stream

import (
	"bytes"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tentameneu/cvm-go/internal/config"
	"github.com/tentameneu/cvm-go/internal/logging"
)

func TestNewStreamGenerator(t *testing.T) {
	t.Run("incremental", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      100,
			"distinct":   10,
			"bufferSize": 100,
			"logLevel":   "info",
			"filePath":   "",
		})
		generator, err := NewStreamGenerator()

		assert.NotNil(t, generator)
		assert.IsType(t, &incrementalStreamGenerator{}, generator)
		assert.Nil(t, err)
	})

	t.Run("Random", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "random",
			"total":      100,
			"distinct":   10,
			"randomMin":  100,
			"randomMax":  1000000,
			"bufferSize": 100,
			"logLevel":   "info",
			"filePath":   "",
		})
		generator, err := NewStreamGenerator()

		assert.NotNil(t, generator)
		assert.IsType(t, &randomStreamGenerator{}, generator)
		assert.Nil(t, err)
	})

	t.Run("File", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "file",
			"total":      100,
			"distinct":   10,
			"bufferSize": 100,
			"logLevel":   "info",
			"filePath":   "./stream_file",
		})
		generator, err := NewStreamGenerator()

		assert.NotNil(t, generator)
		assert.IsType(t, &fileStreamGenerator{}, generator)
		assert.Nil(t, err)
	})

	t.Run("Unknown", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "unknown",
			"total":      100,
			"distinct":   10,
			"bufferSize": 100,
			"logLevel":   "info",
			"filePath":   "",
		})
		generator, err := NewStreamGenerator()

		assert.Nil(t, generator)
		assert.Error(t, err, ("unknown generator type 'unknown'"))
	})
}

func TestStreamGenerate(t *testing.T) {
	t.Run("incremental", func(t *testing.T) {
		err := config.SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      10,
			"distinct":   5,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})
		assert.Nil(t, err)
		generator, err := NewStreamGenerator()
		assert.NotNil(t, generator)
		assert.Nil(t, err)

		stream, err := generator.Generate()
		assert.Nil(t, err)
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
		err := config.SetConfig(map[string]any{
			"streamType": "random",
			"total":      10,
			"distinct":   5,
			"randomMin":  10,
			"randomMax":  25,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})
		assert.Nil(t, err)
		generator, err := NewStreamGenerator()
		assert.NotNil(t, generator)
		assert.Nil(t, err)

		stream, err := generator.Generate()
		assert.Nil(t, err)
		assert.Equal(t, 10, len(stream))
		assert.Condition(t, func() (success bool) { return config.RandomMin() <= stream[0] && stream[0] <= config.RandomMax() })
		assert.Condition(t, func() (success bool) { return config.RandomMin() <= stream[1] && stream[1] <= config.RandomMax() })
		assert.Condition(t, func() (success bool) { return config.RandomMin() <= stream[2] && stream[2] <= config.RandomMax() })
		assert.Condition(t, func() (success bool) { return config.RandomMin() <= stream[3] && stream[3] <= config.RandomMax() })
		assert.Condition(t, func() (success bool) { return config.RandomMin() <= stream[4] && stream[4] <= config.RandomMax() })
		assert.Equal(t, stream[0], stream[5])
		assert.Equal(t, stream[1], stream[6])
		assert.Equal(t, stream[2], stream[7])
		assert.Equal(t, stream[3], stream[8])
		assert.Equal(t, stream[4], stream[9])
	})

	t.Run("File", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			dirPath, err := os.Getwd()
			assert.Nil(t, err)
			config.SetConfig(map[string]any{
				"streamType": "file",
				"total":      10,
				"distinct":   5,
				"bufferSize": 10,
				"logLevel":   "info",
				"filePath":   path.Join(dirPath, "test_files", "valid"),
			})
			err = logging.InitializeLogger(os.Stdout)
			assert.Nil(t, err)
			generator, err := NewStreamGenerator()
			assert.NotNil(t, generator)
			assert.Nil(t, err)

			stream, err := generator.Generate()
			assert.Nil(t, err)
			assert.Equal(t, 10, len(stream))
			assert.Equal(t, 5, stream[0])
			assert.Equal(t, 6, stream[1])
			assert.Equal(t, 1, stream[2])
			assert.Equal(t, 3, stream[3])
			assert.Equal(t, 8, stream[4])
			assert.Equal(t, 9, stream[5])
			assert.Equal(t, 4, stream[6])
			assert.Equal(t, 7, stream[7])
			assert.Equal(t, 2, stream[8])
			assert.Equal(t, 5, stream[9])
		})

		t.Run("ContainsNotNumber", func(t *testing.T) {
			dirPath, err := os.Getwd()
			assert.Nil(t, err)
			config.SetConfig(map[string]any{
				"streamType": "file",
				"total":      10,
				"distinct":   5,
				"bufferSize": 10,
				"logLevel":   "info",
				"filePath":   path.Join(dirPath, "test_files", "contains_not_number"),
			})
			err = logging.InitializeLogger(os.Stdout)
			assert.Nil(t, err)
			generator, err := NewStreamGenerator()
			assert.NotNil(t, generator)
			assert.Nil(t, err)

			stream, err := generator.Generate()
			assert.Nil(t, stream)
			assert.EqualError(t, err, "strconv.Atoi: parsing \"test\": invalid syntax")
		})

		t.Run("ContainsDoubleSpacing", func(t *testing.T) {
			dirPath, err := os.Getwd()
			assert.Nil(t, err)
			config.SetConfig(map[string]any{
				"streamType": "file",
				"total":      10,
				"distinct":   5,
				"bufferSize": 10,
				"logLevel":   "info",
				"filePath":   path.Join(dirPath, "test_files", "contains_double_spacing"),
			})
			err = logging.InitializeLogger(os.Stdout)
			assert.Nil(t, err)
			generator, err := NewStreamGenerator()
			assert.NotNil(t, generator)
			assert.Nil(t, err)

			stream, err := generator.Generate()
			assert.Nil(t, stream)
			assert.EqualError(t, err, "strconv.Atoi: parsing \"\": invalid syntax")
		})
	})
}

type streamGenerateBenchmark struct {
	fileName string
	total    int
}

func BenchmarkStreamGenerate(b *testing.B) {
	dirPath, err := os.Getwd()
	assert.Nil(b, err)

	tests := []streamGenerateBenchmark{
		{fileName: "1000_elements", total: 1000},
		{fileName: "100000_elements", total: 100000},
		{fileName: "10000000_elements", total: 10000000},
	}
	for _, test := range tests {
		b.Run(test.fileName, func(b *testing.B) {
			config.SetConfig(map[string]any{
				"streamType": "file",
				"total":      10,
				"distinct":   5,
				"bufferSize": 1000,
				"logLevel":   "info",
				"filePath":   path.Join(dirPath, "test_files", "benchmark", test.fileName),
			})
			err = logging.InitializeLogger(new(bytes.Buffer))
			assert.Nil(b, err)
			generator, err := NewStreamGenerator()
			assert.NotNil(b, generator)
			assert.Nil(b, err)
			b.ResetTimer()

			stream, err := generator.Generate()
			assert.Nil(b, err)
			assert.Equal(b, test.total, len(stream))
		})
	}
}
