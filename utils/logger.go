package utils

import (
	"io"
	"log"
)

var (
	// Trace Log
	Trace *log.Logger
	// Info Log
	Info *log.Logger
	// Warning Log
	Warning *log.Logger
	//Error Log
	Error *log.Logger
)

// https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html
// InitLogger would initiate the Logger functions
func InitLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
