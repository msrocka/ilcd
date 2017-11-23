package ilcd

import (
	"encoding/xml"
)

// Method contains the information of an ILCD LCIA method data set.
type Method struct {
	XMLName     xml.Name           `xml:"LCIAMethodDataSet"`
	Info        *MethodInfo        `xml:"LCIAMethodInformation>dataSetInformation"`
	RefQuantity *Ref               `xml:"LCIAMethodInformation>quantitativeReference>referenceQuantity"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
	Factors     []ImpactFactor     `xml:"characterisationFactors>factor"`
}

// UUID returns the UUID of the data set.
func (m *Method) UUID() string {
	if m == nil || m.Info == nil {
		return ""
	}
	return m.Info.UUID
}

// Version returns the version of the data set.
func (m *Method) Version() string {
	if m == nil || m.Publication == nil {
		return ""
	}
	return m.Publication.Version
}

// MethodInfo :<dataSetInformation>
type MethodInfo struct {
	UUID            string     `xml:"UUID"`
	Name            LangString `xml:"name"`
	Methodology     string     `xml:"methodology"`
	ImpactCategory  string     `xml:"impactCategory"`
	ImpactIndicator string     `xml:"impactIndicator"`
	Comment         LangString `xml:"generalComment"`
	ExternalDocs    []Ref      `xml:"referenceToExternalDocumentation"`
}

// ImpactFactor :<characterisationFactors/factor>
type ImpactFactor struct {
	Flow           *Ref    `xml:"referenceToFlowDataSet"`
	Direction      string  `xml:"exchangeDirection"`
	MeanValue      float64 `xml:"meanValue"`
	DataDerivation string  `xml:"dataDerivationTypeStatus"`
	Location       string  `xml:"location"`
}
