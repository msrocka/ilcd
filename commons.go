package ilcd

type LangStrings []LangString

// LangString represents an ILCD multi-language string
type LangString struct {
	Value string `xml:",chardata"`
	Lang  string `xml:"lang,attr"`
}

func (list LangStrings) GetValue(lang string) string {
	if list == nil {
		return ""
	}
	for _, ls := range list {
		if ls.Lang == lang {
			return ls.Value
		}
	}
	return ""
}

// Ref is a data set reference to an ILCD data set.
type Ref struct {
	UUID    string `xml:"refObjectId,attr"`
	Type    string `xml:"type,attr"`
	URI     string `xml:"uri,attr"`
	Version string `xml:"version,attr"`
	Name    string `xml:"shortDescription"`
}
