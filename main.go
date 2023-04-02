package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/snksoft/crc"
	"github.com/tarm/serial"
)

const (
	ACv string = "db/ACv.rrd" //AC Voltage IN OUT
	ACf string = "db/ACf.rrd" //AC Freq IN OUT
	OPa string = "db/OPa.rrd" //Outpet Power Active Apparnt
	OPp string = "db/OPp.rrd" //Output Power %
	PV  string = "db/PV.rrd"  //PV Volt Amper IN
	BC  string = "db/BC.rrd"  //Battery chatging discharging
	Bv  string = "db/Bv.rrd"  //Battery Voltage
	Bc  string = "db/Bc.rrd"  //Battery capacity %
)

var (
	Mode string
)

func init() {
	log.Println("INIT")
	rcv, err := getData("QID")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("QID\t", rcv)
	rcv, err = getData("QVFW")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("QVFW\t", rcv)
	rcv, err = getData("QVFW2")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("QVFW2\t", rcv)

	fmt.Println("Check db dir.")
	_, err = os.Stat("db")
	if err != nil {
		err := os.Mkdir("db", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Check ACv file.")
	_, err = os.Stat(ACv)
	if err != nil {
		cmd := exec.Command("/usr/bin/rrdtool", "create", ACv, "--step", "60",
			"DS:Vi:GAUGE:120:0:250",
			"DS:Vo:GAUGE:120:0:250",
			"RRA:AVERAGE:0.5:1:1440",
			"RRA:AVERAGE:0.5:10:144",
			"RRA:AVERAGE:0.5:60:168",
			"RRA:AVERAGE:0.5:360:87600",
			"RRA:AVERAGE:0.5:720:87600")
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Check ACf file.")
	_, err = os.Stat(ACf)
	if err != nil {
		cmd := exec.Command("/usr/bin/rrdtool", "create", ACf, "--step", "60",
			"DS:Fi:GAUGE:120:0:60",
			"DS:Fo:GAUGE:120:0:60",
			"RRA:AVERAGE:0.5:1:1440",
			"RRA:AVERAGE:0.5:10:144",
			"RRA:AVERAGE:0.5:60:168",
			"RRA:AVERAGE:0.5:360:87600",
			"RRA:AVERAGE:0.5:720:87600")
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Check OPa file.")
	_, err = os.Stat(OPa)
	if err != nil {
		cmd := exec.Command("/usr/bin/rrdtool", "create", OPa, "--step", "60",
			"DS:Ac:GAUGE:120:0:U",
			"DS:Ap:GAUGE:120:0:U",
			"RRA:AVERAGE:0.5:1:1440",
			"RRA:AVERAGE:0.5:10:144",
			"RRA:AVERAGE:0.5:60:168",
			"RRA:AVERAGE:0.5:360:87600",
			"RRA:AVERAGE:0.5:720:87600")
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Check OPp file.")
	_, err = os.Stat(OPp)
	if err != nil {
		cmd := exec.Command("/usr/bin/rrdtool", "create", OPp, "--step", "60",
			"DS:Op:GAUGE:120:0:100",
			"RRA:AVERAGE:0.5:1:1440",
			"RRA:AVERAGE:0.5:10:144",
			"RRA:AVERAGE:0.5:60:168",
			"RRA:AVERAGE:0.5:360:87600",
			"RRA:AVERAGE:0.5:720:87600")
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Check PV file.")
	_, err = os.Stat(PV)
	if err != nil {
		cmd := exec.Command("/usr/bin/rrdtool", "create", PV, "--step", "60",
			"DS:V:GAUGE:120:0:50",
			"DS:A:GAUGE:120:0:25",
			"RRA:AVERAGE:0.5:1:1440",
			"RRA:AVERAGE:0.5:10:144",
			"RRA:AVERAGE:0.5:60:168",
			"RRA:AVERAGE:0.5:360:87600",
			"RRA:AVERAGE:0.5:720:87600")
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Check BC file.")
	_, err = os.Stat(BC)
	if err != nil {
		cmd := exec.Command("/usr/bin/rrdtool", "create", BC, "--step", "60",
			"DS:Bc:GAUGE:120:0:U",
			"DS:Bd:GAUGE:120:0:U",
			"RRA:AVERAGE:0.5:1:1440",
			"RRA:AVERAGE:0.5:10:144",
			"RRA:AVERAGE:0.5:60:168",
			"RRA:AVERAGE:0.5:360:87600",
			"RRA:AVERAGE:0.5:720:87600")
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Check Bv file.")
	_, err = os.Stat(Bv)
	if err != nil {
		cmd := exec.Command("/usr/bin/rrdtool", "create", Bv, "--step", "60",
			"DS:Bv:GAUGE:120:0:36",
			"RRA:AVERAGE:0.5:1:1440",
			"RRA:AVERAGE:0.5:10:144",
			"RRA:AVERAGE:0.5:60:168",
			"RRA:AVERAGE:0.5:360:87600",
			"RRA:AVERAGE:0.5:720:87600")
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Check Bc file.")
	_, err = os.Stat(Bc)
	if err != nil {
		cmd := exec.Command("/usr/bin/rrdtool", "create", Bc, "--step", "60",
			"DS:Bc:GAUGE:120:0:100",
			"RRA:AVERAGE:0.5:1:1440",
			"RRA:AVERAGE:0.5:10:144",
			"RRA:AVERAGE:0.5:60:168",
			"RRA:AVERAGE:0.5:360:87600",
			"RRA:AVERAGE:0.5:720:87600")
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	fmt.Println("START")
	tt := time.NewTicker(time.Second * 60)
	go func() {
		for range tt.C {
			getStatus()
			go graphACv()
			go graphACf()
			go graphOPa()
			go graphOPp()
			go graphPV()
			go graphBC()
			go graphBv()
			go graphBc()
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1", "localhost", "192.168.88.0/24", "10.11.12.0/24"})
	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept", "X-Requested-With", "Accept-Encoding", "User-Agent", "Referrer", "Host", "Token"},
	}
	r.Use(cors.New(corsConfig))
	r.Static("/", "./html")
	err := r.Run(":80")
	if err != nil {
		log.Fatalln(err)
	}
}

func getData(rq string) (string, error) {
	log.Println("getData")
	config := serial.Config{
		Name:     "/dev/ttyUSB0",
		Baud:     2400,
		Size:     8,
		Parity:   serial.ParityNone,
		StopBits: 1,
	}
	serial, err := serial.OpenPort(&config)
	if err != nil {
		return "", err
	}
	defer serial.Close()
	data := []byte(rq)
	Crc := crc.CalculateCRC(crc.XMODEM, data)
	cb := make([]byte, 2)
	binary.BigEndian.PutUint16(cb, uint16(Crc))
	data = append(data, cb...)
	data = append(data, 0x0d)
	_, err = serial.Write(data)
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(serial)
	reply, err := reader.ReadBytes('\x0d')
	if err != nil {
		return "", err
	}
	l := len(reply)
	d := reply[:l-3]
	s := reply[l-3 : l-1]
	sm := binary.BigEndian.Uint16(s)
	sd := crc.CalculateCRC(crc.XMODEM, d)
	if uint16(sd) != sm {
		return "", fmt.Errorf("CRC Error")
	}
	dd := d[1:]
	log.Printf("Data % X\n", dd)
	return string(dd), nil
}

func getStatus() {
	log.Println("getStatus")
	/*
	   'QPI'		=> "515049beac0d", ##  Device Protocol ID Inquiry
	   	'QID'		=> "514944d6ea0d", ## Device Serial Number Inquiry
	   	'QVFW'		=> "5156465762990d", ## Main CPU Firmware version Inqiry
	   	'QVFW2'		=> "5156465732c3f50d", ## Another CPU Firmware version Inqiry
	   	'QPIRI'		=> "5150495249f8540d", ## Device rating information Inquiry
	   	'QFLAG'		=> "51464c414798740d", ##Device Flag status Inquiry
	   	'QPIGS'		=> "5150494753b7a90d", ## Device general Status parameters Inquiry
	   	'QPIWS'		=> "5150495753b4da0d", ##Device Warning Status Inquiry
	   	'QDI'		=> "514449711b0d", ## Default Setting Value Information - default settings - needed to restore defaults, in the software
	  	'QMCHGCR'	=> "514d4348474352d8550d", ## Enquiry selectable value about max charging current - needed for creating the dropdown in the software
	   	'QMUCHGCR'	=> "514d55434847435226340d", ##Enquiry selectable value about max utility charging current - needed for creating the dropdown in the software
	   	'QBOOT'		=> "51424f4f540a88", ## Enquiry DSP has bootstrap or not
	   	'QOPM'		=> "514f504da5c50d", ## Enquiry output mode (For 4000/5000)
	   	'QPGS0'		=> "51504753303fda0d", ## Parallel Information Inquiry. same values as in QPIGS
	   	'QRST'		=> "5152535472bc0d", ## nicht dokumentiert, NAKss
	   	'QMN'		=> "514d4ebb640d", ##nicht dokumentiert, NAKss
	   	'QGMNI'		=> "51474d4e49290d", ##  nicht dokumentiert
	   	'QSID'		=> "51534944bb050d", ## nicht dokumentiert
	   	'QBEQI'		=> "51424851492ea90d", ## nicht dokumentiert VERMUTUNG: Equalisation function - liefert keine Antwort
	   	'QBEGI'		=> "51424551492ea90d", ## nicht dokumentierti
	   	'QMOD'		=> "514d4f4449c10d" ## Device Mode inquiry
	*/

	rcv, err := getData("QMOD")
	if err != nil {
		log.Println(err)
		return
	}
	Mode = rcv
	log.Println("Mode -", Mode)

	rcv, err = getData("QPIGS")
	if err != nil {
		log.Println(err)
		rcv, err = getData("QPIGS")
		if err != nil {
			log.Println(err)
			return
		}
	}

	ss := strings.Split(rcv, " ")
	log.Printf("Status %v\n", ss)
	ACvoltage := ss[0]
	acvi, _ := strconv.ParseFloat(ACvoltage, 32)
	Outputvoltage := ss[2]
	acvo, _ := strconv.ParseFloat(Outputvoltage, 32)
	acvupd := fmt.Sprintf("N:%.2f:%.2f", acvi, acvo)
	log.Println("acvupd", acvupd)
	acv := exec.Command("/usr/bin/rrdtool", "update", ACv, acvupd)
	err = acv.Run()
	if err != nil {
		log.Println(err)
	}

	ACfrequency := ss[1]
	acfi, _ := strconv.ParseFloat(ACfrequency, 32)
	Outputfrequency := ss[3]
	acfo, _ := strconv.ParseFloat(Outputfrequency, 32)
	acfupd := fmt.Sprintf("N:%.2f:%.2f", acfi, acfo)
	log.Println("acfupd", acfupd)
	acf := exec.Command("/usr/bin/rrdtool", "update", ACf, acfupd)
	err = acf.Run()
	if err != nil {
		log.Println(err)
	}

	OutputActivePower := ss[4]
	oac, _ := strconv.Atoi(OutputActivePower)
	OutputApparntPower := ss[5]
	oap, _ := strconv.Atoi(OutputApparntPower)
	oapupd := fmt.Sprintf("N:%d:%d", oac, oap)
	log.Println("oapupd", oapupd)
	oapu := exec.Command("/usr/bin/rrdtool", "update", OPa, oapupd)
	err = oapu.Run()
	if err != nil {
		log.Println(err)
	}

	OutputPower := ss[6]
	oapp, _ := strconv.Atoi(OutputPower)
	oappupd := fmt.Sprintf("N:%d", oapp)
	log.Println("oappupd", oappupd)
	oappu := exec.Command("/usr/bin/rrdtool", "update", OPp, oappupd)
	err = oappu.Run()
	if err != nil {
		log.Println(err)
	}

	PVInputCurrent := ss[12]
	pva, _ := strconv.Atoi(PVInputCurrent)
	PVInputVoltage := ss[13]
	pvv, _ := strconv.ParseFloat(PVInputVoltage, 32)
	pvu := fmt.Sprintf("N:%.2f:%d", pvv, pva)
	log.Println("pvu", pvu)
	pvupd := exec.Command("/usr/bin/rrdtool", "update", PV, pvu)
	err = pvupd.Run()
	if err != nil {
		log.Println(err)
	}

	ChargingCurrent := ss[9]
	bcc, _ := strconv.Atoi(ChargingCurrent)
	BatteryDischargeCurrent := ss[15]
	bdc, _ := strconv.Atoi(BatteryDischargeCurrent)
	bcdu := fmt.Sprintf("N:%d:%d", bcc, bdc)
	log.Println("bcdu", bcdu)
	bcdup := exec.Command("/usr/bin/rrdtool", "update", BC, bcdu)
	err = bcdup.Run()
	if err != nil {
		log.Println(err)
	}

	BatteryVoltage := ss[8]
	bv, _ := strconv.ParseFloat(BatteryVoltage, 32)
	bvu := fmt.Sprintf("N:%.2f", bv)
	log.Println("bvu", bvu)
	bvup := exec.Command("/usr/bin/rrdtool", "update", Bv, bvu)
	err = bvup.Run()
	if err != nil {
		log.Println(err)
	}

	BatteryCapacity := ss[17]
	bc, _ := strconv.Atoi(BatteryCapacity)
	bcu := fmt.Sprintf("N:%d", bc)
	log.Println("bcu", bcu)
	bcup := exec.Command("/usr/bin/rrdtool", "update", Bc, bcu)
	err = bcup.Run()
	if err != nil {
		log.Println(err)
	}
}
