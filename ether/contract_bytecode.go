package ether 

import (

	"context"
	"encoding/hex"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"

)

func ContractBytecode () {
	//call SetUpParams() in smart.go
	client, _, _ := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")


	contractAddress := common.HexToAddress("0x0035Ba6E027cCdfBb861616eD6FE5e5f86ef1092")
	byteCode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(byteCode))
}