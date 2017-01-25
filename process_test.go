package ilcd

import "testing"

func TestProcessName(t *testing.T) {
	p, err := ReadProcessFile("sample_data/process.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(p.Info.Name.BaseName) != 2 {
		t.Fatal("there should be 2 base names")
	}
}
