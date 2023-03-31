package main

import (
	"log"
	"os/exec"
)

func graphBc() {
	Bc30mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bc30m.png",
		"-enow",
		"-snow-30m",
		"-w600",
		"-h300",
		"-tBatteryCapacity",
		"DEF:Bc=db/Bc.rrd:Bc:AVERAGE",
		"LINE1:Bc#00FF00:%",
	)
	err := Bc30mcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bc1hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bc1h.png",
		"-enow",
		"-snow-1h",
		"-w600",
		"-h300",
		"-tBatteryCapacity",
		"DEF:Bc=db/Bc.rrd:Bc:AVERAGE",
		"LINE1:Bc#00FF00:%",
	)
	err = Bc1hcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bc3hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bc3h.png",
		"-enow",
		"-snow-3h",
		"-w600",
		"-h300",
		"-tBatteryCapacity",
		"DEF:Bc=db/Bc.rrd:Bc:AVERAGE",
		"LINE1:Bc#00FF00:%",
	)
	err = Bc3hcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bc1wcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bc1w.png",
		"-enow",
		"-snow-1w",
		"-w600",
		"-h300",
		"-tBatteryCapacity",
		"DEF:Bc=db/Bc.rrd:Bc:AVERAGE",
		"LINE1:Bc#00FF00:%",
	)
	err = Bc1wcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bc1mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bc1m.png",
		"-enow",
		"-snow-1m",
		"-w600",
		"-h300",
		"-tBatteryCapacity",
		"DEF:Bc=db/Bc.rrd:Bc:AVERAGE",
		"LINE1:Bc#00FF00:%",
	)
	err = Bc1mcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bc1ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bc1y.png",
		"-enow",
		"-snow-1y",
		"-w600",
		"-h300",
		"-tBatteryCapacity",
		"DEF:Bc=db/Bc.rrd:Bc:AVERAGE",
		"LINE1:Bc#00FF00:%",
	)
	err = Bc1ycom.Run()
	if err != nil {
		log.Println(err)
	}
	Bc2ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bc2y.png",
		"-enow",
		"-snow-2y",
		"-w600",
		"-h300",
		"-tBatteryCapacity",
		"DEF:Bc=db/Bc.rrd:Bc:AVERAGE",
		"LINE1:Bc#00FF00:%",
	)
	err = Bc2ycom.Run()
	if err != nil {
		log.Println(err)
	}
}
