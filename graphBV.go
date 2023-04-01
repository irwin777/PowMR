package main

import (
	"log"
	"os/exec"
)

func graphBv() {
	Bv30mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bv30m.png",
		"-enow",
		"-snow-30m",
		"-w600",
		"-h300",
		"-tBatteryVoltage",
		"-vVolt",
		"-E",
		"DEF:Bv=db/Bv.rrd:Bv:AVERAGE",
		"LINE1:Bv#00FF00:V",
	)
	err := Bv30mcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bv1hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bv1h.png",
		"-enow",
		"-snow-1h",
		"-w600",
		"-h300",
		"-tBatteryVoltage",
		"-vVolt",
		"-E",
		"DEF:Bv=db/Bv.rrd:Bv:AVERAGE",
		"LINE1:Bv#00FF00:V",
	)
	err = Bv1hcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bv3hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bv3h.png",
		"-enow",
		"-snow-3h",
		"-w600",
		"-h300",
		"-tBatteryVoltage",
		"-vVolt",
		"-E",
		"DEF:Bv=db/Bv.rrd:Bv:AVERAGE",
		"LINE1:Bv#00FF00:V",
	)
	err = Bv3hcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bv1dcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bv1d.png",
		"-enow",
		"-snow-1d",
		"-w600",
		"-h300",
		"-tBatteryVoltage",
		"-vVolt",
		"-E",
		"DEF:Bv=db/Bv.rrd:Bv:AVERAGE",
		"LINE1:Bv#00FF00:V",
	)
	err = Bv1dcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bv1wcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bv1w.png",
		"-enow",
		"-snow-1w",
		"-w600",
		"-h300",
		"-tBatteryVoltage",
		"-vVolt",
		"-E",
		"DEF:Bv=db/Bv.rrd:Bv:AVERAGE",
		"LINE1:Bv#00FF00:V",
	)
	err = Bv1wcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bv1mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bv1m.png",
		"-enow",
		"-snow-1m",
		"-w600",
		"-h300",
		"-tBatteryVoltage",
		"-vVolt",
		"-E",
		"DEF:Bv=db/Bv.rrd:Bv:AVERAGE",
		"LINE1:Bv#00FF00:V",
	)
	err = Bv1mcom.Run()
	if err != nil {
		log.Println(err)
	}
	Bv1ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bv1y.png",
		"-enow",
		"-snow-1y",
		"-w600",
		"-h300",
		"-tBatteryVoltage",
		"-vVolt",
		"-E",
		"DEF:Bv=db/Bv.rrd:Bv:AVERAGE",
		"LINE1:Bv#00FF00:V",
	)
	err = Bv1ycom.Run()
	if err != nil {
		log.Println(err)
	}
	Bv2ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/Bv2y.png",
		"-enow",
		"-snow-2y",
		"-w600",
		"-h300",
		"-tBatteryVoltage",
		"-vVolt",
		"-E",
		"DEF:Bv=db/Bv.rrd:Bv:AVERAGE",
		"LINE1:Bv#00FF00:V",
	)
	err = Bv2ycom.Run()
	if err != nil {
		log.Println(err)
	}
}
