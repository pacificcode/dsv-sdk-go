package main

import (
	"os"

	"github.com/DelineaXPM/dsv-sdk-go/v2/auth"
	"github.com/DelineaXPM/dsv-sdk-go/v2/vault"
	"github.com/rs/zerolog/log"
)

func main() {
	var exitCode = 1
	dsv, err := vault.New(vault.Configuration{
		Credentials: vault.ClientCredential{
			ClientID: os.Getenv("AZURE_CLIENT_ID"), // CLIENT_ID of the MSI identity you wish to use
		},
		Tenant:   os.Getenv("DSV_TENANT"), // your tenant name
		TLD:      os.Getenv("DSV_TLD"),    // defaults to com change if your domain is au, eu etc
		Provider: auth.AZURE,              // required to enable Azure authentication
	})
	if err != nil {
		log.Printf("failed to configure vault: %v", err)
		os.Exit(exitCode)
	}

	secret, err := dsv.Secret("<secret path or ID")
	if err != nil {
		log.Printf("failed to fetch secret: %v", err)
		os.Exit(exitCode)
	}
	log.Printf("\nsecret data: %v\n\n", secret.Data)
}
