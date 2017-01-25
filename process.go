package ilcd

import (
	"encoding/xml"
)

// Process represents an ILCD process data set
type Process struct {
	XMLName   xml.Name     `xml:"processDataSet"`
	Info      *ProcessInfo `xml:"processInformation>dataSetInformation"`
	Exchanges []Exchange   `xml:"exchanges>exchange"`
}

// ProcessInfo contains the general process information
type ProcessInfo struct {
	UUID     string       `xml:"baseName"`
	Name     *ProcessName `xml:"name"`
	Synonyms []LangString `xml:"synonyms"`
}

// ProcessName contains the name fields of a process.
type ProcessName struct {
	BaseName       LangStrings  `xml:"baseName"`
	Treatment      []LangString `xml:"treatmentStandardsRoutes"`
	MixAndLocation []LangString `xml:"mixAndLocationTypes"`
	Properties     []LangString `xml:"functionalUnitFlowProperties"`
}

// Exchange is an input or output of an ILCD process data set.
type Exchange struct {
	InternalID      int     `xml:"dataSetInternalID,attr"`
	Flow            Ref     `xml:"referenceToFlowDataSet"`
	Direction       string  `xml:"exchangeDirection"`
	MeanAmount      float64 `xml:"meanAmount"`
	ResultingAmount float64 `xml:"resultingAmount"`
}
