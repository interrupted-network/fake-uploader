package log

import (
	"fmt"
	"log"
)

type Logger interface {
	WithPrefix(prefix string) Logger
	Debugf(format string, a ...any)
	Errorf(format string, a ...any)
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

func (l *logger) Debugf(format string, a ...any) {
	l.Println(fmt.Sprintf(format, a...))
}

func (l *logger) Errorf(format string, a ...any) {
	l.Println(fmt.Sprintf(format, a...))
}
