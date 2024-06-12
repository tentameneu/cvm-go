package logging

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

const LevelDeep = slog.Level(-8)

type slogHandler struct {
	handler slog.Handler
	writer  io.Writer
}

func newHandler(writer io.Writer, opts *slog.HandlerOptions) slogHandler {
	return slogHandler{
		handler: slog.NewTextHandler(writer, opts),
		writer:  writer,
	}
}

func (slogHndl slogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return slogHndl.handler.Enabled(ctx, level)
}

func (slogHndl slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return slogHandler{
		handler: slogHndl.handler.WithAttrs(attrs),
		writer:  slogHndl.writer,
	}
}

func (slogHndl slogHandler) Handle(ctx context.Context, record slog.Record) error {
	formattedTime := record.Time.Format("15:04:05.000")

	level := record.Level.String()
	if record.Level == LevelDeep {
		level = "DEEP"
	}

	strs := []string{formattedTime, "||", level, "||", record.Message}

	if record.NumAttrs() != 0 {
		record.Attrs(func(a slog.Attr) bool {
			strs = append(strs, fmt.Sprintf("%s=%s", a.Key, a.Value.String()))
			return true
		})
	}

	result := strings.Join(strs, " ") + "\n"
	b := []byte(result)

	_, err := slogHndl.writer.Write(b)

	return err
}

func (slogHndl slogHandler) WithGroup(name string) slog.Handler {
	return slogHandler{
		handler: slogHndl.handler.WithGroup(name),
		writer:  slogHndl.writer,
	}
}
