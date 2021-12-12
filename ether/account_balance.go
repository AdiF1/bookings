package ether

import (
	"context"
	"fmt"
	"log"
	"math"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func etherAccountBalances() {
    client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
    if err != nil {
        log.Fatal(err)
    }

    account := common.HexToAddress("0x2355B93c4551a235315e14c1e76AcDDe0Bfcc4C9")
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println("Balance:" ,balance) // 25893180161173005034

   /*  blockNumber := big.NewInt(5532993)
    balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balanceAt)
	*/

    fbalance := new(big.Float)
    fbalance.SetString(balance.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
    fmt.Println("Eth Value:", ethValue)
	
	
    //pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
    //fmt.Println("Pending:",pendingBalance) // 25729324269165216042 */
}