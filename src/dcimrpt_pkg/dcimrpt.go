package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	//var wg sync.WaitGroup

	// Variable definition
	var Server string
	var Report string
	//Defenitiion of List of availabale flags

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s [-C|-D|-P|-I] [option of search in case of -I: LName|Label|SID] [optoinal -X and/or -disposed] [optional -S]:\n", os.Args[0])
		flag.PrintDefaults()
	}

	Cptr := flag.Bool("C", false, "cabinet report")
	Dptr := flag.Bool("D", false, "Department report")
	Iptr := flag.Bool("I", false, "Inventory report")
	Pptr := flag.Bool("P", false, "People report")
	Disptr := flag.Bool("disposed", false, "Show disposed HW")
	Xptr := flag.Bool("X", false, "Expired warranty report")
	Sptr := flag.String("S", "dcim.cscs.ch", "server")
	IDptr := flag.String("SID", "", "Specify the InventoryID you want to search for")
	LNameptr := flag.String("LName", "", "Specify the Owner LastName you want to search for")
	Label := flag.String("Label", "", "Specify the Label you want to search for")
	Serial := flag.String("Serial", "", "Specify the Serial number you want to search for")

	//ProgName := os.Args[0]
	//fmt.Printf("Progam Name: %s\n", ProgName)

	//Parsing the flags

	flag.Parse()

	//Server definition
	Server = *Sptr
	//fmt.Println(Server)

	//Analize which flags report has been requested

	switch {
	case *Cptr:
		Report = "cabinet"
		body := DCIM_req(Server, Report)
		_ = parse_Cabinet(body)
	case *Dptr:
		Report = "department"
		body := DCIM_req(Server, Report)
		_ = parse_Dpt(body)
	case *Iptr:

		var body1 []byte
		var body2 []byte
		var body3 []byte
		var body4 []byte
		var reportP reportP
		var reportC reportC
		var reportD reportD
		var reportDpt reportDpt

		Report1 := "people"
		Report2 := "cabinet"
		Report3 := "department"
		Report4 := "device"
		//wg.Add(4)
		//go func() { body1 = DCIM_req(Server, Report1) }()
		body1 = DCIM_req(Server, Report1)
		body2 = DCIM_req(Server, Report2)
		body3 = DCIM_req(Server, Report3)
		body4 = DCIM_req(Server, Report4)
		//wg.Wait()
		reportP = parse_People(body1)
		reportC = parse_Cabinet(body2)
		reportDpt = parse_Dpt(body3)
		reportD = parse_Device(body4)

		ReportI(reportC, reportP, reportDpt, reportD, *IDptr, *LNameptr, *Disptr, *Xptr, *Label, *Serial)

	case *Pptr:
		Report = "people"
		body := DCIM_req(Server, Report)
		_ = parse_People(body)

	default:
		fmt.Println("No Report reqiested!")
		os.Exit(0)
	}

}
