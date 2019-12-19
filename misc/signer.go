package misc

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	signer "github.com/ethereum/go-ethereum/signer/core"
)

func Challenge() (*signer.TypedData, error) {
	// Replace this with the address of the user's wallet
	walletAddress := "0x61e0499cF10d341A5E45FA9c211aF3Ba9A2b50ef"
	salt := "some-random-string-or-hash-here"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Generate a random nonce to include in our challenge
	nonceBytes := make([]byte, 32)
	n, err := rand.Read(nonceBytes)
	if n != 32 {
		return nil, errors.New("nonce: n != 64 (bytes)")
	} else if err != nil {
		return nil, err
	}
	nonce := hex.EncodeToString(nonceBytes)

	signerData := signer.TypedData{
		Types: signer.Types{
			"Challenge": []signer.Type{
				{Name: "address", Type: "address"},
				{Name: "nonce", Type: "string"},
				{Name: "timestamp", Type: "string"},
			},
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "version", Type: "string"},
				{Name: "salt", Type: "string"},
			},
		},
		PrimaryType: "Challenge",
		Domain: signer.TypedDataDomain{
			Name:    "ETHChallenger",
			Version: "1",
			Salt:    salt,
			ChainId: math.NewHexOrDecimal256(1),
		},
		Message: signer.TypedDataMessage{
			"timestamp": timestamp,
			"address":   walletAddress,
			"nonce":     nonce,
		},
	}
	return &signerData, nil
}
func Verify(storedChallengeHash []byte, userAddress common.Address, incomingMetamaskSignature string) bool {

	// Decode the hex-encoded signature from metamask.
	signature, _ := hex.DecodeString(incomingMetamaskSignature)

	if len(signature) != 65 {
		return false
	}

	if signature[64] != 27 && signature[64] != 28 {
		return false
	}
	signature[64] -= 27

	pubKeyRaw, err := crypto.Ecrecover(storedChallengeHash, signature)
	if err != nil {
		return false
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if !bytes.Equal(userAddress.Bytes(), recoveredAddr.Bytes()) {
		return false
	}
	return true
}
