package editor

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const (
	RTColorscheme = "colorscheme"
	RTSyntax      = "syntax"
)

func readDir(fs http.FileSystem, path string) ([]os.FileInfo, error) {
	if fs == nil {
		return nil, os.ErrNotExist
	}

	dir, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	return dir.Readdir(0)
}

func readFile(fs http.FileSystem, path string) ([]byte, error) {
	if fs == nil {
		return nil, os.ErrNotExist
	}

	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// RuntimeFile allows the program to read runtime data like colorschemes or syntax files
type RuntimeFile interface {
	// Name returns a name of the file without paths or extensions
	Name() string
	// Data returns the content of the file.
	Data() ([]byte, error)
}

// allFiles contains all available files, mapped by filetype
var allFiles map[string][]RuntimeFile

// some file on filesystem
type realFile struct {
	fs   http.FileSystem
	path string
}

// some file on filesystem but with a different name
type namedFile struct {
	realFile
	name string
}

// a file with the data stored in memory
type memoryFile struct {
	name string
	data []byte
}

func (mf memoryFile) Name() string {
	return mf.name
}
func (mf memoryFile) Data() ([]byte, error) {
	return mf.data, nil
}

func (rf realFile) Name() string {
	fn := filepath.Base(rf.path)
	return fn[:len(fn)-len(filepath.Ext(fn))]
}

func (rf realFile) Data() ([]byte, error) {
	return readFile(rf.fs, rf.path)
}

func (nf namedFile) Name() string {
	return nf.name
}

// RuntimeFiles tracks a set of runtime files.
type RuntimeFiles struct {
	allFiles map[string][]RuntimeFile
}

// NewRuntimeFiles creates a new set of runtime files from the colorscheme and syntax files present in the given
// http.Filesystme. Colorschemes should be located under "/colorschemes" and must end with a "micro" extension.
// Syntax files should be located under "/syntax" and must end with a "yaml" extension.
func NewRuntimeFiles(fs http.FileSystem) *RuntimeFiles {
	rfs := &RuntimeFiles{}
	rfs.AddFilesFromDirectory(fs, RTColorscheme, "/colorschemes", "*.micro")
	rfs.AddFilesFromDirectory(fs, RTSyntax, "/syntax", "*.yaml")
	return rfs
}

// AddRuntimeFile registers a file for the given filetype
func (rfs *RuntimeFiles) AddFile(fileType string, file RuntimeFile) {
	if rfs.allFiles == nil {
		rfs.allFiles = make(map[string][]RuntimeFile)
	}
	rfs.allFiles[fileType] = append(rfs.allFiles[fileType], file)
}

// AddFilesFromDirectory registers each file from the given directory for
// the filetype which matches the file-pattern
func (rfs *RuntimeFiles) AddFilesFromDirectory(fs http.FileSystem, fileType, directory, pattern string) {
	files, _ := readDir(fs, directory)
	for _, f := range files {
		if ok, _ := filepath.Match(pattern, f.Name()); !f.IsDir() && ok {
			fullPath := filepath.Join(directory, f.Name())
			rfs.AddFile(fileType, realFile{fs, fullPath})
		}
	}
}

// FindFile finds a runtime file of the given filetype and name
// will return nil if no file was found
func (rfs RuntimeFiles) FindFile(fileType, name string) RuntimeFile {
	for _, f := range rfs.ListRuntimeFiles(fileType) {
		if f.Name() == name {
			return f
		}
	}
	return nil
}

// ListRuntimeFiles lists all known runtime files for the given filetype
func (rfs RuntimeFiles) ListRuntimeFiles(fileType string) []RuntimeFile {
	if files, ok := rfs.allFiles[fileType]; ok {
		return files
	}
	return []RuntimeFile{}
}
