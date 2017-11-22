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
	if ref.FlowProperty.UUID != "93a60a56-a3c8-11da-a746-0800200b9a66" {
		t.Fatal("could not get reference flow property")
	}
}
