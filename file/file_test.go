package file

import "testing"

var ntpYaml = `
ntp:
  source_interface: Vlan10
  ntp_servers:
    - ipv4_host: 10.0.0.106
      priority: false
    - ipv4_host: 10.0.0.200
      priority: true
`

var ntpJson = "{\"ntp\":{\"source_interface\":\"Vlan10\",\"ntp_servers\":[{\"ipv_4_host\":\"10.0.0.106\",\"priority\":false},{\"ipv_4_host\":\"10.0.0.200\",\"priority\":true}]}}\n"

var ntpXml = "<NTP><ntp><source_interface>Vlan10</source_interface><ntp_servers><ipv4_host>10.0.0.106</ipv4_host><priority>false</priority></ntp_servers><ntp_servers><ipv4_host>10.0.0.200</ipv4_host><priority>true</priority></ntp_servers></ntp></NTP>"

func TestReadWriteYAMLFile(t *testing.T) {
	tempDirectory := t.TempDir() + "temp.yml"
	err := WriteYAMLFile(tempDirectory, ntpYaml)

	if err != nil {
		t.Errorf("ERROR: WriteYAMLFile %v", tempDirectory)
	}

	data, err := ReadYAMLFile(tempDirectory)

	if err != nil {
		t.Errorf("ERROR: ReadYAMLFile %v", err)
	}

	if data != ntpYaml {
		t.Errorf("FAILED: ReadYAMLFile expected %v received %v", data, ntpYaml)
	} else {
		t.Logf("PASSED: ReadYAMLFile using tempDirectory %v", tempDirectory)
	}

}

func TestReadYAMLFileBad(t *testing.T) {
	var scenarios = []struct {
		name string
		file string
	}{
		{
			name: "TEST Bad File Name",
			file: "./fake.txt",
		},
		{
			name: "TEST Bad File Location Good Name .yml",
			file: "./fake.yml",
		},
		{
			name: "TEST Bad File Location Good Name .yaml",
			file: "fake.yaml",
		},
	}

	for _, scenario := range scenarios {
		_, err := ReadYAMLFile(scenario.file)

		if err == nil {
			t.Errorf("FAILED: %v ReadYAMLFile error %v", scenario.name, err)
		} else {
			t.Logf("PASSED: %v ReadYAMLFile error %v", scenario.name, err)
		}
	}

}

func TestWriteYAMLFileBad(t *testing.T) {
	var scenarios = []struct {
		name string
		file string
	}{
		{
			name: "TEST Bad File Name",
			file: "./fake.txt",
		},
		{
			name: "TEST Bad File Location Good Name .yml",
			file: "./moo/fake.yml",
		},
		{
			name: "TEST Bad File Location Good Name .yaml",
			file: "./moo/fake.yaml",
		},
	}

	for _, scenario := range scenarios {
		err := WriteYAMLFile(scenario.file, ntpYaml)

		if err == nil {
			t.Errorf("FAILED: %v WriteYAMLFile error %v", scenario.name, err)
		} else {
			t.Logf("PASSED: %v WriteYAMLFile error %v", scenario.name, err)
		}
	}
}

func TestReadWriteJSONFile(t *testing.T) {
	tempDirectory := t.TempDir() + "temp.json"
	err := WriteJSONFile(tempDirectory, ntpJson)

	if err != nil {
		t.Errorf("ERROR: WriteJSONFile %v", tempDirectory)
	}

	data, err := ReadJSONFile(tempDirectory)

	if err != nil {
		t.Errorf("ERROR: ReadJSONFile %v", err)
	}

	if data != ntpJson {
		t.Errorf("FAILED: ReadJSONFile expected %v received %v", data, ntpYaml)
	} else {
		t.Logf("PASSED: ReadJSONFile using tempDirectory %v", tempDirectory)
	}

}

func TestReadJSONFileBad(t *testing.T) {
	var scenarios = []struct {
		name string
		file string
	}{
		{
			name: "TEST Bad File Name",
			file: "./fake.txt",
		},
		{
			name: "TEST Bad File Location Good Name .json",
			file: "./fake.json",
		},
	}

	for _, scenario := range scenarios {
		_, err := ReadJSONFile(scenario.file)

		if err == nil {
			t.Errorf("FAILED: %v ReadJSONFile error %v", scenario.name, err)
		} else {
			t.Logf("PASSED: %v ReadJSONFile error %v", scenario.name, err)
		}
	}

}

func TestWriteJSONFileBad(t *testing.T) {
	var scenarios = []struct {
		name string
		file string
	}{
		{
			name: "TEST Bad File Name",
			file: "./fake.txt",
		},
		{
			name: "TEST Bad File Location Good Name .json",
			file: "./moo/fake.json",
		},
	}

	for _, scenario := range scenarios {
		err := WriteJSONFile(scenario.file, ntpJson)

		if err == nil {
			t.Errorf("FAILED: %v WriteJSONFile error %v", scenario.name, err)
		} else {
			t.Logf("PASSED: %v WriteJSONFile error %v", scenario.name, err)
		}
	}
}

func TestReadWriteXMLFile(t *testing.T) {
	tempDirectory := t.TempDir() + "temp.xml"
	err := WriteXMLFile(tempDirectory, ntpXml)

	if err != nil {
		t.Errorf("ERROR: WriteXMLFile %v", tempDirectory)
	}

	data, err := ReadXMLFile(tempDirectory)

	if err != nil {
		t.Errorf("ERROR: ReadXMLFile %v", err)
	}

	if data != ntpXml {
		t.Errorf("FAILED: ReadXMLFile expected %v received %v", data, ntpXml)
	} else {
		t.Logf("PASSED: ReadXMLFile using tempDirectory %v", tempDirectory)
	}
}

func TestReadXMLFileBad(t *testing.T) {
	var scenarios = []struct {
		name string
		file string
	}{
		{
			name: "TEST Bad File Name",
			file: "./fake.txt",
		},
		{
			name: "TEST Bad File Location Good Name .xml",
			file: "./fake.xml",
		},
	}

	for _, scenario := range scenarios {
		_, err := ReadXMLFile(scenario.file)

		if err == nil {
			t.Errorf("FAILED: %v ReadXMLFile error %v", scenario.name, err)
		} else {
			t.Logf("PASSED: %v ReadXMLFile error %v", scenario.name, err)
		}
	}
}

func TestWriteXMLFileBad(t *testing.T) {
	var scenarios = []struct {
		name string
		file string
	}{
		{
			name: "TEST Bad File Name",
			file: "./fake.txt",
		},
		{
			name: "TEST Bad File Location Good Name .xml",
			file: "./moo/fake.xml",
		},
	}

	for _, scenario := range scenarios {
		err := WriteXMLFile(scenario.file, ntpXml)

		if err == nil {
			t.Errorf("FAILED: %v WriteXMLFile error %v", scenario.name, err)
		} else {
			t.Logf("PASSED: %v WriteXMLFile error %v", scenario.name, err)
		}
	}
}
