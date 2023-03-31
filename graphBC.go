package main

import (
	"log"
	"os/exec"
)

func graphBC() {
	BC30mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/BC30m.png",
		"-enow",
		"-snow-30m",
		"-w600",
		"-h300",
		"-tBC_30_m",
		"DEF:BCc=db/BC.rrd:Bc:AVERAGE",
		"DEF:BCd=db/BC.rrd:Bd:AVERAGE",
		"LINE1:BCc#0000FF:Charge",
		"LINE1:BCd#00FF00:Discharge",
	)
	err := BC30mcom.Run()
	if err != nil {
		log.Println(err)
	}
	BC1hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/BC1h.png",
		"-enow",
		"-snow-1h",
		"-w600",
		"-h300",
		"-tBC_1_h",
		"DEF:BCc=db/BC.rrd:Bc:AVERAGE",
		"DEF:BCd=db/BC.rrd:Bd:AVERAGE",
		"LINE1:BCc#0000FF:Charge",
		"LINE1:BCd#00FF00:Discharge",
	)
	err = BC1hcom.Run()
	if err != nil {
		log.Println(err)
	}
	BC3hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/BC3h.png",
		"-enow",
		"-snow-3h",
		"-w600",
		"-h300",
		"-tBC_3_h",
		"DEF:BCc=db/BC.rrd:Bc:AVERAGE",
		"DEF:BCd=db/BC.rrd:Bd:AVERAGE",
		"LINE1:BCc#0000FF:Charge",
		"LINE1:BCd#00FF00:Discharge",
	)
	err = BC3hcom.Run()
	if err != nil {
		log.Println(err)
	}
	BC1wcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/BC1w.png",
		"-enow",
		"-snow-1w",
		"-w600",
		"-h300",
		"-tBC_1_w",
		"DEF:BCc=db/BC.rrd:Bc:AVERAGE",
		"DEF:BCd=db/BC.rrd:Bd:AVERAGE",
		"LINE1:BCc#0000FF:Charge",
		"LINE1:BCd#00FF00:Discharge",
	)
	err = BC1wcom.Run()
	if err != nil {
		log.Println(err)
	}
	BC1mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/BC1m.png",
		"-enow",
		"-snow-1m",
		"-w600",
		"-h300",
		"-tBC_1_m",
		"DEF:BCc=db/BC.rrd:Bc:AVERAGE",
		"DEF:BCd=db/BC.rrd:Bd:AVERAGE",
		"LINE1:BCc#0000FF:Charge",
		"LINE1:BCd#00FF00:Discharge",
	)
	err = BC1mcom.Run()
	if err != nil {
		log.Println(err)
	}
	BC1ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/BC1y.png",
		"-enow",
		"-snow-1y",
		"-w600",
		"-h300",
		"-tBC_1_y",
		"DEF:BCc=db/BC.rrd:Bc:AVERAGE",
		"DEF:BCd=db/BC.rrd:Bd:AVERAGE",
		"LINE1:BCc#0000FF:Charge",
		"LINE1:BCd#00FF00:Discharge",
	)
	err = BC1ycom.Run()
	if err != nil {
		log.Println(err)
	}
	BC2ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/BC2y.png",
		"-enow",
		"-snow-2y",
		"-w600",
		"-h300",
		"-tBC_2_y",
		"DEF:BCc=db/BC.rrd:Bc:AVERAGE",
		"DEF:BCd=db/BC.rrd:Bd:AVERAGE",
		"LINE1:BCc#0000FF:Charge",
		"LINE1:BCd#00FF00:Discharge",
	)
	err = BC2ycom.Run()
	if err != nil {
		log.Println(err)
	}
}
