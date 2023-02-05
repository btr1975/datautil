package yaml

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

var ntpYaml = `
ntp:
  source_interface: Vlan10
  ntp_servers:
    - ipv4_host: 10.0.0.106
      priority: false
    - ipv4_host: 10.0.0.200
      priority: true
`

func TestYAMLToStruct(t *testing.T) {
	var tempVar ntp

	err := YAMLToStruct(&tempVar, ntpYaml)

	if err != nil {
		t.Errorf("FAILED: YAMLToStruct error %v", err)
	} else {
		t.Logf("PASSED: YAMLToStruct error %v", err)

	}
}

func TestYAMLToStructBad(t *testing.T) {
	var tempVar ntp

	err := YAMLToStruct(tempVar, ntpYaml)

	if err == nil {
		t.Errorf("FAILED: YAMLToStruct error %v", err)
	} else {
		t.Logf("PASSED: YAMLToStruct error %v", err)

	}
}

func TestStructToYAML(t *testing.T) {
	yamlData, err := StructToYAML(ntp{})

	if err != nil {
		t.Errorf("FAILED: TestStructToYAML error %v", err)
	} else {
		t.Logf("PASSED: TestStructToYAML %v", yamlData)

	}
}
