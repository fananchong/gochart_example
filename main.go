package main

import (
	"flag"
	"github.com/fananchong/gochart"
	"log"
)

const start = `version: 1.0
http://localhost:8000`

const (
	DEFAULT_REFRESH_TIME = 5
	DEFAULT_SAMPLE_NUM   = int(3600 / DEFAULT_REFRESH_TIME)
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	println(start)

	s := &gochart.ChartServer{}
	s.AddChart("cpu", NewChartCPU(), true)
	s.AddChart("mem", NewChartMemory(), true)
	s.AddChart("net", NewChartNetwork(), true)

	println(s.ListenAndServe(":8000").Error())
}
