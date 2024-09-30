package logger

import (
	"free5gclib/logger_conf"
	"free5gclib/logger_util"
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var NasMsgLog *logrus.Entry
var ConvertLog *logrus.Entry
var SecurityLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	free5gcLogHook, err := logger_util.NewFileHook(logger_conf.Free5gcLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(free5gcLogHook)
	}

	selfLogHook, err := logger_util.NewFileHook(logger_conf.LibLogDir+"nas.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	NasMsgLog = log.WithFields(logrus.Fields{"component": "NAS", "category": "Message"})
	ConvertLog = log.WithFields(logrus.Fields{"component": "NAS", "category": "Convert"})
	SecurityLog = log.WithFields(logrus.Fields{"component": "NAS", "category": "Security"})
}
