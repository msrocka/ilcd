package ilcd

import (
	"archive/zip"
	"fmt"
	"io"
)

type zipFile struct {
	f      *zip.File
	uuid   string
	dsType DataSetType
}

func newZipFile(f *zip.File) *zipFile {
	path := xmlPathOf(f.Name)
	if path == nil {
		return &zipFile{f, "", Asset}
	} else {
		return &zipFile{f, path.uuid, path.dsType}
	}
}

func (f *zipFile) String() string {
	return f.f.Name
}

type ZipReader struct {
	r *zip.ReadCloser
}

func NewZipReader(filePath string) (*ZipReader, error) {
	r, err := zip.OpenReader(filePath)
	return &ZipReader{r: r}, err
}

func (r *ZipReader) Close() error {
	return r.r.Close()
}

func (r *ZipReader) ReadSource(uuid string) (*Source, error) {
	return findParseZipEntry(r, SourceDataSet, uuid, ReadSource)
}

func (r *ZipReader) ReadContact(uuid string) (*Contact, error) {
	return findParseZipEntry(r, ContactDataSet, uuid, ReadContact)
}

func (r *ZipReader) ReadUnitGroup(uuid string) (*UnitGroup, error) {
	return findParseZipEntry(r, UnitGroupDataSet, uuid, ReadUnitGroup)
}

func (r *ZipReader) ReadFlowProperty(uuid string) (*FlowProperty, error) {
	return findParseZipEntry(r, FlowPropertyDataSet, uuid, ReadFlowProperty)
}

func (r *ZipReader) ReadFlow(uuid string) (*Flow, error) {
	return findParseZipEntry(r, FlowDataSet, uuid, ReadFlow)
}

func (r *ZipReader) ReadProcess(uuid string) (*Process, error) {
	return findParseZipEntry(r, ProcessDataSet, uuid, ReadProcess)
}

func (r *ZipReader) ReadModel(uuid string) (*Model, error) {
	return findParseZipEntry(r, ModelDataSet, uuid, ReadModel)
}

func (r *ZipReader) ReadMethod(uuid string) (*Method, error) {
	return findParseZipEntry(r, MethodDataSet, uuid, ReadMethod)
}

func (r *ZipReader) EachModel(fn func(*Model) bool) error {
	return parseEachZipEntry(r, ModelDataSet, ReadModel, fn)
}

func (r *ZipReader) EachMethod(fn func(*Method) bool) error {
	return parseEachZipEntry(r, MethodDataSet, ReadMethod, fn)
}

func (r *ZipReader) EachProcess(fn func(*Process) bool) error {
	return parseEachZipEntry(r, ProcessDataSet, ReadProcess, fn)
}

func (r *ZipReader) EachFlow(fn func(*Flow) bool) error {
	return parseEachZipEntry(r, FlowDataSet, ReadFlow, fn)
}

func (r *ZipReader) EachFlowProperty(fn func(*FlowProperty) bool) error {
	return parseEachZipEntry(r, FlowPropertyDataSet, ReadFlowProperty, fn)
}

func (r *ZipReader) EachUnitGroup(fn func(*UnitGroup) bool) error {
	return parseEachZipEntry(r, UnitGroupDataSet, ReadUnitGroup, fn)
}

func (r *ZipReader) EachSource(fn func(*Source) bool) error {
	return parseEachZipEntry(r, SourceDataSet, ReadSource, fn)
}

func (r *ZipReader) EachContact(fn func(*Contact) bool) error {
	return parseEachZipEntry(r, ContactDataSet, ReadContact, fn)
}

func findParseZipEntry[T any](
	r *ZipReader,
	dsType DataSetType,
	uuid string,
	parser func([]byte) (*T, error),
) (*T, error) {
	files := r.r.File
	for i := range files {
		zf := newZipFile(files[i])
		if zf.dsType == dsType && zf.uuid == uuid {
			return parseZipEntry(zf, parser)
		}
	}
	return nil, ErrDataSetNotFound
}

func parseEachZipEntry[T any](
	r *ZipReader,
	dsType DataSetType,
	parser func([]byte) (*T, error),
	callback func(*T) bool,
) error {

	files := r.r.File
	for i := range files {
		file := files[i]
		if file.FileInfo().IsDir() {
			continue
		}
		entry := newZipFile(file)
		if entry.dsType != dsType {
			continue
		}

		dataSet, err := parseZipEntry(entry, parser)
		if err != nil {
			return err
		}

		if !callback(dataSet) {
			break
		}
	}
	return nil
}

func parseZipEntry[T any](
	entry *zipFile, parser func([]byte) (*T, error),
) (*T, error) {

	reader, err := entry.f.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", entry, err)
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		reader.Close()
		return nil, fmt.Errorf("failed to read %s: %w", entry, err)
	}

	if err := reader.Close(); err != nil {
		return nil, fmt.Errorf("failed to close %s: %w", entry, err)
	}

	return parser(data)
}
