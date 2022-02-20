package ether

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockInfo() {
    client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
    if err != nil {
        log.Fatal(err)
    }

    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(header.Number.String()) // 5671744

    blockNumber := big.NewInt(5671744)
    block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("block.Number:", block.Number().Uint64())     // 5671744
    fmt.Println("block.Time:", block.Time())       // 1527211625
    fmt.Println("block.Difficulty:", block.Difficulty().Uint64()) // 3217000136609065
    fmt.Println("block.Hash:", block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
    fmt.Println("block.Transactions:", len(block.Transactions()))   // 144

    count, err := client.TransactionCount(context.Background(), block.Hash())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("count:", count) // 144	  
}