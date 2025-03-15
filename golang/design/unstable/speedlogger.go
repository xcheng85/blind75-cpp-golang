package unstable

import (
	"fmt"

	"github.com/xcheng85/blind75-cpp-golang/golang/design/stable"
)

type SpeedLoggerConsoleImp struct {
}

func (l *SpeedLoggerConsoleImp) Info(s ...string) error {
	fmt.Println("[INFO]", s)
	return nil
}

func (l *SpeedLoggerConsoleImp) Warn(s ...string) error {
	fmt.Println("[WARN]", s)
	return nil
}

func (l *SpeedLoggerConsoleImp) Error(s ...string) error {
	fmt.Println("[ERROR]", s)
	return nil
}

type SpeedLogger struct {
	payload string
	pimpl   stable.ILoggerImpl
}

func (l *SpeedLogger) Info() error {
	return l.pimpl.Info(l.payload)
}

func (l *SpeedLogger) Error() error {
	return l.pimpl.Error(l.payload)
}

func (l *SpeedLogger) Warn() error {
	return l.pimpl.Warn(l.payload)
}
