package service

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "magic-infra/component" // for demo
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := store.NewComponentCaller(address, client)
	issuer := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	identity := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	var data []byte = []byte("valid investor")
	array := common.BytesToHash(data).Bytes()
	var array32 [32]byte
	copy(array32[:], array)

	instance.GetClaim(&bind.CallOpts{}, issuer, identity, array32)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
