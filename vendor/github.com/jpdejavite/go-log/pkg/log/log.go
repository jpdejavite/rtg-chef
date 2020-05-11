package log

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	// DefaultFormat default logging format
	DefaultFormat   = "%s [%s] %s %s %s"
	debugLogLevel   = "DEBUG"
	errorLogLevel   = "ERROR"
	fatalLogLevel   = "FATAL"
	infoLogLevel    = "INFO"
	warningLogLevel = "WARN"
)

// Debug print in log file with debug level
func Debug(tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, debugLogLevel, tag, message, metadata, coi)
}

// Debugf print in log file with custom format and debug level
func Debugf(format string, tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, debugLogLevel, tag, message, metadata, coi)
}

// Error print in log file with error level
func Error(tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, errorLogLevel, tag, message, metadata, coi)
}

// Errorf print in log file with custom format and error level
func Errorf(format string, tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, errorLogLevel, tag, message, metadata, coi)
}

// Fatal print in log file with fatal level
func Fatal(tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, fatalLogLevel, tag, message, metadata, coi)
}

// Fatalf print in log file with custom format and fatal level
func Fatalf(format string, tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, fatalLogLevel, tag, message, metadata, coi)
}

// Info print in log file with info level
func Info(tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, infoLogLevel, tag, message, metadata, coi)
}

// Infof print in log file with custom format and info level
func Infof(format string, tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, infoLogLevel, tag, message, metadata, coi)
}

// Warn print in log file with warning level
func Warn(tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, warningLogLevel, tag, message, metadata, coi)
}

// Warnf print in log file with custom format and warning level
func Warnf(format string, tag string, message string, metadata interface{}, coi string) {
	print(DefaultFormat, warningLogLevel, tag, message, metadata, coi)
}

func print(format string, level string, tag string, message string, metadata interface{}, coi string) {
	meta := "{}"
	if metadata != nil {
		jsonByteArray, _ := json.Marshal(metadata)
		meta = string(jsonByteArray)
	}
	log.Print(fmt.Sprintf(DefaultFormat, level, coi, tag, message, meta))
}
