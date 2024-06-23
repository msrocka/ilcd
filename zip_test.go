package ilcd

import (
	"path/filepath"
	"testing"
)

func TestWriteRead(t *testing.T) {
	zipPath := filepath.Join(t.TempDir(), "ilcd.zip")
	w, err := NewZipWriter(zipPath)
	if err != nil {
		t.Fatal("failed to create zip writer", err)
	}

	err = w.WriteDataSet(&Process{
		Info: &ProcessInfo{
			UUID: "abc",
			Name: &ProcessName{
				BaseName: LangStringOf("process A", "en"),
			},
		},
	})
	if err != nil {
		t.Fatal("failed to write process")
	}

	if err := w.Close(); err != nil {
		t.Fatal("failed to close writer", err)
	}

	r, err := NewZipReader(zipPath)
	if err != nil {
		t.Fatal("failed to create reader")
	}

	p, err := r.ReadProcess("abc")
	if err != nil || p.Info.Name.BaseName.Default() != "process A" {
		t.Fatal("failed to read process")
	}

	count := 0
	err = r.EachProcess(func(p *Process) bool {
		count++
		return true
	})
	if err != nil || count != 1 {
		t.Fatal("failed to iterate over processes")
	}

	if err := r.Close(); err != nil {
		t.Fatal("failed to close reader")
	}
}
