package api

import (
	"fmt"
	"net/http"

	"github.com/HachimanHiki/zkrpApi/service"
	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {

	add, err := service.DeployContract("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5", 
		"0x19be17213c2ee781defeef4abc2d6964f1418177f8d54418c55412e198eebf2107c6e8cc1a43e9be4e7f331f7b729e53a2e14e37aa874383628cf89be3cd0ef6",
		"d248f0355b955dad7d88be03cf279654bd8ebbbbc8d6302ae19ff34c26143eae",
	)

	fmt.Println(add)

	test, err := service.GetCommitment(add)

	fmt.Println(err)

	if(err == nil){
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": test,
		})

	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": err,
		})
	}
}