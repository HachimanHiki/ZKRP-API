package api

import (
	"bytes"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"io/ioutil"
	"image/png"
	//"context"
	//"log"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
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

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", nil)
}

func NewProve (c *gin.Context) {
	proveRequired := selftype.ProveRequired{}

	if c.BindJSON(&proveRequired) == nil {
		const layout = "20060102" // time format
		var ProvePackages []selftype.ProvePackage

		upperbound, _ := strconv.ParseInt(time.Now().Format(layout), 10, 64) // today

		for _, hospitalInfo := range proveRequired.HospitalInfo {
			t, _ := time.Parse(layout, time.Now().Format(layout))
			t = t.AddDate(0, 0, -28)
			lowerbound, _ := strconv.ParseInt(t.Format(layout), 10, 64) // "hospitalInfo.DuringTime" day before today

			number, _ := strconv.ParseInt(hospitalInfo.OutHospitalDate, 10, 64)

			commitmentPackage := zsl.Committer(int(number))

			ProvePackage := selftype.ProvePackage {
				Type: "DiseaseID",
				Code: hospitalInfo.DiseaseID,
				Prove: zsl.Prover(number, lowerbound, upperbound, commitmentPackage.Confounding),
				Lowerbound: lowerbound,
				Upperbound: upperbound,
				Commitment: commitmentPackage.Commitment,
			}
			
			ProvePackages = append(ProvePackages, ProvePackage)
		}

		for _, hospitalInfo := range proveRequired.HospitalInfo {
			t, _ := time.Parse(layout, time.Now().Format(layout))
			t = t.AddDate(0, 0, -28)
			lowerbound, _ := strconv.ParseInt(t.Format(layout), 10, 64) // "hospitalInfo.DuringTime" day before today

			number, _ := strconv.ParseInt(hospitalInfo.OutHospitalDate, 10, 64)

			commitmentPackage := zsl.Committer(int(number))

			ProvePackage := selftype.ProvePackage {
				Type: "OperationID",
				Code: hospitalInfo.OperationID,
				Prove: zsl.Prover(number, lowerbound, upperbound, commitmentPackage.Confounding),
				Lowerbound: lowerbound,
				Upperbound: upperbound,
				Commitment: commitmentPackage.Commitment,
			}
			
			ProvePackages = append(ProvePackages, ProvePackage)
		}

		for _, westernMedicine := range proveRequired.WesternMedicine {
			t, _ := time.Parse(layout, time.Now().Format(layout))
			t = t.AddDate(0, 0, -28)
			lowerbound, _ := strconv.ParseInt(t.Format(layout), 10, 64) // "procedure.DuringTime" day before today

			number, _ := strconv.ParseInt(westernMedicine.MedicalDate, 10, 64)

			commitmentPackage := zsl.Committer(int(number))

			ProvePackage := selftype.ProvePackage {
				Type: "DiseaseID",
				Code: westernMedicine.DiseaseID,
				Prove: zsl.Prover(number, lowerbound, upperbound, commitmentPackage.Confounding),
				Lowerbound: lowerbound,
				Upperbound: upperbound,
				Commitment: commitmentPackage.Commitment,
			}

			ProvePackages = append(ProvePackages, ProvePackage) 
		}

		fmt.Println(ProvePackages)

		jsonValue, _ := json.Marshal(ProvePackages)
		jsonString := `{"provePackages":` + string(jsonValue) + "}"
		jsonValue = []byte(jsonString)

		res, _ := http.Post("http://localhost:8080/verify", "application/json", bytes.NewBuffer(jsonValue))
		fmt.Println(res.Body)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": string(body),
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
// @Param userName body string true "userName"
// @Success 200 {object} selftype.JSONResponse
// @Failure 400 {object} selftype.JSONResponse
// @Router /verify [post]
func PostVerify(c *gin.Context) {
	verify := selftype.Verify{}
	verify.Commitment = "0x100b92ff311c1f05dfbe31dbcd8f416245bb0ab7bd800afe7ecc402a5c2480fc4ee9f329d7d7ddb8ddab7f99fb25e3f439144046782d95795bf8c0726a24c3ec"

	if c.BindJSON(&verify) == nil {
	// mongodb
	/*	
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

		if err != nil {
			log.Fatal(err)
		}
		// Check the connection
		err = client.Ping(context.Background(), nil)
	
		if err != nil {
			log.Fatal(err)
		}
		
		type db struct {
			Commitment string
			UserName string 
		}
		collection := client.Database("test").Collection("test1")
		filter := bson.M{"userName": verify.UserName}
	
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil { log.Fatal(err) }
		defer cur.Close(ctx)
		//for cur.Next(ctx) {
		   var result db

		   err = collection.FindOne(context.Background(), filter).Decode(&result)
		   if err != nil { log.Fatal(err) }

		   fmt.Println(result)
		   verify.Commitment = result.Commitment
		//}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
*/
		if verify.UserName == "\u738b\u6625\u5b0c" {
			verify.Commitment = "0x19be17213c2ee781defeef4abc2d6964f1418177f8d54418c55412e198eebf2107c6e8cc1a43e9be4e7f331f7b729e53a2e14e37aa874383628cf89be3cd0ef6"

		} else if verify.UserName == "\u674e\u5c0f\u8c6a" {
			verify.Commitment = "0x1d68c88fe825373b1c40afb557e0e0ae62580b22f8a1ae93198eca1e64048405ff3bd5fd5725a5d42ae541e6187861fe12aac833d8bf5f889f59b71936039181"

		} else {
			verify.Commitment = "0x1edbe504d08a8da27424f7815dc6bfb0ff18746fabc38f0b7ec05915da0d9b302cf345a6f4523163ac3254155d45329f1f5c8ac835b2fd558db2f5bd543aefb8"
		}

		const layout = "20060102" // time format

		if zsl.Verifier(verify.Commitment, verify.Lowerbound, verify.Upperbound, []byte(verify.Prove)) {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "Verify success with user name: " + verify.UserName,
			})

		}else {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "Verify fail with user name: " + verify.UserName,
			})
		}
/*
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
					//"message": "Verify fail with user name: " + verify.UserName,
					"message": "false",
				})
				return
			}
		}*/
/*
		c.HTML(http.StatusOK, "new.tmpl", gin.H{
			"title": "IT HOME again",
		})*/
/*
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			//"message": "Verify success with user name: " + verify.UserName,
			"message": "true",
		})
*/
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
	const verifyURLPrefix = "http://localhost:8080/verify"
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