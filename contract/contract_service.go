package contract

import (
	"log"

	"github.com/idhub-did-plus/magic-infra/config"
	"github.com/idhub-did-plus/magic-infra/misc/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractServiceType struct {
	address780  common.Address
	address1484 common.Address
	identity    common.Address
	issuer      common.Address
	url         string
	client      *ethclient.Client

	claimCaller    *ClaimCaller
	identityCaller *IdentityCaller
}

var ContractService ContractServiceType

func init() {
	ContractService = ContractServiceType{}
	ContractService.address780 = common.HexToAddress(config.Config.Contract.Address780)
	ContractService.address1484 = common.HexToAddress(config.Config.Contract.Address1484)
	ContractService.issuer = common.HexToAddress(config.Config.Contract.Issuer)
	ContractService.identity = common.HexToAddress(config.Config.Contract.Identity)
	client, err := ethclient.Dial(config.Config.Contract.Provider)
	if err != nil {
		log.Fatal(err)
	}
	ContractService.client = client
	ContractService.claimCaller, _ = NewClaimCaller(ContractService.address780, ContractService.client)
	ContractService.identityCaller, _ = NewIdentityCaller(ContractService.address1484, ContractService.client)

}

func (cs *ContractServiceType) HasClaim(identity string, key string, value string) bool {

	return true

}

func (cs *ContractServiceType) GetClaim(identity string, key string) ([32]byte, error) {
	id := common.HexToAddress(identity)
	var claimType = utils.Sha3(key)

	claim, err := cs.claimCaller.GetClaim(&bind.CallOpts{}, cs.issuer, id, claimType)

	return claim, err

}
