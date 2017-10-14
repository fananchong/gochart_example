package main

import (
	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
)

type ChartCPU struct {
	gochart.ChartTime
	cpus map[int][]int
}

func NewChartCPU() *ChartCPU {
	cc, _ := cpu.Percent(0, false)
	inst := &ChartCPU{cpus: make(map[int][]int)}
	for i := 0; i < len(cc); i++ {
		inst.cpus[i] = make([]int, DEFAULT_SAMPLE_NUM)
	}

	inst.RefreshTime = strconv.Itoa(DEFAULT_REFRESH_TIME)
	inst.ChartType = "line"
	inst.Title = "CPU占用"
	inst.SubTitle = ""
	inst.YAxisText = "cpu"
	inst.YMax = "100"
	inst.ValueSuffix = "%"

	return inst
}

func (this *ChartCPU) Update(now int64) []interface{} {
	this.updateData()
	datas := make([]interface{}, 0)
	for i := 0; i < len(this.cpus); i++ {
		json := this.AddData("cpu"+strconv.Itoa(i), this.cpus[i], now, DEFAULT_SAMPLE_NUM, DEFAULT_REFRESH_TIME)
		datas = append(datas, json)
	}
	return datas
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
