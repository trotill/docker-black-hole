package app

import (
	"docker-black-hole/internal/types"
	"fmt"
	"sync"
)

type JobMapType struct {
	sync.RWMutex
	m map[string]*types.JobListItem
}

var JobMap = JobMapType{m: make(map[string]*types.JobListItem)}

func (jMap JobMapType) GetJob(key string) *types.JobListItem {
	jMap.RLock()
	defer jMap.RUnlock()
	return jMap.m[key]
}

func (jMap JobMapType) SetJob(key string, value *types.JobListItem) {
	jMap.Lock()
	defer jMap.Unlock()
	jMap.m[key] = value
}

func (jMap JobMapType) DeleteJob(key string) {
	jMap.Lock()
	defer jMap.Unlock()
	delete(jMap.m, key)
}

func (jMap JobMapType) GetAllJobs() map[string]*types.JobListItem {
	jMap.RLock()
	defer jMap.RUnlock()
	return jMap.m
}

func (jMap JobMapType) Dump() {
	jMap.RLock()
	defer jMap.RUnlock()
	fmt.Printf("MAP %+v\n", jMap.m)
}
