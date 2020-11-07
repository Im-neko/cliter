package logger

import (
	"log"
)

// Error is export log msg with timestamp
func Error(msg string, v ...interface{}) {
	log.Printf(
		"[Error]"+msg,
		v,
	)
}

// Warn is export log msg with timestamp
func Warn(msg string, v ...interface{}) {
	log.Printf(
		"[Warning]"+msg,
		v,
	)
}

// Debug is export log msg with timestamp
func Debug(msg string, v ...interface{}) {
	log.Printf(
		"[Debug]"+msg,
		v,
	)
}

// Info is export log msg with timestamp
func Info(msg string, v ...interface{}) {
	log.Printf(
		"[Info]"+msg,
		v,
	)
}
