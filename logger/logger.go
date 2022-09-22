package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type CategoryLogger interface {
	WithUuid(uuid string)
	AddSubCategory(subCategory string) CategoryLogger
	Debug(message string, v ...interface{})
	Info(message string, v ...interface{})
	Warning(message string, v ...interface{})
	Error(message string, v ...interface{})
	Fatal(message string, v ...interface{})
}

type StreamInterface interface {
	debugStream() *log.Logger
	infoStream() *log.Logger
	warningStream() *log.Logger
	errorStream() *log.Logger
	fatalStream() *log.Logger
}

type StreamLog struct {
	debugLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	fatalLog   *log.Logger
}

func (st *StreamLog) debugStream() *log.Logger {
	return st.debugLog
}

func (st *StreamLog) infoStream() *log.Logger {
	return st.infoLog
}

func (st *StreamLog) warningStream() *log.Logger {
	return st.warningLog
}

func (st *StreamLog) errorStream() *log.Logger {
	return st.errorLog
}

func (st *StreamLog) fatalStream() *log.Logger {
	return st.fatalLog
}

type ImplCategoryLogger struct {
	category string
	uuid     string
	stream   StreamInterface
}

func (l *ImplCategoryLogger) WithUuid(uuid string) {
	l.uuid = uuid
}

func (l *ImplCategoryLogger) AddSubCategory(subCategory string) CategoryLogger {
	c := fmt.Sprintf("%s/%s", l.category, subCategory)
	return NewCategoryLogger(c, l.uuid, l.stream)
}

func (l *ImplCategoryLogger) Debug(message string, v ...interface{}) {
	l.stream.debugStream().Println(l.format(message, v...))
}

func (l *ImplCategoryLogger) Info(message string, v ...interface{}) {
	l.stream.infoStream().Println(l.format(message, v...))
}

func (l *ImplCategoryLogger) Warning(message string, v ...interface{}) {
	l.stream.warningStream().Println(l.format(message, v...))
}

func (l *ImplCategoryLogger) Error(message string, v ...interface{}) {
	l.stream.errorStream().Println(l.format(message, v...))
}

func (l *ImplCategoryLogger) Fatal(message string, v ...interface{}) {
	l.stream.fatalStream().Println(l.format(message, v...))
}

func (l *ImplCategoryLogger) format(message string, v ...interface{}) string {
	m := fmt.Sprintf(message, v...)
	t := time.Now()
	return fmt.Sprintf("[%s][%s][%s] %s", t.Format("2006-01-02 15:04:05"), l.uuid, l.category, m)
}

func NewStreamLog() *StreamLog {
	return &StreamLog{
		debugLog:   log.New(os.Stdout, "[DEBUG]", 0),
		infoLog:    log.New(os.Stdout, "[INFO]", 0),
		warningLog: log.New(os.Stdout, "[WARNING]", 0),
		errorLog:   log.New(os.Stderr, "[ERROR]", 0),
		fatalLog:   log.New(os.Stderr, "[FATAL]", 0),
	}
}

func NewCategoryLogger(category string, uuid string, stream StreamInterface) CategoryLogger {
	l := &ImplCategoryLogger{
		stream:   stream,
		category: category,
		uuid:     uuid,
	}
	return l
}
