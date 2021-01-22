package valueobject

import (
	"io"
	"strings"
)

// File static file helper object, contains useful attributes about a file to be used anywhere
type File struct {
	Name      string
	Size      int64
	Extension string
	File      io.Reader
}

// File creates and populates required data by File
func NewFile(filename string, size int64, file io.Reader) *File {
	name, ext := getFileNameExtension(filename)
	return &File{
		Name:      name,
		Size:      size,
		Extension: ext,
		File:      file,
	}
}

// FileNameExtension returns the file name and extension from a raw file name
//	Note: Merged both name and extension to avoid multiple string splitting
//
//	e.g "foo.png -> {"foo", "png"}
func getFileNameExtension(name string) (string, string) {
	sl := strings.Split(name, ".")
	if len(sl) < 2 {
		return name, ""
	}
	return sl[0], strings.ToLower(sl[len(sl)-1])
}
