package api

import (
	"net/http"
	
	"github.com/HachimanHiki/zkrpApi/selftype"
	"github.com/gin-gonic/gin"
)

// PostEvent godoc
// @Tags Event 
// @Summary Add/Update event
// @Description null
// @Accept  application/json
// @Produce  application/json
// @Param eventName body string true "EventName"
// @Param deagnosis body []selftype.Deagnosis true "[]Deagnosis" 
// @Param procedure body []selftype.Procedure true "[]Procedure" 
// @Success 200 {object} selftype.EventResponse
// @Failure 400 {object} selftype.JSONResponse
// @Router /event [post]
func PostEvent (c *gin.Context) {
	event := selftype.Event{}
	const verifyURLPrefix = "http://140.119.19.121:8080/verify"
	event.VarifyURL = verifyURLPrefix

	if c.BindJSON(&event) == nil {
		if allEvent == nil {
			allEvent = make(map[string]selftype.Event)
		}

		allEvent[event.EventName] = event

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": allEvent[event.EventName],
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}

// GetEvent godoc
// @Tags Event 
// @Summary Get an event
// @Description Get an event from string
// @Accept  json
// @Produce  json
// @Param eventName query string true "event search by eventName"
// @Success 200 {object} selftype.EventResponse
// @Failure 400 {object} selftype.JSONResponse
// @Router /event/{eventName} [get]
func GetEvent (c *gin.Context) {
	eventName := c.Query("eventName")

	// for demo
	if allEvent == nil {
		allEvent = make(map[string]selftype.Event)
	}
/*
	var diseaseInfoSlice []selftype.DiseaseInfo
	diseaseInfo := selftype.DiseaseInfo {
		DiseaseID: "R",
		DiseaseName: "\u75f0\u7570\u5e38",
		RequireDays: 28,
	}
*/
	allEvent["第1屆 指南馬拉松"] = selftype.Event{
		EventName: "第1屆 指南馬拉松",
		EventInfo: "我們將從您的健康存摺加密處理後交給馬拉松中心進行第三方驗證 確定您是否符合馬拉松資格",
		EventType: "zkrp",
		EventRequired: append([]string{}, "西醫門診", "住院/手術資料"),
		EventRequiredDetail: append([]string{}, "7天內無呼吸道症狀、無腹瀉症狀", "365天內無冠心病症狀、無貧血症狀、無高血壓症狀"),
		//DiseaseInfo: append(diseaseInfoSlice, diseaseInfo),
		VarifyURL: "http://140.119.19.121:8080/verify",
	}

	allEvent["指南臨床試驗"] = selftype.Event{
		EventName: "指南臨床試驗",
		EventInfo: "指南臨床試驗 政府立案 字號 NCTXXXXXX",
		EventType: "merkletree",
		EventRequired: append([]string{}, "自行選擇提供"),
		EventRequiredDetail: append([]string{}, "無"),
		VarifyURL: "http://140.119.19.121:8080/verify",
	}
	//

	if len(eventName) > 0 {
		if _, ok := allEvent[eventName]; ok {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data": allEvent[eventName],
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

// DelEvent godoc
// @Tags Event
// @Summary Delete an event
// @Accept  json
// @Produce  json
// @Param eventName query string true "event delete by eventName"
// @Success 200 {object} selftype.JSONResponse
// @Failure 400 {object} selftype.JSONResponse
// @Failure 404 {object} selftype.JSONResponse
// @Router /event/{eventName} [delete]
func DelEvent (c *gin.Context) {
	eventName := c.Query("eventName")

	if len(eventName) > 0 {
		if _, ok := allEvent[eventName]; ok {
			delete(allEvent, eventName)

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
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