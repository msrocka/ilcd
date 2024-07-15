package ilcd

import (
	"archive/zip"
	"encoding/xml"
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

// Path calculates the path of the zip entry of the given data set.
func (w *ZipWriter) Path(ds DataSet) string {
	if ds == nil {
		return ""
	}
	file := ds.UUID() + "_" + ds.Version() + ".xml"
	dsType := Type(ds)
	return "ILCD/" + dsType.Folder() + "/" + file
}

// WriteDataSet writes the given data set to the zip package.
func (w *ZipWriter) WriteDataSet(ds DataSet) error {
	if ds == nil {
		return nil
	}
	data, err := xml.MarshalIndent(ds, "", "  ")
	if err != nil {
		return err
	}
	return w.Write(w.Path(ds), data)
}

// Write writes the given data under the given path into the zip package.
func (w *ZipWriter) Write(path string, data []byte) error {
	if path == "" || len(data) == 0 {
		return nil
	}
	writer, err := w.w.Create(path)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}
