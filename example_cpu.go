package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
	"time"
)

type ExampleCPU struct {
	gochart.ChartTime
	cpus     map[int][]int
	lenlimit int
}

func NewExampleCPU() *ExampleCPU {
	lenlimit := DEFAULT_SAMPLE_NUM
	cc, _ := cpu.Percent(0, false)
	inst := &ExampleCPU{cpus: make(map[int][]int), lenlimit: lenlimit}
	for i := 0; i < len(cc); i++ {
		inst.cpus[i] = make([]int, lenlimit)
	}

	inst.RefreshTime = "1"
	inst.ChartType = "line"
	inst.Title = "CPU占用"
	inst.SubTitle = ""
	inst.YAxisText = "cpu"
	inst.YMax = "100"
	inst.ValueSuffix = "%"

	return inst
}

func (this *ExampleCPU) Update() {
	this.updateData()

	endtime := 1000 * int(8*60*60+time.Now().Unix())
	begintime := endtime - 1000*this.lenlimit

	datas := make([]interface{}, 0)
	var json *simplejson.Json
	for i := 0; i < len(this.cpus); i++ {
		json = simplejson.New()
		json.Set("name", "cpu"+strconv.Itoa(i))
		json.Set("data", this.cpus[i])
		json.Set("pointInterval", 1000)
		json.Set("pointStart", begintime)
		json.Set("pointEnd", endtime)
		datas = append(datas, json)
	}
	json = simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	this.DataArray = string(b)
}

func (this *ExampleCPU) updateData() {
	cc, _ := cpu.Percent(0, false)
	for i := 0; i < len(cc); i++ {
		this.cpus[i] = append(this.cpus[i], int(cc[i]))
		if len(this.cpus[i]) > DEFAULT_SAMPLE_NUM {
			this.cpus[i] = this.cpus[i][1:]
		}
	}
}
