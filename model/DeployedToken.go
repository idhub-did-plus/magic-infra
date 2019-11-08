package model

type DeployedToken struct {
	Name            string   `json:"name"`
	Symbol          string   `json:"symbol"`
	Decimals        string   `json:"decimals"`
	ContractAddress string   `json:"contractAddress"`
	DeployAccount   string   `json:"deployAccount"`
	TokenType       string   `json:"tokenType"`
	Controllers     []string `json:"controllers"`
	RegistryAddress string   `json:"registryAddress"`
}
