package main

import (
	"encoding/json"
)

type reportDpt struct {
	Error      bool `json:"error"`
	Errorcode  int  `json:"errorcode"`
	Department []struct {
		Deptid         string `json:"DeptID"`
		Name           string `json:"Name"`
		Execsponsor    string `json:"ExecSponsor"`
		Sdm            string `json:"SDM"`
		Classification string `json:"Classification"`
		Deptcolor      string `json:"DeptColor"`
	} `json:"department"`
}

func parse_Dpt(body []byte) reportDpt {

	//fmt.Println(string(body))

	// we initialize our Users array
	Dptreport := reportDpt{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	_ = json.Unmarshal(body, &Dptreport)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	/*if len(Dptreport.Department) > 0 {
		fmt.Println("Department Report")
		fmt.Println("--------------------------------------------------------------------------------------------------------------")
		fmt.Println(" ID  Name                       Service Delivery Manager             Executive Sponsor          Classification")
		fmt.Println("---  -------------------------  -----------------------------------  -------------------------  --------------")
	}*/
	//for i := 0; i < len(Dptreport.Department); i++ {
	//fmt.Printf("%3s  %-25.25s  %-35.35s  %-25.25s  %s\n", Dptreport.Department[i].Deptid, Dptreport.Department[i].Name, Dptreport.Department[i].Sdm, Dptreport.Department[i].Execsponsor, Dptreport.Department[i].Classification)
	//}
	return Dptreport
}
