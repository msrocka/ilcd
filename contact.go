package ilcd

import "encoding/xml"

// Contact represents an ILCD contact data set
type Contact struct {
	XMLName     xml.Name           `xml:"contactDataSet"`
	Info        *ContactInfo       `xml:"contactInformation>dataSetInformation"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
}

// UUID returns the UUID of the data set.
func (c *Contact) UUID() string {
	if c == nil || c.Info == nil {
		return ""
	}
	return c.Info.UUID
}

// Version returns the version of the data set.
func (c *Contact) Version() string {
	if c == nil || c.Publication == nil {
		return ""
	}
	return c.Publication.Version
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
