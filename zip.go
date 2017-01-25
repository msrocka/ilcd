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
func (r *ZipReader) GetProcess(uuid string) *Process {
	file := r.findXML("processes", uuid)
	if file == nil {
		return nil
	}
	p := &Process{}
	err := unmarshal(file, p)
	if err != nil {
		return nil
	}
	return p
}

// GetFlow returns the flow with the given UUID from the zip.
func (r *ZipReader) GetFlow(uuid string) *Flow {
	file := r.findXML("flows", uuid)
	if file == nil {
		return nil
	}
	f := &Flow{}
	err := unmarshal(file, f)
	if err != nil {
		return nil
	}
	return f
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
	reader, err := file.Open()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, ds)
	return err
}

// EachProcess iterates over each process data set in the ILCD package and calls
// the given handler with the respective process data set.
func (r *ZipReader) EachProcess(handler func(*Process) bool) error {
	for _, f := range r.r.File {
		name := f.Name
		if strings.Contains(name, "processes/") && strings.HasSuffix(name, ".xml") {
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
