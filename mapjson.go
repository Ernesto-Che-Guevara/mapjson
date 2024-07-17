package mapjson

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func parseJsonString(inpS string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(inpS), &result)
	if err != nil {
		log.Print(inpS)
		log.Fatal("mapjson :: parseJsonString - ", err)
	}
	return result
}

func parseJsonFile(path string) map[string]interface{} {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print(err)
	}

	return parseJsonString(string(fileData))
}