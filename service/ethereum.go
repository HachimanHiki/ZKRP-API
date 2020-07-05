package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	health "github.com/HachimanHiki/zkrpApi/contract"
)

// DeployContract return contract address
func DeployContract(private, c, m string) (string, error){
	
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}	

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	//inputC := "0x19be17213c2ee781defeef4abc2d6964f1418177f8d54418c55412e198eebf2107c6e8cc1a43e9be4e7f331f7b729e53a2e14e37aa874383628cf89be3cd0ef6"
	//inputM := "d248f0355b955dad7d88be03cf279654bd8ebbbbc8d6302ae19ff34c26143eae"
	address, _, _, err := health.DeployHealth(auth, client, c, m)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return address.Hex(), nil
}

func GetCommitment(add string) (string, error){

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(add)
	instance, err := health.NewHealth(address, client)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	return instance.Commitment(nil)
}

func GetMerkleTreeRoot(add string) (string, error){

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(add)
	instance, err := health.NewHealth(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance.MerkleTreeRoot(nil)
}