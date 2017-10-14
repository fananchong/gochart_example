package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/mem"
	"math"
	"strconv"
	"time"
)

type ExampleMemory struct {
	gochart.ChartTime
	mem      []int
	lenlimit int
}

func NewExampleMemory() *ExampleMemory {
	lenlimit := DEFAULT_SAMPLE_NUM
	inst := &ExampleMemory{mem: make([]int, lenlimit), lenlimit: lenlimit}

	m, _ := mem.VirtualMemory()
	inst.RefreshTime = strconv.Itoa(DEFAULT_REFRESH_TIME)
	inst.ChartType = "line"
	inst.Title = "内存占用"
	inst.SubTitle = "内存大小: " + strconv.Itoa(int(math.Ceil(float64(m.Total)/float64(1024*1024*1024)))) + "GB"
	inst.YAxisText = "memory"
	inst.YMax = "100"
	inst.ValueSuffix = "%"
	inst.TickInterval = strconv.Itoa(DEFAULT_REFRESH_TIME * 1000)

	return inst
}

func (this *ExampleMemory) Update() {
	this.updateData()

	endtime := 1000 * int(8*60*60+time.Now().Unix())
	begintime := endtime - 1000*this.lenlimit*DEFAULT_REFRESH_TIME

	datas := make([]interface{}, 0)

	var json *simplejson.Json
	json = simplejson.New()
	json.Set("name", "memory")
	json.Set("data", this.mem)
	json.Set("pointStart", begintime)
	json.Set("pointEnd", endtime)

	datas = append(datas, json)
	json = simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	this.DataArray = string(b)
}

func (this *ExampleMemory) updateData() {
	m, _ := mem.VirtualMemory()
	this.mem = append(this.mem, int(m.UsedPercent))
	if len(this.mem) > DEFAULT_SAMPLE_NUM {
		this.mem = this.mem[1:]
	}
}
