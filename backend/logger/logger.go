package logger

import (
	"log"
	"os"
)

type Logger struct {
	verbosity int
}

type ILogger interface {
	SetVerbosity(level int)
	GetVerbosity() int
	InfoLog(message string)
	ErrorLog(message string)
	DebugLog(message string)
}

func (l *Logger) SetVerbosity(level int) {
	l.verbosity = level
}

func (l *Logger) GetVerbosity() int {
	return 0
}

func (l *Logger) InfoLog(message string) {
	if l.verbosity > 0 {
		log.SetOutput(os.Stdout)
		log.SetPrefix("INFO: ")
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println(message)
	}
}

func (l *Logger) ErrorLog(message string) {
	if l.verbosity >= 1 {
		log.SetOutput(os.Stdout)
		log.SetPrefix("ERROR: ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Println(message)
	}
}

func (l *Logger) DebugLog(message string) {
	if l.verbosity >= 2 {
		log.SetOutput(os.Stdout)
		log.SetPrefix("DEBUG: ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println(message)
	}
}
