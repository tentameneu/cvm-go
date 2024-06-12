package logging

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/tentameneu/cvm-go/internal/config"
)

var logger *slog.Logger

func InitializeLogger(writer io.Writer, conf *config.Config) error {
	var level slog.Level

	if conf != nil {
		switch conf.GetLogLevel() {
		case "info":
			level = slog.LevelInfo
		case "debug":
			level = slog.LevelDebug
		default:
			return fmt.Errorf("invalid logging level '%s'", conf.GetLogLevel())
		}
	} else {
		level = slog.LevelInfo
	}

	options := &slog.HandlerOptions{
		Level: level,
	}

	handler := newHandler(writer, options)
	logger = slog.New(handler)

	return nil
}

func Logger() *slog.Logger {
	return logger
}
