package app

import (
	//"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

// https://habr.com/ru/articles/758662/#3
// https://gist.github.com/h3ssan/9510fbb2291d41b090cf52adb2edd1c4
// https://app.studyraid.com/en/read/5926/130190/using-gin-with-databases

// @Summary      Get app
// @Description  Returns app
// @Tags         app
// @Accept       json
// @Produce      json
// @Param        request  body  UsersGetAllRequest  true  "Users request with filter"
// @Success      200  {object}  map[string]string
// @Router       /app [get]
func getUsersHandler(ctx *gin.Context, db *goqu.Database) {
	users := GetAllUsers(db)
	ctx.JSON(200, users)
}
func Controller(g *gin.Engine) {
	g.POST("/job", func(ctx *gin.Context) {
		SetJob(ctx)
	})
}
