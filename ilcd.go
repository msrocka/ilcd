package ilcd

// DataSetType is an enumeration type of the different ILCD data set types.
type DataSetType int

// The ILCD data set types
const (
	ProcessDataSet DataSetType = iota + 1
	MethodDataSet
	FlowDataSet
	FlowPropertyDataSet
	UnitGroupDataSet
	SourceDataSet
	ContactDataSet
	ExternalDoc
)

// DataSetTypes returns a list with all possible data set types.
func DataSetTypes() []DataSetType {
	return []DataSetType{
		ProcessDataSet,
		MethodDataSet,
		FlowDataSet,
		FlowPropertyDataSet,
		UnitGroupDataSet,
		SourceDataSet,
		ContactDataSet,
		ExternalDoc}
}

func (t DataSetType) String() string {
	switch t {
	case ProcessDataSet:
		return "process data set"
	case SourceDataSet:
		return "source data set"
	case ContactDataSet:
		return "contact data set"
	case FlowDataSet:
		return "flow data set"
	case FlowPropertyDataSet:
		return "flow property data set"
	case UnitGroupDataSet:
		return "unit group data set"
	case MethodDataSet:
		return "LCIA method data set"
	case ExternalDoc:
		return "other external file"
	default:
		return "unknown?"
	}
}

// Folder returns the name of the folder where data sets of the given type are
// stored in an ILCD package.
func (t DataSetType) Folder() string {
	switch t {
	case ProcessDataSet:
		return "processes"
	case SourceDataSet:
		return "sources"
	case ContactDataSet:
		return "contacts"
	case FlowDataSet:
		return "flows"
	case FlowPropertyDataSet:
		return "flowproperties"
	case UnitGroupDataSet:
		return "unitgroups"
	case MethodDataSet:
		return "lciamethods"
	case ExternalDoc:
		return "external_docs"
	default:
		return "other"
	}
}

// DataSet contains the common functions of all data set types.
type DataSet interface {
	UUID() string
	Version() string
}

// FlowType is an enumeration type of the different ILCD flow types.
type FlowType int

// Enumaration constants for the ILCD flow types.
const (
	ElementaryFlow FlowType = iota + 1
	ProductFlow
	WasteFlow
	OtherFlow
)

func (ft FlowType) String() string {
	switch ft {
	case ElementaryFlow:
		return "Elementary flow"
	case ProductFlow:
		return "Product flow"
	case WasteFlow:
		return "Waste flow"
	default:
		return "Other flow"
	}
}
