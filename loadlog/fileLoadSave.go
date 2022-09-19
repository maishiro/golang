package main

import (
	"io/ioutil"
	"os"

	"github.com/go-gota/gota/dataframe"
)

type FileLoadSave interface {
	LoadFile(path string) (string, error)
	SaveFileCSV(df *dataframe.DataFrame, pathFile string)
}

func NewFileLoadSave() FileLoadSave {
	return &fileLoadSave{}
}

type fileLoadSave struct {
}

func (file *fileLoadSave) LoadFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (file *fileLoadSave) SaveFileCSV(df *dataframe.DataFrame, pathFile string) {
	csvFile, _ := os.Create(pathFile)
	defer csvFile.Close()

	df.WriteCSV(csvFile)
}
