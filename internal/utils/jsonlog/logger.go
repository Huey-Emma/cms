package jsonlog

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

type Level int8

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
)

func (level Level) String() string {
	switch level {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

type Logger struct {
	w        io.Writer
	minLevel Level
	mu       *sync.Mutex
}

func New(w io.Writer, minLevel Level) *Logger {
	mu := new(sync.Mutex)
	return &Logger{w, minLevel, mu}
}

func (l *Logger) print(level Level, message string, info any) (int, error) {
	aux := struct {
		Level   string `json:"level"`
		Time    string `json:"time"`
		Message string `json:"message"`
		Info    any    `json:"info,omitempty"`
		Trace   string `json:"trace,omitempty"`
	}{
		Level:   level.String(),
		Time:    time.Now().UTC().Format(time.RFC3339),
		Message: message,
		Info:    info,
	}

	if level > LevelInfo {
		aux.Trace = string(debug.Stack())
	}

	b, err := json.Marshal(aux)
	if err != nil {
		b = []byte(LevelError.String() + ": could not marshal json: " + err.Error())
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	return l.w.Write(append(b, '\n'))
}

func (l *Logger) PrintInfo(message string, info any) {
	l.print(LevelInfo, message, info)
}

func (l *Logger) PrintError(err error, info any) {
	l.print(LevelError, err.Error(), info)
}

func (l *Logger) PrintFatal(err error, info any) {
	l.print(LevelFatal, err.Error(), info)
	os.Exit(1)
}
