package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type FormatInterface interface {
	CommonInterface
	Format(category string, uuid string, message string, v ...interface{}) *CategoryFormat // @TODO change to interface
}

type CommonInterface interface {
	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)
	Fatal(message string)
}

type PrepareInterface interface {
	Debug()
	Info()
	Warning()
	Error()
	Fatal()
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

type CategoryFormat struct {
	logger  *CategoryLogger
	message string
}

func (cf *CategoryFormat) Debug() {
	cf.logger.Debug(cf.message)
}

func (cf *CategoryFormat) Info() {
	cf.logger.Info(cf.message)
}

func (cf *CategoryFormat) Warning() {
	cf.logger.Warning(cf.message)
}

func (cf *CategoryFormat) Error() {
	cf.logger.Error(cf.message)
}

func (cf *CategoryFormat) Fatal() {
	cf.logger.Fatal(cf.message)
}

type CategoryLogger struct {
	stream       StreamInterface
	rootCategory string
}

func (cl *CategoryLogger) Format(category string, uuid string, message string, v ...interface{}) *CategoryFormat {
	msg := fmt.Sprintf(message, v...)
	t := time.Now()
	return &CategoryFormat{
		logger:  cl,
		message: fmt.Sprintf("[%s][%s][%s/%s] %s", t.Format("2006-01-02 15:04:05"), uuid, cl.rootCategory, category, msg),
	}
}

func (cl *CategoryLogger) Debug(message string) {
	cl.stream.debugStream().Println(message)
}

func (cl *CategoryLogger) Info(message string) {
	cl.stream.infoStream().Println(message)
}

func (cl *CategoryLogger) Warning(message string) {
	cl.stream.warningStream().Println(message)
}

func (cl *CategoryLogger) Error(message string) {
	cl.stream.errorStream().Println(message)
}

func (cl *CategoryLogger) Fatal(message string) {
	cl.stream.fatalStream().Println(message)
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

func NewCategoryLogger(rootCategory string, stream StreamInterface) *CategoryLogger {
	l := &CategoryLogger{
		stream:       stream,
		rootCategory: rootCategory,
	}
	return l
}
