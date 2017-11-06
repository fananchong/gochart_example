package main

import (
	"fmt"
	"github.com/fananchong/gochart"
)

type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (this *DefaultLogger) Info(args ...interface{}) {
	fmt.Print(args)
}

func (this *DefaultLogger) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (this *DefaultLogger) Infoln(args ...interface{}) {
	fmt.Println(args)
}

func (this *DefaultLogger) Warning(args ...interface{}) {
	fmt.Print(args)
}

func (this *DefaultLogger) Warningf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (this *DefaultLogger) Warningln(args ...interface{}) {
	fmt.Println(args)
}

func (this *DefaultLogger) Error(args ...interface{}) {
	fmt.Print(args)
}

func (this *DefaultLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (this *DefaultLogger) Errorln(args ...interface{}) {
	fmt.Println(args)
}

func (this *DefaultLogger) Fatal(args ...interface{}) {
	fmt.Print(args)
}

func (this *DefaultLogger) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (this *DefaultLogger) Fatalln(args ...interface{}) {
	fmt.Println(args)
}

var (
	glog gochart.ILogger = NewDefaultLogger()
)
