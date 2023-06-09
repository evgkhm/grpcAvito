package logging

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}
	return err
}

func (hook writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func GetLogger() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	l.SetReportCaller(true)

	err := os.MkdirAll("./logs", 0o644)
	if err != nil {
		panic(err)
	}
	logFile, err := os.OpenFile("./logs/all.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		l.Printf("Couldn't open logfile. %s", err.Error())
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{logFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	return l
}
