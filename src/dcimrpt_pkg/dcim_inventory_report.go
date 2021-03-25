package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type reportS struct {
	Deviceid       int
	Label          string
	Serialno       string
	Owner          string
	Primarycontact string
	Cabinet        int
	Position       int
	Height         int
	Installdate    string
	Warrantyco     string
	Warrantyexpire string
	Status         string
	CscsInventarNr string
}

func ReportI(Creport reportC, Preport reportP, Dptreport reportDpt, Dreport reportD, Inv_ID string, Name string, D bool, X bool) {

	var Sreport = make([]reportS, 0)
	var Sreport_temp reportS
	var ID_to_check int
	var name_to_check string

	today := time.Now()
	//fmt.Printf("Time %s", time.Now())
	layout := "2006-01-02"

	if Inv_ID != "" {
		i := 0
		for v := 0; v < len(Dreport.Device); v++ {
			if Dreport.Device[v].CscsInventarNr == Inv_ID {

				Sreport_temp.Deviceid = Dreport.Device[v].Deviceid
				Sreport_temp.Label = Dreport.Device[v].Label
				Sreport_temp.Serialno = Dreport.Device[v].Serialno
				for z := 0; z < len(Preport.People); z++ {
					//Convert the Owner and Primary contact ID in Names
					if Preport.People[z].Personid == Dreport.Device[v].Owner {
						Sreport_temp.Owner = Preport.People[z].Firstname + " " + Preport.People[z].Lastname
					}
					if Preport.People[z].Personid == Dreport.Device[v].Primarycontact {
						Sreport_temp.Primarycontact = Preport.People[z].Firstname + " " + Preport.People[z].Lastname
					}
				}
				Sreport_temp.Cabinet = Dreport.Device[v].Cabinet
				Sreport_temp.Position = Dreport.Device[v].Position
				Sreport_temp.Height = Dreport.Device[v].Height
				Sreport_temp.Installdate = Dreport.Device[v].Installdate
				Sreport_temp.Warrantyco = Dreport.Device[v].Warrantyco
				Sreport_temp.Warrantyexpire = Dreport.Device[v].Warrantyexpire
				Sreport_temp.Status = Dreport.Device[v].Status
				Sreport_temp.CscsInventarNr = Dreport.Device[v].CscsInventarNr
				if i == len(Sreport) {
					if Sreport_temp.Status != "Disposed" || D {
						//Add the entry to the output Slice
						Sreport = append(Sreport, Sreport_temp)
						i++
					}
				}
			}
		}
	}
	if Name != "" {
		i := 0
		for z := 0; z < len(Preport.People); z++ {
			name_to_check = Preport.People[z].Firstname + " " + Preport.People[z].Lastname
			//Checking the string all uppercase
			if strings.Contains(strings.ToUpper(name_to_check), strings.ToUpper(Name)) {
				ID_to_check = Preport.People[z].Personid
				//fmt.Printf("ID %d\n", ID_to_check)
			}
		}
		for v := 0; v < len(Dreport.Device); v++ {
			if Dreport.Device[v].Primarycontact == ID_to_check {
				log.Println("Primary contact Name matched")
				Sreport_temp.Deviceid = Dreport.Device[v].Deviceid
				Sreport_temp.Label = Dreport.Device[v].Label
				Sreport_temp.Serialno = Dreport.Device[v].Serialno
				for z := 0; z < len(Preport.People); z++ {
					//Convert the Owner and Primary contact ID in Names
					if Preport.People[z].Personid == Dreport.Device[v].Owner {
						Sreport_temp.Owner = Preport.People[z].Firstname + " " + Preport.People[z].Lastname
					}
					if Preport.People[z].Personid == Dreport.Device[v].Primarycontact {
						Sreport_temp.Primarycontact = Preport.People[z].Firstname + " " + Preport.People[z].Lastname
					}
				}
				Sreport_temp.Cabinet = Dreport.Device[v].Cabinet
				Sreport_temp.Position = Dreport.Device[v].Position
				Sreport_temp.Height = Dreport.Device[v].Height
				Sreport_temp.Installdate = Dreport.Device[v].Installdate
				Sreport_temp.Warrantyco = Dreport.Device[v].Warrantyco
				Sreport_temp.Warrantyexpire = Dreport.Device[v].Warrantyexpire
				Sreport_temp.Status = Dreport.Device[v].Status
				Sreport_temp.CscsInventarNr = Dreport.Device[v].CscsInventarNr
				log.Printf("len Srep %d", len(Sreport))
				if i == len(Sreport) {
					if Sreport_temp.Status != "Disposed" || D {
						//Add the entry to the output Slice
						Sreport = append(Sreport, Sreport_temp)
						i++
					}
				}
			}
		}
	}

	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("  ID  Label                                Serial Number              Inventory # Installed   Warranty    Contact                    Department            Location              Position  Status")
	fmt.Println("----  -----------------------------------  -------------------------  ----------- ----------  ----------  -------------------------  --------------------  --------------------  --------  --------")
	for i := 0; i < len(Sreport); i++ {
		//log.Printf("Date warranty %s %s", today, layout)
		data, _ := time.Parse(layout, Sreport[i].Warrantyexpire)
		//log.Printf("Date warranty %s", Sreport[i].Warrantyexpire)
		if today.Before(data) {
			log.Print("Printing a line in Warranty")
			fmt.Printf("%4d  %-35s  %-25.25s  %-11.11s %-10s  %-10s  %-25s  %-20s  %-20d   P%2.2dH%2.2d   %s\n",
				Sreport[i].Deviceid, Sreport[i].Label, Sreport[i].Serialno, Sreport[i].CscsInventarNr, Sreport[i].Installdate,
				Sreport[i].Warrantyexpire, Sreport[i].Primarycontact, Sreport[i].Owner,
				Sreport[i].Cabinet, Sreport[i].Position, Sreport[i].Height, Sreport[i].Status)

		} else if X {
			log.Print("Printing a line out Warranty")
			fmt.Printf("%4d  %-35s  %-25.25s  %-11.11s %-10s  %-10s  %-25s  %-20s  %-20d   P%2.2dH%2.2d   %s\n",
				Sreport[i].Deviceid, Sreport[i].Label, Sreport[i].Serialno, Sreport[i].CscsInventarNr, Sreport[i].Installdate,
				Sreport[i].Warrantyexpire, Sreport[i].Primarycontact, Sreport[i].Owner,
				Sreport[i].Cabinet, Sreport[i].Position, Sreport[i].Height, Sreport[i].Status)

		} else {
			log.Print("printing no line")
			log.Printf("Date warranty %s", data)
		}
		/*
			fmt.Printf("%4d  %-35s  %-25.25s  %-11.11s %-10s  %-10s  %-25s  %-20s  %-20d   P%2.2dH%2.2d   %s\n",
				Sreport[i].Deviceid, Sreport[i].Label, Sreport[i].Serialno, Sreport[i].CscsInventarNr, Sreport[i].Installdate,
				Sreport[i].Warrantyexpire, Sreport[i].Primarycontact, Sreport[i].Owner,
				Sreport[i].Cabinet, Sreport[i].Position, Sreport[i].Height, Sreport[i].Status)
		*/
	}
}
