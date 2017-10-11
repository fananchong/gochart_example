package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/fananchong/gochart"
	"time"
)

type TimeChartExample struct {
	gochart.ChartTime
}

var (
	playernum []int = []int{1, 5, 3, 6, 8, 0, 10, 11, 6, 10, 11}
	roomnum   []int = []int{11, 7, 6, 3, 9, 2, 4, 1, 9, 3, 2}
)

func (this *TimeChartExample) Update() {
	this.RefreshTime = "1"

	this.ChartType = "line"
	this.Title = "x时间轴例子"
	this.SubTitle = ""
	this.YAxisText = "Num"
	this.ValueSuffix = ""

	endtime := 1000 * uint(8*60*60+time.Now().Unix())
	begintime := endtime - 1000*uint(len(playernum))

	datas := make([]interface{}, 0)

	json := simplejson.New()
	json.Set("name", "player")
	json.Set("data", playernum)
	json.Set("pointInterval", 1000)
	json.Set("pointStart", begintime)
	json.Set("pointEnd", endtime)
	datas = append(datas, json)

	json = simplejson.New()
	json.Set("name", "room")
	json.Set("data", roomnum)
	json.Set("pointInterval", 1000)
	json.Set("pointStart", begintime)
	json.Set("pointEnd", endtime)
	datas = append(datas, json)

	json = simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	this.DataArray = string(b)
}
