package ilcd

import "encoding/xml"

// UnitGroup represents an ILCD unit group data set
type UnitGroup struct {
	XMLName     xml.Name           `xml:"unitGroupDataSet"`
	Info        *UnitGroupInfo     `xml:"unitGroupInformation>dataSetInformation"`
	QRef        int                `xml:"unitGroupInformation>quantitativeReference>referenceToReferenceUnit"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
	Units       []Unit             `xml:"units>unit"`
}

// UUID returns the UUID of the data set.
func (ug *UnitGroup) UUID() string {
	if ug == nil || ug.Info == nil {
		return ""
	}
	return ug.Info.UUID
}

// Version returns the version of the data set.
func (ug *UnitGroup) Version() string {
	if ug == nil || ug.Publication == nil {
		return ""
	}
	return ug.Publication.Version
}

// ReferenceUnit returns the reference unit of an unit group.
func (ug *UnitGroup) ReferenceUnit() *Unit {
	if ug == nil {
		return nil
	}
	for i := range ug.Units {
		if ug.Units[i].InternalID == ug.QRef {
			return &ug.Units[i]
		}
	}
	return nil
}

// UnitGroupInfo <dataSetInformation>
type UnitGroupInfo struct {
	UUID            string           `xml:"UUID"`
	Name            LangString       `xml:"name"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	Comment         LangString       `xml:"generalComment"`
}

// Unit contains the information of a <unit> element in an unit group data set.
type Unit struct {
	InternalID int     `xml:"dataSetInternalID,attr"`
	Name       string  `xml:"name"`
	Factor     float64 `xml:"meanValue"`
}
