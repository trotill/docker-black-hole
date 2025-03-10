package routine

import (
	"docker-black-hole/internal/types"
	"docker-black-hole/internal/utils"
	"fmt"
	"time"
)

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
	time.Sleep(10 * time.Second)
}
func ExecRoutine(job *types.JobListItem) {
	job.Result.Status = utils.JOB_STATUS_RUN
	if job.Payload.Type == utils.JOB_TYPE_EMBEDDED {
		RunEmbedded(job)
	}
}
