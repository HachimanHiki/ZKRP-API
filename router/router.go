package router

import (
    "github.com/zkrpApi/api"

    "github.com/swaggo/gin-swagger" // gin-swagger middleware
    "github.com/swaggo/files" // swagger embed files

    "github.com/gin-gonic/gin"
)

// InitRouter initialize a router
func InitRouter() {

    //gin.SetMode(gin.ReleaseMode)
	//使用gin的Default方法建立一個路由handler
	
    router := gin.Default()

    router.NoRoute(api.NotFound)
    
    router.POST("/prove", api.PostUserProve)
    router.POST("/verify", api.PostVerify)

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    router.Run(":8080")
}