package types

type JobRequest struct {
	Id        string   `json:"id" binding:"required,min=2,max=100"`
	Action    string   `json:"action" binding:"required,min=1,max=2048"`
	Arguments []string `json:"arguments" binding:"required,dive,min=1,max=100"`
	Type      string   `json:"type" enums:"related,absolute" binding:"required,oneof=related absolute"`
	Timeout   uint32   `json:"timeout" default:"1000" binding:"required"`
}

type ExecRequest struct {
	Path    string
	Args    []string
	Timeout uint32
}

type JobError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type JobResult struct {
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exit_code"`
}
type JobResponse struct {
	Status string     `json:"status" enums:"run,error,finish"`
	Error  *JobError  `json:"error,omitempty"`
	Result *JobResult `json:"result,omitempty"`
}

type JobListItem struct {
	Id        string      `json:"id"`
	Payload   *JobRequest `json:"payload"`
	Result    JobResponse `json:"result"`
	CreatedAt int64       `json:"created_at"`
}
