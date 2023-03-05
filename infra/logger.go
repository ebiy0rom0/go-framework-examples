package infra

import (
	"context"
	"errors"
	"examples/entity"
	"fmt"
	"io"
	"os"
)

var loggerKey = struct{}{}

type logger struct {
	out io.Writer
}

type logDetail struct {
	level entity.LogLevel
	msg   string
}

func InitLogger(w io.Writer) {
	entity.Logger = &logger{out: w}
}

func NewLogContext(ctx context.Context) context.Context {
	l := &logDetail{}
	return context.WithValue(ctx, &loggerKey, l)
}

func getLogDetail(ctx context.Context) (*logDetail, error) {
	v := ctx.Value(&loggerKey)
	ins, ok := v.(*logDetail)
	if !ok {
		return nil, errors.New("logger cannot used")
	}
	return ins, nil
}

func (l *logger) Log(ctx context.Context, msg string, level entity.LogLevel) error {
	detail, err := getLogDetail(ctx)
	if err != nil {
		return err
	}

	if detail.level < level {
		detail.level = level
	}
	if len(detail.msg) == 0 {
		detail.msg = msg
	}
	return nil
}

func (l *logger) Info(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, entity.LogLevelInfo)
}

func (l *logger) Warn(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, entity.LogLevelWarn)
}

func (l *logger) Err(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, entity.LogLevelErr)
}

func (l *logger) Debug(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, entity.LogLevelDebug)
}

func (l *logger) Fatal(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, entity.LogLevelFatal)
}

func (l *logger) Send(ctx context.Context) error {
	detail, err := getLogDetail(ctx)
	if err != nil {
		return err
	}

	if len(detail.msg) > 0 {
		fmt.Printf("[%s]%s\n", detail.level, detail.msg)
	}
	return nil
}

var _ entity.Logging = (*logger)(nil)

func init() {
	InitLogger(os.Stdout)
}
