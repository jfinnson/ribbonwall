package logging

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"runtime/debug"
	"sync"

	"github.com/ribbonwall/common/metadata"
	log "github.com/sirupsen/logrus"
)

var once sync.Once
var mutex = &sync.Mutex{}
var logger *log.Logger

const ErrorLevel = log.ErrorLevel
const InfoLevel = log.InfoLevel

type Fields = log.Fields
type Logger = log.Logger
type LevelHooks = log.LevelHooks
type TextFormatter = log.TextFormatter

// GetLogger ---
func GetLogger() *log.Logger {
	return logger
}

// NewLogger create new logger.
func NewLogger() *log.Logger {

	once.Do(func() {
		logger = log.New()
	})

	return logger
}

// WithContext attach metadata of context given on the structure payload.
func WithContext(
	ctx context.Context,
) *log.Entry {

	if ctx == nil {
		return log.NewEntry(logger)
	}

	fields := make(log.Fields)

	// Incoming Context
	md, ok := metadata.FromContext(ctx)
	if ok {
		for k, v := range md {
			fields[k] = v
		}
	}

	return logger.WithFields(fields)
}

// WithError creates an entry from the logrus logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *log.Entry {
	return logger.WithField(log.ErrorKey, err)
}

// WithField creates an entry from the logrus logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *log.Entry {
	return logger.WithField(key, value)
}

// Writer creates an io.Writer
func Writer() *io.PipeWriter {
	return logger.Writer()
}

// WithFields creates an entry from the logrus logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields log.Fields) *log.Entry {
	return logger.WithFields(fields)
}

// Debug logs a message at level Debug on the logrus logger.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Print logs a message at level Info on the logrus logger.
func Print(args ...interface{}) {
	logger.Print(args...)
}

// Info logs a message at level Info on the logrus logger.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn logs a message at level Warn on the logrus logger.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warning logs a message at level Warn on the logrus logger.
func Warning(args ...interface{}) {
	logger.Warning(args...)
}

// Error logs a message at level Error on the logrus logger.
func Error(args ...interface{}) {

	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Error(args...)
}

// Panic logs a message at level Panic on the logrus logger.
func Panic(args ...interface{}) {
	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Panic(args...)
}

// Fatal logs a message at level Fatal on the logrus logger.
func Fatal(args ...interface{}) {
	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Fatal(args...)
}

// Debugf logs a message at level Debug on the logrus logger.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Printf logs a message at level Info on the logrus logger.
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Infof logs a message at level Info on the logrus logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warnf logs a message at level Warn on the logrus logger.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the logrus logger.
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

// Errorf logs a message at level Error on the logrus logger.
func Errorf(format string, args ...interface{}) {

	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the logrus logger.
func Panicf(format string, args ...interface{}) {

	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the logrus logger.
func Fatalf(format string, args ...interface{}) {

	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Fatalf(format, args...)
}

// Debugln logs a message at level Debug on the logrus logger.
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Println logs a message at level Info on the logrus logger.
func Println(args ...interface{}) {
	logger.Println(args...)
}

// Infoln logs a message at level Info on the logrus logger.
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Warnln logs a message at level Warn on the logrus logger.
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

// Warningln logs a message at level Warn on the logrus logger.
func Warningln(args ...interface{}) {
	logger.Warningln(args...)
}

// Errorln logs a message at level Error on the logrus logger.
func Errorln(args ...interface{}) {

	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Errorln(args...)
}

// Panicln logs a message at level Panic on the logrus logger.
func Panicln(args ...interface{}) {

	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the logrus logger.
func Fatalln(args ...interface{}) {

	// Attach trace stack to the message
	args = append(args, fmt.Sprintf(" \n%s", string(debug.Stack())))

	logger.Fatalln(args...)
}

// Disable all logs
func Disable() {
	log.SetOutput(ioutil.Discard)
}
