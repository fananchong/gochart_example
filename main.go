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
	DEFAULT_SAMPLE_NUM   = 3600 / 5
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	println(start)

	s := &gochart.ChartServer{}
	s.AddChart("area", &SplineChart{filename: "area.chart"})
	s.AddChart("bar", &SplineChart{filename: "bar.chart"})
	s.AddChart("column", &SplineChart{filename: "column.chart"})
	s.AddChart("line", &SplineChart{filename: "line.chart"})
	s.AddChart("pie", &PieChart{filename: "pie.chart"})
	s.AddChart("spline", &SplineChart{filename: "spline.chart"})
	s.AddChart("time", &TimeChartExample{})
	s.AddChart("cpu", NewExampleCPU())
	s.AddChart("memory", NewExampleMemory())
	s.AddChart("net", NewExampleNetwork())

	println(s.ListenAndServe(":8000").Error())
}
