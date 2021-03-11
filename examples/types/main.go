package main

import (
	"context"
	"fmt"
	"os"

	linode "github.com/displague/linode-go/v1"
)

func main() {
	token := os.Getenv("LINODE_TOKEN")

	configuration := linode.NewConfiguration()
	api_client := linode.NewAPIClient(configuration)
	resp, r, err := api_client.LinodeTypesApi.GetLinodeTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeTypesApi.GetLinodeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeTypes`: InlineResponse20023
	fmt.Fprintf(os.Stdout, "Response from `LinodeTypesApi.GetLinodeTypes`: %v\n", resp)
}
