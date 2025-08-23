package logger

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

const FATAL int = 0
const ERROR int = 1
const WARN int = 2
const DEBUG int = 3
const INFO int = 4

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
	Depth int
	Severity int
	Ctx context.Context
	Err error
}

func (c *LoggingContext) Init() {
	ctx := context.Background()
	if c.Depth == 0 {
		c.Depth = 2
	}
	c.Ctx = ctx
}

func (c LoggingContext) Flogger(msg string, vars ...any) LoggingContext {
	c.Err = c.Logger(fmt.Sprintf(msg, vars...))
	return c
}

// func (c LoggingContext) GetDatabase() *sql.DB {
// 	return database.GetPQDatabase()
// }

func (c LoggingContext) GetEchoCtx() context.Context {
	if c.Ctx == nil {
		c.Init()
	}
	return c.Ctx
}

func (c LoggingContext) Logger(msg string) error {
	cnf := LogConfig{Msg: msg}
	sev := ""
	filename := "./log.txt"
	switch c.Severity {
	case 0:
		sev = "fatal"
		filename = "./error.txt"
	case 1:
		sev = "error"
		filename = "./error.txt"
	case 2:
		sev = "warn"
		filename = "./warn.txt"
	case 3:
		sev = "debug"
		filename = "./debug.txt"
	case 4:
		sev = "info"
		filename = "./info.txt"
	default:
		sev = "info"
		filename = "./info.txt"
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, file, line, _ := runtime.Caller(c.Depth)
	cnf.File = file
	cnf.LineNumber = line
	log := fmt.Sprintf("%s %s:%d %s: %s\n", sev, cnf.File, cnf.LineNumber, time.Now().Format(time.RFC3339Nano), cnf.Msg)
	if _, err := f.WriteString(log); err != nil {
		return err
	}
	return nil
}