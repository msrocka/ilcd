package ilcd

// DataSetType is an enumeration type of the different ILCD data set types.
type DataSetType int

// The ILCD data set types
const (
	ProcessType DataSetType = iota + 1
	MethodType
	FlowType
	FlowPropertyType
	UnitGroupType
	SourceType
	ContactType
	ExternalDocType
	UnknownType
)

func (t DataSetType) String() string {
	switch t {
	case ProcessType:
		return "process data set"
	case SourceType:
		return "source data set"
	case ContactType:
		return "contact data set"
	case FlowType:
		return "flow data set"
	case FlowPropertyType:
		return "flow property data set"
	case UnitGroupType:
		return "unit group data set"
	case MethodType:
		return "LCIA method data set"
	case ExternalDocType:
		return "other external file"
	case UnknownType:
		return "unknown?"
	default:
		return "unknown?"
	}
}

// Folder returns the name of the folder where data sets of the given type are
// stored in an ILCD package.
func (t DataSetType) Folder() string {
	switch t {
	case ProcessType:
		return "processes"
	case SourceType:
		return "sources"
	case ContactType:
		return "contacts"
	case FlowType:
		return "flows"
	case FlowPropertyType:
		return "flowproperties"
	case UnitGroupType:
		return "unitgroups"
	case MethodType:
		return "lciamethods"
	case ExternalDocType:
		return "external_docs"
	case UnknownType:
		return "other"
	default:
		return "other"
	}
}
