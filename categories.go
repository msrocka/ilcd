package ilcd

import (
	"encoding/xml"
)

// CategorySystem contains categories that can be used in the data sets.
type CategorySystem struct {
	XMLName       xml.Name       `xml:"CategorySystem"`
	Name          string         `xml:"name,attr,omitempty"`
	CategoryLists []CategoryList `xml:"categories"`
}

// CategoryList contains the root categories of a category system for a data
// set type.
type CategoryList struct {
	DataType   string     `xml:"dataType,attr,omitempty"`
	Categories []Category `xml:"category"`
}

// Category contains the information of a category in a category system.
type Category struct {
	ID     string     `xml:"id,attr,omitempty"`
	Name   string     `xml:"name,attr,omitempty"`
	Childs []Category `xml:"category"`
}
