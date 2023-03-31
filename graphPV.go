package main

import (
	"log"
	"os/exec"
)

func graphPV() {
	pv30mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/PV30m.png",
		"-enow",
		"-snow-30m",
		"-w600",
		"-h300",
		"-tPV",
		"-vV/A",
		"DEF:PVV=db/PV.rrd:V:AVERAGE",
		"DEF:PVA=db/PV.rrd:A:AVERAGE",
		"LINE1:PVV#0000FF:V",
		"LINE1:PVA#00FF00:A",
	)
	err := pv30mcom.Run()
	if err != nil {
		log.Println(err)
	}
	pv1hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/PV1h.png",
		"-enow",
		"-snow-1h",
		"-w600",
		"-h300",
		"-tPV",
		"-vV/A",
		"DEF:PVV=db/PV.rrd:V:AVERAGE",
		"DEF:PVA=db/PV.rrd:A:AVERAGE",
		"LINE1:PVV#0000FF:V",
		"LINE1:PVA#00FF00:A",
	)
	err = pv1hcom.Run()
	if err != nil {
		log.Println(err)
	}
	pv3hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/PV3h.png",
		"-enow",
		"-snow-3h",
		"-w600",
		"-h300",
		"-tPV",
		"-vV/A",
		"DEF:PVV=db/PV.rrd:V:AVERAGE",
		"DEF:PVA=db/PV.rrd:A:AVERAGE",
		"LINE1:PVV#0000FF:V",
		"LINE1:PVA#00FF00:A",
	)
	err = pv3hcom.Run()
	if err != nil {
		log.Println(err)
	}
	pv1dcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/PV1d.png",
		"-enow",
		"-snow-1d",
		"-w600",
		"-h300",
		"-tPV",
		"-vV/A",
		"DEF:PVV=db/PV.rrd:V:AVERAGE",
		"DEF:PVA=db/PV.rrd:A:AVERAGE",
		"LINE1:PVV#0000FF:V",
		"LINE1:PVA#00FF00:A",
	)
	err = pv1dcom.Run()
	if err != nil {
		log.Println(err)
	}
	pv1wcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/PV1w.png",
		"-enow",
		"-snow-1w",
		"-w600",
		"-h300",
		"-tPV",
		"-vV/A",
		"DEF:PVV=db/PV.rrd:V:AVERAGE",
		"DEF:PVA=db/PV.rrd:A:AVERAGE",
		"LINE1:PVV#0000FF:V",
		"LINE1:PVA#00FF00:A",
	)
	err = pv1wcom.Run()
	if err != nil {
		log.Println(err)
	}
	pv1mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/PV1m.png",
		"-enow",
		"-snow-1m",
		"-w600",
		"-h300",
		"-tPV",
		"-vV/A",
		"DEF:PVV=db/PV.rrd:V:AVERAGE",
		"DEF:PVA=db/PV.rrd:A:AVERAGE",
		"LINE1:PVV#0000FF:V",
		"LINE1:PVA#00FF00:A",
	)
	err = pv1mcom.Run()
	if err != nil {
		log.Println(err)
	}
	pv1ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/PV1y.png",
		"-enow",
		"-snow-1y",
		"-w600",
		"-h300",
		"-tPV",
		"-vV/A",
		"DEF:PVV=db/PV.rrd:V:AVERAGE",
		"DEF:PVA=db/PV.rrd:A:AVERAGE",
		"LINE1:PVV#0000FF:V",
		"LINE1:PVA#00FF00:A",
	)
	err = pv1ycom.Run()
	if err != nil {
		log.Println(err)
	}
	pv2ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/PV2y.png",
		"-enow",
		"-snow-2y",
		"-w600",
		"-h300",
		"-tPV",
		"-vV/A",
		"DEF:PVV=db/PV.rrd:V:AVERAGE",
		"DEF:PVA=db/PV.rrd:A:AVERAGE",
		"LINE1:PVV#0000FF:V",
		"LINE1:PVA#00FF00:A",
	)
	err = pv2ycom.Run()
	if err != nil {
		log.Println(err)
	}
}
