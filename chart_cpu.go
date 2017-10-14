package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
	"time"
)

type ChartCPU struct {
	gochart.ChartTime
	cpus     map[int][]int
	lenlimit int
}

func NewChartCPU() *ChartCPU {
	lenlimit := DEFAULT_SAMPLE_NUM
	cc, _ := cpu.Percent(0, false)
	inst := &ChartCPU{cpus: make(map[int][]int), lenlimit: lenlimit}
	for i := 0; i < len(cc); i++ {
		inst.cpus[i] = make([]int, lenlimit)
	}

	inst.RefreshTime = strconv.Itoa(DEFAULT_REFRESH_TIME)
	inst.ChartType = "line"
	inst.Title = "CPU占用"
	inst.SubTitle = ""
	inst.YAxisText = "cpu"
	inst.YMax = "100"
	inst.ValueSuffix = "%"
	inst.TickInterval = strconv.Itoa(DEFAULT_REFRESH_TIME * 1000)

	return inst
}

func (this *ChartCPU) Update() {
	this.updateData()

	endtime := 1000 * int(8*60*60+time.Now().Unix())
	begintime := endtime - 1000*this.lenlimit*DEFAULT_REFRESH_TIME

	datas := make([]interface{}, 0)
	var json *simplejson.Json
	for i := 0; i < len(this.cpus); i++ {
		json = simplejson.New()
		json.Set("name", "cpu"+strconv.Itoa(i))
		json.Set("data", this.cpus[i])
		json.Set("pointStart", begintime)
		json.Set("pointEnd", endtime)
		datas = append(datas, json)
	}
	json = simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	this.DataArray = string(b)
}

func (this *ChartCPU) updateData() {
	cc, _ := cpu.Percent(0, false)
	for i := 0; i < len(cc); i++ {
		this.cpus[i] = append(this.cpus[i], int(cc[i]))
		if len(this.cpus[i]) > DEFAULT_SAMPLE_NUM {
			this.cpus[i] = this.cpus[i][1:]
		}
	}
}
