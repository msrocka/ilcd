package ilcd

import "encoding/xml"

// LangString is an ILCD multi-language string
type LangString []LangStringItem

// LangStringItem represents an item in an ILCD multi-language string
type LangStringItem struct {
	Value string `xml:",chardata"`
	Lang  string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
}

// Get returns the value for the given language code
func (ls LangString) Get(lang string) string {
	if ls == nil {
		return ""
	}
	for _, item := range ls {
		if item.Lang == lang {
			return item.Value
		}
	}
	return ""
}

// Default returns the default value of a multi-language string which is the
// English string (code = "en") or the first string in the list.
func (ls LangString) Default() string {
	if len(ls) == 0 {
		return ""
	}
	for _, item := range ls {
		if item.Lang == "en" {
			return item.Value
		}
	}
	return ls[0].Value
}

func (ls LangString) GetOrDefault(lang string) string {
	s := ""
	for _, item := range ls {
		if item.Lang == lang {
			return item.Value
		}
		if item.Lang == "en" {
			s = item.Value
		}
	}
	return s
}

func LangStringOf(value string, lang string) LangString {
	return []LangStringItem{{Value: value, Lang: lang}}
}

// Ref is a data set reference to an ILCD data set.
type Ref struct {
	UUID    string     `xml:"refObjectId,attr"`
	Type    string     `xml:"type,attr"`
	URI     string     `xml:"uri,attr,omitempty"`
	Version string     `xml:"version,attr,omitempty"`
	Name    LangString `xml:"http://lca.jrc.it/ILCD/Common shortDescription"`
}

// DataSetType returns the type of the data set referenced as defined in the ILCD
// schema.
func (ref *Ref) DataSetType() DataSetType {
	if ref == nil {
		return 0
	}
	switch ref.Type {
	case "source data set":
		return SourceDataSet
	case "process data set":
		return ProcessDataSet
	case "flow data set":
		return FlowDataSet
	case "flow property data set":
		return FlowPropertyDataSet
	case "unit group data set":
		return UnitGroupDataSet
	case "contact data set":
		return ContactDataSet
	case "LCIA method data set":
		return MethodDataSet
	case "other external file":
		return ExternalDoc
	default:
		return ExternalDoc
	}
}

// Classification describes an ILCD classification entry in a data set
type Classification struct {
	XMLName xml.Name `xml:"http://lca.jrc.it/ILCD/Common classification"`
	Name    string   `xml:"name,attr,omitempty"`
	Classes []Class  `xml:"class"`
}

// GetClass returns the class with the given level from the classification.
func (c *Classification) GetClass(level int) *Class {
	if c == nil || c.Classes == nil {
		return nil
	}
	for i, class := range c.Classes {
		if class.Level == level {
			return &c.Classes[i]
		}
	}
	return nil
}

// Class is a category in an ILCD data set classification.
type Class struct {
	Level int    `xml:"level,attr"`
	ID    string `xml:"classId,attr,omitempty"`
	Name  string `xml:",chardata"`
}

// CommonDataEntry <dataEntryBy>
type CommonDataEntry struct {
	TimeStamp   string `xml:"http://lca.jrc.it/ILCD/Common timeStamp"`
	DataFormats []Ref  `xml:"http://lca.jrc.it/ILCD/Common referenceToDataSetFormat"`
}

// CommonPublication <publicationAndOwnership>
type CommonPublication struct {
	Version           string `xml:"http://lca.jrc.it/ILCD/Common dataSetVersion"`
	PrecedingVersions []Ref  `xml:"http://lca.jrc.it/ILCD/Common referenceToPrecedingDataSetVersion"`
	URI               string `xml:"http://lca.jrc.it/ILCD/Common permanentDataSetURI"`
	Owner             *Ref   `xml:"http://lca.jrc.it/ILCD/Common referenceToOwnershipOfDataSet"`
}
