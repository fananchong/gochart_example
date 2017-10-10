package main

import (
	"github.com/fananchong/gochart"
	"github.com/golang/glog"
	"github.com/zieckey/goini"
	"strings"
)

type PieChart struct {
	gochart.ChartPie
	filename string
}

func (c *PieChart) Update() {
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

func (c *PieChart) Parse(ini *goini.INI, file string) error {
	c.ChartType, _ = ini.Get("ChartType")
	c.Title, _ = ini.Get("Title")
	c.SubTitle, _ = ini.Get("SubTitle")
	c.SeriesName, _ = ini.Get("SeriesName")

	/* Generate DataArray:
	   [
	       ['Firefox',   45.0],
	       ['IE',       26.8],
	       ['Chrome',  12.8],
	       ['Safari',    8.5],
	       ['Opera',     6.2],
	       ['Others',   0.7]
	   ]
	*/

	DataArray := "[\n"
	mapkeys, kvmap, err := LoadConfGetOrderMap(file)
	if err != nil {
		return err
	}
	for _, k := range mapkeys {
		if !strings.HasPrefix(k, DataPrefix) {
			continue
		}

		key := k[len(DataPrefix):]
		DataArray = DataArray + "['" + key + "' , " + kvmap[k] + "],\n"
	}

	DataArray = DataArray + "]"

	c.DataArray = DataArray
	return nil
}
