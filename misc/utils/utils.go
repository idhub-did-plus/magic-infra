package utils

import (
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func Sha3(value string) [32]byte {

	hash := solsha3.SoliditySHA3(solsha3.Bytes32(value))

	var result [32]byte
	copy(result[:], hash)
	return result
}
