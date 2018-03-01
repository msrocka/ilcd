package ilcd

import (
	"encoding/xml"
	"strings"
)

// Process represents an ILCD process data set
type Process struct {
	XMLName     xml.Name           `xml:"processDataSet"`
	Info        *ProcessInfo       `xml:"processInformation>dataSetInformation"`
	QRefs       []int              `xml:"processInformation>quantitativeReference>referenceToReferenceFlow"`
	Location    *ProcessLocation   `xml:"processInformation>geography>locationOfOperationSupplyOrProduction"`
	DataEntry   *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
	Exchanges   []Exchange         `xml:"exchanges>exchange"`
}

// UUID returns the UUID of the data set.
func (p *Process) UUID() string {
	if p == nil || p.Info == nil {
		return ""
	}
	return p.Info.UUID
}

// Version returns the version of the data set.
func (p *Process) Version() string {
	if p == nil || p.Publication == nil {
		return ""
	}
	return p.Publication.Version
}

// FullName returns the full name of the process for the given language whith
// all name parts concatenated to a single string.
func (p *Process) FullName(lang string) string {
	if p == nil || p.Info == nil || p.Info.Name == nil {
		return ""
	}
	return p.Info.Name.concat(lang)
}

func (name *ProcessName) concat(lang string) string {
	if name == nil {
		return ""
	}
	parts := [4]string{name.BaseName.Get(lang),
		name.Treatment.Get(lang),
		name.MixAndLocation.Get(lang),
		name.Properties.Get(lang)}
	n := ""
	for _, part := range parts {
		p := strings.TrimSpace(part)
		if p == "" {
			continue
		}
		if n == "" {
			n = p
		} else {
			n += "; " + p
		}
	}
	return n
}

// RefFlows returns the exchanges that are defined as quantitative refeferences
// of the process. In most cases this should be just one exchange.
func (p *Process) RefFlows() []*Exchange {
	if len(p.QRefs) == 0 {
		return nil
	}
	n := 0
	var refs []*Exchange
	for _, e := range p.Exchanges {
		for _, id := range p.QRefs {
			if id != e.InternalID {
				continue
			}
			refs = append(refs, &e)
			n++
			if n >= len(p.QRefs) {
				return refs
			}
		}
	}
	return refs
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
	Flow            *Ref    `xml:"referenceToFlowDataSet"`
	Direction       string  `xml:"exchangeDirection"`
	MeanAmount      float64 `xml:"meanAmount"`
	ResultingAmount float64 `xml:"resultingAmount"`
	Location        string  `xml:"location"`
}
