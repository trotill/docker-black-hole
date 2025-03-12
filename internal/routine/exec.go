package routine

import (
	"bytes"
	"context"
	"docker-black-hole/internal/types"
	"docker-black-hole/internal/utils"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func execute(ExecPath string, ExecArgs []string) types.JobResult {
	timeout := 10 * time.Second
	cmdCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	command := ExecPath
	if len(ExecArgs) > 0 {
		command += " " + strings.Join(ExecArgs, " ")
	}

	//cmd := exec.CommandContext(cmdCtx, "nsenter", "--target", "1", "--mount", "--uts", "--ipc", "--net", "--pid", "bash", "-c", command)

	cmd := exec.CommandContext(cmdCtx, "bash", "-c", command)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	exitCode := 0
	runErr := cmd.Run()
	if runErr != nil {
		if exitError, ok := runErr.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
			log.Printf("Terminated <%+v %+v> with exit code %d\n", ExecPath, ExecArgs, exitCode)
		} else {
			log.Printf("Error run command <%v %v> %v\n", ExecPath, ExecArgs, runErr)
		}
	} else {
		log.Printf("Finish <%+v %+v>\n", ExecPath, ExecArgs)
	}
	result := types.JobResult{}
	result.Stderr = stderr.String()
	result.Stdout = stdout.String()
	result.ExitCode = exitCode
	return result
}

func RunAbsolute(job *types.JobListItem) {
	result := execute(job.Payload.Action, job.Payload.Arguments)
	job.Result.Result = &result
}

func RunRelated(job *types.JobListItem) {
	basename := filepath.Base(job.Payload.Action)
	path := "/opt/" + basename
	result := execute(path, job.Payload.Arguments)

	job.Result.Result = &result
}
func RunEmbedded(job *types.JobListItem) {
	switch utils.EmbeddedActions(job.Payload.Action) {
	case utils.RebootHost:
		fmt.Println("Embedded:RebootHost")
	case utils.ComposeRestart:
		fmt.Println("Embedded:ComposeRestart")
	case utils.DockerRestart:
		fmt.Println("Embedded:DockerRestart")
	default:
		fmt.Println("Embedded:Unknown")
	}

	fmt.Printf("End %s\n", job.Payload.Action)
}
func ExecRoutine(job *types.JobListItem) {

	job.Result.Status = utils.JOB_STATUS_RUN
	switch job.Payload.Type {
	case utils.JOB_TYPE_EMBEDDED:
		RunEmbedded(job)
		job.Result.Status = utils.JOB_STATUS_FINISH
	case utils.JOB_TYPE_RELATED:
		RunRelated(job)
		job.Result.Status = utils.JOB_STATUS_FINISH
	case utils.JOB_TYPE_ABSOLUTE:
		RunAbsolute(job)
		job.Result.Status = utils.JOB_STATUS_FINISH
	default:
		job.Result.Status = utils.JOB_STATUS_ERROR
		job.Result.Error = &types.JobError{Code: "unknownType", Description: "unknown type"}
	}
	log.Printf("End %+v\n", job)
}
