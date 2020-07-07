package api

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"image/png"

	"github.com/HachimanHiki/zkrpApi/service"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/HachimanHiki/zkrpApi/selftype"
	"github.com/gin-gonic/gin"
)

var (
	allEvent map[string]selftype.Event
	zkrpResultStatus bool = false
	zkrpResultMessage []string
	shareResultStatus bool = false
	shareResultMessage []string
	contractAddress map[string]string
	userPrivateKey map[string]string
)

func init() {

	if contractAddress == nil {
		contractAddress = make(map[string]string)
	}

	if userPrivateKey == nil {
		userPrivateKey = make(map[string]string)
	}

	add, _ := service.DeployContract("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5", 
		"0x19be17213c2ee781defeef4abc2d6964f1418177f8d54418c55412e198eebf2107c6e8cc1a43e9be4e7f331f7b729e53a2e14e37aa874383628cf89be3cd0ef6", 
		"d248f0355b955dad7d88be03cf279654bd8ebbbbc8d6302ae19ff34c26143eae")

	contractAddress["\u738b\u6625\u5b0c"] = add
	userPrivateKey["\u738b\u6625\u5b0c"] = "f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5"
	
	add, _ = service.DeployContract("91821f9af458d612362136648fc8552a47d8289c0f25a8a1bf0860510332cef9", 
		"0x159e049899d7b4e444459feb5e832c1d589b97f348ff7a37bd84d2cdc4b2bb5e653c3163346c4055f17d9d27c01016869a64bb6bb3a91e2558b8b3183d3f4720", 
		"8aa08b98e39aacc3f4d3846a46e440789869ed6dc401ea1595bb69db11bbd8e6")

	contractAddress["\u674e\u5c0f\u8c6a"] = add
	userPrivateKey["\u674e\u5c0f\u8c6a"] = "91821f9af458d612362136648fc8552a47d8289c0f25a8a1bf0860510332cef9"


	add, _ = service.DeployContract("bb32062807c162a5243dc9bcf21d8114cb636c376596e1cf2895ec9e5e3e0a68", 
		"0x1edbe504d08a8da27424f7815dc6bfb0ff18746fabc38f0b7ec05915da0d9b302cf345a6f4523163ac3254155d45329f1f5c8ac835b2fd558db2f5bd543aefb8", 
		"724a9dbd73a2c83a4ec8fb1244e2661e5acd77963ebbfdad65df08f3260e6a21")
	
	contractAddress["\u5f35\u5fd7\u660e"] = add
	userPrivateKey["\u5f35\u5fd7\u660e"] = "bb32062807c162a5243dc9bcf21d8114cb636c376596e1cf2895ec9e5e3e0a68"
}

// NotFound godoc
func NotFound(c *gin.Context) {
    c.JSON(http.StatusNotFound, gin.H{
        "status": "fail",
        "error":  "Page not exists!",
    })
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", nil)
}

func GetMarathon(c *gin.Context) {
	c.HTML(http.StatusOK, "marathon.html", nil)
}

func GetShare(c *gin.Context) {
	c.HTML(http.StatusOK, "share.html", nil)
}

func GetZkrpResult(c *gin.Context) {

	if(zkrpResultStatus == true){
		c.JSON(http.StatusOK, gin.H{
			"status": zkrpResultStatus,
			"message": zkrpResultMessage,
		})

		zkrpResultStatus = false
		zkrpResultMessage = zkrpResultMessage[:0]

	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}

func GetShareResult(c *gin.Context) {

	if(shareResultStatus == true){
		c.JSON(http.StatusOK, gin.H{
			"status": shareResultStatus,
			"message": shareResultMessage,
		})

		shareResultStatus = false
		shareResultMessage = shareResultMessage[:0]

	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}

// GetQRcode godoc
// @Tags Event 
// @Summary Get an qrcode
// @Description Return a qrcode which encodes event info
// @Accept  json
// @Produce  json
// @Param eventName query string true "event search by eventName"
// @Success 200 {object} selftype.EventResponse
// @Failure 400 {object} selftype.JSONResponse
// @Router /event_qr/{eventName} [get]
func GetQRcode (c *gin.Context) {
	eventName := c.Query("eventName")
	if len(eventName) > 0 {
		// for demo
		if allEvent == nil {
			allEvent = make(map[string]selftype.Event)
		}
		allEvent["第1屆 指南馬拉松"] = selftype.Event{}
		allEvent["指南臨床試驗 政府立案 字號 NCTXXXXXX"] = selftype.Event{}
		//
		if _, ok := allEvent[eventName]; ok {
			
			qrCode, _ := qr.Encode("http://140.119.19.121:8080/event?eventName=" + eventName, qr.M, qr.Auto)

			// Scale the barcode to 200x200 pixels
			qrCode, _ = barcode.Scale(qrCode, 200, 200)

			var buf bytes.Buffer
			err :=png.Encode(&buf, qrCode)
			if err != nil {
				panic(err)
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data": allEvent[eventName],
				"qrcode": base64.StdEncoding.EncodeToString(buf.Bytes()),
			})

		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"message": "Cannot find this event",
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}