package main

import (
	"github.com/fananchong/gochart"
	"log"
)

const start = `version: 1.0
http://localhost:8000`

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	println(start)

	s := &gochart.ChartServer{}
	s.AddChart("area", &SplineChart{filename: "area.chart"})
	s.AddChart("bar", &SplineChart{filename: "bar.chart"})
	s.AddChart("column", &SplineChart{filename: "column.chart"})
	s.AddChart("line", &SplineChart{filename: "line.chart"})
	s.AddChart("pie", &PieChart{filename: "pie.chart"})
	s.AddChart("spline", &SplineChart{filename: "spline.chart"})
	println(s.ListenAndServe(":8000").Error())
}
