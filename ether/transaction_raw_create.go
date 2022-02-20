package ether // page 87 

import (
    "context"
    "encoding/hex"
    "fmt"
    "log"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func TransactionRawCreate() {

    client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")

    value := big.NewInt(1000000000000000000) // in wei (1 eth)
    gasLimit := uint64(21000)                // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
		log.Fatal(err)
    }

    toAddress := common.HexToAddress("0x4Cd481B57B13C14988498BfF28974eAC919e3862")
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    ts := types.Transactions{signedTx}
    //rawTxBytes := ts.GetRlp(0)

	rawTxBytes, err := rlp.EncodeToBytes(ts[0])
		if err != nil {
			panic(err)
		}
    rawTxHex := hex.EncodeToString(rawTxBytes)

    fmt.Printf(rawTxHex) // f86...772”

	//.func (r Receipts) GetRlp(i int) []byte {
		
		//return bytes
	//}
	
}