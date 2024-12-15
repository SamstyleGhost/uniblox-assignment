package jsonfileworker

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
)

func GetAllObjects(pathToFile string, v interface{}) error {

	absPathToFile, _ := filepath.Abs(pathToFile)
	pathFile, err := os.Open(absPathToFile)
	if err != nil {
		return err
	}

	defer pathFile.Close()

	byteValue, err := io.ReadAll(pathFile)
	if err != nil {
		return err
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(byteValue, v)
	if err != nil {
		return err
	}

	return nil
}

// Will see if I can work on updating modularly if possible
func SetAllObjects(pathToFile string, v interface{}) error {
	absPathToFile, _ := filepath.Abs(pathToFile)

	pathFile, err := os.OpenFile(absPathToFile, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer pathFile.Close()

	updatedContents, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	_, err = pathFile.Write(updatedContents)
	if err != nil {
		return err
	}

	return nil
}
