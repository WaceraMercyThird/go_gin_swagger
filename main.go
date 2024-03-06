package main

import (
	"go_api/handlers"
	"log"
	
	"github.com/gin-gonic/gin"
	
	swaggerFiles "github.com/swaggo/files"
	
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go_api/docs"
)

// @title Your API Title
// @version 1.0
// @description Your API Description
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html




func main() {
    // Create a new Gin router
    r := gin.Default()

    // Define API groups
    v1 := r.Group("/api/v1")
    user := v1.Group("/users")
    {
        user.GET("/", handlers.GetUsers)
    }

    // Configure Swagger UI endpoint
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("doc.json")))

    // Run the application
    err := r.Run()

    if err != nil {
        log.Fatal(err)
    }
}
