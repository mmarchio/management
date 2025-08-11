package logger

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/mmarchio/management/database"
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
	Ctx context.Context
	Err error
}

func (c *LoggingContext) Init() {
	ctx := context.Background()
	ctx = database.GetPQContext(ctx)
	c.Ctx = ctx
}

func (c LoggingContext) Flogger(msg string, vars ...any) LoggingContext {
	c.Err = Logger(fmt.Sprintf(msg, vars...))
	return c
}

func (c LoggingContext) GetEchoCtx() context.Context {
	if c.Ctx == nil {
		c.Init()
	}
	return c.Ctx
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
	log := fmt.Sprintf("%s:%d %s: %s\n", cnf.File, cnf.LineNumber, time.Now().Format(time.RFC3339Nano), cnf.Msg)
	if _, err := f.WriteString(log); err != nil {
		return err
	}
	return nil
}