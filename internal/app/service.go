package app

import (
	"docker-black-hole/internal/routine"
	"docker-black-hole/internal/types"
	"docker-black-hole/internal/utils"
	"encoding/json"
	"fmt"
)

var jobList map[string]*types.JobListItem

func GetJob(id string) *types.JobListItem {
	//DumpJobList()
	return JobMap.GetJob(id)
}
func DumpJobList() {
	js, _ := json.Marshal(jobList)
	fmt.Printf("json %+v\n", string(js))
}

func SetJob(job *types.JobRequest) bool {
	var rewrite = false
	gotJob := JobMap.GetJob(job.Id)
	if gotJob == nil || gotJob.Result.Status == utils.JOB_STATUS_FINISH {
		rewrite = true
	}
	if rewrite {
		JobMap.SetJob(job.Id, &types.JobListItem{Id: job.Id, Payload: job, CreatedAt: utils.GetUnixTimestamp(), Result: types.JobResponse{Status: utils.JOB_STATUS_UNKNOWN, Result: nil, Error: nil}})
		go routine.ExecRoutine(JobMap.GetJob(job.Id))
		JobMap.Dump()
		return true
	}
	return false
}
