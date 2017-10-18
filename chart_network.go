package main

import (
	"fmt"
	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/net"
	"strconv"
)

type ChartNetwork struct {
	gochart.ChartTime
	send    []float64
	recv    []float64
	presend uint64
	prerecv uint64
}

func NewChartNetwork() *ChartNetwork {
	inst := &ChartNetwork{send: make([]float64, DEFAULT_SAMPLE_NUM), recv: make([]float64, DEFAULT_SAMPLE_NUM)}
	inst.RefreshTime = strconv.Itoa(DEFAULT_REFRESH_TIME)
	inst.ChartType = "line"
	inst.Title = "网络带宽"
	inst.SubTitle = ""
	inst.YAxisText = "net"
	inst.YMax = "1000"
	inst.ValueSuffix = "Mbps"
	return inst
}

func (this *ChartNetwork) Update(now int64) []interface{} {
	this.updateData()
	datas := make([]interface{}, 0)
	json1 := this.AddData("Sent", this.send, now, DEFAULT_SAMPLE_NUM, DEFAULT_REFRESH_TIME)
	datas = append(datas, json1)
	json2 := this.AddData("Recv", this.recv, now, DEFAULT_SAMPLE_NUM, DEFAULT_REFRESH_TIME)
	datas = append(datas, json2)
	return datas
}

func (this *ChartNetwork) updateData() {
	nv, _ := net.IOCounters(false)

	if this.presend == 0 {
		this.presend = nv[0].BytesSent
	}
	if this.prerecv == 0 {
		this.prerecv = nv[0].BytesRecv
	}

	v1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(nv[0].BytesSent-this.presend)*8/float64(1024*1024)), 64)
	this.send = append(this.send, v1)
	for len(this.send) > DEFAULT_SAMPLE_NUM {
		this.send = this.send[1:]
	}
	v2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(nv[0].BytesRecv-this.prerecv)*8/float64(1024*1024)), 64)
	this.recv = append(this.recv, v2)
	if len(this.recv) > DEFAULT_SAMPLE_NUM {
		this.recv = this.recv[1:]
	}

	this.presend = nv[0].BytesSent
	this.prerecv = nv[0].BytesRecv
}
