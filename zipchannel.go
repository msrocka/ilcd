package ilcd

import (
	"archive/zip"
)

type zDataEntry struct {
	path string
	data []byte
}

//
func (r *ZipReader) MapReduce(
	fn func(file *zip.File) (string, []byte), w *ZipWriter) {
	if w == nil {
		return
	}
	c := make(chan *zDataEntry)
	go func() {
		files := r.r.File
		for i := range files {
			path, data := fn(files[i])
			if path == "" || len(data) == 0 {
				continue
			}
			c <- &zDataEntry{path, data}
		}
	}()
	go func() {
		for e := range c {
			w.WriteEntry(e.path, e.data)
		}
	}()
}
