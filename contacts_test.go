package ilcd

import (
	"testing"
)

func TestContactName(t *testing.T) {
	c, _ := ReadContactFile("sample_data/contact.xml")
	if c.Info.ShortName.Get("en") != "JRC" {
		t.Fatal("wrong short name")
	}
}

func TestContactVersion(t *testing.T) {
	c, _ := ReadContactFile("sample_data/contact.xml")
	if c.Publication.Version != "01.01.000" {
		t.Fatal("wrong version")
	}
}
