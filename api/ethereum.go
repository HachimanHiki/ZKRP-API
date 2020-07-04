package api

import (
	"fmt"
	"net/http"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {

	client, err := ethclient.Dial("http://localhost:8545")

	fmt.Println(*client)
	fmt.Println(err)

	account := common.HexToAddress("0xe280029a7867ba5c9154434886c241775ea87e53")
	balance, err := client.BalanceAt(context.Background(), account, nil)

	if(err == nil){
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": balance,
		})

	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": err,
		})
	}
}