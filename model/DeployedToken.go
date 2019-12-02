package model

type DeployedToken struct {
	Name            string
	Symbol          string
	Decimals        string
	ContractAddress string
	DeployAccount   string
	TokenType       string
	Controllers     []string
	RegistryAddress string
}
