package questlib

import (
	"fmt"
)
type QuestDef struct {
	Id int `json:"id"`;
    Description string `json:"description"`;
    Ongoing string `json:"ongoing"`;
    Onfinished string `json:"onfinished"`;
    Prerequisites []int `json:"prerequisites"`;
    Touch  []int `json:"touch"`;
    Questgivers []int `json:"questgivers"`;
    Questrequest []QuestRequest `json:"required"`;
}
func newQuestDef(id int, description string, ongoing string, onfinished string, prerequisites []int, touch []int, questgivers []int, questrequest []QuestRequest) *QuestDef {
    q := QuestDef{Id: id, Description: description, Ongoing: ongoing, Onfinished: onfinished, Prerequisites: prerequisites, Touch : touch, Questgivers: questgivers, Questrequest: questrequest}
    return &q
}

func (q QuestDef) String() string {
	return fmt.Sprintf("QuestDef{id=%d, description='%s', ongoing='%s', onfinished='%s'}", q.Id, q.Description, q.Ongoing, q.Onfinished)
}