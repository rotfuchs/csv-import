package file_reader

import (
	"log"
	"os"
)

type FileReader struct {
	filePath string
	fileHandler *os.File
}

type FileReaderInterface interface {
	SetPath(path string)
	GetPath() string
}

func (f *FileReader) SetPath(path string) {
	if isPathValid(path) {
		f.filePath = path
	}
}

func (f *FileReader) GetPath() string {
	return f.filePath
}

func (f *FileReader) getFileHandler() *os.File {
	if f.fileHandler == nil {
		f.fileHandler = getNewFileHandler(f.filePath)
	}

	return f.fileHandler
}

func getNewFileHandler(path string) *os.File {
	fileHandler, err := os.Open(path)

	if err != nil {
		log.Fatalln("Couldn't open the file", err)
	}

	return fileHandler
}

func isPathValid(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}