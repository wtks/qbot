package logging

import "log"

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type NullLogger struct{}

func (n NullLogger) Debugf(format string, args ...interface{}) {
	return
}

func (n NullLogger) Infof(format string, args ...interface{}) {
	return
}

func (n NullLogger) Errorf(format string, args ...interface{}) {
	return
}

func (n NullLogger) Fatalf(format string, args ...interface{}) {
	return
}

func (n NullLogger) Debug(args ...interface{}) {
	return
}

func (n NullLogger) Info(args ...interface{}) {
	return
}

func (n NullLogger) Error(args ...interface{}) {
	return
}

func (n NullLogger) Fatal(args ...interface{}) {
	return
}

type StdLogger struct{}

func (s StdLogger) Debugf(format string, args ...interface{}) {
	return
}

func (s StdLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (s StdLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (s StdLogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func (s StdLogger) Debug(args ...interface{}) {
	return
}

func (s StdLogger) Info(args ...interface{}) {
	log.Print(args...)
}

func (s StdLogger) Error(args ...interface{}) {
	log.Print(args...)
}

func (s StdLogger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}
