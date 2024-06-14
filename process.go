package ilcd

import (
	"encoding/xml"
	"strings"
)

// Process represents an ILCD process data set
type Process struct {
	XMLName          xml.Name `xml:"http://lca.jrc.it/ILCD/Process processDataSet"`
	EpdFormatVersion string   `xml:"http://www.indata.network/EPD/2019 epd-version,attr"`

	Info      *ProcessInfo      `xml:"processInformation>dataSetInformation"`
	QuantRef  *ProcessQuantRef  `xml:"processInformation>quantitativeReference"`
	Time      *ProcessTime      `xml:"processInformation>time"`
	Location  *ProcessLocation  `xml:"processInformation>geography>locationOfOperationSupplyOrProduction"`
	Tech      *ProcessTech      `xml:"processInformation>technology"`
	MathModel *ProcessMathModel `xml:"processInformation>mathematicalRelations"`
	Modelling *ProcessModelling `xml:"modellingAndValidation"`

	DataEntry     *CommonDataEntry   `xml:"administrativeInformation>dataEntryBy"`
	Publication   *CommonPublication `xml:"administrativeInformation>publicationAndOwnership"`
	Exchanges     []Exchange         `xml:"exchanges>exchange"`
	ImpactResults []ImpactResult     `xml:"LCIAResults>LCIAResult"`
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
	if p.QuantRef == nil || len(p.QuantRef.RefFlows) == 0 {
		return nil
	}
	n := 0
	var refs []*Exchange
	for _, e := range p.Exchanges {
		for _, id := range p.QuantRef.RefFlows {
			if id != e.InternalID {
				continue
			}
			refs = append(refs, &e)
			n++
			if n >= len(p.QuantRef.RefFlows) {
				return refs
			}
		}
	}
	return refs
}

// ProcessInfo contains the general process information
type ProcessInfo struct {
	UUID            string           `xml:"http://lca.jrc.it/ILCD/Common UUID"`
	Name            *ProcessName     `xml:"name"`
	Synonyms        LangString       `xml:"synonyms"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	Comment         LangString       `xml:"http://lca.jrc.it/ILCD/Common generalComment"`
	ExternalDocs    []Ref            `xml:"referenceToExternalDocumentation"`
	EpdExt          *EpdInfoExt      `xml:"http://lca.jrc.it/ILCD/Common other"`
}

// ProcessName contains the name fields of a process.
type ProcessName struct {
	BaseName       LangString `xml:"baseName"`
	Treatment      LangString `xml:"treatmentStandardsRoutes"`
	MixAndLocation LangString `xml:"mixAndLocationTypes"`
	Properties     LangString `xml:"functionalUnitFlowProperties"`
}

// ProcessQuantRef contains the quantitative reference information of a process.
type ProcessQuantRef struct {
	Type     string `xml:"type,attr,omitempty"`
	RefFlows []int  `xml:"referenceToReferenceFlow"`
}

type ProcessTime struct {
	ReferenceYear *int        `xml:"http://lca.jrc.it/ILCD/Common referenceYear"`
	ValidUntil    *int        `xml:"http://lca.jrc.it/ILCD/Common dataSetValidUntil"`
	EpdExt        *EpdTimeExt `xml:"http://lca.jrc.it/ILCD/Common other"`
}

// ProcessLocation contains the information of a process location.
type ProcessLocation struct {
	Code        string     `xml:"location,attr,omitempty"`
	LatLong     string     `xml:"latitudeAndLongitude,attr,omitempty"`
	Description LangString `xml:"descriptionOfRestrictions"`
}

type ProcessTech struct {
	Description   LangString `xml:"technologyDescriptionAndIncludedProcesses"`
	Applicability LangString `xml:"technologicalApplicability"`
	Pictogramme   *Ref       `xml:"referenceToTechnologyPictogramme"`
	FlowDiagrams  []Ref      `xml:"referenceToTechnologyFlowDiagrammOrPicture"`
}

type ProcessMathModel struct {
	Description LangString  `xml:"modelDescription"`
	Parameters  []Parameter `xml:"variableParameter"`
}

// Parameter contains the information of a process parameter or variable under
// the tag <variableParameter>
type Parameter struct {
	Name    string     `xml:"name,attr"`
	Formula string     `xml:"formula,omitempty"`
	Value   float64    `xml:"meanValue"`
	SD95    float64    `xml:"relativeStandardDeviation95In"`
	Comment LangString `xml:"comment"`
}

// Exchange is an input or output of an ILCD process data set. Note that an
// exchange has a MeanAmount and ResultingAmount. Both values are the same if
// the exchange has no reference to a variable. Otherwise the ResultingAmount
// is calculated via the formula: ResultingAmount = MeanAmount * Variable.
type Exchange struct {
	InternalID      int           `xml:"dataSetInternalID,attr"`
	Flow            *Ref          `xml:"referenceToFlowDataSet"`
	Direction       string        `xml:"exchangeDirection"`
	MeanAmount      float64       `xml:"meanAmount"`
	Variable        string        `xml:"referenceToVariable,omitempty"`
	ResultingAmount float64       `xml:"resultingAmount"`
	Location        string        `xml:"location"`
	EpdExt          *EpdResultExt `xml:"http://lca.jrc.it/ILCD/Common other"`
}

type ImpactResult struct {
	Method     *Ref          `xml:"referenceToLCIAMethodDataSet"`
	MeanAmount float64       `xml:"meanAmount"`
	Comment    LangString    `xml:"generalComment"`
	EpdExt     *EpdResultExt `xml:"http://lca.jrc.it/ILCD/Common other"`
}
