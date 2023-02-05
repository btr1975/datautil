package json

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
		SourceInterface string `yaml:"source_interface" json:"source_interface" xml:"source_interface"`
		NTPServerKey    []struct {
			IPv4Host string `yaml:"ipv4_host" json:"ipv4_host" xml:"ipv4_host"`
			Priority bool   `yaml:"priority" json:"priority" xml:"priority"`
		} `yaml:"ntp_servers" json:"ntp_servers" xml:"ntp_servers"`
	} `yaml:"ntp" json:"ntp" xml:"ntp"`
}

var ntpJson = "{\"ntp\":{\"source_interface\":\"Vlan10\",\"ntp_servers\":[{\"ipv_4_host\":\"10.0.0.106\",\"priority\":false},{\"ipv_4_host\":\"10.0.0.200\",\"priority\":true}]}}\n"

func TestJSONToStruct(t *testing.T) {
	var tempVar ntp

	err := JSONToStruct(&tempVar, ntpJson)

	if err != nil {
		t.Errorf("FAILED: YAMLToStruct error %v", err)
	} else {
		t.Logf("PASSED: YAMLToStruct error %v", err)

	}
}

func TestJSONToStructBad(t *testing.T) {
	var tempVar ntp

	err := JSONToStruct(tempVar, ntpJson)

	if err == nil {
		t.Errorf("FAILED: YAMLToStruct error %v", err)
	} else {
		t.Logf("PASSED: YAMLToStruct error %v", err)

	}
}

func TestStructToJSON(t *testing.T) {
	jsonData, err := StructToJSON(ntp{})

	if err != nil {
		t.Errorf("FAILED: StructToJSON error %v", err)
	} else {
		t.Logf("PASSED: StructToJSON %v", jsonData)

	}
}
