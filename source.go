package ilcd

import "encoding/xml"

// Source represents an ILCD source data set
type Source struct {
	XMLName     xml.Name           `xml:"sourceDataSet"`
	Info        *SourceInfo        `xml:"sourceInformation>dataSetInformation"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
}

// UUID returns the UUID of the data set.
func (s *Source) UUID() string {
	if s == nil || s.Info == nil {
		return ""
	}
	return s.Info.UUID
}

// Version returns the version of the data set.
func (s *Source) Version() string {
	if s == nil || s.Publication == nil {
		return ""
	}
	return s.Publication.Version
}

// SourceInfo <dataSetInformation>
type SourceInfo struct {
	UUID            string           `xml:"UUID"`
	ShortName       LangString       `xml:"shortName"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	Citation        string           `xml:"sourceCitation,omitempty"`
	PublicationType string           `xml:"publicationType,omitempty"`
	FileRefs        []SourceFileRef  `xml:"referenceToDigitalFile"`
}

type SourceFileRef struct {
	URI string `xml:"uri,attr,omitempty"`
}
