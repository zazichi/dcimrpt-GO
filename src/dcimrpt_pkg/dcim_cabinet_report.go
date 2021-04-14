package main

import (
	"encoding/json"
)

type reportC struct {
	Error     bool `json:"error"`
	Errorcode int  `json:"errorcode"`
	Cabinet   []struct {
		Cabinetid        string `json:"CabinetID"`
		Datacenterid     string `json:"DataCenterID"`
		Location         string `json:"Location"`
		Locationsortable string `json:"LocationSortable"`
		Assignedto       string `json:"AssignedTo"`
		Zoneid           string `json:"ZoneID"`
		Cabrowid         string `json:"CabRowID"`
		Cabinetheight    string `json:"CabinetHeight"`
		Model            string `json:"Model"`
		Keylock          string `json:"Keylock"`
		Maxkw            string `json:"MaxKW"`
		Maxweight        string `json:"MaxWeight"`
		Installationdate string `json:"InstallationDate"`
		Mapx1            string `json:"MapX1"`
		Mapy1            string `json:"MapY1"`
		Mapx2            string `json:"MapX2"`
		Mapy2            string `json:"MapY2"`
		Frontedge        string `json:"FrontEdge"`
		Notes            string `json:"Notes"`
		U1Position       string `json:"U1Position"`
		Rights           string `json:"Rights"`
		Datacentername   string `json:"DataCenterName"`
	} `json:"cabinet"`
}

func parse_Cabinet(body []byte) reportC {

	//fmt.Println(string(body))

	// we initialize our Users array
	Creport := reportC{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	_ = json.Unmarshal(body, &Creport)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	/*if len(Creport.Cabinet) > 0 {
		fmt.Println("Cabinet Report")
		fmt.Println("------------------------------------------------")
		fmt.Println("  ID  Location              Model")
		fmt.Println("----  --------------------  --------------------")
	}
	for i := 0; i < len(Creport.Cabinet); i++ {
		fmt.Printf("%4s  %-20s  %-20s\n", Creport.Cabinet[i].Cabinetid, Creport.Cabinet[i].Location, Creport.Cabinet[i].Model)
	}*/
	return Creport
}
