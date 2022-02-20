package ether // page 69

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "golang.org/x/crypto/sha3"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/core/types"
)

func TransferTokens() {
    //call SetUpParams() in smart.go
    client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")

    // Set the value of the transaction to 0.
    // This value is the amount of ETH to be transferred for this transaction, which should 
    // be 0 since we're transferring ERC-20 Tokens and not ETH. We'll set the value of Tokens 
    // to be transferred in the data field later.
    value := big.NewInt(0) // in wei (0 eth)



	gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    // Then, store the address you'll be sending tokens to in a variable.
    toAddress := common.HexToAddress("0x4Cd481B57B13C14988498BfF28974eAC919e3862")

    // First, let's assign the token contract address to a variable.
    tokenAddress := common.HexToAddress("0x1a72581c2DdC3bb87B03Bd192c660fF04d6F8d74")


// method set up...
    transferFnSignature := []byte("transfer(address,uint256)")
    hash := sha3.NewLegacyKeccak256()
    hash.Write(transferFnSignature)
    methodID := hash.Sum(nil)[:4]
    fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

    paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
    fmt.Println("paddedAddress: ", hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

    amount := new(big.Int)
    amount.SetString("950000000000000000000", 10) // sets the value to 1000 tokens, in the token denomination

    paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
    fmt.Println("paddedAmount: ", hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

    var data []byte
    data = append(data, methodID...)
    data = append(data, paddedAddress...)
    data = append(data, paddedAmount...)

    /* gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
        To:   &tokenAddress,
        Data: data,
    })
    if err != nil {
		fmt.Println("1: ")
        log.Fatal(err)
    }
    fmt.Println("gasLimit: ", gasLimit) // 23256 */

    tx := types.NewTransaction(nonce, tokenAddress, value, 2223256, gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal("2: ", err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal("3: ", err)
    }
	fmt.Printf("token tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}

// 0x5a86f0cafd4ef3ba4f0344c138afcc84bd1ed222 Holder of all SCX1 tokens