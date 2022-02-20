package ether

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "github.com/AdiF1/solidity/bookings/contracts"
)

func ContractQuery() {
    client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress("0x0035Ba6E027cCdfBb861616eD6FE5e5f86ef1092")
    instance, err := adf2.NewAdf2(address, client)
    if err != nil {
        log.Fatalf("INSTANCE ERR ", err)
    }

    fmt.Println("instance: ", instance.Adf2Transactor)

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatalf("adf ERR ", err)
	}

	fmt.Println("message: ", version)

    // _ = instance

}

// 0x65D416060CA6E384CceAC8c6C6aB6d33BeBB3c42

// 0x1ae92b10E8BA728C375F0d3a69070d2b0B5d85B0
