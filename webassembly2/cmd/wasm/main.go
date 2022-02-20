package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"syscall/js"
	"time"

	//"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"

)
func main(){
    fmt.Println("Go WebAssembly2!")
    js.Global().Set("swap", swapEther())
    js.Global().Set("testBalance", testAccountBalances())
    js.Global().Set("accounts", etherAccountBalances())
    js.Global().Set("skinnyBal", skinnyBalance())
    js.Global().Set("skinnyCon", skinnyContribute())

    //js.Global().Set("contribute", ContractBalance())

	<-make(chan bool)
}

// BaseCampaignCreator returns address on creator from instance
func skinnyContribute() js.Func{
    skinnyContributeFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) != 1 {
            result := map[string]interface{}{
                "error": "Invalid no of arguments passed",
            }
            return result
        }
        jsDoc := js.Global().Get("document")
        if !jsDoc.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get document object",
            }
            return result
        }
        jsonOutputTextArea := jsDoc.Call("getElementById", "someoutput")
        if !jsonOutputTextArea.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get output text area 2",
            }
            return result
        }
        //inputJSON := args[0].String()

        var contractAddress = common.HexToAddress("0x995a5B8e3ccf74Ec37804B264482c41aC6522c3b")

        client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
        if err != nil {
            result := map[string]interface{}{
                "error": "client dial 0",
            }
            return result
        }

       

        go goSkinnyContribute(client, contractAddress)

        return nil

    })

    return skinnyContributeFunc
}

func goSkinnyContribute(client *ethclient.Client, contractAddress common.Address) {

     // the first step is to load your private key. (acc #1)
     privateKey, err := crypto.HexToECDSA("94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")
     if err != nil {
         log.Fatal(err)
     }

     publicKey := privateKey.Public()
     publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
     if !ok {
         log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
     }
     fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)


     nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
     //nonce = uint64(252)
     if err != nil {
         log.Fatal(err)
     }

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

    instance, err := NewCampaign(contractAddress, client)
	    if err != nil {
        log.Fatal(err)
    }

	tx, err := instance.Contribute(opts)

	fmt.Println("tx value: ", tx.Value())
}

// CampaignMinimumContribution returns the amount set on creation
func skinnyBalance() js.Func{
    skinnyBalanceFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) != 1 {
            result := map[string]interface{}{
                "error": "Invalid no of arguments passed",
            }
            return result
        }
        jsDoc := js.Global().Get("document")
        if !jsDoc.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get document object",
            }
            return result
        }
        jsonOutputTextArea := jsDoc.Call("getElementById", "someoutput")
        if !jsonOutputTextArea.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get output text area 2",
            }
            return result
        }
        //inputJSON := args[0].String()

        var contractAddress = common.HexToAddress("0x995a5B8e3ccf74Ec37804B264482c41aC6522c3b")

        client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
        if err != nil {
            result := map[string]interface{}{
                "error": "client dial 0",
            }
            return result
        }

        go goGetContractBalance(*client, contractAddress)

        return nil
    })
    return skinnyBalanceFunc
}

func testAccountBalances() js.Func { 
    fmt.Printf("\nin testAccountBalances\n")

    accTestFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

    

        if len(args) != 1 {
            result := map[string]interface{}{
                "error": "Invalid no of arguments passed",
            }
            return result
        }

        jsDoc := js.Global().Get("document")
        if !jsDoc.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get document object",
            }
            return result
        }

        jsonOutputTextArea := jsDoc.Call("getElementById", "someoutput")
        if !jsonOutputTextArea.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get output text area 1",
            }
            return result
        }
        inputJSON := args[0].String()
        fmt.Printf("testAccountBalances input: %s\n", inputJSON)

        jsonOutputTextArea.Set("value", inputJSON)


        return nil
    })
    return accTestFunc
  }

func etherAccountBalances() js.Func { 
    fmt.Printf("\nin etherAccountBalances\n")

    accBalanceFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) != 1 {
            result := map[string]interface{}{
                "error": "Invalid no of arguments passed",
            }
            return result
        }
        jsDoc := js.Global().Get("document")
        if !jsDoc.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get document object",
            }
            return result
        }
        jsonOutputTextArea := jsDoc.Call("getElementById", "someoutput")
        if !jsonOutputTextArea.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get output text area 2",
            }
            return result
        }
        inputJSON := args[0].String()
        
        client, err := ethclient.Dial(inputJSON)
        if err != nil {
            result := map[string]interface{}{
                "error": "client dial 0",
            }
            return result
        }

        account := common.HexToAddress("0x2355B93c4551a235315e14c1e76AcDDe0Bfcc4C9")

        go goGetAccountBalance(*client, account)

        return nil    
  })

  return accBalanceFunc
}

