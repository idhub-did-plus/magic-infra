package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

func Sha3(value string) [32]byte {
	var data []byte = []byte(value)
	array := common.BytesToHash(data).Bytes()
	var result [32]byte
	copy(result[:], array)
	return result
}
