package ilcd

import "testing"

func TestRefFlowProperty(t *testing.T) {
	f, _ := ReadFlowFile("sample_data/flow.xml")
	ref := f.ReferenceFlowProperty()
	if ref == nil {
		t.Fatal("could not get reference flow property")
	}
	if ref.FlowProperty.Type != "flow property data set" {
		t.Fatal("could not get reference flow property")
	}
	if ref.FlowProperty.UUID != "00000000-0000-0000-0000-000000000000" {
		t.Fatal("could not get reference flow property")
	}
}
