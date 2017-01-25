package ilcd

import "testing"

func TestProcessInfo(t *testing.T) {
	p, _ := ReadProcessFile("sample_data/process.xml")
	if p.Info.Classifications[1].Classes[1].Name != "class3" {
		t.Fatal("Could not get expected class name")
	}
	if p.Info.UUID != "00000000-0000-0000-0000-000000000000" {
		t.Fatal("uuid != 00000000-0000-0000-0000-000000000000")
	}
}

func TestProcessName(t *testing.T) {
	p, _ := ReadProcessFile("sample_data/process.xml")
	if len(p.Info.Name.BaseName) != 2 {
		t.Fatal("there should be 2 base names")
	}
	if p.Info.Name.BaseName.Get("en") != "baseName0" {
		t.Fatal("baseName@en != 'baseName0'")
	}
	if p.Info.Name.BaseName.Get("de") != "baseName1" {
		t.Fatal("baseName@de != baseName1")
	}
}

func TestProcessSynonyms(t *testing.T) {
	p, _ := ReadProcessFile("sample_data/process.xml")
	synonyms := p.Info.Synonyms
	if len(synonyms) != 2 {
		t.Fatal("there should be 2 synonyms")
	}
	if synonyms.Get("en") != "synonyms0" {
		t.Fatal("synonyms@en != 'synonyms0'")
	}
	if synonyms.Get("de") != "synonyms1" {
		t.Fatal("synonyms@de != 'synonyms1'")
	}
}
