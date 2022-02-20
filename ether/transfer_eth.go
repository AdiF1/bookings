package ether // page 61

import (
    "context"
    //"crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    //"github.com/ethereum/go-ethereum/crypto"
    //"github.com/ethereum/go-ethereum/ethclient"
)

func TransferEther() {
    //call SetUpParams() in smart.go
    client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")

    /* client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
    if err != nil {
        log.Fatal(err)
    } */

	// the first step is to load your private key. (acc #1)
    /* privateKey, err := crypto.HexToECDSA("94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")
    if err != nil {
        log.Fatal(err)
} */
	/* 	Afterwards we need to get the account nonce. Every transaction requires a nonce. A nonce by definition 
		is a number that is only used once. If it's a new account sending out a transaction then the nonce will 
		be 0. Every new transaction from an account must have a nonce that the previous nonce incremented by 1. 
		It's hard to keep manual track of all the nonces so the ethereum client provides a helper method 
		PendingNonceAt that will return the next nonce you should use.

       The function requires the public address of the account we're sending from -- which we can derive from the private key.
	*/
    /* publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) */

	/* 	privateKey.Public() returns an interface that contains our public key. We perform a type assertion with publicKey.
		(<expectedType>) to explicitly set the type of our publicKey variable, and assign it to publicKeyECDSA. This allows 
		us to use it where our program expects an input of type *ecdsa.PublicKey.

		Now we can read the nonce that we should use for the account's transaction.
	*/
    /* nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    } */

	/*	The next step is to set the amount of ETH that we'll be transferring. However we must convert ether to wei s
		ince that's what the Ethereum blockchain uses. Ether supports up to 18 decimal places so 1 ETH is 1 plus 18 zeros. 
		Here's a little tool to help you convert between ETH and wei: https://etherconverter.netlify.com
		The gas limit for a standard ETH transfer is 21000 units.
	*/
    value := big.NewInt(100000000000000000) // in wei (0.1 eth)
    gasLimit := uint64(21000)                // in units

	/*	Gas prices are always fluctuating based on market demand and what users are willing to pay, so hardcoding a 
		gas price is sometimes not ideal. The go-ethereum client provides the SuggestGasPrice function for getting 
		the average gas price based on x number of previous blocks.
	*/
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

	// figure out who we're sending the ETH to... (acc #3)
    toAddress := common.HexToAddress("0xF9a20B1F7995B652313dC64D0FeEB451FD146522")

	/* Now we can finally generate our unsigned ethereum transaction by importing the go-ethereum core/types package 
		and invoking NewTransaction which takes in the nonce, to address, value, gas limit, gas price, and optional data. 
		The data field is nil for just sending ETH. We'll be using the data field when it comes to interacting with smart 
		contracts.
	*/
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	
	/* sign the transaction with the private key of the sender. To do this we call the SignTx method that takes in the 
		unsigned transaction and the private key that we constructed earlier. The SignTx method requires the EIP155 signer, 
		which we derive the chain ID from the client.
	*/
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

	/* Now we are finally ready to broadcast the transaction to the entire network by calling SendTransaction on the client 
		which takes in the signed transaction. 
	*/
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}