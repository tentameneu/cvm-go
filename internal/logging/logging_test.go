package logging

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tentameneu/cvm-go/internal/config"
)

func TestInitializeLogging(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      10,
			"distinct":   10,
			"bufferSize": 10,
			"logLevel":   "info",
			"filePath":   "",
		})
		err := InitializeLogger(os.Stdout)
		assert.Nil(t, err)
		assert.True(t, logger.Handler().Enabled(context.Background(), slog.LevelInfo))
		assert.False(t, logger.Handler().Enabled(context.Background(), slog.LevelDebug))
		assert.False(t, logger.Handler().Enabled(context.Background(), LevelDeep))
	})

	t.Run("Debug", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      10,
			"distinct":   10,
			"bufferSize": 10,
			"logLevel":   "debug",
			"filePath":   "",
		})
		err := InitializeLogger(os.Stdout)
		assert.Nil(t, err)
		assert.True(t, logger.Handler().Enabled(context.Background(), slog.LevelInfo))
		assert.True(t, logger.Handler().Enabled(context.Background(), slog.LevelDebug))
		assert.False(t, logger.Handler().Enabled(context.Background(), LevelDeep))
	})

	t.Run("Deep", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      10,
			"distinct":   10,
			"bufferSize": 10,
			"logLevel":   "deep",
			"filePath":   "",
		})
		err := InitializeLogger(os.Stdout)
		assert.Nil(t, err)
		assert.True(t, logger.Handler().Enabled(context.Background(), slog.LevelInfo))
		assert.True(t, logger.Handler().Enabled(context.Background(), slog.LevelDebug))
		assert.True(t, logger.Handler().Enabled(context.Background(), LevelDeep))
	})

	t.Run("InvalidLevel", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      10,
			"distinct":   10,
			"bufferSize": 10,
			"logLevel":   "unknown",
			"filePath":   "",
		})
		err := InitializeLogger(os.Stdout)
		assert.EqualError(t, err, "invalid logging level 'unknown'")
	})

	t.Run("DefaultLevelOnNilConfig", func(t *testing.T) {
		config.Conf = nil
		err := InitializeLogger(os.Stdout)
		assert.Nil(t, err)
		assert.True(t, logger.Handler().Enabled(context.Background(), slog.LevelInfo))
		assert.False(t, logger.Handler().Enabled(context.Background(), slog.LevelDebug))
	})
}
