package ilcd

import (
	"archive/zip"
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

// FindDataSet searches for a data set of the give type and with the given
// uuid and returns the corresponding zip file. If nothing is found, it returns
// nil.
func (r *ZipReader) FindDataSet(dsType DataSetType, uuid string) *ZipFile {
	if r == nil {
		return nil
	}
	dsFolder := dsType.Folder()
	files := r.r.File
	for i := range files {
		f := files[i]
		path := strings.ToLower(f.Name)
		if !strings.Contains(path, dsFolder) {
			continue
		}
		if !strings.HasSuffix(path, ".xml") {
			continue
		}
		if strings.Contains(path, uuid) {
			return newZipFile(f)
		}
	}
	return nil
}

func (r *ZipReader) ReadSource(uuid string) (*Source, error) {
	entry := r.FindDataSet(SourceDataSet, uuid)
	if entry == nil {
		return nil, ErrDataSetNotFound
	}
	return entry.ReadSource()
}

func (r *ZipReader) ReadContact(uuid string) (*Contact, error) {
	entry := r.FindDataSet(ContactDataSet, uuid)
	if entry == nil {
		return nil, ErrDataSetNotFound
	}
	return entry.ReadContact()
}

func (r *ZipReader) ReadUnitGroup(uuid string) (*UnitGroup, error) {
	entry := r.FindDataSet(UnitGroupDataSet, uuid)
	if entry == nil {
		return nil, ErrDataSetNotFound
	}
	return entry.ReadUnitGroup()
}

func (r *ZipReader) ReadFlowProperty(uuid string) (*FlowProperty, error) {
	entry := r.FindDataSet(FlowPropertyDataSet, uuid)
	if entry == nil {
		return nil, ErrDataSetNotFound
	}
	return entry.ReadFlowProperty()
}

func (r *ZipReader) ReadFlow(uuid string) (*Flow, error) {
	entry := r.FindDataSet(FlowDataSet, uuid)
	if entry == nil {
		return nil, ErrDataSetNotFound
	}
	return entry.ReadFlow()
}

func (r *ZipReader) ReadProcess(uuid string) (*Process, error) {
	entry := r.FindDataSet(ProcessDataSet, uuid)
	if entry == nil {
		return nil, ErrDataSetNotFound
	}
	return entry.ReadProcess()
}

func (r *ZipReader) ReadModel(uuid string) (*Model, error) {
	entry := r.FindDataSet(ModelDataSet, uuid)
	if entry == nil {
		return nil, ErrDataSetNotFound
	}
	return entry.ReadModel()
}

func (r *ZipReader) ReadMethod(uuid string) (*Method, error) {
	entry := r.FindDataSet(MethodDataSet, uuid)
	if entry == nil {
		return nil, ErrDataSetNotFound
	}
	return entry.ReadMethod()
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
		file := files[i]
		if file.FileInfo().IsDir() {
			continue
		}
		zf := newZipFile(file)
		if !fn(zf) {
			break
		}
	}
}

type zDataEntry struct {
	path string
	data []byte
}

// Map applies the given function to all entries in the zip file and writes
// the function's output to the given writer.
//
// The result of a function call is the path and the data that should be written
// to the writer. If the path or data are empty, nothing will be written. The
// given function is executed in a separate Go routine.
func (r *ZipReader) Map(w *ZipWriter, fn func(file *ZipFile) (string, []byte)) {
	if w == nil {
		return
	}
	c := make(chan *zDataEntry)
	go func() {
		r.EachFile(func(zf *ZipFile) bool {
			path, data := fn(zf)
			if path != "" && len(data) > 0 {
				c <- &zDataEntry{path, data}
			}
			return true
		})
		close(c)
	}()
	for {
		entry, more := <-c
		if !more {
			break
		}
		w.Write(entry.path, entry.data)
	}
}
