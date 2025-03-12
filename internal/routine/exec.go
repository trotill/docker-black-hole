package routine

import (
	"context"
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
	time.Sleep(1 * time.Second)
	fmt.Printf("End %s\n", job.Payload.Action)
}
func ExecRoutine(ctx context.Context, job *types.JobListItem) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout", job.Id)
			return
		default:
			job.Result.Status = utils.JOB_STATUS_RUN
			if job.Payload.Type == utils.JOB_TYPE_EMBEDDED {
				RunEmbedded(job)
			}
			//time.Sleep(500 * time.Millisecond)
			//fmt.Println("Текущее число Фибоначчи:")
		}
	}
}
