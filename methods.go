package ilcd

import (
	"encoding/xml"
)

// Method contains the information of an ILCD LCIA method data set.
type Method struct {
	XMLName     xml.Name       `xml:"LCIAMethodDataSet"`
	Info        *MethodInfo    `xml:"LCIAMethodInformation>dataSetInformation"`
	RefQuantity *Ref           `xml:"LCIAMethodInformation>quantitativeReference>referenceQuantity"`
	Factors     []ImpactFactor `xml:"characterisationFactors>factor"`
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
	Flow           Ref     `xml:"referenceToFlowDataSet"`
	Direction      string  `xml:"exchangeDirection"`
	MeanValue      float64 `xml:"meanValue"`
	DataDerivation string  `xml:"dataDerivationTypeStatus"`
}
