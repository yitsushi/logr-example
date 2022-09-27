package main

import (
	"github.com/bombsimon/logrusr/v3"
	"github.com/go-logr/logr"
	"github.com/sirupsen/logrus"
)

type Context struct {
	log logr.Logger
}

func NewContext(logger logr.Logger) Context {
	return Context{
		log: logger,
	}
}

func (c *Context) Logger() logr.Logger {
	return c.log
}

func main() {
	fancy()

	withLogrus()
}

func fancy() {
	ctx := NewContext(NewFancyLogger())

	log := ctx.Logger().WithName("my realm")
	log.Info("start process")

	processlog := log.WithValues("step", "download")
	processlog.Info("message")
}

func withLogrus() {
	ctx := NewContext(logrusr.New(logrus.New()))

	log := ctx.Logger().WithName("my realm")
	log.Info("start process")

	processlog := log.WithValues("step", "download")
	processlog.Info("message")
}
