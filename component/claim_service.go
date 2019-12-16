package component

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ClaimServiceType struct {
	contractAddress common.Address
	identity        common.Address
	issuer          common.Address
	url             string
	client          *ethclient.Client

	caller *ComponentCaller
}

var ClaimService ClaimServiceType

func init() {
	ClaimService = ClaimServiceType{}
	ClaimService.contractAddress = common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")

	ClaimService.issuer = common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	ClaimService.identity = common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	ClaimService.client = client
	ClaimService.caller, _ = NewComponentCaller(ClaimService.contractAddress, ClaimService.client)

}

func (cs *ClaimServiceType) GetClaim(key string) ([32]byte, error) {

	var data []byte = []byte(key)
	array := common.BytesToHash(data).Bytes()
	var claimType [32]byte
	copy(claimType[:], array)

	claim, err := cs.caller.GetClaim(&bind.CallOpts{}, cs.issuer, cs.identity, claimType)

	return claim, err

}
