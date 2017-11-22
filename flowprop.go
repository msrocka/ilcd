package ilcd

import (
	"encoding/xml"
)

// FlowProperty represents an ILCD flow property data set.
type FlowProperty struct {
	XMLName     xml.Name           `xml:"flowPropertyDataSet"`
	Info        *FlowPropertyInfo  `xml:"flowPropertiesInformation>dataSetInformation"`
	UnitGroup   *Ref               `xml:"flowPropertiesInformation>quantitativeReference>referenceToReferenceUnitGroup"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
}

// UUID returns the UUID of the data set.
func (fp *FlowProperty) UUID() string {
	if fp == nil || fp.Info == nil {
		return ""
	}
	return fp.Info.UUID
}

// Version returns the version of the data set.
func (fp *FlowProperty) Version() string {
	if fp == nil || fp.Publication == nil {
		return ""
	}
	return fp.Publication.Version
}

// FlowPropertyInfo contains the general flow property information
type FlowPropertyInfo struct {
	UUID            string           `xml:"UUID"`
	Name            LangString       `xml:"name"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	Comment         LangString       `xml:"generalComment"`
}
