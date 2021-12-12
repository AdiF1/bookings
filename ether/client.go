package ether

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"

)


func etherNet() (infuraClient, ganacheClient *ethclient.Client){

	infuraClient, err := ethclient.Dial("https://mainnet.infura.io") 
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("We have an Infura connection!")
	 // we'll use this in the upcoming sections 


	ganacheClient, err = ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("We have a Ganache connection!")
	return infuraClient, ganacheClient
}