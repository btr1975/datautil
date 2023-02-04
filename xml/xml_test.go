package xml

import "testing"

// ntp is the full NTP YAML schema
//
//	Example YAML Data with source interface
//
//	ntp_servers:
//	  source_interface: Vlan10
//	  ntp_servers:
//	  - ipv4_host: 10.0.0.106
//	  - ipv4_host: 10.0.0.200
//	    priority: true
//	  - ipv4_host: 10.0.0.205
//	    priority: false
//
//	Example YAML Data without source interface
//
//	ntp_servers:
//	  ntp_servers:
//	  - ipv4_host: 10.0.0.106
type ntp struct {
	NTPServersData struct {
		SourceInterface string `yaml,json:"source_interface" xml:"source_interface"`
		NTPServerKey    []struct {
			IPv4Host string `yaml,json:"ipv4_host" xml:"ipv4_host"`
			Priority bool   `yaml,json:"priority" xml:"priority"`
		} `yaml,json:"ntp_servers" xml:"ntp_servers"`
	} `yaml,json:"ntp" xml:"ntp"`
}

var ntpXml = "<NTP><ntp><source_interface>Vlan10</source_interface><ntp_servers><ipv4_host>10.0.0.106</ipv4_host><priority>false</priority></ntp_servers><ntp_servers><ipv4_host>10.0.0.200</ipv4_host><priority>true</priority></ntp_servers></ntp></NTP>"

func TestXMLToStruct(t *testing.T) {
	structData, err := XMLToStruct(ntp{}, ntpXml)

	if err != nil {
		t.Errorf("FAILED: XMLToStruct error %v", err)
	} else {
		t.Logf("PASSED: XMLToStruct %v", structData)

	}
}

func TestStructToXML(t *testing.T) {
	xmlData, err := StructToXML(ntp{})

	if err != nil {
		t.Errorf("FAILED: StructToXML error %v", err)
	} else {
		t.Logf("PASSED: StructToXML %v", xmlData)

	}
}
