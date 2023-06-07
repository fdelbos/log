package zerolog

import (
	"os"

	"github.com/fdelbos/log"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type ZeroLogger struct {
	Logger zerolog.Logger
}

func (zl ZeroLogger) Debug(msg string, keyvals ...interface{}) {
	zl.Logger.Debug().Fields(keyvals).Msg(msg)
}

func (zl ZeroLogger) Info(msg string, keyvals ...interface{}) {
	zl.Logger.Info().Fields(keyvals).Msg(msg)
}

func (zl ZeroLogger) Error(msg string, keyvals ...interface{}) {
	zl.Logger.Error().Fields(keyvals).Msg(msg)
}

func (zl ZeroLogger) With(keyvals ...interface{}) log.Logger {
	return ZeroLogger{zl.Logger.With().Fields(keyvals).Logger()}
}

func NewLogger() log.Logger {
	lg := zlog.Logger
	lg = lg.With().
		Timestamp().
		CallerWithSkipFrameCount(3).
		Logger()

	lg = lg.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	return ZeroLogger{Logger: lg}
}
