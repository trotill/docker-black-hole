package app

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"net/http"
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

type JobRequestJob struct {
	Action    string   `json:"action"`
	Arguments []string `json:"arguments"`
	Type      string   `json:"type"`
	Timeout   uint32   `json:"timeout"`
}

func SetJob(ctx *gin.Context) {
	var json JobRequestJob
	if err := ctx.ShouldBindBodyWithJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("json %+v\n", json)
	//c.JSON(http.StatusOK, gin.H{"id": id})
}
