package ilcd

import "encoding/xml"

type EpdInfoExt struct {
	XMLName       xml.Name          `xml:"http://lca.jrc.it/ILCD/Common other"`
	ScenarioList  *EpdScenarioList  `xml:"http://www.iai.kit.edu/EPD/2013 scenarios"`
	SafetyMargins *EpdSafetyMargins `xml:"http://www.iai.kit.edu/EPD/2013 safetyMargins"`
}

type EpdSafetyMargins struct {
	Value       *float64   `xml:"margins"`
	Description LangString `xml:"description"`
}

// EpdScenarioList is just a wrapper for the list of scenarios. We need to
// add this wrapper so that the namespace serialization is correct.
type EpdScenarioList struct {
	Entries []EpdScenario `xml:"scenario"`
}

type EpdScenario struct {
	Name        string     `xml:"name,attr,omitempty"`
	IsDefault   bool       `xml:"default,attr"`
	Group       string     `xml:"group,attr,omitempty"`
	Description LangString `xml:"description"`
}

type EpdTimeExt struct {
	PublicationDate string `xml:"http://www.indata.network/EPD/2019 publicationDateOfEPD"`
}

type EpdInventoryMethodExt struct {
	EpdType string `xml:"http://www.iai.kit.edu/EPD/2013 subType"`
}

type EpdSourcesExt struct {
	OriginalEpds []Ref `xml:"referenceToOriginalEPD"`
}

type EpdResultExt struct {
	Values    []EpdValue `xml:"http://www.iai.kit.edu/EPD/2013 amount"`
	UnitGroup *Ref       `xml:"http://www.iai.kit.edu/EPD/2013 referenceToUnitGroupDataSet"`
}

type EpdValue struct {
	Amount   *float64 `xml:",chardata"`
	Module   string   `xml:"http://www.iai.kit.edu/EPD/2013 module,attr,omitempty"`
	Scenario string   `xml:"http://www.iai.kit.edu/EPD/2013 scenario,attr,omitempty"`
}

type EpdPublicationExt struct {
	Publishers []Ref `xml:"http://www.indata.network/EPD/2019 referenceToPublisher"`
}
