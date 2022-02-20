package ether

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	//"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EtherAccountBalances() {

    client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
    if err != nil {
        log.Fatal(err)
    }

    account := common.HexToAddress("0x2355B93c4551a235315e14c1e76AcDDe0Bfcc4C9")
    /* balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    } */
    //fmt.Println("Balance:" ,balance) // 25893180161173005034

   /*  blockNumber := big.NewInt(5532993)
    balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balanceAt)
	*/

    _balance := make(chan *big.Int, 1)
    go getBalance(*client, account, _balance)
    balance := <-_balance

    fbalance := new(big.Float)
    fbalance.SetString(balance.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
    fmt.Println("Eth Value:", ethValue)
	
	
    //pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
    //fmt.Println("Pending:",pendingBalance) // 25729324269165216042 */
}




func getBalance(client ethclient.Client, account common.Address, returnValue chan *big.Int) {

balance, err := client.BalanceAt(context.Background(), account, nil)
if err != nil {
result := map[string]interface{}{
    "error": "balance error 1",
}
fmt.Println(result)
}

fmt.Printf("4-Balance: %s\n", balance)

returnValue <- balance

}