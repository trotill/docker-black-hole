package types

type JobRequest struct {
	// Job ID. If more than one application is running at the same time, the job ID is used to separate their responses
	// example: myLs
	Id string `json:"id" binding:"required,min=2,max=100"`
	// The name of the application/script or the path to the application/script. If ALLOW_ABSOLUTE_MODE=0, the path is replaced by SCRIPT_PATH
	// example: ls
	Action string `json:"action" binding:"required,min=1,max=2048"`
	// Array of application/script string arguments
	// example: ["-la"]
	Arguments []string `json:"arguments" binding:"required,dive,min=1,max=100"`
	// Execution method, by absolute path or from SCRIPT_PATH folder
	// example: "absolute"
	Type string `json:"type" enums:"related,absolute" binding:"required,oneof=related absolute"`
	// Maximum execution time (in seconds)
	// example: 60
	Timeout uint32 `json:"timeout" default:"1000" binding:"required"`
}

type ExecRequest struct {
	Path    string
	Args    []string
	Timeout uint32
}

type JobError struct {
	// The error code of the job
	Code string `json:"code"`
	// The error description of the job
	Description string `json:"description"`
}

type JobResult struct {
	// Job stdout
	Stdout string `json:"stdout"`
	// Job stderr
	Stderr string `json:"stderr"`
	// Job exit code
	ExitCode int `json:"exit_code"`
}
type JobResponse struct {
	// The status of the job
	// example: finish
	Status string `json:"status" enums:"run,error,finish"`
	// The error of the job execution
	Error *JobError `json:"error,omitempty"`
	// The result of the job execution
	Result *JobResult `json:"result,omitempty"`
}

type JobListItem struct {
	// Job ID
	Id string `json:"id"`
	// Job request
	Payload *JobRequest `json:"payload"`
	// Job result
	Result JobResponse `json:"result"`
	// Job create time
	CreatedAt int64 `json:"created_at"`
}
