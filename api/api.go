package api

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/HachimanHiki/zkrpApi/selftype"
	"github.com/gin-gonic/gin"
)

var (
	allEvent map[string]selftype.Event
	userHashRoot map[string]string
	zkrpResultStatus bool = false
	zkrpResultMessage []string
	shareResultStatus bool = false
	shareResultMessage []string
)

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
		allEvent["marathon"] = selftype.Event{}
		allEvent["share"] = selftype.Event{}
		//
		if _, ok := allEvent[eventName]; ok {
			
			qrCode, _ := qr.Encode("http://localhost:8080/event?eventName=" + eventName, qr.M, qr.Auto)

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