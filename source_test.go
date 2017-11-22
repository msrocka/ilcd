package ilcd

import "testing"

func TestSourceInfo(t *testing.T) {
	s, _ := ReadSourceFile("sample_data/source.xml")
	if s.Info.ShortName.Get("en") != "blank.JPG" {
		t.Fatal("wrong name")
	}
	if s.Info.Citation != "GaBi database" {
		t.Fatal("wrong citation")
	}
}
