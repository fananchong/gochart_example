package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/net"
	"time"
)

type ExampleNetwork struct {
	gochart.ChartTime
	send     map[string][]float64
	recv     map[string][]float64
	lenlimit int
}

func NewExampleNetwork() *ExampleNetwork {
	lenlimit := 12
	nv, _ := net.IOCounters(true)
	inst := &ExampleNetwork{send: make(map[string][]float64), recv: make(map[string][]float64), lenlimit: lenlimit}
	for i := 0; i < len(nv); i++ {
		inst.send[nv[i].Name] = make([]float64, lenlimit)
		inst.recv[nv[i].Name] = make([]float64, lenlimit)
	}

	inst.RefreshTime = "1"
	inst.ChartType = "line"
	inst.Title = "网络带宽"
	inst.SubTitle = ""
	inst.YAxisText = "net"
	inst.YMax = "100"
	inst.ValueSuffix = "M"

	return inst
}

func (this *ExampleNetwork) Update() {
	this.updateData()

	endtime := 1000 * int(8*60*60+time.Now().Unix())
	begintime := endtime - 1000*this.lenlimit

	datas := make([]interface{}, 0)

	var json *simplejson.Json
	for k, v := range this.send {
		json = simplejson.New()
		json.Set("name", k+"(Sent)")
		json.Set("data", v)
		json.Set("pointInterval", 1000)
		json.Set("pointStart", begintime)
		json.Set("pointEnd", endtime)
		datas = append(datas, json)
	}
	for k, v := range this.recv {
		json = simplejson.New()
		json.Set("name", k+"(Recv)")
		json.Set("data", v)
		json.Set("pointInterval", 1000)
		json.Set("pointStart", begintime)
		json.Set("pointEnd", endtime)
		datas = append(datas, json)
	}

	datas = append(datas, json)
	json = simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	this.DataArray = string(b)
}

func (this *ExampleNetwork) updateData() {
	nv, _ := net.IOCounters(true)
	for i := 0; i < len(nv); i++ {
		this.send[nv[i].Name] = append(this.send[nv[i].Name], float64(nv[i].BytesSent)/(1024*1024))
		this.send[nv[i].Name] = this.send[nv[i].Name][1:]
	}
	for i := 0; i < len(nv); i++ {
		this.recv[nv[i].Name] = append(this.recv[nv[i].Name], float64(nv[i].BytesRecv)/(1024*1024))
		this.recv[nv[i].Name] = this.recv[nv[i].Name][1:]
	}
}
