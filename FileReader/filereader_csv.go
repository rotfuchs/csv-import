package filereader

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
)

type Csv struct {
	fileReader   *FileReader
	colSeparator rune
	columns      []string
	rowCounter   int
	reader       *csv.Reader
}

func NewCsvFileReader(filePath string) *Csv {
	fileReader := &FileReader{}
	fileReader.SetPath(filePath)

	return &Csv{
		fileReader:   fileReader,
		colSeparator: ';',
		rowCounter:   0,
	}
}

func (c *Csv) SetColSeparator(colSep rune) {
	c.colSeparator = colSep
}

func (c *Csv) GetHeader() []string {
	r := c.getNewCsvReader()

	columns, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	return columns
}

func (c *Csv) ReadAll() ([][]string, error) {
	if c.reader == nil {
		c.reader = c.getNewCsvReader()
	}
	c.reader.Comma = c.colSeparator

	return c.reader.ReadAll()
}

func (c *Csv) GetNextDataSet() ([]string, error) {
	if c.reader == nil {
		c.reader = c.getNewCsvReader()
	}
	//c.reader.Comma = c.colSeparator

	return c.reader.Read()
}

func (c *Csv) Count() (int, error) {
	file := c.fileReader.getFileHandler()
	buf := make([]byte, 32*1024)

	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func (c *Csv) getNewCsvReader() *csv.Reader {
	reader := csv.NewReader(getNewFileHandler(c.fileReader.filePath))
	reader.Comma = c.colSeparator

	return reader
}
