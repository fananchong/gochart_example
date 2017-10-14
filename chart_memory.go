package main

import (
	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/mem"
	"math"
	"strconv"
)

type ChartMemory struct {
	gochart.ChartTime
	mem []int
}

func NewChartMemory() *ChartMemory {
	inst := &ChartMemory{mem: make([]int, DEFAULT_SAMPLE_NUM)}

	m, _ := mem.VirtualMemory()
	inst.RefreshTime = strconv.Itoa(DEFAULT_REFRESH_TIME)
	inst.ChartType = "line"
	inst.Title = "内存占用"
	inst.SubTitle = "内存大小: " + strconv.Itoa(int(math.Ceil(float64(m.Total)/float64(1024*1024*1024)))) + "GB"
	inst.YAxisText = "memory"
	inst.YMax = "100"
	inst.ValueSuffix = "%"

	return inst
}

func (this *ChartMemory) Update(now int64) []interface{} {
	this.updateData()
	datas := make([]interface{}, 0)
	json := this.AddData("memory", this.mem, now, DEFAULT_SAMPLE_NUM, DEFAULT_REFRESH_TIME)
	datas = append(datas, json)
	return datas
}

func (this *ChartMemory) updateData() {
	m, _ := mem.VirtualMemory()
	this.mem = append(this.mem, int(m.UsedPercent))
	if len(this.mem) > DEFAULT_SAMPLE_NUM {
		this.mem = this.mem[1:]
	}
}
