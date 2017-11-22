package ilcd

import "encoding/xml"

// Contact represents an ILCD contact data set
type Contact struct {
	XMLName     xml.Name           `xml:"contactDataSet"`
	Info        *ContactInfo       `xml:"contactInformation>dataSetInformation"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
}

// ContactInfo <dataSetInformation>
type ContactInfo struct {
	UUID            string           `xml:"UUID"`
	ShortName       LangString       `xml:"shortName"`
	Name            LangString       `xml:"name"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	Address         LangString       `xml:"contactAddress"`
	Email           string           `xml:"email,omitempty"`
	URL             string           `xml:"WWWAddress,omitempty"`
	Comment         LangString       `xml:"contactDescriptionOrComment"`
}
