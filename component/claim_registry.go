// claim_registry
package component

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	//solsha3 "github.com/miguelmota/go-solidity-sha3"
)

const abiraw = `[
	{
		"constant": true,
		"inputs": [
		  {
			"name": "issuer",
			"type": "address"
		  },
		  {
			"name": "subject",
			"type": "address"
		  },
		  {
			"name": "key",
			"type": "bytes32"
		  }
		],
		"name": "getClaim",
		"outputs": [
		  {
			"name": "",
			"type": "bytes32"
		  }
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]`

func main() {
	abi, err := abi.JSON(strings.NewReader(abiraw))
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress("0xca35b7d915458ef540ade6068dfe2f44e8fa733c")
	out, err := abi.Pack("addrTy", address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", out[4:])
}
