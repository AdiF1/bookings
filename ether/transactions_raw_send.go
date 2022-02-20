package ether

import (
    "context"
    "encoding/hex"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/rlp"
)

func TransactionRawSend() {
    client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
	   if err != nil {
        	log.Fatal(err)
    	}

    rawTx := "f86c81f3843b9aca09825208944cd481b57b13c14988498bff28974eac919e3862880de0b6b3a7640000802ba083e7606ff72c203925fc9d179b1ca2e29742e847a113a4760231eda1b0a27658a00713965ad2d827d3eb939ee1869dbdcc4b8146f77be6bfb40274dee7fbd08f98"

    rawTxBytes, err := hex.DecodeString(rawTx)

    tx := new(types.Transaction)
    rlp.DecodeBytes(rawTxBytes, &tx)

    err = client.SendTransaction(context.Background(), tx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
