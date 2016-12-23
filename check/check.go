package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/glennmate/go-azure/helpers"
  "github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
)

func main() {
	rgname := os.Getenv("AZURE_RESOURCE_GROUP_NAME")

	creds := map[string]string{
		"AZURE_CLIENT_ID":       os.Getenv("AZURE_CLIENT_ID"),
		"AZURE_CLIENT_SECRET":   os.Getenv("AZURE_CLIENT_SECRET"),
		"AZURE_SUBSCRIPTION_ID": os.Getenv("AZURE_SUBSCRIPTION_ID"),
		"AZURE_TENANT_ID":       os.Getenv("AZURE_TENANT_ID")}
	if err := checkEnvVar(&creds); err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	serviceprincipaltoken, err := helpers.NewToken(creds, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	accountcreds := resources.NewGroupsClient(creds["AZURE_SUBSCRIPTION_ID"])
	accountcreds.Authorizer = serviceprincipaltoken

	accountcreds.Sender = autorest.CreateSender(
		autorest.WithLogging(log.New(os.Stdout, "example: ", log.LstdFlags)))

	accountcreds.RequestInspector = helpers.withInspection()
	accountcreds.ResponseInspector = helpers.byInspecting()
	crg, err := accountcreds.CheckExistence(rgname)

	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}


func checkEnvVar(envVars *map[string]string) error {
	var missingVars []string
	for varName, value := range *envVars {
		if value == "" {
			missingVars = append(missingVars, varName)
		}
	}
	if len(missingVars) > 0 {
		return fmt.Errorf("Missing environment variables %v", missingVars)
	}
	return nil
}
