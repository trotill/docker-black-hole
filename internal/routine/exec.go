package routine

import (
	"bytes"
	"context"
	"docker-black-hole/internal/env"
	"docker-black-hole/internal/types"
	"docker-black-hole/internal/utils"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func execute(payload *types.ExecRequest) types.JobResult {
	execPath := payload.Path
	execArgs := payload.Args

	config := env.GetEnv()
	var timeout time.Duration
	if (payload.Timeout == 0) || (int(payload.Timeout) > config.ExecuteMaxTimeoutSec) {
		timeout = time.Duration(config.ExecuteMaxTimeoutSec * int(time.Second))
	} else {
		timeout = time.Duration(int(payload.Timeout) * int(time.Second))
	}

	cmdCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	command := execPath
	if len(execArgs) > 0 {
		command += " " + strings.Join(execArgs, " ")
	}
	var cmd *exec.Cmd

	if config.Docker == 1 {
		cmd = exec.CommandContext(cmdCtx, "nsenter",
			"-t", "1", "-m", "-u", "-i", "-n", "-p",
			"su", "-", config.ExecuteFromUser, "-c",
			config.ShellPath, "-c", command)
	} else {
		cmd = exec.CommandContext(cmdCtx, config.ShellPath, "-c", command)
	}
	log.Printf("cmd: %v\n", cmd.String())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	exitCode := 0
	runErr := cmd.Run()
	if runErr != nil {
		if exitError, ok := runErr.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
			log.Printf("Terminated <%+v %+v> with exit code %d\n", execPath, execArgs, exitCode)
		} else {
			log.Printf("Error run command <%v %v> %v\n", execPath, execArgs, runErr)
		}
	} else {
		log.Printf("Finish <%+v %+v>\n", execPath, execArgs)
	}
	result := types.JobResult{}
	result.Stderr = stderr.String()
	result.Stdout = stdout.String()
	result.ExitCode = exitCode
	return result
}

func RunAbsolute(job *types.JobListItem) {
	execParam := &types.ExecRequest{Path: job.Payload.Action, Args: job.Payload.Arguments, Timeout: job.Payload.Timeout}
	result := execute(execParam)
	job.Result.Result = &result
}

func RunRelated(job *types.JobListItem) {
	basename := filepath.Base(job.Payload.Action)
	config := env.GetEnv()
	path := config.ScriptPath + basename
	execParam := &types.ExecRequest{Path: path, Args: job.Payload.Arguments, Timeout: job.Payload.Timeout}
	result := execute(execParam)

	job.Result.Result = &result
}

func ExecRoutine(job *types.JobListItem) {
	config := env.GetEnv()
	job.Result.Status = utils.JOB_STATUS_RUN
	switch job.Payload.Type {
	case utils.JOB_TYPE_RELATED:
		RunRelated(job)
		job.Result.Status = utils.JOB_STATUS_FINISH
	case utils.JOB_TYPE_ABSOLUTE:
		if config.AllowAbsoluteMode != 1 {
			job.Result.Status = utils.JOB_STATUS_ERROR
			job.Result.Error = &types.JobError{Code: "restrictedType", Description: "restricted type"}
			break
		}
		RunAbsolute(job)
		job.Result.Status = utils.JOB_STATUS_FINISH
	default:
		job.Result.Status = utils.JOB_STATUS_ERROR
		job.Result.Error = &types.JobError{Code: "unknownType", Description: "unknown type"}
	}
	log.Printf("End %+v\n", job)
}
