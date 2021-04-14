package main

import (
	"encoding/json"
)

type reportD struct {
	Error     bool `json:"error"`
	Errorcode int  `json:"errorcode"`
	Device    []struct {
		Deviceid         int         `json:"DeviceID"`
		Label            string      `json:"Label"`
		Serialno         string      `json:"SerialNo"`
		Assettag         string      `json:"AssetTag"`
		Primaryip        string      `json:"PrimaryIP"`
		Snmpversion      string      `json:"SNMPVersion"`
		V3Securitylevel  string      `json:"v3SecurityLevel"`
		V3Authprotocol   string      `json:"v3AuthProtocol"`
		V3Authpassphrase string      `json:"v3AuthPassphrase"`
		V3Privprotocol   string      `json:"v3PrivProtocol"`
		V3Privpassphrase string      `json:"v3PrivPassphrase"`
		Snmpcommunity    string      `json:"SNMPCommunity"`
		Snmpfailurecount int         `json:"SNMPFailureCount"`
		Hypervisor       string      `json:"Hypervisor"`
		Apiusername      string      `json:"APIUsername"`
		Apipassword      string      `json:"APIPassword"`
		Apiport          int         `json:"APIPort"`
		Proxmoxrealm     string      `json:"ProxMoxRealm"`
		Owner            int         `json:"Owner"`
		Escalationtimeid int         `json:"EscalationTimeID"`
		Escalationid     int         `json:"EscalationID"`
		Primarycontact   int         `json:"PrimaryContact"`
		Cabinet          int         `json:"Cabinet"`
		Position         int         `json:"Position"`
		Height           int         `json:"Height"`
		Ports            int         `json:"Ports"`
		Firstportnum     int         `json:"FirstPortNum"`
		Templateid       int         `json:"TemplateID"`
		Nominalwatts     int         `json:"NominalWatts"`
		Powersupplycount int         `json:"PowerSupplyCount"`
		Devicetype       string      `json:"DeviceType"`
		Chassisslots     int         `json:"ChassisSlots"`
		Rearchassisslots int         `json:"RearChassisSlots"`
		Parentdevice     int         `json:"ParentDevice"`
		Mfgdate          string      `json:"MfgDate"`
		Installdate      string      `json:"InstallDate"`
		Warrantyco       string      `json:"WarrantyCo"`
		Warrantyexpire   string      `json:"WarrantyExpire"`
		Notes            string      `json:"Notes"`
		Status           string      `json:"Status"`
		Rights           string      `json:"Rights"`
		Halfdepth        int         `json:"HalfDepth"`
		Backside         int         `json:"BackSide"`
		Auditstamp       string      `json:"AuditStamp"`
		Customvalues     interface{} `json:"CustomValues"`
		Weight           int         `json:"Weight"`
		CscsInventarNr   string      `json:"CSCS_inventar_nr"`
	} `json:"device"`
}

func parse_Device(body []byte) reportD {

	//fmt.Println(string(body))

	// we initialize our Users array
	Dreport := reportD{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	_ = json.Unmarshal(body, &Dreport)

	return Dreport
}
