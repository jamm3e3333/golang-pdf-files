package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	l *zerolog.Logger
}

func NewLogger() *Logger {
	loc, _ := time.LoadLocation("Europe/Prague")

	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFieldName = "@timestamp"
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(loc)
	}
	lg := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{
		l: &lg,
	}
}

func (l *Logger) Info(msg string) {
	l.l.Info().Msg(msg)
}

func (l *Logger) Error(err error) {
	l.l.Error().Err(err)
}

func (l *Logger) Debug(msg string) {
	l.l.Debug().Msg(msg)
}

func (l *Logger) Warn(msg string) {
	l.l.Warn().Msg(msg)
}
