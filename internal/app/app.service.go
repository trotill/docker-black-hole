package app

import (
	"docker-black-hole/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/fatih/structs"
	"time"
)

type UserItem struct {
	Id         int       `db:"id" json:"id,omitempty"`
	Name       string    `db:"name"`
	Updated_at time.Time `db:"updated_at"`
	Created_at time.Time `db:"created_at"`
}

type UsersResponseStruct struct {
	Count uint32     `json:"count"`
	Items []UserItem `json:"items"`
}

type UsersGetAllRequest struct {
	Filter struct {
		Name string   `json:"name,omitempty"`
		Ids  []string `json:"ids,omitempty"`
	}
	Offset uint64 `json:"offset,default=0"`
	Limit  uint64 `json:"limit,default=20"`
}

func GetAllUsers(db *goqu.Database) map[string]interface{} {

	var usersDbResult []UserItem

	count, err := db.From("app").Count()
	if err != nil {
		fmt.Println(err.Error())
	}

	ds, err := db.From("app").Executor().Scanner()
	if err != nil {
		fmt.Println("Read app", err.Error())
	}
	if ds == nil {
		return structs.Map(UsersResponseStruct{Count: 0, Items: usersDbResult})
	}
	defer ds.Close()
	scanErr := ds.ScanStructs(&usersDbResult)
	if scanErr != nil {
		fmt.Println("Error: Read item", scanErr.Error())
	}
	fmt.Printf("GOQU results df %+v\n", usersDbResult)

	return structs.Map(UsersResponseStruct{Count: uint32(count), Items: usersDbResult})
}

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

var jobList = []JobListItem{}

func FindJobById(jobList []JobListItem, id string) *JobListItem {
	for _, job := range jobList {
		if job.Id == id {
			return &job
		}
	}
	return nil
}

func GetJob(id string) *JobListItem {
	fmt.Println("GetJob", id)
	foundJob := FindJobById(jobList, id)
	DumpJobList()
	if foundJob == nil {
		return nil
	}
	return foundJob
}
func DumpJobList() {
	js, _ := json.Marshal(jobList)
	fmt.Printf("json %+v\n", string(js))
}
func SetJob(job *JobRequest) bool {
	foundJob := FindJobById(jobList, job.Id)
	if foundJob == nil {
		jobList = append(jobList, JobListItem{Id: job.Id, Payload: job, CreatedAt: utils.GetUnixTimestamp()})
		return true
	}
	return false
	//c.JSON(http.StatusOK, gin.H{"id": id})
}
