package log

type (
	Logger interface {
		Debug(msg string, keyvals ...interface{})
		Info(msg string, keyvals ...interface{})
		Error(msg string, keyvals ...interface{})
		With(keyvals ...interface{}) Logger
	}

	NoOp struct{}
)

func (l NoOp) Debug(msg string, keyvals ...interface{}) {
}

func (l NoOp) Info(msg string, keyvals ...interface{}) {
}

func (l NoOp) Error(msg string, keyvals ...interface{}) {
}

func (l NoOp) With(keyvals ...interface{}) Logger {
	return l
}
