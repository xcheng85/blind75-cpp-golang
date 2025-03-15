package stable

type ILogger interface {
	Info() error
	Error() error
	Warn() error
}

type ILoggerImpl interface {
	// variadic func
	Info(...string) error
	Error(...string) error
	Warn(...string) error
}
