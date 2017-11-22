package ilcd

import "encoding/xml"

// Source represents an ILCD source data set
type Source struct {
	XMLName     xml.Name           `xml:"sourceDataSet"`
	Info        *SourceInfo        `xml:"sourceInformation>dataSetInformation"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
}

// SourceInfo <dataSetInformation>
type SourceInfo struct {
	UUID            string           `xml:"UUID"`
	ShortName       LangString       `xml:"shortName"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	Citation        string           `xml:"sourceCitation,omitempty"`
	PublicationType string           `xml:"publicationType,omitempty"`
}
