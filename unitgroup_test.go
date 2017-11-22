package ilcd

import "testing"

func TestUnitGroupInfo(t *testing.T) {
	ug, _ := ReadUnitGroupFile("sample_data/unitgroup.xml")
	if ug.Info.Name.Get("en") != "Unit of mass" {
		t.Fatal("wrong name")
	}
}

func TestRefUnit(t *testing.T) {
	ug, _ := ReadUnitGroupFile("sample_data/unitgroup.xml")
	if len(ug.Units) < 5 {
		t.Fatal("failed to get units")
	}
	if ug.ReferenceUnit().Name != "kg" {
		t.Fatal("wrong reference unit")
	}
}
