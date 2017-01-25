package ilcd

import (
	"archive/zip"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

// PackReader can read ILCD packages.
type PackReader struct {
	zipReader *zip.ReadCloser
}

// NewPackReader creates a new package reader.
func NewPackReader(filePath string) (*PackReader, error) {
	reader, err := zip.OpenReader(filePath)
	return &PackReader{zipReader: reader}, err
}

// Close closes the pack reader.
func (r *PackReader) Close() error {
	return r.zipReader.Close()
}

// EachProcess iterates over each process data set in the ILCD package and calls
// the given handler with the respective process data set.
func (r *PackReader) EachProcess(handler func(*Process) bool) error {
	for _, f := range r.zipReader.File {
		name := f.Name
		if strings.Contains(name, "processes/") && strings.HasSuffix(name, ".xml") {
			reader, err := f.Open()
			if err != nil {
				return err
			}
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				return err
			}
			err = reader.Close()
			if err != nil {
				return err
			}
			process := &Process{}
			err = xml.Unmarshal(data, process)
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
