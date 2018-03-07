package ilcd

import "encoding/xml"

// Model represents a life cycle model data set of the extended ILCD (eILCD)
// format.
type Model struct {
	XMLName     xml.Name           `xml:"lifeCycleModelDataSet"`
	Info        *ProcessInfo       `xml:"lifeCycleModelInformation>dataSetInformation"`
	QRef        *int               `xml:"lifeCycleModelInformation>quantitativeReference>referenceToReferenceProcess"`
	Processes   []ProcessInstance  `xml:"lifeCycleModelInformation>technology>processes>processInstance"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
}

// UUID returns the UUID of the data set.
func (m *Model) UUID() string {
	if m == nil || m.Info == nil {
		return ""
	}
	return m.Info.UUID
}

// Version returns the version of the data set.
func (m *Model) Version() string {
	if m == nil || m.Publication == nil {
		return ""
	}
	return m.Publication.Version
}

// FullName returns the full name of the life cylce model for the given language
// whith all name parts concatenated to a single string.
func (m *Model) FullName(lang string) string {
	if m == nil || m.Info == nil || m.Info.Name == nil {
		return ""
	}
	return m.Info.Name.concat(lang)
}

// ProcessInstance describes a process reference together with its connections
// in a life cycle model.
type ProcessInstance struct {
	InternalID  int                 `xml:"dataSetInternalID,attr"`
	Process     *Ref                `xml:"referenceToProcess"`
	Connections []ProcessConnection `xml:"connections>outputExchange"`
	Parameters  []ModelParameter    `xml:"parameters>parameter"`
}

// ProcessConnection describes a connection between two processes in a life
// cycle model.
type ProcessConnection struct {
	OutputFlow string           `xml:"flowUUID,attr"`
	IsDominant *bool            `xml:"dominant,attr,omitempty"`
	Links      []DownstreamLink `xml:"downstreamProcess"`
}

// A DownstreamLink links the output of a process to an input of another process
// in a life cylce model.
type DownstreamLink struct {
	InputFlow  string `xml:"flowUUID,attr"`
	ProcessID  int    `xml:"id,attr"`
	Location   string `xml:"location,attr,omitempty"`
	IsDominant *bool  `xml:"dominant,attr,omitempty"`
}

// A ModelParameter is a parameter of a process instance in a life cycle model.
type ModelParameter struct {
	Name  string  `xml:"name,attr"`
	Value float64 `xml:",chardata"`
}
