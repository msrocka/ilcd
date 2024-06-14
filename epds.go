package ilcd

import "encoding/xml"

type EpdInfoExtension struct {
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

type EpdResultExtension struct {
	Values    []EpdValue `xml:"http://www.iai.kit.edu/EPD/2013 amount"`
	UnitGroup *Ref       `xml:"http://www.iai.kit.edu/EPD/2013 referenceToUnitGroupDataSet"`
}

type EpdValue struct {
	Amount   *float64 `xml:",chardata"`
	Module   string   `xml:"module,attr,omitempty"`
	Scenario string   `xml:"scenario,attr,omitempty"`
}
