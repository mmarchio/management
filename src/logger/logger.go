package logger

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
)

type LoggerKeyT int64

type LoggerFuncT func(string)error

var LoggerKey LoggerKeyT = 4

type LogConfig struct {
	File string
	LineNumber int
	DateTime time.Time
	Msg string
}

type LoggingContext struct {
	echo.Context
}

func (c LoggingContext) Flogger(msg string, vars ...any) error {
	return Logger(fmt.Sprintf(msg, vars...))
}

func (c LoggingContext) GetEchoCtx() context.Context {
	ctxInterface := c.Get("context")
	if ctx, ok := ctxInterface.(context.Context); ok {
		return ctx
	}
	fmt.Printf("GetEchoCtx not working\n")
	return nil
}



func Logger(msg string) error {
	cnf := LogConfig{Msg: msg}
	filename := "./log.txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, file, line, _ := runtime.Caller(2)
	cnf.File = file
	cnf.LineNumber = line
	log := fmt.Sprintf("%s:%d %s: %s", cnf.File, cnf.LineNumber, cnf.DateTime.Format(time.RFC3339), cnf.Msg)
	if _, err := f.WriteString(log); err != nil {
		return err
	}
	return nil
}