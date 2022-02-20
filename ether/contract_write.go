package ether

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/AdiF1/solidity/bookings/contracts"
)

func ContractWrite() {
	//call SetUpParams() in smart.go
	client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")

	address := common.HexToAddress("0x0035Ba6E027cCdfBb861616eD6FE5e5f86ef1092")
    instance, err := adf2.NewAdf2(address, client)
    if err != nil {
        log.Fatalf("INSTANCE ERR ", err)
    }
	fmt.Println("\n\ncontract is loaded!")


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
	fmt.Println("\n\nnonce: ", auth.Nonce)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	//auth.GasPrice = big.NewInt(10)
	auth.GasPrice = gasPrice

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("Steve"))
	copy(value[:], []byte("Jobs"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\ntx sent: %s", tx.Hash()) // tx sent

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\nresult: ",string(result[:]))

}