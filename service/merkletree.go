package service 

import (
	"crypto/sha256"
	"fmt"
	"encoding/json"

	"github.com/HachimanHiki/zkrpApi/selftype"
)

func GenerateHashFromJsonString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
	//return h.Sum(nil)
}

func GenerateHashFromHash(a, b string) string {
	h := sha256.New()
	h.Write(append([]byte(a), []byte(b)...))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateMerkleTreeRoot(hashArray []string) string {

	if len(hashArray) == 1 {
		return hashArray[0]
	}
	
	if len(hashArray) % 2 != 0 {
		hashArray = append(hashArray, hashArray[len(hashArray)-1])
	}

	var newArray []string
	
	for i := 0 ; i < len(hashArray) ; i += 2 {
		newArray = append(newArray, GenerateHashFromHash(hashArray[i], hashArray[i+1]))
	}

	return GenerateMerkleTreeRoot(newArray)
	
}

func GenerateHashFromStructAndHash(informathionArray interface{}, hashArray []string) string{
	var tmpHashArray []string

	switch iType := informathionArray.(type) {
		case []selftype.MedicineUsage:
			for i := len(iType) - 1 ; i >= 0 ; i-- {
		
				informathion := iType[i]		
		
				for len(tmpHashArray) != (informathion.ID - 1) {
					tmpHashArray = append(tmpHashArray, hashArray[len(hashArray)-1])
					hashArray = hashArray[:len(hashArray)-1]
				}
				j, _ := json.Marshal(informathion)
				tmpHashArray = append(tmpHashArray, GenerateHashFromJsonString(string(j)))
			} 
		
			for len(hashArray) != 0 {
				tmpHashArray = append(tmpHashArray, hashArray[len(hashArray)-1])
				hashArray = hashArray[:len(hashArray)-1]
			}

		case []selftype.WesternMedicine:
			for i := len(iType) - 1 ; i >= 0 ; i-- {
		
				informathion := iType[i]		
		
				for len(tmpHashArray) != (informathion.ID - 1) {
					tmpHashArray = append(tmpHashArray, hashArray[len(hashArray)-1])
					hashArray = hashArray[:len(hashArray)-1]
				}
				j, _ := json.Marshal(informathion)
				tmpHashArray = append(tmpHashArray, GenerateHashFromJsonString(string(j)))
			} 
		
			for len(hashArray) != 0 {
				tmpHashArray = append(tmpHashArray, hashArray[len(hashArray)-1])
				hashArray = hashArray[:len(hashArray)-1]
			}
	}

	return GenerateMerkleTreeRoot(tmpHashArray)
}