package api

import (
	"net/http"
	"strconv"
	"time"
	//"fmt"

	//"github.com/zkrpApi/zsl/zkrp"
	"github.com/HachimanHiki/zkrpApi/selftype"
	"github.com/HachimanHiki/zkrpApi/zsl"
    "github.com/gin-gonic/gin"
)

var (
	allEvent map[string]selftype.Event
)

// NotFound godoc
func NotFound(c *gin.Context) {
    c.JSON(http.StatusNotFound, gin.H{
        "status": "fail",
        "error":  "Page not exists!",
    })
}

// PostUserProve godoc
// @Tags ZKRP 
// @Summary Prove ZKRP
// @Description Give a prove of ZKRP
// @Accept  application/json
// @Produce  application/json
// @Param requestTime body string true "requestTime"
// @Param deagnosis body []selftype.DeagnosisDate true "[]Deagnosis" 
// @Param procedure body []selftype.ProcedureDate true "[]Procedure" 
// @Success 200 {object} selftype.ProvePackagesResponse
// @Failure 400 {object} selftype.JSONResponse
// @Router /prove [post]
func PostUserProve(c *gin.Context) {
	proveRequired := selftype.ProveRequired{}

	if c.BindJSON(&proveRequired) == nil {
		const layout = "20060102" // time format
		var ProvePackages []selftype.ProvePackage

		upperbound, _ := strconv.ParseInt(proveRequired.RequestTime, 10, 64) // today

		for _, deagnosis := range proveRequired.Deagnosis {
			t, _ := time.Parse(layout, proveRequired.RequestTime)
			t = t.AddDate(0, 0, -28)
			lowerbound, _ := strconv.ParseInt(t.Format(layout), 10, 64) // 28 day before today

			number, _ := strconv.ParseInt(deagnosis.Date, 10, 64)

			commitmentPackage := zsl.Committer(int(number))

			ProvePackage := selftype.ProvePackage {
				Type: "Deagnosis",
				Code: deagnosis.DeagnosisCode,
				Prove: zsl.Prover(number, lowerbound, upperbound, commitmentPackage.Confounding),
				Lowerbound: lowerbound,
				Upperbound: upperbound,
				Commitment: commitmentPackage.Commitment,
			}
			
			ProvePackages = append(ProvePackages, ProvePackage)
		}

		for _, procedure := range proveRequired.Procedure {
			t, _ := time.Parse(layout, proveRequired.RequestTime)
			t = t.AddDate(0, 0, -28) 
			lowerbound, _ := strconv.ParseInt(t.Format(layout), 10, 64) // 28 day before today

			number, _ := strconv.ParseInt(procedure.Date, 10, 64)

			commitmentPackage := zsl.Committer(int(number))

			ProvePackage := selftype.ProvePackage {
				Type: "Procedure",
				Code: procedure.ProcedureCode,
				Prove: zsl.Prover(number, lowerbound, upperbound, commitmentPackage.Confounding),
				Lowerbound: lowerbound,
				Upperbound: upperbound,
				Commitment: commitmentPackage.Commitment,
			}

			ProvePackages = append(ProvePackages, ProvePackage) 
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": ProvePackages,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}

// PostVerify godoc
// @Tags ZKRP 
// @Summary Verify ZKRP
// @Description Verify a prove of ZKRP
// @Accept  application/json
// @Produce  application/json
// @Param provePackages body []selftype.ProvePackage true "[]ProvePackage" 
// @Success 200 {object} selftype.JSONResponse
// @Failure 400 {object} selftype.JSONResponse
// @Router /verify [post]
func PostVerify(c *gin.Context) {
	verify := selftype.Verify{}

	if c.BindJSON(&verify) == nil {
		const layout = "20060102" // time format

		result := make(chan bool, len(verify.ProvePackages))

		for _, provePackage := range verify.ProvePackages {
			go func(provePackage selftype.ProvePackage) {
				result <- zsl.Verifier(provePackage.Commitment, provePackage.Lowerbound, provePackage.Upperbound, []byte(provePackage.Prove))
			} (provePackage)
		}

		for i := 0; i < len(verify.ProvePackages); i++ {
			if ! <-result {
				c.JSON(http.StatusOK, gin.H{
					"status": "success",
					"message": "verify fail",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message": "verify success",
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}

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
	//eventName := selftype.EventName{}
	eventName := c.Query("eventName")
	
	if len(eventName) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": allEvent[eventName],
		})

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
				"message": "null",
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"message": "Cannot find eventName",
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}