package model

type Account struct {
	Name     string `json:"name"`
	Identity string `json:"identity"`
}
type ClaimRequest struct {
	Name     string `json:"name"`
	Identity string `json:"identity"`
}
