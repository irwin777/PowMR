package main

import (
	"log"
	"os/exec"
)

func graphACv() {
	ACv30mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACv30m.png",
		"-enow",
		"-snow-30m",
		"-w600",
		"-h300",
		"-tACVoltage",
		"-vVolt",
		"-E",
		"DEF:ACvi=db/ACv.rrd:Vi:AVERAGE",
		"DEF:ACvo=db/ACv.rrd:Vo:AVERAGE",
		"LINE1:ACvi#0000FF:In",
		"LINE1:ACvo#00FF00:Out",
		"GPRINT:ACvi:LAST:Current\\:%4.1lf",
		"GPRINT:ACvo:LAST:Current\\:%4.1lf",
	)
	err := ACv30mcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACv1hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACv1h.png",
		"-enow",
		"-snow-1h",
		"-w600",
		"-h300",
		"-tACVoltage",
		"-vVolt",
		"-E",
		"DEF:ACvi=db/ACv.rrd:Vi:AVERAGE",
		"DEF:ACvo=db/ACv.rrd:Vo:AVERAGE",
		"LINE1:ACvi#0000FF:In",
		"LINE1:ACvo#00FF00:Out",
		"GPRINT:ACvi:LAST:Current\\:%4.1lf",
		"GPRINT:ACvo:LAST:Current\\:%4.1lf",
	)
	err = ACv1hcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACv3hcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACv3h.png",
		"-enow",
		"-snow-3h",
		"-w600",
		"-h300",
		"-tACVoltage",
		"-vVolt",
		"-E",
		"DEF:ACvi=db/ACv.rrd:Vi:AVERAGE",
		"DEF:ACvo=db/ACv.rrd:Vo:AVERAGE",
		"LINE1:ACvi#0000FF:In",
		"LINE1:ACvo#00FF00:Out",
		"GPRINT:ACvi:LAST:Current\\:%4.1lf",
		"GPRINT:ACvo:LAST:Current\\:%4.1lf",
	)
	err = ACv3hcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACv1dcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACv1d.png",
		"-enow",
		"-snow-1d",
		"-w600",
		"-h300",
		"-tACVoltage",
		"-vVolt",
		"-E",
		"DEF:ACvi=db/ACv.rrd:Vi:AVERAGE",
		"DEF:ACvo=db/ACv.rrd:Vo:AVERAGE",
		"LINE1:ACvi#0000FF:In",
		"LINE1:ACvo#00FF00:Out",
		"GPRINT:ACvi:LAST:Current\\:%4.1lf",
		"GPRINT:ACvo:LAST:Current\\:%4.1lf",
	)
	err = ACv1dcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACv1wcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACv1w.png",
		"-enow",
		"-snow-1w",
		"-w600",
		"-h300",
		"-tACVoltage",
		"-vVolt",
		"-E",
		"DEF:ACvi=db/ACv.rrd:Vi:AVERAGE",
		"DEF:ACvo=db/ACv.rrd:Vo:AVERAGE",
		"LINE1:ACvi#0000FF:In",
		"LINE1:ACvo#00FF00:Out",
		"GPRINT:ACvi:LAST:Current\\:%4.1lf",
		"GPRINT:ACvo:LAST:Current\\:%4.1lf",
	)
	err = ACv1wcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACv1mcom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACv1m.png",
		"-enow",
		"-snow-1m",
		"-w600",
		"-h300",
		"-tACVoltage",
		"-vVolt",
		"-E",
		"DEF:ACvi=db/ACv.rrd:Vi:AVERAGE",
		"DEF:ACvo=db/ACv.rrd:Vo:AVERAGE",
		"LINE1:ACvi#0000FF:In",
		"LINE1:ACvo#00FF00:Out",
		"GPRINT:ACvi:LAST:Current\\:%4.1lf",
		"GPRINT:ACvo:LAST:Current\\:%4.1lf",
	)
	err = ACv1mcom.Run()
	if err != nil {
		log.Println(err)
	}
	ACv1ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACv1y.png",
		"-enow",
		"-snow-1y",
		"-w600",
		"-h300",
		"-tACVoltage",
		"-vVolt",
		"-E",
		"DEF:ACvi=db/ACv.rrd:Vi:AVERAGE",
		"DEF:ACvo=db/ACv.rrd:Vo:AVERAGE",
		"LINE1:ACvi#0000FF:In",
		"LINE1:ACvo#00FF00:Out",
		"GPRINT:ACvi:LAST:Current\\:%4.1lf",
		"GPRINT:ACvo:LAST:Current\\:%4.1lf",
	)
	err = ACv1ycom.Run()
	if err != nil {
		log.Println(err)
	}
	ACv2ycom := exec.Command("/usr/bin/rrdtool", "graph",
		"html/rrd/ACv2y.png",
		"-enow",
		"-snow-2y",
		"-w600",
		"-h300",
		"-tACVoltage",
		"-vVolt",
		"-E",
		"DEF:ACvi=db/ACv.rrd:Vi:AVERAGE",
		"DEF:ACvo=db/ACv.rrd:Vo:AVERAGE",
		"LINE1:ACvi#0000FF:In",
		"LINE1:ACvo#00FF00:Out",
		"GPRINT:ACvi:LAST:Current\\:%4.1lf",
		"GPRINT:ACvo:LAST:Current\\:%4.1lf",
	)
	err = ACv2ycom.Run()
	if err != nil {
		log.Println(err)
	}
}
