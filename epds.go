package ilcd

type EpdInfoExtension struct {
	SafetyMargins *EpdSafetyMargins `xml:"http://www.iai.kit.edu/EPD/2013 safetyMargins"`
	Scenarios     []EpdScenario     `xml:"http://www.iai.kit.edu/EPD/2013 scenarios>scenario"`
}

type EpdSafetyMargins struct {
	Value       *float64   `xml:"http://www.iai.kit.edu/EPD/2013 margins"`
	Description LangString `xml:"http://www.iai.kit.edu/EPD/2013 description"`
}

type EpdScenario struct {
	Name        string     `xml:"name,attr,omitempty"`
	IsDefault   bool       `xml:"default,attr"`
	Group       string     `xml:"group,attr,omitempty"`
	Description LangString `xml:"http://www.iai.kit.edu/EPD/2013 description"`
}

type EpdResultExtension struct {
	Values    []EpdValue `xml:"http://www.iai.kit.edu/EPD/2013 amount"`
	UnitGroup *Ref       `xml:"http://www.iai.kit.edu/EPD/2013 referenceToUnitGroupDataSet"`
}

type EpdValue struct {
	Amount   *float64 `xml:",chardata,omitempty"`
	Module   string   `xml:"module,attr,omitempty"`
	Scenario string   `xml:"scenario,attr,omitempty"`
}
