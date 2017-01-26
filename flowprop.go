package ilcd

import (
	"encoding/xml"
)

// FlowProperty represents an ILCD flow property data set.
type FlowProperty struct {
	XMLName   xml.Name          `xml:"flowPropertyDataSet"`
	Info      *FlowPropertyInfo `xml:"flowPropertiesInformation>dataSetInformation"`
	UnitGroup *Ref              `xml:"flowPropertiesInformation>quantitativeReference>referenceToReferenceUnitGroup"`
}

// FlowPropertyInfo contains the general flow property information
type FlowPropertyInfo struct {
	UUID            string           `xml:"UUID"`
	Name            LangString       `xml:"name"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	Comment         LangString       `xml:"generalComment"`
}
