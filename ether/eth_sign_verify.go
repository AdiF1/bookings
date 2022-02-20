package ether

import (
	"fmt"
	"bytes"
	"log"
	"crypto/ecdsa"	
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"encoding/hex"
)

// 0x2355B93c4551a235315e14c1e76AcDDe0Bfcc4C9
// 94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7

// from book... 	signature_verify.go 	Page 234
func CheckIt() {

	privateKey, err := crypto.HexToECDSA("94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")
    if err != nil {
        log.Fatal(err)
    }

	//fmt.Println("FromECDSA: ", crypto.(privateKey))

    publicKey := privateKey.Public()

    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	fmt.Println("PubkeyToAddress: ", crypto.PubkeyToAddress(*publicKeyECDSA))

	if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

    data := []byte("hello")
    hash := crypto.Keccak256Hash(data)

    signature, err := crypto.Sign(hash.Bytes(), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Println("hexutil.Encode: " ,hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

    sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
    if err != nil {
        log.Fatal(err)
    }
	//fmt.Println("sigPublicKey: " ,sigPublicKey)

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
    fmt.Println("matches1: " ,matches) // true

    sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
    if err != nil {
        log.Fatal(err)
    }

    sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
    matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
    fmt.Println("matches2: " , matches) // true

    signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id



	dst := make([]byte, hex.EncodedLen(len(signatureNoRecoverID)))
	hex.Encode(dst, signatureNoRecoverID)

	fmt.Printf("%s\n", dst)


	//fmt.Println("signatureNoRecoverID: " , signatureNoRecoverID)
    verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
    fmt.Println("verified: " , verified) // true

}


// -------------------------------------------

// from eth_sign_verify.go 	https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go

func SignatureVerify() {
	fmt.Println(verifySignature(
		"0x829814B6E4dfeC4b703F2c6fDba28F1724094D11",
		"0x53edb561b0c1719e46e1e6bbbd3d82ff798762a66d0282a9adf47a114e32cbc600c248c247ee1f0fb3a6136a05f0b776db4ac82180442d3a80f3d67dde8290811c",
		[]byte("hello"),
	))
}

func verifySignature(from, signatureHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)

	signature := hexutil.MustDecode(signatureHex)
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if signature[64] != 27 && signature[64] != 28 {
		return false
	}
	signature[64] -= 27

	pubKey, err := crypto.SigToPub(signHash(msg), signature)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr
}

// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L404
// signHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}