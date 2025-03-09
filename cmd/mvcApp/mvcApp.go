// @title           My API
// @version         1.0
// @description     This is a sample API with Swagger documentation.
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
	//"github.com/doug-martin/goqu/v9"
	_ "docker-black-hole/docs"
	"docker-black-hole/internal/app"
	"docker-black-hole/internal/swagger"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	ginCtx := gin.Default()
	swagger.Controller(ginCtx)
	app.Controller(ginCtx)
	port := os.Getenv("PORT")
	ginCtx.Run(":" + port)
}
