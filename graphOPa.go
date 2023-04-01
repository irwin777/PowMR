package main

import (
	"log"
	"os/exec"
)

func graphOPa() {
	OPa30mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPa30m.png",
		"-enow",
		"-snow-30m",
		"-w600",
		"-h300",
		"-tOutPower",
		"-vVA/W",
		"-E",
		"DEF:OPac=db/OPa.rrd:Ac:AVERAGE",
		"DEF:OPap=db/OPa.rrd:Ap:AVERAGE",
		"LINE1:OPac#0000FF:ActivePower",
		"LINE1:OPap#00FF00:ApparntPower",
	)
	err := OPa30mcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPa1hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPa1h.png",
		"-enow",
		"-snow-1h",
		"-w600",
		"-h300",
		"-tOutPower",
		"-vVA/W",
		"-E",
		"DEF:OPac=db/OPa.rrd:Ac:AVERAGE",
		"DEF:OPap=db/OPa.rrd:Ap:AVERAGE",
		"LINE1:OPac#0000FF:ActivePower",
		"LINE1:OPap#00FF00:ApparntPower",
	)
	err = OPa1hcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPa3hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPa3h.png",
		"-enow",
		"-snow-3h",
		"-w600",
		"-h300",
		"-tOutPower",
		"-vVA/W",
		"-E",
		"DEF:OPac=db/OPa.rrd:Ac:AVERAGE",
		"DEF:OPap=db/OPa.rrd:Ap:AVERAGE",
		"LINE1:OPac#0000FF:ActivePower",
		"LINE1:OPap#00FF00:ApparntPower",
	)
	err = OPa3hcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPa1dcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPa1d.png",
		"-enow",
		"-snow-1d",
		"-w600",
		"-h300",
		"-tOutPower",
		"-vVA/W",
		"-E",
		"DEF:OPac=db/OPa.rrd:Ac:AVERAGE",
		"DEF:OPap=db/OPa.rrd:Ap:AVERAGE",
		"LINE1:OPac#0000FF:ActivePower",
		"LINE1:OPap#00FF00:ApparntPower",
	)
	err = OPa1dcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPa1wcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPa1w.png",
		"-enow",
		"-snow-1w",
		"-w600",
		"-h300",
		"-tOutPower",
		"-vVA/W",
		"-E",
		"DEF:OPac=db/OPa.rrd:Ac:AVERAGE",
		"DEF:OPap=db/OPa.rrd:Ap:AVERAGE",
		"LINE1:OPac#0000FF:ActivePower",
		"LINE1:OPap#00FF00:ApparntPower",
	)
	err = OPa1wcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPa1mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPa1m.png",
		"-enow",
		"-snow-1m",
		"-w600",
		"-h300",
		"-tOutPower",
		"-vVA/W",
		"-E",
		"DEF:OPac=db/OPa.rrd:Ac:AVERAGE",
		"DEF:OPap=db/OPa.rrd:Ap:AVERAGE",
		"LINE1:OPac#0000FF:ActivePower",
		"LINE1:OPap#00FF00:ApparntPower",
	)
	err = OPa1mcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPa1ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPa1y.png",
		"-enow",
		"-snow-1y",
		"-w600",
		"-h300",
		"-tOutPower",
		"-vVA/W",
		"-E",
		"DEF:OPac=db/OPa.rrd:Ac:AVERAGE",
		"DEF:OPap=db/OPa.rrd:Ap:AVERAGE",
		"LINE1:OPac#0000FF:ActivePower",
		"LINE1:OPap#00FF00:ApparntPower",
	)
	err = OPa1ycom.Run()
	if err != nil {
		log.Println(err)
	}
	OPa2ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPa2y.png",
		"-enow",
		"-snow-2y",
		"-w600",
		"-h300",
		"-tOutPower",
		"-vVA/W",
		"-E",
		"DEF:OPac=db/OPa.rrd:Ac:AVERAGE",
		"DEF:OPap=db/OPa.rrd:Ap:AVERAGE",
		"LINE1:OPac#0000FF:ActivePower",
		"LINE1:OPap#00FF00:ApparntPower",
	)
	err = OPa2ycom.Run()
	if err != nil {
		log.Println(err)
	}
}
