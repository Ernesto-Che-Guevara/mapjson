package mapjson

import (
	"fmt"
	"log"
	"testing"
)

func TestParseJsonString(t *testing.T) {
	log.Println("testing TestParseJsonString")
	inpS := ` {
		"smth1" : 123,
		"smth2" : [4,5,6],
		"smth3" : "string7",
		"smth4" : true
	}
	`
	correctAns := make(map[string]interface{})
	correctAns["smth1"] = 123
	correctAns["smth2"] = []int{4, 5, 6}
	correctAns["smth3"] = "string7"
	correctAns["smth4"] = true

	if fmt.Sprint(ParseJsonString(inpS)) != fmt.Sprint(correctAns) {
		t.Fatalf("plan %s \n \t    but fact %s", fmt.Sprint(correctAns), fmt.Sprint(ParseJsonString(inpS)))
	}
}

func TestParseJsonFile(t *testing.T) {
	log.Println("testing TestParseJsonFile")
	path := "./test/test.json"

	correctAns := make(map[string]interface{})
	correctAns["smth1"] = 123
	correctAns["smth2"] = []int{4, 5, 6}
	correctAns["smth3"] = "string7"
	correctAns["smth4"] = true

	if fmt.Sprint(ParseJsonFile(path)) != fmt.Sprint(correctAns) {
		t.Fatalf("plan %s \n \t    but fact %s", fmt.Sprint(correctAns), fmt.Sprint(ParseJsonString(path)))
	}
}
