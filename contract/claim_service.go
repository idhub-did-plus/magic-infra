package contract

import (
	"log"
	"magic-infra/config"
	"magic-infra/misc/utils"

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

	caller *ClaimCaller
}

var ClaimService ClaimServiceType

func init() {
	ClaimService = ClaimServiceType{}
	ClaimService.contractAddress = common.HexToAddress(config.Config.Contract.ContractAddress)

	ClaimService.issuer = common.HexToAddress(config.Config.Contract.Issuer)
	ClaimService.identity = common.HexToAddress(config.Config.Contract.Identity)
	client, err := ethclient.Dial(config.Config.Contract.Provider)
	if err != nil {
		log.Fatal(err)
	}
	ClaimService.client = client
	ClaimService.caller, _ = NewClaimCaller(ClaimService.contractAddress, ClaimService.client)

}

func (cs *ClaimServiceType) GetClaim(key string) ([32]byte, error) {

	var claimType = utils.Sha3(key)

	claim, err := cs.caller.GetClaim(&bind.CallOpts{}, cs.issuer, cs.identity, claimType)

	return claim, err

}
