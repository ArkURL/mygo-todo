package logger

import "go.uber.org/zap"

var log = zap.NewNop()

func Init() error {
	var err error
	log, err = zap.NewDevelopment()
	return err
}

func L() *zap.Logger {
	return log
}

func Sync() {
	_ = log.Sync()
}
