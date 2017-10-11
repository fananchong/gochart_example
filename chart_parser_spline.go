package main

import (
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/fananchong/gochart"
	"github.com/golang/glog"
	"github.com/zieckey/goini"
)

type SplineChart struct {
	gochart.ChartSpline
	filename string
}

func (c *SplineChart) Update() {
	file := "./res/" + c.filename
	ini := goini.New()
	err := ini.ParseFile(file)
	if err != nil {
		glog.Errorln(err)
		return
	}
	err = c.Parse(ini, file)
	if err != nil {
		glog.Errorln(err)
		return
	}
}

func (c *SplineChart) Parse(ini *goini.INI, file string) error {

	c.RefreshTime = "1"

	c.ChartType, _ = ini.Get("ChartType")
	c.Title, _ = ini.Get("Title")
	c.SubTitle, _ = ini.Get("SubTitle")
	c.YAxisText, _ = ini.Get("YAxisText")
	c.XAxisNumbers, _ = ini.Get("XAxisNumbers")
	c.ValueSuffix, _ = ini.Get("ValueSuffix")
	c.Height, _ = ini.Get("Height")

	datas := make([]interface{}, 0)

	mapkeys, kvmap, err := LoadConfGetOrderMap(file)
	if err != nil {
		return err
	}

	for _, key := range mapkeys {
		if !strings.HasPrefix(key, DataPrefix) {
			continue
		}

		dd := strings.Split(kvmap[key], ",")
		jd := make([]interface{}, 0)
		for _, d := range dd {
			d = strings.TrimSpace(d)
			val, err := strconv.ParseFloat(d, 64)
			if err == nil {
				jd = append(jd, val)
			}
		}
		json := simplejson.New()
		json.Set("name", key[len(DataPrefix):])
		json.Set("data", jd)
		datas = append(datas, json)
	}

	json := simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	c.DataArray = string(b)
	return nil
}
