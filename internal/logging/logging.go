package logging

import (
	"io"
	"log/slog"
)

var logger *slog.Logger

func InitializeLogger(writer io.Writer) {
	options := &slog.HandlerOptions{}
	handler := newHandler(writer, options)
	logger = slog.New(handler)
}

func Logger() *slog.Logger {
	return logger
}
