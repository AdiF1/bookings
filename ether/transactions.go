package ether

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

func GetTransactions() {
    client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
    if err != nil {
        log.Fatal(err)
    }

    blockNumber := big.NewInt(5671744)
    block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal(err)
    }

	// We can read the transactions in a block by calling the Transactions method which returns 
	// a list of Transaction type. It's then trivial to iterate over the collection and retrieve 
	// any information regarding the transaction.
    for _, tx := range block.Transactions() {
        fmt.Println("1: ", tx.Hash().Hex())        // 
        fmt.Println("2: ", tx.Value().String())    // 
        fmt.Println("3: ", tx.Gas())               // 
        fmt.Println("4: ", tx.GasPrice().Uint64()) // 
        fmt.Println("5: ", tx.Nonce())             // 
        fmt.Println("6: ", tx.Data())              // 
        fmt.Println("7: ", tx.To().Hex())          // 
	

		// In order to read the sender address, we need to call AsMessage on the transaction 
		// which returns a Message type containing a function to return the sender (from) 
		// address. The AsMessage method requires the EIP155 signer, which we derive the chain 
		// ID from the client.
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err == nil {
			fmt.Println("8: ", msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		// Each transaction has a receipt which contains the result of the execution of the transaction, 
		// such as any return values and logs, as well as the status which will be 1 (success) or 0 (fail)
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("9: ", receipt.Status) // 1
		fmt.Println("9a: ", receipt.Logs)
	}
	// Another way to iterate over transaction without fetching the block is to call the 
	// client's TransactionInBlock method. This method accepts only the block hash and the 
	// index of the transaction within the block. You can call TransactionCount to know 
	// how many transactions there are in the block.
    blockHash := common.HexToHash("0xa4851499c7c704163eae5bbb1c49423746fe2c566fc23793bd94c6247933e90b")
    count, err := client.TransactionCount(context.Background(), blockHash)
    if err != nil {
        log.Fatal("!", err)
    }

    for idx := uint(0); idx < count; idx++ {
        tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
        if err != nil {
            log.Fatal("!!", err)
        }

        fmt.Println("10: ", tx.Hash().Hex()) // 
    }

	// “query for a single transaction directly given the transaction hash by using TransactionByHash
    txHash := common.HexToHash("0xec355e5f4f67ef07fc69ca91b51d1552f7e6b02a22a2105dbaf6a7e7e775f660")
    tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
    if err != nil {
        log.Fatal("!!!", err)
    }

    fmt.Println("11: ", tx.Hash().Hex())
    fmt.Println("12: ", isPending) 
}