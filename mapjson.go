package mapjson

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ParseJsonString(inpS string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(inpS), &result)
	if err != nil {
		log.Fatal("mapjson :: ParseJsonString - ", err)
	}
	return result
}

func ParseJsonFile(path string) map[string]interface{} {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("mapjson :: ParseJsonFile - ", err)
	}

	return ParseJsonString(string(fileData))
}
