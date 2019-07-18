package vips

import (
	"io"
	"os"
)

// LazyFile is a lazy reader or writer
type LazyFile struct {
	name string
	file *os.File
}

func LazyOpen(file string) io.Reader {
	return &LazyFile{name: file}
}

func LazyCreate(file string) io.Writer {
	return &LazyFile{name: file}
}

func (r *LazyFile) Read(p []byte) (n int, err error) {
	if r.file == nil {
		f, err := os.Open(r.name)
		if err != nil {
			return 0, err
		}
		r.file = f
	}
	return r.file.Read(p)
}

func (r *LazyFile) Close() error {
	if r.file != nil {
		_ = r.file.Close()
		r.file = nil
	}
	return nil
}

func (r *LazyFile) Write(p []byte) (n int, err error) {
	if r.file == nil {
		f, err := os.Create(r.name)
		if err != nil {
			return 0, err
		}
		r.file = f
	}
	return r.file.Write(p)
}
