package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/net"
	"strconv"
	"time"
)

type ExampleNetwork struct {
	gochart.ChartTime
	send     []float64
	recv     []float64
	presend  uint64
	prerecv  uint64
	lenlimit int
}

func NewExampleNetwork() *ExampleNetwork {
	lenlimit := 1800
	inst := &ExampleNetwork{send: make([]float64, lenlimit), recv: make([]float64, lenlimit), lenlimit: lenlimit}
	inst.RefreshTime = "1"
	inst.ChartType = "line"
	inst.Title = "网络带宽"
	inst.SubTitle = ""
	inst.YAxisText = "net"
	inst.YMax = "1024000"
	inst.ValueSuffix = "K"
	return inst
}

func (this *ExampleNetwork) Update() {
	this.updateData()

	endtime := 1000 * int(8*60*60+time.Now().Unix())
	begintime := endtime - 1000*this.lenlimit

	datas := make([]interface{}, 0)

	var json *simplejson.Json

	json = simplejson.New()
	json.Set("name", "Sent")
	json.Set("data", this.send)
	json.Set("pointInterval", 1000)
	json.Set("pointStart", begintime)
	json.Set("pointEnd", endtime)
	datas = append(datas, json)

	json = simplejson.New()
	json.Set("name", "Recv")
	json.Set("data", this.recv)
	json.Set("pointInterval", 1000)
	json.Set("pointStart", begintime)
	json.Set("pointEnd", endtime)
	datas = append(datas, json)

	json = simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	this.DataArray = string(b)
}

func (this *ExampleNetwork) updateData() {
	nv, _ := net.IOCounters(false)

	if this.presend == 0 {
		this.presend = nv[0].BytesSent
	}
	if this.prerecv == 0 {
		this.prerecv = nv[0].BytesRecv
	}

	v1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(nv[0].BytesSent-this.presend)/float64(1024)), 64)
	this.send = append(this.send, v1)
	this.send = this.send[1:]
	v2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(nv[0].BytesRecv-this.prerecv)/float64(1024)), 64)
	this.recv = append(this.recv, v2)
	this.recv = this.recv[1:]

	this.presend = nv[0].BytesSent
	this.prerecv = nv[0].BytesRecv
}
