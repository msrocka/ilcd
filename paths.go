package ilcd

import (
	"regexp"
	"strings"
)

var uuidRegex *regexp.Regexp

// FindUUID returns the UUID from the given path or an empty string if it cannot
// find it.
func FindUUID(path string) string {
	if uuidRegex == nil {
		pattern := "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
		uuidRegex = regexp.MustCompile(pattern)
	}
	return uuidRegex.FindString(path)
}

// IsModelPath returns true if the given file path or zip entry name is
// probably a life cycle model data set (of the extended ILCD format).
func IsModelPath(path string) bool {
	return isXMLInFolder(path, "lifecyclemodels")
}

// IsMethodPath returns true if the given file path or zip entry name is
// probably a LCIA method data set.
func IsMethodPath(path string) bool {
	return isXMLInFolder(path, "lciamethods")
}

// IsProcessPath returns true if the given file path or zip entry name is
// probably a process data set.
func IsProcessPath(path string) bool {
	return isXMLInFolder(path, "processes")
}

// IsFlowPath returns true if the given file path or zip entry name is
// probably a flow data set.
func IsFlowPath(path string) bool {
	return isXMLInFolder(path, "flows")
}

// IsFlowPropertyPath returns true if the given file path or zip entry name is
// probably a flow property data set.
func IsFlowPropertyPath(path string) bool {
	return isXMLInFolder(path, "flowproperties")
}

// IsUnitGroupPath returns true if the given file path or zip entry name is
// probably a unit group data set.
func IsUnitGroupPath(path string) bool {
	return isXMLInFolder(path, "unitgroups")
}

// IsSourcePath returns true if the given file path or zip entry name is
// probably a source data set.
func IsSourcePath(path string) bool {
	return isXMLInFolder(path, "sources")
}

// IsContactPath returns true if the given file path or zip entry name is
// probably a contact data set.
func IsContactPath(path string) bool {
	return isXMLInFolder(path, "contacts")
}

// IsExternalDoc returns true if the given path is something in the
// `external_docs` folder.
func IsExternalDoc(path string) bool {
	p := strings.ToLower(path)
	if strings.HasSuffix(p, "external_docs") {
		// we a searching something *in* the external doc folder, not the folder
		// itself
		return false
	}
	return strings.Contains(p, "external_docs")
}

func isXMLInFolder(path, folder string) bool {
	p := strings.ToLower(path)
	if !strings.Contains(p, folder) {
		return false
	}
	return strings.HasSuffix(p, ".xml")
}

type xmlPath struct {
	uuid   string
	dsType DataSetType
}

func xmlPathOf(s string) *xmlPath {

	// the file extension must be ".xml"
	sLen := len(s)
	if sLen < 5 {
		return nil
	}
	ext := strings.ToLower(s[sLen-4 : sLen])
	if ext != ".xml" {
		return nil
	}

	// the segment before the file name must
	// indicate the dataset type
	parts := strings.Split(s[0:sLen-4], "/")
	n := len(parts)
	if n < 2 {
		return nil
	}
	dsType := dsTypeOfFolder(parts[n-2])

	// the can be an optional version suffix
	// after the uuid
	uuid := parts[n-1]
	splitPos := -1
	for i := 0; i < len(uuid); i++ {
		if uuid[i] == '_' {
			splitPos = i
			break
		}
	}
	if splitPos > 0 {
		uuid = uuid[0:splitPos]
	}

	return &xmlPath{uuid, dsType}
}

func dsTypeOfFolder(folder string) DataSetType {
	s := strings.ToLower(strings.TrimSpace(folder))
	switch s {
	case "lifecyclemodels":
		return ModelDataSet
	case "processes":
		return ProcessDataSet
	case "sources":
		return SourceDataSet
	case "contacts":
		return ContactDataSet
	case "flows":
		return FlowDataSet
	case "flowproperties":
		return FlowPropertyDataSet
	case "unitgroups":
		return UnitGroupDataSet
	case "lciamethods":
		return MethodDataSet
	case "external_docs":
		return ExternalDoc
	default:
		return Asset
	}
}
