package types

type JobRequest struct {
	Id        string   `json:"id" binding:"required,min=2,max=100"`
	Action    string   `json:"action" binding:"required,min=1,max=2048"`
	Arguments []string `json:"arguments" binding:"required,dive,min=2,max=100"`
	Type      string   `json:"type" enums:"embedded,related,absolute" binding:"required,oneof=embedded related absolute"`
	Timeout   uint32   `json:"timeout" default:"1000" binding:"required"`
}

type JobError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type JobResponse struct {
	Status string    `json:"status" enums:"run,error,finish"`
	Error  *JobError `json:"error,omitempty"`
}

type JobListItem struct {
	Id        string       `json:"id"`
	Payload   *JobRequest  `json:"payload"`
	Result    *JobResponse `json:"result"`
	CreatedAt int64        `json:"created_at"`
}
