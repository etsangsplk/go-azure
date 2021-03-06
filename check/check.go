package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/glennmate/go-azure/helpers"
  "github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

func main() {
	// var resourceGroupName = os.Getenv("AZURE_RESOURCE_GROUP_NAME")

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
		autorest.WithLogging(log.New(os.Stdout, "Logging send request:\n", log.LstdFlags)))

	accountcreds.RequestInspector = helpers.WithInspection()
	accountcreds.ResponseInspector = helpers.ByInspecting()
	crg, err := accountcreds.CheckExistence(os.Getenv("AZURE_RESOURCE_GROUP_NAME"))

	if err != nil {
		log.Fatalf("Error: %v\n", err)
		return
	}


responsemessage := fmt.Sprintln("%s", crg.Response)
responsecode, _	:= regexp.MatchString("2[0-9]{2}|299", responsemessage)

 if responsecode == true {
	 fmt.Printf("Resource group exists\n")
 } else {  log.Fatalf("Bad HTTP response code, result is '%s'\n", crg.Response)
	 }

}


func checkEnvVar(envVars *map[string]string) error {
	var missingVars []string
	for varName, value := range *envVars {
		if value == "" {
			missingVars = append(missingVars, varName)
		}
	}
	if len(missingVars) > 0 {
		return fmt.Errorf("Missing environment variables %v\n", missingVars)
	}
	return nil
}
