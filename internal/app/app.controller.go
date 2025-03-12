package app

import (
	"context"
	"docker-black-hole/internal/types"
	"docker-black-hole/internal/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"net/http"
)

// https://habr.com/ru/articles/758662/#3
// https://gist.github.com/h3ssan/9510fbb2291d41b090cf52adb2edd1c4
// https://app.studyraid.com/en/read/5926/130190/using-gin-with-databases

func getUsersHandler(ctx *gin.Context, db *goqu.Database) {
	users := GetAllUsers(db)
	ctx.JSON(200, users)
}

// SetJobController
// @Summary      Run job
// @Description  Returns job id
// @Tags         job
// @Accept       json
// @Produce      json
// @Param        request  body  types.JobRequest  true  "Job request"
// @Success      201
// @Failure      409 {object} utils.HttpError
// @Router       /job [post]
func SetJobController(ctx *gin.Context, goCtx context.Context) {
	var json types.JobRequest
	if err := ctx.ShouldBindBodyWithJSON(&json); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, utils.HttpError{Code: "validation", Description: "validation error", Validation: err})
		return
	}

	if SetJob(&json, goCtx) {
		ctx.JSON(http.StatusCreated, nil)
	} else {
		utils.ErrorResponse(ctx, http.StatusConflict, utils.HttpError{Code: "jobExists", Description: "job already exists"})
	}
}

// GetJobController
// @Summary      Get job
// @Description  Returns job info
// @Tags         job
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Job id"
// @Success      200  {object} types.JobListItem
// @Router       /job/{id} [get]
func GetJobController(ctx *gin.Context) {
	id := ctx.Param("id")
	job := GetJob(id)
	if job == nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, utils.HttpError{Code: "jobNotFound", Description: "job not found"})
		return
	}
	ctx.JSON(http.StatusOK, job)
}
func Controller(g *gin.Engine, goCtx context.Context) {

	g.POST("/job", func(ctx *gin.Context) {
		SetJobController(ctx, goCtx)
	})

	g.GET("/job/:id", func(ctx *gin.Context) {
		GetJobController(ctx)
	})
}
