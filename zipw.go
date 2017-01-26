package ilcd

import (
	"archive/zip"
	"os"
)

// ZipWriter provides functions to write ILCD zip packages
type ZipWriter struct {
	w *zip.Writer
}

// NewZipWriter creates a new ZipWriter.
func NewZipWriter(filePath string) (*ZipWriter, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	writer := &ZipWriter{w: zip.NewWriter(file)}
	return writer, nil
}

// Close closes the underlying zip file.
func (w *ZipWriter) Close() error {
	return w.w.Close()
}

// PutProcessData adds the given process data set
func (w *ZipWriter) PutProcessData(uuid string, data []byte) error {
	return w.PutData("ILCD/processes/"+uuid+".xml", data)
}

// PutFlowData adds the given flow data set
func (w *ZipWriter) PutFlowData(uuid string, data []byte) error {
	return w.PutData("ILCD/flows/"+uuid+".xml", data)
}

// PutFlowPropertyData adds the given flow property data set
func (w *ZipWriter) PutFlowPropertyData(uuid string, data []byte) error {
	return w.PutData("ILCD/flowproperties/"+uuid+".xml", data)
}

// PutUnitGroupData adds the given unit group data set
func (w *ZipWriter) PutUnitGroupData(uuid string, data []byte) error {
	return w.PutData("ILCD/unitgroups/"+uuid+".xml", data)
}

// PutSourceData adds the given source data set
func (w *ZipWriter) PutSourceData(uuid string, data []byte) error {
	return w.PutData("ILCD/sources/"+uuid+".xml", data)
}

// PutContactData adds the given contact data set
func (w *ZipWriter) PutContactData(uuid string, data []byte) error {
	return w.PutData("ILCD/contacts/"+uuid+".xml", data)
}

// PutData stores the given data under the given file name in the zip.
func (w *ZipWriter) PutData(filePath string, data []byte) error {
	writer, err := w.w.Create(filePath)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}
