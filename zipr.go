package ilcd

import (
	"archive/zip"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

// ZipReader can read data sets from ILCD packages.
type ZipReader struct {
	r *zip.ReadCloser
}

// NewZipReader creates a new package reader.
func NewZipReader(filePath string) (*ZipReader, error) {
	r, err := zip.OpenReader(filePath)
	return &ZipReader{r: r}, err
}

// Close closes the pack reader.
func (r *ZipReader) Close() error {
	return r.r.Close()
}

// GetProcess returns the process with the given UUID from the zip.
func (r *ZipReader) GetProcess(uuid string) (*Process, error) {
	file := r.findXML("processes", uuid)
	if file == nil {
		return nil, ErrDataSetNotFound
	}
	p := &Process{}
	err := unmarshal(file, p)
	return p, err
}

// GetProcessData returns the process data set with the given UUID as byte array.
func (r *ZipReader) GetProcessData(uuid string) ([]byte, error) {
	return r.xmlData("processes", uuid)
}

// GetFlow returns the flow with the given UUID from the zip.
func (r *ZipReader) GetFlow(uuid string) (*Flow, error) {
	file := r.findXML("flows", uuid)
	if file == nil {
		return nil, ErrDataSetNotFound
	}
	f := &Flow{}
	err := unmarshal(file, f)
	return f, err
}

// GetFlowData returns the flow data set with the given UUID as byte array.
func (r *ZipReader) GetFlowData(uuid string) ([]byte, error) {
	return r.xmlData("flows", uuid)
}

// GetFlowProperty returns the flow property with the given UUID from the zip.
func (r *ZipReader) GetFlowProperty(uuid string) (*FlowProperty, error) {
	file := r.findXML("flowproperties", uuid)
	if file == nil {
		return nil, ErrDataSetNotFound
	}
	fp := &FlowProperty{}
	err := unmarshal(file, fp)
	return fp, err
}

// GetFlowPropertyData returns the flow property data set with the given UUID
// as byte array.
func (r *ZipReader) GetFlowPropertyData(uuid string) ([]byte, error) {
	return r.xmlData("flowproperties", uuid)
}

// GetUnitGroupData returns the unit group data set with the given UUID as byte array.
func (r *ZipReader) GetUnitGroupData(uuid string) ([]byte, error) {
	return r.xmlData("unitgroups", uuid)
}

// GetSource returns the source data set with the given UUID as byte array.
func (r *ZipReader) GetSource(uuid string) ([]byte, error) {
	return r.xmlData("sources", uuid)
}

// GetContact returns the contact data set with the given UUID as byte array.
func (r *ZipReader) GetContact(uuid string) ([]byte, error) {
	return r.xmlData("contacts", uuid)
}

func (r *ZipReader) xmlData(path, uuid string) ([]byte, error) {
	file := r.findXML(path, uuid)
	if file == nil {
		return nil, ErrDataSetNotFound
	}
	return readData(file)
}

func (r *ZipReader) findXML(path, uuid string) *zip.File {
	for _, f := range r.r.File {
		name := f.Name
		if !strings.Contains(name, path) {
			continue
		}
		if !strings.HasSuffix(name, ".xml") {
			continue
		}
		if strings.Contains(name, uuid) {
			return f
		}
	}
	return nil
}

// EachModel iterates over each life cycle model in the package unless
// the given handler returns false.
func (r *ZipReader) EachModel(fn func(*Model) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsModelPath(f.Path()) {
			return true
		}
		val, err := f.ReadModel()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachMethod iterates over each Method data set in the package unless
// the given handler returns false.
func (r *ZipReader) EachMethod(fn func(*Method) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsMethodPath(f.Path()) {
			return true
		}
		val, err := f.ReadMethod()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachProcess iterates over each Process data set in the package unless
// the given handler returns false.
func (r *ZipReader) EachProcess(fn func(*Process) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsProcessPath(f.Path()) {
			return true
		}
		val, err := f.ReadProcess()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachFlow iterates over each Flow data set in the package unless
// the given handler returns false.
func (r *ZipReader) EachFlow(fn func(*Flow) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsFlowPath(f.Path()) {
			return true
		}
		val, err := f.ReadFlow()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachFlowProperty iterates over each FlowProperty data set in the package unless
// the given handler returns false.
func (r *ZipReader) EachFlowProperty(fn func(*FlowProperty) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsFlowPropertyPath(f.Path()) {
			return true
		}
		val, err := f.ReadFlowProperty()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachUnitGroup iterates over each UnitGroup data set in the package unless
// the given handler returns false.
func (r *ZipReader) EachUnitGroup(fn func(*UnitGroup) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsUnitGroupPath(f.Path()) {
			return true
		}
		val, err := f.ReadUnitGroup()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachSource iterates over each Source data set in the package unless
// the given handler returns false.
func (r *ZipReader) EachSource(fn func(*Source) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsSourcePath(f.Path()) {
			return true
		}
		val, err := f.ReadSource()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachContact iterates over each Contact data set in the package unless
// the given handler returns false.
func (r *ZipReader) EachContact(fn func(*Contact) bool) error {
	var gerr error
	r.EachFile(func(f *ZipFile) bool {
		if !IsContactPath(f.Path()) {
			return true
		}
		val, err := f.ReadContact()
		if err != nil {
			gerr = err
			return false
		}
		return fn(val)
	})
	return gerr
}

// EachFile calls the given function for each file in the zip package. It stops
// when the function returns false or when there are no more files in the
// package.
func (r *ZipReader) EachFile(fn func(f *ZipFile) bool) {
	files := r.r.File
	for i := range files {
		zf := newZipFile(files[i])
		if !fn(zf) {
			break
		}
	}
}

// EachEntry calls the given function with the name and data of each entry in
// the zip file.
func (r *ZipReader) EachEntry(fn func(name string, data []byte) error) error {
	for _, f := range r.r.File {
		data, err := readData(f)
		if err != nil {
			return err
		}
		err = fn(f.Name, data)
		if err != nil {
			return err
		}
	}
	return nil
}
