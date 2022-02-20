package ether

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/AdiF1/solidity/bookings/contracts"
)

func ContractDeploy() {
	//call SetUpParams() in smart.go
	client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	//auth.GasPrice = big.NewInt(3000000)
	auth.GasPrice = gasPrice
	
	//auth.Value  = "1.15"

	input := "1.0"
	address, tx, instance, err := adf2.DeployAdf2(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\naddress: ", address.Hex())

	fmt.Println("\n\ntx.Hash: ", tx.Hash().Hex())

	fmt.Println("\n\nnonce: ", auth.Nonce)

	/* whatup, err := address.Value()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("tx ChainID: ", whatup) */

	//fmt.Println(instance)

	_, _ = instance, tx

}