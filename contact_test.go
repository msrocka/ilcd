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

func TestContactTime(t *testing.T) {
	c, _ := ReadContactFile("sample_data/contact.xml")
	if c.DataEntry.TimeStamp != "2012-01-04T15:42:24.609+01:00" {
		t.Fatal("wrong time")
	}
}

func TestEmptyName(t *testing.T) {
	p := &Process{Info: &ProcessInfo{Name: &ProcessName{}}}
	name := p.Info.Name.BaseName.Default()
	if name != "" {
		t.Fatal("empty name expected")
	}
}
