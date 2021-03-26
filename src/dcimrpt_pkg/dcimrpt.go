package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	// Variable definition
	var Server string
	var Report string
	//Defenitiion of List of availabale flags

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	Cptr := flag.Bool("C", false, "cabinet report")
	Dptr := flag.Bool("D", false, "Department report")
	Iptr := flag.Bool("I", false, "Inventory report")
	Pptr := flag.Bool("P", false, "People report")
	Disptr := flag.Bool("disposed", false, "Show disposed HW")
	Xptr := flag.Bool("X", false, "Expired warranty report")
	Sptr := flag.String("S", "dcim.cscs.ch", "server")
	IDptr := flag.String("SID", "", "Specify the InventoryID you want to search for")
	LNameptr := flag.String("LName", "", "Specify the Owner LastName you want to search for")

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

		Report = "people"
		body := DCIM_req(Server, Report)
		reportP := parse_People(body)
		Report = "cabinet"
		body = DCIM_req(Server, Report)
		reportC := parse_Cabinet(body)
		Report = "department"
		body = DCIM_req(Server, Report)
		reportDpt := parse_Dpt(body)
		Report = "device"
		body = DCIM_req(Server, Report)
		reportD := parse_Device(body)
		ReportI(reportC, reportP, reportDpt, reportD, *IDptr, *LNameptr, *Disptr, *Xptr)

	case *Pptr:
		Report = "people"
		body := DCIM_req(Server, Report)
		_ = parse_People(body)

	default:
		fmt.Println("No Report reqiested!")
		os.Exit(0)
	}

}
