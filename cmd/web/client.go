/* package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func openClient() {
	client, err := ethclient.Dial("https://mainnet.infura.io") 
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections 
} */