package router

import (
    "github.com/HachimanHiki/zkrpApi/api"

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
    
    router.LoadHTMLGlob("view/html/*")
    router.Static("/css", "view/asset")
    router.GET("/", api.GetIndex)
    router.GET("/marathon", api.GetMarathon)
    router.GET("/share", api.GetShare)
    router.GET("/result/zkrp", api.GetZkrpResult)
    router.GET("/result/share", api.GetShareResult)

    router.GET("account", api.GetAccount)
    //router.POST("/merkletree", api.PostMerkleTreeRoot)
    router.POST("/merkletree", api.VerifyMerkleTreeRoot)
    router.POST("/prove", api.NewProve)
    router.POST("/verify", api.PostVerify)
    router.POST("/event", api.PostEvent)
    router.GET("/event_qr", api.GetQRcode)
    router.GET("/event", api.GetEvent)
    router.DELETE("/event", api.DelEvent)

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    router.Run(":8080")
}