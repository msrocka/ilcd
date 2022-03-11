package ilcd

// LangString is an ILCD multi-language string
type LangString []LangStringItem

// LangStringItem represents an item in an ILCD multi-language string
type LangStringItem struct {
	Value string `xml:",chardata"`
	Lang  string `xml:"lang,attr"`
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

// Ref is a data set reference to an ILCD data set.
type Ref struct {
	UUID    string     `xml:"refObjectId,attr"`
	Type    string     `xml:"type,attr"`
	URI     string     `xml:"uri,attr"`
	Version string     `xml:"version,attr"`
	Name    LangString `xml:"shortDescription"`
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
	Name    string  `xml:"name,attr"`
	Classes []Class `xml:"class"`
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
	ID    string `xml:"classId1,attr"`
	Name  string `xml:",chardata"`
}

// CommonDataEntry <dataEntryBy>
type CommonDataEntry struct {
	TimeStamp   string `xml:"timeStamp"`
	DataFormats []Ref  `xml:"referenceToDataSetFormat"`
}

// CommonPublication <publicationAndOwnership>
type CommonPublication struct {
	Version string `xml:"dataSetVersion"`
	URI     string `xml:"permanentDataSetURI"`
}
