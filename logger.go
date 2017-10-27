package main

import (
	"os"
)

type Logger struct {
	file *os.File
}

func (lg *Logger) Open(path string) {
	var err error
	lg.file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
}

func (lg *Logger) Log(line string) {
	_, err := lg.file.WriteString(line + "\n")
	if err != nil {
	    panic(err)
	}
}

func (lg *Logger) Close() {
	lg.file.Close()
}