package api

import (
	"fmt"
	"net/http"
	
	"github.com/HachimanHiki/zkrpApi/service"
	"github.com/HachimanHiki/zkrpApi/selftype"
	"github.com/gin-gonic/gin"
)

// PostMerkleTreeRoot godoc
// @Tags MerkleTree 
// @Summary Add/Update user's merkleTreeRoot
// @Description Return user's merkleTreeRoot
// @Accept  application/json
// @Produce  application/json
// @Param medicineUsages body []selftype.MedicineUsage true "[]MedicineUsage"
// @Param userName body string true "userName"
// @Success 200 {object} selftype.JSONResponse
// @Failure 400 {object} selftype.JSONResponse
// @Router /merkletree [post]
func PostMerkleTreeRoot(c *gin.Context) {
	merkleTreeRequire := selftype.MerkleTreeRequire{}

	if c.BindJSON(&merkleTreeRequire) == nil {
		var hashArray []string

		for _, medicineUsage := range merkleTreeRequire.MedicineUsages {
			hashArray = append(hashArray, service.GenerateHashFromStruct(medicineUsage))
		}

		if userHashRoot == nil {
			userHashRoot = make(map[string]string)
		}
		fmt.Println(hashArray)
		userHashRoot[merkleTreeRequire.UserName] = service.GenerateMerkleTreeRoot(hashArray)
		
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message": userHashRoot[merkleTreeRequire.UserName],
		})

	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}

// VerifyMerkleTreeRoot godoc
// @Tags MerkleTree 
// @Summary Verify merkle tree root
// @Description Return result of verification
// @Accept  json
// @Produce  json
// @Param medicineUsages body []selftype.MedicineUsage true "[]MedicineUsage"
// @Param userName body string true "userName"
// @Param hashArray body []string true "[]hashArray"
// @Success 200 {object} selftype.EventResponse
// @Failure 400 {object} selftype.JSONResponse
// @Router /merkletree [get]
func VerifyMerkleTreeRoot(c *gin.Context) {
	verifyMerkleTree := selftype.VerifyMerkleTree{}

	if c.BindJSON(&verifyMerkleTree) == nil {
		// for demo use
		if userHashRoot == nil {
			userHashRoot = make(map[string]string)
		}
		userHashRoot["\u738b\u6625\u5b0c"] = "144395288b617a54e0eda5706c41857ecbb39b113bad83515c01fef20a6b3eb6"
		userHashRoot["\u674e\u5c0f\u8c6a"] = "fca3a82e4f649975626bb40a45442c381b2ee9f38699e6ff0af03c649b90d8a7"
		userHashRoot["\u5f35\u5fd7\u660e"] = "40a388b578be1a7443ba1258ac761f09bfd821aebb694e5109ee9035c30b3fef"
		// 

		var hashArray []string

		for i := len(verifyMerkleTree.MedicineUsages) - 1 ; i >= 0 ; i-- {
			medicineUsage := verifyMerkleTree.MedicineUsages[i]

			for len(hashArray) != (medicineUsage.ID - 1) {
				hashArray = append(hashArray, verifyMerkleTree.HashArray[len(verifyMerkleTree.HashArray)-1])
				verifyMerkleTree.HashArray = verifyMerkleTree.HashArray[:len(verifyMerkleTree.HashArray)-1]
			}
			hashArray = append(hashArray, service.GenerateHashFromStruct(medicineUsage))
		} 

		for len(verifyMerkleTree.HashArray) != 0 {
			hashArray = append(hashArray, verifyMerkleTree.HashArray[len(verifyMerkleTree.HashArray)-1])
			verifyMerkleTree.HashArray = verifyMerkleTree.HashArray[:len(verifyMerkleTree.HashArray)-1]
		}

		if userHashRoot[verifyMerkleTree.UserName] == service.GenerateMerkleTreeRoot(hashArray) {
			shareResultStatus = true
			shareResultMessage = append(shareResultMessage, "Merkle tree verify successful with user name: " + verifyMerkleTree.UserName)

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "verify success with user name: " + verifyMerkleTree.UserName,
			})

		} else {
			shareResultStatus = true
			shareResultMessage = append(shareResultMessage, "Merkle tree verify failure with user name: " + verifyMerkleTree.UserName)

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "verify failure with user name: " + verifyMerkleTree.UserName,
			})
		}

	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Bad request!",
		})
	}
}