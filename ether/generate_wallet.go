package ether

import (
    "crypto/ecdsa"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
    "golang.org/x/crypto/sha3"
)

func GenerateWallet() {
    // import the go-ethereum crypto package that provides the GenerateKey method for 
	// generating a random private key.
	privateKey, err := crypto.GenerateKey()
    if err != nil {
        log.Fatal(err)
    }
	// convert it to bytes by importing the golang crypto/ecdsa package and 
	// using the FromECDSA method
    privateKeyBytes := crypto.FromECDSA(privateKey)

    // convert it to a hexadecimal string by using the go-ethereum hexutil package 
	// which provides the Encode method which takes a byte slice. 
	// Then we strip off the 0x after it's hex encoded.
	fmt.Println("Private key:", hexutil.Encode(privateKeyBytes)[2:])

	// Since the public key is derived from the private key, go-ethereum's 
	// crypto private key has a Public method that will return the public key
	publicKey := privateKey.Public()

	// “Converting it to hex is a similar process that we went through with the private key. 
	// We strip off the 0x and the first 2 characters 04 which is always the EC prefix 
	// and is not required.”
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
    fmt.Println("Public key bytes:" ,hexutil.Encode(publicKeyBytes)[4:])

    // Now that we have the public key we can easily generate the public address 
	// which is what you're used to seeing. In order to do that, the go-ethereum 
	// crypto package has a PubkeyToAddress method which accepts an ECDSA public key, 
	// and returns the public address.
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
    fmt.Println("Public key address:" ,address)

	// The public address is simply the Keccak-256 hash of the public key, and then we take 
	// the last 40 characters (20 bytes) and prefix it with 0x. Here's how you can do it 
	// manually using the crypto/sha3 keccak256 function.
    hash := sha3.NewLegacyKeccak256()
    hash.Write(publicKeyBytes[1:])
    fmt.Println("Public key hash:" ,hexutil.Encode(hash.Sum(nil)[12:]))
}