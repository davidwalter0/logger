package logger

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Logger *log.Logger

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Plain   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	plainHandle io.Writer) {

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

	Plain = log.New(plainHandle, "", 0)
}

// func (log *log.Logger) Log(args interface{}) string {
// 	return log.Println(args)
// }

func Args2String(args []interface{}) string {
	return fmt.Sprintf("%v", args...)
}

func init() {
	flag.Parse()
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr, os.Stdout)
}

// func init() {
// 	var Info *log.Logger = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
// 	var Error *log.Logger = log.New(os.Stdout, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
// }
