package main

import (
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"
	"./client"
	"./client/operations"
)

func main() {
	petClient := client.NewHTTPClient(strfmt.Default)
	resp, err := petClient.Operations.All(operations.AllParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
