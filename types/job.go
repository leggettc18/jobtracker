package types

type Job struct {
    ID int `json:"id"`
    Position string `json:"position"`
    Company string `json:"company"`
    Listing string `json:"listing"`
    CreatedAt string `json:"createdAt"`
    Status string `json:"status"`
    InterviewDate string `json:"interviewDate"`
    AcceptedDate string `json:"acceptedDate"`
}
