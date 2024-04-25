package logger_util

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type FileHook struct {
	file      *os.File
	flag      int
	chmod     os.FileMode
	formatter *logrus.TextFormatter
}

func NewFileHook(file string, flag int, chmod os.FileMode) (*FileHook, error) {
	plainFormatter := &logrus.TextFormatter{DisableColors: true}
	logFile, err := os.OpenFile(file, flag, chmod)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to write file on filehook %v", err)
		return nil, err
	}

	return &FileHook{logFile, flag, chmod, plainFormatter}, err
}

// Fire event
func (hook *FileHook) Fire(entry *logrus.Entry) error {
	var line string
	if plainformat, err := hook.formatter.Format(entry); err != nil {
		log.Printf("Formatter error: %+v", err)
		return err
	} else {
		line = string(plainformat)
	}
	if _, err := hook.file.WriteString(line); err != nil {
		fmt.Fprintf(os.Stderr, "unable to write file on filehook(entry.String)%v", err)
		return err
	}

	return nil
}

func (hook *FileHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
