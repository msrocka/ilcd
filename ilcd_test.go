package ilcd

import "testing"

func TestCommonDataSetFields(t *testing.T) {
	dataSets := [7]DataSet{}
	dataSets[0], _ = ReadContactFile("sample_data/contact.xml")
	dataSets[1], _ = ReadFlowFile("sample_data/flow.xml")
	dataSets[2], _ = ReadFlowPropertyFile("sample_data/flowprop.xml")
	dataSets[3], _ = ReadMethodFile("sample_data/method.xml")
	dataSets[4], _ = ReadProcessFile("sample_data/process.xml")
	dataSets[5], _ = ReadSourceFile("sample_data/source.xml")
	dataSets[6], _ = ReadUnitGroupFile("sample_data/unitgroup.xml")
	for _, ds := range dataSets {
		if ds.UUID() == "" || ds.Version() == "" {
			t.Fatal("UUID or version missing in", ds)
		}
	}
}
