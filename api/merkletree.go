package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	
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
		var hashArray, hashArray1, hashArray2 []string
		for _, medicineUsage := range merkleTreeRequire.MedicineUsages {
			j, _ := json.Marshal(medicineUsage)
			fmt.Println(string(j))
			hashArray1 = append(hashArray1, service.GenerateHashFromJsonString(string(j)))
		}

		for _, medicineUsage := range merkleTreeRequire.WesternMedicines {
			j, _ := json.Marshal(medicineUsage)
			hashArray2 = append(hashArray2, service.GenerateHashFromJsonString(string(j)))
		}

		if userHashRoot == nil {
			userHashRoot = make(map[string]string)
		}
		
		hashArray = append(hashArray, service.GenerateMerkleTreeRoot(hashArray1))
		hashArray = append(hashArray, service.GenerateMerkleTreeRoot(hashArray2))

		fmt.Println(hashArray1)
		fmt.Println(hashArray2)
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
		/* medicine
		userHashRoot["\u738b\u6625\u5b0c"] = "144395288b617a54e0eda5706c41857ecbb39b113bad83515c01fef20a6b3eb6"
		userHashRoot["\u674e\u5c0f\u8c6a"] = "fca3a82e4f649975626bb40a45442c381b2ee9f38699e6ff0af03c649b90d8a7"
		userHashRoot["\u5f35\u5fd7\u660e"] = "40a388b578be1a7443ba1258ac761f09bfd821aebb694e5109ee9035c30b3fef"
		*/
		/* western
		userHashRoot["\u738b\u6625\u5b0c"] = "ab254dcc9954682081ee568367ad62c4e6348574f510eed0a6efdabc4852c833"
		userHashRoot["\u674e\u5c0f\u8c6a"] = "01a39831ceb9d6a01a4a84a5b4d09abc8fd4eb1411e8fd460d616ce2ce1d7fae"
		userHashRoot["\u5f35\u5fd7\u660e"] = "2c25458e174351f4220f572a633f97e2970541c3d5659a11cab8e682858d3a96"
		*/
		/* total
		userHashRoot["\u738b\u6625\u5b0c"] = "d248f0355b955dad7d88be03cf279654bd8ebbbbc8d6302ae19ff34c26143eae"
		userHashRoot["\u674e\u5c0f\u8c6a"] = "8aa08b98e39aacc3f4d3846a46e440789869ed6dc401ea1595bb69db11bbd8e6"
		userHashRoot["\u5f35\u5fd7\u660e"] = "724a9dbd73a2c83a4ec8fb1244e2661e5acd77963ebbfdad65df08f3260e6a21"
		*/

		merkleTreeRoot, _ := service.GetMerkleTreeRoot(contractAddress[verifyMerkleTree.UserName])
		userHashRoot[verifyMerkleTree.UserName] = merkleTreeRoot
		fmt.Println(userHashRoot[verifyMerkleTree.UserName])
		//

		var hashArray []string
/*
		hashArray = append(hashArray, "40a388b578be1a7443ba1258ac761f09bfd821aebb694e5109ee9035c30b3fef")
		hashArray = append(hashArray, "2c25458e174351f4220f572a633f97e2970541c3d5659a11cab8e682858d3a96")
		fmt.Println(service.GenerateMerkleTreeRoot(hashArray))
*/

		hashArray = append(hashArray, service.GenerateHashFromStructAndHash(verifyMerkleTree.MedicineUsagesWithHashArray.MedicineUsages, verifyMerkleTree.MedicineUsagesWithHashArray.MedicineHashArray))
		hashArray = append(hashArray, service.GenerateHashFromStructAndHash(verifyMerkleTree.WesternMedicineWithHashArray.WesternMedicines, verifyMerkleTree.WesternMedicineWithHashArray.WesternHashArray))
		
		if userHashRoot[verifyMerkleTree.UserName] == service.GenerateMerkleTreeRoot(hashArray) {
			shareResultStatus = true
			shareResultMessage = append(shareResultMessage, "MerkleTree verify successful with user name: " + verifyMerkleTree.UserName)

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"message": "verify successful with user name: " + verifyMerkleTree.UserName,
			})

		} else {
			shareResultStatus = true
			shareResultMessage = append(shareResultMessage, "MerkleTree verify failure with user name: " + verifyMerkleTree.UserName)

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