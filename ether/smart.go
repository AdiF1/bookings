package ether

import (
	"context"
	"crypto/ecdsa"
	"log"
	//"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// MetaMaskWallets: 

	// 0x2355B93c4551a235315e14c1e76AcDDe0Bfcc4C9
	// 94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7

	// 0x4Cd481B57B13C14988498BfF28974eAC919e3862
	// 782336a0704da5f8feee75f66639a1adeb8203e57d14b176b44a032ecf3d38d1


	// 0xF9a20B1F7995B652313dC64D0FeEB451FD146522
	// 65a2f158fb1bbea7843b0be327dcbaed00f7021b34b6c1b4cb49606dad9239ef

	// apparent owner of all SCX1 tokens:
	// 0x5A86f0cafD4ef3ba4f0344C138afcC84bd1ED222


func SetUpParams (dial string, pky string) (cl *ethclient.Client, pk *ecdsa.PrivateKey, non uint64) {

	client, err := ethclient.Dial(dial)
	if err != nil {
		log.Fatal(err)
	}

	//crypto.Ecrecover()
	//crypto.PubkeyToAddress()

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	} else {
		//fmt.Println("header:" ,header.Number.String())
		_ = header
	}

	// the first step is to load your private key. (acc #1)
	privateKey, err := crypto.HexToECDSA(pky)
	if err != nil {
		log.Fatal(err)
	}
	/* 	Afterwards we need to get the account nonce. Every transaction requires a nonce. A nonce by definition 
		is a number that is only used once. If it's a new account sending out a transaction then the nonce will 
		be 0. Every new transaction from an account must have a nonce that the previous nonce incremented by 1. 
		It's hard to keep manual track of all the nonces so the ethereum client provides a helper method 
		PendingNonceAt that will return the next nonce you should use.

       The function requires the public address of the account we're sending from -- which we can derive from the private key.
	*/
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	/* 	privateKey.Public() returns an interface that contains our public key. We perform a type assertion with publicKey.
		(<expectedType>) to explicitly set the type of our publicKey variable, and assign it to publicKeyECDSA. This allows 
		us to use it where our program expects an input of type *ecdsa.PublicKey.

		Now we can read the nonce that we should use for the account's transaction.
	*/
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	//nonce = uint64(252)
	if err != nil {
		log.Fatal(err)
	}

	/* gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
  		log.Fatal(err)
	} */

	return client, privateKey, nonce

}