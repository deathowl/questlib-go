package questlib

type RequestType int


const(
	KILL RequestType =  iota
    GATHER
    USE
    VISIT
    TALK
)

type QuestRequest struct {
        SubjectId int `json:"id"`;
        RequestType RequestType `json:"type"`;
        Count int `json:"count"`;
} 

func newQuestRequest(subjectId int, requesttype RequestType, count int) *QuestRequest {
    qr := QuestRequest{SubjectId: subjectId, RequestType: requesttype, Count: count}
    return &qr
}