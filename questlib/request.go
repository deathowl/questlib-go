package questlib

import "fmt"

type RequestType int


const(
	KILL RequestType =  iota
    GATHER
    USE
    VISIT
    TALK
)

func (rt RequestType) String() string {
    return [...]string{"kill", "gather", "use", "visit", "talk"}[rt]
}

type QuestRequest struct {
        SubjectId int `json:"id"`;
        RequestType RequestType `json:"type"`;
        Count int `json:"count"`;
} 

func newQuestRequest(subjectId int, requesttype RequestType, count int) *QuestRequest {
    qr := QuestRequest{SubjectId: subjectId, RequestType: requesttype, Count: count}
    return &qr
}

func (qr QuestRequest) String() string {
	return fmt.Sprintf("QuestRequest{subject=%d, action='%s', amount='%d'}", qr.SubjectId, qr.RequestType, qr.Count)
}