func swapEther() js.Func { 
    fmt.Printf("\nin swapEther\n")

    accSwapFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        fmt.Printf("in accSwapFunc")
          if len(args) != 1 {
              result := map[string]interface{}{
                  "error": "Invalid no of arguments passed",
              }
              return result
          }
          jsDoc := js.Global().Get("document")
          if !jsDoc.Truthy() {
              result := map[string]interface{}{
                  "error": "Unable to get document object",
              }
              return result
          }
          jsonOutputTextArea := jsDoc.Call("getElementById", "someoutput")
          if !jsonOutputTextArea.Truthy() {
              result := map[string]interface{}{
                  "error": "Unable to get output text area 2",
              }
              return result
          }
          inputAddress := args[0].String()
          
          client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
          if err != nil {
              result := map[string]interface{}{
                  "error": "client dial 0",
              }
              return result
          }
          fmt.Printf("in dialed\n")
  
          //toAddress := common.HexToAddress(inputAddress)
  
          go goTransferEther(*client, inputAddress)
  
          return nil    
    })
  
    return accSwapFunc
  }

func goGetAccountBalance(client ethclient.Client, account common.Address) {

    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        result := map[string]interface{}{
            "error": "balance error 1",
        }
        fmt.Println(result)
    }

    fbalance := new(big.Float)
    fbalance.SetString(balance.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
    fmt.Println("Eth Value:", ethValue) 
}

func goGetContractBalance(client ethclient.Client, contractAddress common.Address) {

    contractBalance, err := client.BalanceAt(context.Background(), contractAddress, nil)
    if err != nil {
        result := map[string]interface{}{
            "error": "balance error 1",
        }
        fmt.Println(result)
    }

    ethBal := new(big.Float)
    ethBal.SetString(contractBalance.String())
    ethValue := new(big.Float).Quo(ethBal, big.NewFloat(math.Pow10(18)))

    fmt.Printf("Balance is %d Wei %f Eth\n",contractBalance,ethValue) 
}

func goTransferEther(client ethclient.Client, _toAddress string) {
    //call SetUpParams() in smart.go  9485...is acc#1
    //client, privateKey, nonce := SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")
    // -------------------
    fmt.Printf("in TransferEther\n")

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	} else {
		//fmt.Println("header:" ,header.Number.String())
		_ = header
	}
    fmt.Printf("in TransferEther 2\n")
	// the first step is to load your private key. (acc #1)
	privateKey, err := crypto.HexToECDSA("94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("in TransferEther 3\n")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    fmt.Printf("in TransferEther 4\n")
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	//nonce = uint64(252)
	if err != nil {
		log.Fatal(err)
	}
    // ----------------------------------------------
    value := big.NewInt(100000000000000000) // in wei (0.1 eth)
    gasLimit := uint64(21000)                // in units
    fmt.Printf("in TransferEther 5\n")
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

	// figure out who we're sending the ETH to... (acc #3)
    toAddress := common.HexToAddress(_toAddress)
    fmt.Printf("in TransferEther 6\n")
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("in TransferEther 7\n")
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }
     go goTransact(client, signedTx)
}

func goTransact(client ethclient.Client, signedTx *types.Transaction) {
    fmt.Printf("in transact 1\n")
    err := client.SendTransaction(context.Background(), signedTx)
    time.Sleep(time.Second * 15)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

/* func SetUpParams (dial string, pky string) (cl *ethclient.Client, pk *ecdsa.PrivateKey, non uint64) {

	client, err := ethclient.Dial(dial)
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	} else {
		//fmt.Println("header:" ,header.Number.String())
		_ = header
	}

	// the first step is to load your private key. (acc #1)
	privateKey, err := crypto.HexToECDSA(pky)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	//nonce = uint64(252)
	if err != nil {
		log.Fatal(err)
	}

	return client, privateKey, nonce

} */

  // -------------------------------------

   //go getBalance(*client, account)

       // _balance := make(chan *big.Int, 1)
       //go getBalance(*client, account, _balance)
       //balance := <-_balance
       //result := make(chan int, 1)
	    //go sum(2, 3, result) 

	    //value := <-result
	    //fmt.Printf("Value: %d\n", value)
	    //close(result)
        // -------------------------------------

/* func getBalance(client ethclient.Client, account common.Address, returnValue chan *big.Int) {

    accBalanceFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) != 2 {
            result := map[string]interface{}{
                "error": "Invalid no of arguments passed",
            }
            return result
        }

    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
    result := map[string]interface{}{
        "error": "balance error 1",
    }
    fmt.Println(result)
    }
    
    fmt.Printf("4-Balance: %s\n", balance)
    
    returnValue <- balance
    return
    
} */




