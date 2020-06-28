package api

import (
	"fmt"
	"net/http"
	
	"github.com/HachimanHiki/zkrpApi/service"
	"github.com/HachimanHiki/zkrpApi/selftype"
	"github.com/gin-gonic/gin"
)

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

func VerifyMerkleTreeRoot(c *gin.Context) {
	verifyMerkleTree := selftype.VerifyMerkleTree{}

	if c.BindJSON(&verifyMerkleTree) == nil {
		// for demo use
		if userHashRoot == nil {
			userHashRoot = make(map[string]string)
		}
		userHashRoot["\u738b\u6625\u5b0c"] = "dffb0a974110005d6fed4e17c8c29619cb50e9df0b3bc0a633cd80b317b9a36d"
		userHashRoot["\u674e\u5c0f\u8c6a"] = "22916c7ae2af50553a8cded598ca4f1917585d1d55ac2f265f7888f0054e5adf"
		userHashRoot["\u5f35\u5fd7\u660e"] = "2eb1eb753e41af3c3c19f0830834db24317fa30d521a7d70e86bb1086f0d628d"
		// 

		var hashArray []string

		for _, medicineUsage := range verifyMerkleTree.MedicineUsages {

			for len(hashArray) != (medicineUsage.ID - 1) {
				hashArray = append(hashArray, verifyMerkleTree.HashArray[0])
				verifyMerkleTree.HashArray = verifyMerkleTree.HashArray[1:]
			}
			hashArray = append(hashArray, service.GenerateHashFromStruct(medicineUsage))
		}

		for len(verifyMerkleTree.HashArray) != 0 {
			hashArray = append(hashArray, verifyMerkleTree.HashArray[0])
			verifyMerkleTree.HashArray = verifyMerkleTree.HashArray[1:]
		}

		if userHashRoot[verifyMerkleTree.UserName] == service.GenerateMerkleTreeRoot(hashArray) {
			resultStatus = true
			resultMessage = append(resultMessage, "Merkle tree verify successful with user name: " + verifyMerkleTree.UserName)

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "verify success with user name: " + verifyMerkleTree.UserName,
			})

		} else {
			resultStatus = true
			resultMessage = append(resultMessage, "Merkle tree verify failure with user name: " + verifyMerkleTree.UserName)

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