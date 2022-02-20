package ether

import (
	//"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"

	token "github.com/AdiF1/solidity/bookings/contracts"
	"github.com/AdiF1/solidity/bookings/internal/models"
)


func ContractReadERC20() models.TokenInfo {
    //call SetUpParams() in smart.go
	client, _, _ := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")


    // SCX007 Address
    tokenAddress := common.HexToAddress("0x1a72581c2ddc3bb87b03bd192c660ff04d6f8d74")
	
	instance, err := token.NewToken(tokenAddress, client)
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress("0x2355B93c4551a235315e14c1e76AcDDe0Bfcc4C9")
    bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
    if err != nil {
        log.Fatal(err)
    }

    name, err := instance.Name(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    symbol, err := instance.Symbol(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    decimals, err := instance.Decimals(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }
	//fmt.Printf("name: %s\n", name)         // "name: Golem Network"
    //fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
    //fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

    //fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

    fbal := new(big.Float)
    fbal.SetString(bal.String())
    value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

    //fmt.Printf("balance: %f\n\n", value) // "balance: 74605500.647409"

    ti := models.TokenInfo{
        Name: name,
        Symbol: symbol,
        Value: value,
    }

    return ti

}
