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

func unmarshal(file *zip.File, ds interface{}) error {
	data, err := readData(file)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, ds)
	return err
}

func readData(file *zip.File) ([]byte, error) {
	reader, err := file.Open()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}

// EachProcess iterates over each process data set in the ILCD package and calls
// the given handler with the respective process data set.
func (r *ZipReader) EachProcess(handler func(*Process) bool) error {
	for _, f := range r.r.File {
		if IsProcessPath(f.Name) {
			process := &Process{}
			err := unmarshal(f, process)
			if err != nil {
				return err
			}
			if !handler(process) {
				break
			}
		}
	}
	return nil
}

// EachFlow iterates over each flow data set in the ILCD package and calls the
// given function with the respective flow data set.
func (r *ZipReader) EachFlow(fn func(*Flow) bool) error {
	for _, f := range r.r.File {
		if IsFlowPath(f.Name) {
			flow := &Flow{}
			err := unmarshal(f, flow)
			if err != nil {
				return err
			}
			if !fn(flow) {
				break
			}
		}
	}
	return nil
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
