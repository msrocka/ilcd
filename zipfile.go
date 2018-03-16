package ilcd

import (
	"archive/zip"
	"io/ioutil"
)

// ZipFile embedds the type `File` from the `archive/zip` package and provides
// additional ILCD specific methods.
type ZipFile struct {
	f *zip.File
}

// newZipFile initializes a new ZipFile from the given archive file.
func newZipFile(f *zip.File) *ZipFile {
	return &ZipFile{f}
}

// Path returns the path of the zip file within the zip package.
func (f *ZipFile) Path() string {
	return f.f.Name
}

// Reads the decompressed data from the zip file.
func (f *ZipFile) Read() ([]byte, error) {
	reader, err := f.f.Open()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if err := reader.Close(); err != nil {
		return nil, err
	}
	return data, nil
}

// ReadModel reads a life cycle model data set from the zip file.
func (f *ZipFile) ReadModel() (*Model, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadModel(data)
}

// ReadMethod reads a Method data set from the zip file.
func (f *ZipFile) ReadMethod() (*Method, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadMethod(data)
}

// ReadProcess reads a Process data set from the zip file.
func (f *ZipFile) ReadProcess() (*Process, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadProcess(data)
}

// ReadFlow reads a Flow data set from the zip file.
func (f *ZipFile) ReadFlow() (*Flow, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadFlow(data)
}

// ReadFlowProperty reads a FlowProperty data set from the zip file.
func (f *ZipFile) ReadFlowProperty() (*FlowProperty, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadFlowProperty(data)
}

// ReadUnitGroup reads a UnitGroup data set from the zip file.
func (f *ZipFile) ReadUnitGroup() (*UnitGroup, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadUnitGroup(data)
}

// ReadSource reads a Source data set from the zip file.
func (f *ZipFile) ReadSource() (*Source, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadSource(data)
}

// ReadContact reads a Contact data set from the zip file.
func (f *ZipFile) ReadContact() (*Contact, error) {
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	return ReadContact(data)
}

/*
func (f *ZipFile) unmarshal(ds interface{}) error {
	data, err := f.Read()
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, ds)
	return err
}
*/

// Type returns the ILCD data set type of the zip file which is inferred from
// the path of the zip file.
func (f *ZipFile) Type() DataSetType {
	path := f.Path()
	if IsModelPath(path) {
		return ModelDataSet
	}
	if IsMethodPath(path) {
		return MethodDataSet
	}
	if IsProcessPath(path) {
		return ProcessDataSet
	}
	if IsFlowPath(path) {
		return FlowDataSet
	}
	if IsFlowPropertyPath(path) {
		return FlowPropertyDataSet
	}
	if IsUnitGroupPath(path) {
		return UnitGroupDataSet
	}
	if IsSourcePath(path) {
		return SourceDataSet
	}
	if IsContactPath(path) {
		return ContactDataSet
	}
	if IsExternalDoc(path) {
		return ExternalDoc
	}
	return Asset
}
