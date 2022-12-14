package log

import (
	"log"
)

type Logger interface {
	WithPrefix(prefix string) Logger
	DebugF(v ...any)
	ErrorF(v ...any)
}

type logger struct {
	log.Logger
}

func New(base *log.Logger) Logger {
	l := &logger{
		Logger: *base,
	}
	return l
}

func (l *logger) clone() *logger {
	l2 := &logger{
		Logger: l.Logger,
	}
	return l2
}

func (l *logger) WithPrefix(prefix string) Logger {
	l2 := l.clone()
	l2.Logger.SetPrefix(l.Prefix() + prefix)
	return l
}

func (l *logger) DebugF(v ...any) {
	l.Println(v...)
}

func (l *logger) ErrorF(v ...any) {
	l.ErrorF(v...)
}
