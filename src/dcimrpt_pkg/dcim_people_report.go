package main

import (
	"encoding/json"
)

type reportP struct {
	Error     bool `json:"error"`
	Errorcode int  `json:"errorcode"`
	People    []struct {
		Personid  int    `json:"PersonID"`
		Firstname string `json:"FirstName"`
		Lastname  string `json:"LastName"`
	} `json:"people"`
}

func parse_People(body []byte) reportP {

	//fmt.Println(string(body))

	// we initialize our Users array
	Preport := reportP{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	_ = json.Unmarshal(body, &Preport)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	/*if len(Preport.People) > 0 {
		fmt.Println("People Report")
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("ID    Name                      Email")
		fmt.Println("---  -------------------------  ------------------------------------")
	}
	for i := 0; i < len(Preport.People); i++ {
		name := Preport.People[i].Firstname + " " + Preport.People[i].Lastname
		email := "email-not-present"
		fmt.Printf("%3d  %-25s  %s\n", Preport.People[i].Personid, name, email)
		//fmt.Printf("User ID: %d", Preport.People[i].Personid)
		//fmt.Println("User First NAme: " + Preport.People[i].Firstname)
		//fmt.Println("User Last Name: " + Preport.People[i].Lastname)
	}*/
	return Preport
}
