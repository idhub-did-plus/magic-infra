package service

import (
    "context"
    "log"
    "fmt"

    _"github.com/ethereum/go-ethereum"
    "github.com/joho/godotenv"
)

var myenv map[string]string

const envLoc = ".env"

func loadEnv() {
    var err error
    if myenv, err = godotenv.Read(envLoc); err != nil {
        log.Printf("could not load env from %s: %v", envLoc, err)
    }
}

func main(){
    loadEnv()

    ctx := context.Background()

    client, err := ethclient.Dial(os.Getenv("GATEWAY"))
    if err != nil {
        log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
    }
    defer client.Close()

    accountAddress := common.HexToAddress("<enter_ethereum_address>")
    balance, _ := client.BalanceAt(ctx, accountAddress, nil)
    fmt.Printf("Balance: %d\n",balance)
}