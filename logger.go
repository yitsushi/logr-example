package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/go-logr/logr"
)

type FancyLogger struct {
	name      string
	keyValues map[string]interface{}
	writer    *tabwriter.Writer
}

func (_ *FancyLogger) Init(info logr.RuntimeInfo) {}

func (_ FancyLogger) Enabled(level int) bool {
	return true
}

func (l FancyLogger) Info(level int, msg string, kvs ...interface{}) {
	fmt.Fprintf(l.writer, "%s\t%s\t", l.name, msg)

	for k, v := range l.keyValues {
		fmt.Fprintf(l.writer, "%s: %+v  ", k, v)
	}

	for i := 0; i < len(kvs); i += 2 {
		fmt.Fprintf(l.writer, "%s: %+v  ", kvs[i], kvs[i+1])
	}

	fmt.Fprintf(l.writer, "\n")
	l.writer.Flush()
}

func (l FancyLogger) Error(err error, msg string, kvs ...interface{}) {
	kvs = append(kvs, "error", err)

	l.Info(0, msg, kvs...)
}

func (l FancyLogger) WithName(name string) logr.LogSink {
	if l.name != "" {
		name = l.name + "." + name
	}

	return &FancyLogger{
		name:      name,
		keyValues: l.keyValues,
		writer:    l.writer,
	}
}

func (l FancyLogger) WithValues(kvs ...interface{}) logr.LogSink {
	newMap := make(map[string]interface{}, len(l.keyValues)+len(kvs)/2)
	for k, v := range l.keyValues {
		newMap[k] = v
	}

	for i := 0; i < len(kvs); i += 2 {
		newMap[kvs[i].(string)] = kvs[i+1]
	}

	return &FancyLogger{
		name:      l.name,
		keyValues: newMap,
		writer:    l.writer,
	}
}

func NewFancyLogger() logr.Logger {
	sink := &FancyLogger{
		writer: tabwriter.NewWriter(os.Stderr, 40, 8, 2, '\t', 0),
	}

	return logr.New(sink)
}
