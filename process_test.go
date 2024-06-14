package ilcd

import "testing"

func TestProcessInfo(t *testing.T) {
	p, _ := ReadProcessFile("sample_data/process.xml")
	if p.Info.Classifications[0].Classes[1].Name != "ELCD" {
		t.Fatal("Could not get expected class name")
	}
	if p.Info.UUID != "c93541fe-0b28-40b8-a890-9948e9f1d41f" {
		t.Fatal("uuid != c93541fe-0b28-40b8-a890-9948e9f1d41f")
	}
}

func TestProcessName(t *testing.T) {
	p, _ := ReadProcessFile("sample_data/process.xml")
	if p.Info.Name.BaseName.Get("en") != "Electricity grid mix 1kV-60kV" {
		t.Fatal("baseName@en != 'Electricity grid mix 1kV-60kV'")
	}
}

func TestProcessSynonyms(t *testing.T) {
	p, _ := ReadProcessFile("sample_data/process.xml")
	synonyms := p.Info.Synonyms
	if synonyms.Get("en") != "Power grid mix" {
		t.Fatal("synonyms@en != 'Power grid mix'")
	}
}

func TestProcessRefFlows(t *testing.T) {
	p, _ := ReadProcessFile("sample_data/process.xml")
	refFlows := p.RefFlows()
	if len(refFlows) != 1 {
		t.Fatal("expected exactly one reference flow")
	}
	if refFlows[0].Flow.Name.Get("en") != "Electricity" {
		t.Fatal("Could not find reference flow")
	}
}

func TestProcessParameters(t *testing.T) {
	p, _ := ReadProcessFile("sample_data/process.xml")
	params := p.MathModel.Parameters
	if len(params) != 1 {
		t.Fatal("There should be a parameter in the example process")
	}
	if params[0].Name != "distance" {
		t.Fatal("The parameter name should be 'distance'")
	}
}
