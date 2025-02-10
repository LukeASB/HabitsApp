package logger

import (
	"log"
	"os"
	"strconv"
)

type Logger struct {
	verbosity int
}

type ILogger interface {
	SetVerbosity(level string) error
	GetVerbosity() int
	InfoLog(functionName, message string)
	ErrorLog(functionName, message string)
	DebugLog(functionName, message string)
}

func (l *Logger) SetVerbosity(level string) error {
	parsedLevel, err := strconv.Atoi(level)

	if err != nil {
		return err
	}

	l.verbosity = parsedLevel

	return nil
}

func (l *Logger) GetVerbosity() int {
	return 0
}

func (l *Logger) InfoLog(functionName, message string) {
	if l.verbosity > 0 {
		log.SetOutput(os.Stdout)
		log.SetPrefix("INFO: ")
		log.SetFlags(log.Ldate | log.Ltime)
		if message == "" {
			log.Println(functionName)
			return
		}
		log.Printf("%s - %s\n", functionName, message)
	}
}

func (l *Logger) ErrorLog(functionName, message string) {
	if l.verbosity >= 1 {
		log.SetOutput(os.Stdout)
		log.SetPrefix("ERROR: ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		if message == "" {
			log.Println(functionName)
			return
		}
		log.Printf("%s - %s\n", functionName, message)
	}
}

func (l *Logger) DebugLog(functionName, message string) {
	if l.verbosity >= 2 {
		log.SetOutput(os.Stdout)
		log.SetPrefix("DEBUG: ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		if message == "" {
			log.Println(functionName)
			return
		}
		log.Printf("%s - %s\n", functionName, message)
	}
}
