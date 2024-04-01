package alphavantage

import (
	"fmt"
	"log"
	"os"
)

var (
	apiKey string
)

const (
	baseUrl    = "https://www.alphavantage.co/query"
	quoteStr   = "GLOBAL_QUOTE"
	secretPath = "secret.txt"
)

func init() {
	buf, err := os.ReadFile(secretPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	apiKey = string(buf)
	fmt.Println(apiKey)
}
