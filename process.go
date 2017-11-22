package ilcd

import (
	"encoding/xml"
)

// Process represents an ILCD process data set
type Process struct {
	XMLName     xml.Name           `xml:"processDataSet"`
	Info        *ProcessInfo       `xml:"processInformation>dataSetInformation"`
	Location    *ProcessLocation   `xml:"processInformation>geography>locationOfOperationSupplyOrProduction"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
	Exchanges   []Exchange         `xml:"exchanges>exchange"`
}

// ProcessInfo contains the general process information
type ProcessInfo struct {
	UUID            string           `xml:"UUID"`
	Name            *ProcessName     `xml:"name"`
	Synonyms        LangString       `xml:"synonyms"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	Comment         LangString       `xml:"generalComment"`
}

// ProcessName contains the name fields of a process.
type ProcessName struct {
	BaseName       LangString `xml:"baseName"`
	Treatment      LangString `xml:"treatmentStandardsRoutes"`
	MixAndLocation LangString `xml:"mixAndLocationTypes"`
	Properties     LangString `xml:"functionalUnitFlowProperties"`
}

// ProcessLocation contains the information of a process location.
type ProcessLocation struct {
	Code        string     `xml:"location,attr"`
	LatLong     string     `xml:"latitudeAndLongitude,attr"`
	Description LangString `xml:"descriptionOfRestrictions"`
}

// Exchange is an input or output of an ILCD process data set.
type Exchange struct {
	InternalID      int     `xml:"dataSetInternalID,attr"`
	Flow            Ref     `xml:"referenceToFlowDataSet"`
	Direction       string  `xml:"exchangeDirection"`
	MeanAmount      float64 `xml:"meanAmount"`
	ResultingAmount float64 `xml:"resultingAmount"`
}
