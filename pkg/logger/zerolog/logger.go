package zerolog

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	log *zerolog.Logger
}

func NewLogger() *Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s***", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	log := zerolog.New(output).With().Timestamp().Caller().Logger()
	return &Logger{log: &log}
}

func (l *Logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *Logger) Trace(msg string) {
	l.log.Trace().Msg(msg)
}

func (l *Logger) Error(msg string) {
	l.log.Error().Msg(msg)
}

func (l *Logger) Fatal(msg string) {
	l.log.Fatal().Msg(msg)
}
