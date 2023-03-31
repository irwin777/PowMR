package main

import (
	"log"
	"os/exec"
)

func graphACf() {
	ACf30mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACf30m.png",
		"-enow",
		"-snow-30m",
		"-w600",
		"-h300",
		"-tACfreqmebcy",
		"DEF:ACfi=db/ACf.rrd:Fi:AVERAGE",
		"DEF:ACfo=db/ACf.rrd:Fo:AVERAGE",
		"LINE1:ACfi#0000FF:In",
		"LINE1:ACfo#00FF00:Out",
	)
	err := ACf30mcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACf1hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACf1h.png",
		"-enow",
		"-snow-1h",
		"-w600",
		"-h300",
		"-tACfreqmebcy",
		"DEF:ACfi=db/ACf.rrd:Fi:AVERAGE",
		"DEF:ACfo=db/ACf.rrd:Fo:AVERAGE",
		"LINE1:ACfi#0000FF:In",
		"LINE1:ACfo#00FF00:Out",
	)
	err = ACf1hcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACf3hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACf3h.png",
		"-enow",
		"-snow-3h",
		"-w600",
		"-h300",
		"-tACfreqmebcy",
		"DEF:ACfi=db/ACf.rrd:Fi:AVERAGE",
		"DEF:ACfo=db/ACf.rrd:Fo:AVERAGE",
		"LINE1:ACfi#0000FF:In",
		"LINE1:ACfo#00FF00:Out",
	)
	err = ACf3hcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACf1wcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACf1w.png",
		"-enow",
		"-snow-1w",
		"-w600",
		"-h300",
		"-tACfreqmebcy",
		"DEF:ACfi=db/ACf.rrd:Fi:AVERAGE",
		"DEF:ACfo=db/ACf.rrd:Fo:AVERAGE",
		"LINE1:ACfi#0000FF:In",
		"LINE1:ACfo#00FF00:Out",
	)
	err = ACf1wcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACf1mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACf1m.png",
		"-enow",
		"-snow-1m",
		"-w600",
		"-h300",
		"-tACfreqmebcy",
		"DEF:ACfi=db/ACf.rrd:Fi:AVERAGE",
		"DEF:ACfo=db/ACf.rrd:Fo:AVERAGE",
		"LINE1:ACfi#0000FF:In",
		"LINE1:ACfo#00FF00:Out",
	)
	err = ACf1mcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACf1ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACf1y.png",
		"-enow",
		"-snow-1y",
		"-w600",
		"-h300",
		"-tACfreqmebcy",
		"DEF:ACfi=db/ACf.rrd:Fi:AVERAGE",
		"DEF:ACfo=db/ACf.rrd:Fo:AVERAGE",
		"LINE1:ACfi#0000FF:In",
		"LINE1:ACfo#00FF00:Out",
	)
	err = ACf1ycom.Run()
	if err != nil {
		log.Println(err)
	}
	ACf2ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACf2y.png",
		"-enow",
		"-snow-2y",
		"-w600",
		"-h300",
		"-tACfreqmebcy",
		"DEF:ACfi=db/ACf.rrd:Fi:AVERAGE",
		"DEF:ACfo=db/ACf.rrd:Fo:AVERAGE",
		"LINE1:ACfi#0000FF:In",
		"LINE1:ACfo#00FF00:Out",
	)
	err = ACf2ycom.Run()
	if err != nil {
		log.Println(err)
	}
}