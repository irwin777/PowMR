package main

import (
	"log"
	"os/exec"
)

func graphOPp() {
	OPp30mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPp30m.png",
		"-enow",
		"-snow-30m",
		"-w600",
		"-h300",
		"-tOutPowerProcent",
		"-vPercent",
		"-E",
		"DEF:OPp=db/OPp.rrd:Op:AVERAGE",
		"AREA:OPp#00FF00:%",
	)
	err := OPp30mcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPp1hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPp1h.png",
		"-enow",
		"-snow-1h",
		"-w600",
		"-h300",
		"-tOutPowerProcent",
		"-vPercent",
		"-E",
		"DEF:OPp=db/OPp.rrd:Op:AVERAGE",
		"AREA:OPp#00FF00:%",
	)
	err = OPp1hcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPp3hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPp3h.png",
		"-enow",
		"-snow-3h",
		"-w600",
		"-h300",
		"-tOutPowerProcent",
		"-vPercent",
		"-E",
		"DEF:OPp=db/OPp.rrd:Op:AVERAGE",
		"AREA:OPp#00FF00:%",
	)
	err = OPp3hcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPp1dcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPp1d.png",
		"-enow",
		"-snow-1d",
		"-w600",
		"-h300",
		"-tOutPowerProcent",
		"-vPercent",
		"-E",
		"DEF:OPp=db/OPp.rrd:Op:AVERAGE",
		"AREA:OPp#00FF00:%",
	)
	err = OPp1dcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPp1wcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPp1w.png",
		"-enow",
		"-snow-1w",
		"-w600",
		"-h300",
		"-tOutPowerProcent",
		"-vPercent",
		"-E",
		"DEF:OPp=db/OPp.rrd:Op:AVERAGE",
		"AREA:OPp#00FF00:%",
	)
	err = OPp1wcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPp1mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPp1m.png",
		"-enow",
		"-snow-1m",
		"-w600",
		"-h300",
		"-tOutPowerProcent",
		"-vPercent",
		"-E",
		"DEF:OPp=db/OPp.rrd:Op:AVERAGE",
		"AREA:OPp#00FF00:%",
	)
	err = OPp1mcom.Run()
	if err != nil {
		log.Println(err)
	}
	OPp1ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPp1y.png",
		"-enow",
		"-snow-1y",
		"-w600",
		"-h300",
		"-tOutPowerProcent",
		"-vPercent",
		"-E",
		"DEF:OPp=db/OPp.rrd:Op:AVERAGE",
		"AREA:OPp#00FF00:%",
	)
	err = OPp1ycom.Run()
	if err != nil {
		log.Println(err)
	}
	OPp2ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/OPp2y.png",
		"-enow",
		"-snow-2y",
		"-w600",
		"-h300",
		"-tOutPowerProcent",
		"-vPercent",
		"-E",
		"DEF:OPp=db/OPp.rrd:Op:AVERAGE",
		"AREA:OPp#00FF00:%",
	)
	err = OPp2ycom.Run()
	if err != nil {
		log.Println(err)
	}
}
