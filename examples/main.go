package main

import (
	"fmt"
	"os"

	eyc "github.com/engineyard/terraform-eyc-sdk"
)

func main() {
	token := os.Getenv("eyc_token")
	fmt.Printf(token)
	c, _ := eyc.NewClient(nil, &token)
	fmt.Printf("%v\n", c)
	body, _ := c.GetEnvVars()
	fmt.Printf("Env Vars: %v\n", body)
}
