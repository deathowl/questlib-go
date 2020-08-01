package questlib

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"fmt"
)
func LoadQuestDefs(jsonPath string) *[]QuestDef {
	
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var quests []QuestDef
	err = json.Unmarshal(byteValue, &quests)
	if err != nil {
		fmt.Println(err)
	}
	return &quests
}