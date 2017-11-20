package ilcd

import (
	"strings"
)

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

func isXMLInFolder(path, folder string) bool {
	p := strings.ToLower(path)
	if !strings.Contains(p, folder) {
		return false
	}
	return strings.HasSuffix(p, ".xml")
}
