// @title           black hole API
// @version         1.0
// @description     This is API with Swagger documentation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9080
// @BasePath  /

package main

import (
	_ "docker-black-hole/docs"
	"docker-black-hole/internal/app"
	"docker-black-hole/internal/env"
	"docker-black-hole/internal/swagger"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/validator/v10"
	"io"
	"log"
)

func main() {
	config := env.GetEnv()
	if config.DisableLogs != 0 {
		log.SetOutput(io.Discard)
		gin.DefaultWriter = io.Discard
	}
	ginCtx := gin.Default()

	swagger.Controller(ginCtx)
	app.Controller(ginCtx)
	log.Printf("Execute from user %v\n", config.ExecuteFromUser)
	log.Printf("Listening on %v\n", config.Port)
	ginCtx.Run(":" + config.Port)
}
