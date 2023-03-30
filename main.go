package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/snksoft/crc"
	"github.com/tarm/serial"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type status struct {
	gorm.Model
	DeviceMode              string
	ACvoltage               float64
	ACfrequency             float64
	OutputVoltage           float64
	OutputFrequency         float64
	OutputActivePower       int
	OutputApparntPower      int
	OutputPower             int
	PVBAT                   float64
	ChargingCurrent         int
	PVinputCurrent          int
	PVinputVoltage          float64
	BatteryVoltage          float64
	BatteryDischargeCurrent int
	BatteryCapacity         int
}

var (
	Db *gorm.DB
)

func init() {
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
	_, err = os.Stat("db")
	if err != nil {
		err := os.Mkdir("db", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
	Db, err = gorm.Open(sqlite.Open("db/solar.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to open database")
	}

	err = Db.AutoMigrate(&status{})
	if err != nil {
		log.Println(err)
	}
}

func main() {
	fmt.Println("START")
	tt := time.NewTicker(time.Second * 60)
	for range tt.C {
		getStatus()
	}
}

func getData(rq string) (string, error) {
	config := serial.Config{
		Name:     "/dev/ttyUSB0",
		Baud:     2400,
		Size:     8,
		Parity:   serial.ParityNone,
		StopBits: 1,
	}
	serial, err := serial.OpenPort(&config)
	if err != nil {
		log.Fatalln(err)
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
	return string(dd), nil
}

func getStatus() {

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
	var lastStatus status
	rcv, err := getData("QMOD")
	if err != nil {
		log.Println(err)
		return
	}
	lastStatus.DeviceMode = rcv

	rcv, err = getData("QPIGS")
	if err != nil {
		log.Println(err)
		return
	}
	ss := strings.Split(rcv, " ")
	ACvoltage := ss[0]
	lastStatus.ACvoltage, _ = strconv.ParseFloat(ACvoltage, 32)
	ACfrequency := ss[1]
	lastStatus.ACfrequency, _ = strconv.ParseFloat(ACfrequency, 32)
	Outputvoltage := ss[2]
	lastStatus.OutputVoltage, _ = strconv.ParseFloat(Outputvoltage, 32)
	Outputfrequency := ss[3]
	lastStatus.OutputFrequency, _ = strconv.ParseFloat(Outputfrequency, 32)
	OutputActivePower := ss[4]
	lastStatus.OutputActivePower, _ = strconv.Atoi(OutputActivePower)
	OutputApparntPower := ss[5]
	lastStatus.OutputApparntPower, _ = strconv.Atoi(OutputApparntPower)
	OutputPower := ss[6]
	lastStatus.OutputPower, _ = strconv.Atoi(OutputPower)
	PVBAT := ss[8]
	lastStatus.PVBAT, _ = strconv.ParseFloat(PVBAT, 32)
	ChargingCurrent := ss[9]
	lastStatus.ChargingCurrent, _ = strconv.Atoi(ChargingCurrent)
	PVInputCurrent := ss[12]
	lastStatus.PVinputCurrent, _ = strconv.Atoi(PVInputCurrent)
	PVInputVoltage := ss[13]
	lastStatus.PVinputVoltage, _ = strconv.ParseFloat(PVInputVoltage, 32)
	BatteryVoltage := ss[14]
	lastStatus.BatteryVoltage, _ = strconv.ParseFloat(BatteryVoltage, 32)
	BatteryDischargeCurrent := ss[15]
	lastStatus.BatteryDischargeCurrent, _ = strconv.Atoi(BatteryDischargeCurrent)
	BatteryCapacity := ss[17]
	lastStatus.BatteryCapacity, _ = strconv.Atoi(BatteryCapacity)
	tx := Db.Create(&lastStatus)
	if tx.Error != nil {
		log.Println(tx.Error)
	}
	jb, _ := json.MarshalIndent(lastStatus, "", " ")
	fmt.Print(string(jb))
}
