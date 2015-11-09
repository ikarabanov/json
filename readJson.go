package readJson

import (
	"encoding/json"
	"io/ioutil"
)

func loadJSON(filePath string, data interface{}) (e error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileData, data)
	if err != nil {
		return err
	}
	return nil
}
func writeJSON(filePath string, data interface{}) (e error) {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, fileData, 0644)
	if err != nil {
		return err
	}
	return nil
}
