package ether

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"github.com/AdiF1/solidity/bookings/contracts"
	"github.com/AdiF1/solidity/bookings/internal/models"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/kylesliu/web3.go"
)

// retrieved campaign contract Address
var contractAddress = common.HexToAddress("0x820cE9F86dDE7367fFE6CE06D642f80FdB63a66f")
var pk1 = "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7"
var pk2 = "782336a0704da5f8feee75f66639a1adeb8203e57d14b176b44a032ecf3d38d1"
var recipAddress = common.HexToAddress("0xF9a20B1F7995B652313dC64D0FeEB451FD146522") // account 3
var contractBalance = big.NewInt(0)
//var web3 := 

func BaseCampaignDeploy() {
	//call SetUpParams() in smart.go
	client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk1)

	//log.Println("client: ", client.NetworkID())
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	initValue := big.NewInt(10)
	address, _, _, err := adf2.DeployBaseCampaign(auth, client, initValue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\ncontract address: ", address.Hex())

	//fmt.Println("\n\ntx.Hash: ", tx.Hash().Hex())

	fmt.Println("\n\nnonce: ", auth.Nonce)

	//_, _ = instance, tx

}

// CampaignCreator returns address on creator from instance
func BaseCampaignCreator() {
	//call SetUpParams() in smart.go
	client, _, _ := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk1)

	

	instance, err := adf2.NewBaseCampaign(contractAddress, client)

	    if err != nil {
        log.Fatal(err)
    }

	creator, err := instance.Creator(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creator: ", creator)

}

// CampaignMinimumContribution returns the amount set on creation
func BaseCampaignMinimumContribution() {
	//call SetUpParams() in smart.go
	client, _, _ := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk1)

	instance, err := adf2.NewBaseCampaign(contractAddress, client)
	    if err != nil {
        log.Fatal(err)
    }

	contribution, err := instance.MinimumContribution(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("contribution: ", contribution)

}

// CampaignMinimumContribution returns the amount set on creation
func BaseCampaignBalance() {
	//call SetUpParams() in smart.go
	client, _, _ := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk1)

    contractBalance, _ = client.BalanceAt(context.Background(), contractAddress, nil)

    ethBal := new(big.Float)
    ethBal.SetString(contractBalance.String())
    ethValue := new(big.Float).Quo(ethBal, big.NewFloat(math.Pow10(18)))

    fmt.Printf("Balance is %d Wei %f Eth\n",contractBalance,ethValue) 

}

// -------------------------------------------------------------
// BaseCampaignCreator returns address on creator from instance
func BaseCampaignContribute() {
	//call SetUpParams() in smart.go
	client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk2)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(101)

	instance, err := adf2.NewCampaign(contractAddress, client)
	    if err != nil {
        log.Fatal(err)
    }

	tx, err := instance.Contribute(opts)

	fmt.Println("tx value: ", tx.Value())

}

// CampaignCreator returns address on creator from instance
func BaseCampaignSetRequest() {
	//call SetUpParams() in smart.go
	client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk1)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(1010101)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)
	opts.GasLimit = uint64(3000000)
	opts.GasPrice = gasPrice

	instance, err := adf2.NewCampaign(contractAddress, client)
	    if err != nil {
        log.Fatal(err)
    }
	
	smallBalance := big.NewInt(100)
	tx, err := instance.CreateRequest(opts, "Vegan Reactor", smallBalance, recipAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx value: ", tx.Value())
}

// BaseCampaignGetRequest
func BaseCampaignGetRequest() {
	//call SetUpParams() in smart.go
	client, _, _ := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk1)

	instance, err := adf2.NewBaseCampaign(contractAddress, client)
	    if err != nil {
        log.Fatal(err)
    }
	request, err := instance.Requests(nil, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	valueAsInt := big.NewInt(request.Value.Int64())
	valueAsString := valueAsInt.String()		

	req := models.Request{
        Description: request.Description,
		Value: valueAsString,
		Complete: request.Complete,
		Address: request.Recipient.Hex(),
		ApprovalCount: request.ApprovalCount.Bit(0),
    }

	log.Println("request: ", req)

}

func BaseCampaignApproveRequest() {
	//call SetUpParams() in smart.go
	client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk2)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(101)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)
	opts.GasLimit = uint64(3000000)
	opts.GasPrice = gasPrice

	instance, err := adf2.NewBaseCampaign(contractAddress, client)
	    if err != nil {
        log.Fatal(err)
    }

	reqId := big.NewInt(0)
	tx, err := instance.ApproveRequest(opts, reqId)

	fmt.Println("tx value: ", tx.Value())

}


func BaseCampaignFinalizeRequest() {
	//call SetUpParams() in smart.go
	client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", pk1)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)
	opts.GasLimit = uint64(3000000)
	opts.GasPrice = gasPrice

	instance, err := adf2.NewBaseCampaign(contractAddress, client)
	    if err != nil {
        log.Fatal(err)
    }

	reqId := big.NewInt(0)
	tx, err := instance.FinalizeRequest(opts, reqId)

	fmt.Println("tx value: ", tx.Value())

}