package util

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-pp/pp"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var globalLogger *logrus.Logger

// GetLogger returns global logger
func GetLogger() *logrus.Logger {

	if globalLogger != nil {
		return globalLogger
	}

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Out = os.Stdout

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logger.Infoln(err)
		level = logrus.DebugLevel
	}
	logger.Level = level
	logger.Infof("Global Log Level is %v", level.String())

	globalLogger = logger

	return logger
}

// InitLogger Loggerについて初期化する
func InitLogger() {
	globalLogger = nil
	envcode := os.Getenv("ENVCODE")
	pp.ColoringEnabled = envcode == "local" || envcode == "test"
}

// DebugLogWriter ログ出力 io.Writer
func DebugLogWriter() *io.PipeWriter {
	return GetLogger().WriterLevel(logrus.DebugLevel)
}

// InfoLogWriter ログ出力 io.Writer
func InfoLogWriter() *io.PipeWriter {
	return GetLogger().WriterLevel(logrus.InfoLevel)
}

// WarnLogWriter ログ出力 io.Writer
func WarnLogWriter() *io.PipeWriter {
	return GetLogger().WriterLevel(logrus.WarnLevel)
}

// ErrorLogWriter ログ出力 io.Writer
func ErrorLogWriter() *io.PipeWriter {
	return GetLogger().WriterLevel(logrus.ErrorLevel)
}

// WriteError エラー情報を出力します。
// `LOG_STACK_TRACE` が設定されているときのみ、追加でスタックトレースを標準出力に出力します。
// それ以外の場合は、エラーメッセージのみを出力します。
func WriteError(err error) {
	logger := GetLogger()
	logger.Error(err)

	if os.Getenv("LOG_STACK_TRACE") != "" {
		logger.Error(fmt.Sprintf("%+v", errors.WithStack(err)))
	}
}

// SentrySimpleHook はlogrusのSentry用のHooks
type SentrySimpleHook struct {
	captureErrorFunc func(err error) error
}

// NewSentrySimpleHook は Sentry用のHooksを生成
func NewSentrySimpleHook(captureErrorFunc func(err error) error) *SentrySimpleHook {
	hook := &SentrySimpleHook{
		captureErrorFunc: captureErrorFunc,
	}
	return hook
}

// Fire is called logrus emitting log
func (hook *SentrySimpleHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		logMessage := makeJSONLogString(info, err.Error())
		log.SetOutput(os.Stdout)
		log.Println(logMessage)
	}

	switch entry.Level {
	case logrus.PanicLevel:
		hook.captureErrorFunc(fmt.Errorf(line))
	case logrus.FatalLevel:
		hook.captureErrorFunc(fmt.Errorf(line))
	case logrus.ErrorLevel:
		hook.captureErrorFunc(fmt.Errorf(line))
	case logrus.WarnLevel:
		hook.captureErrorFunc(fmt.Errorf(line))
	case logrus.InfoLevel, logrus.DebugLevel, logrus.TraceLevel:
	default:
	}
	return nil
}

// Levels returns supoort levels
func (hook *SentrySimpleHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
