package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

func Init() {
	w := os.Stderr

	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
		}),
	))
}
