package main
import (
	"github.com/deathowl/questlib-go/questlib"
	"fmt"
)
func main(){
	var quests []questlib.QuestDef
	quests = *questlib.LoadQuestDefs("./test.json")
	for _, q := range quests {
		fmt.Println(q)
	}}