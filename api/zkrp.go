package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"io/ioutil"
	"bytes"
	"fmt"
	
	"github.com/HachimanHiki/zkrpApi/service"
	"github.com/HachimanHiki/zkrpApi/service/zsl"
	"github.com/HachimanHiki/zkrpApi/selftype"
	"github.com/gin-gonic/gin"
)

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
		/*
		if verify.UserName == "\u738b\u6625\u5b0c" {
			verify.Commitment = "0x19be17213c2ee781defeef4abc2d6964f1418177f8d54418c55412e198eebf2107c6e8cc1a43e9be4e7f331f7b729e53a2e14e37aa874383628cf89be3cd0ef6"

		} else if verify.UserName == "\u674e\u5c0f\u8c6a" {
			verify.Commitment = "0x159e049899d7b4e444459feb5e832c1d589b97f348ff7a37bd84d2cdc4b2bb5e653c3163346c4055f17d9d27c01016869a64bb6bb3a91e2558b8b3183d3f4720"

		} else {
			verify.Commitment = "0x1edbe504d08a8da27424f7815dc6bfb0ff18746fabc38f0b7ec05915da0d9b302cf345a6f4523163ac3254155d45329f1f5c8ac835b2fd558db2f5bd543aefb8"
		}
		*/
		verify.Commitment, _ = service.GetCommitment(contractAddress[verify.UserName])

		const layout = "20060102" // time format

		if zsl.Verifier(verify.Commitment, verify.Lowerbound, verify.Upperbound, []byte(verify.Prove)) {
			zkrpResultStatus = true
			zkrpResultMessage = append(zkrpResultMessage, "ZKRP verify failure with user name: " + verify.UserName) 

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "Verify failure with user name: " + verify.UserName,
				"data": verify.Code,
			})

		}else {
			zkrpResultStatus = true
			zkrpResultMessage = append(zkrpResultMessage, "ZKRP verify successful with user name: " + verify.UserName) 

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "Verify successful with user name: " + verify.UserName,
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
			//"message": "Verify successful with user name: " + verify.UserName,
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