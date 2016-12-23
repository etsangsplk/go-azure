package helpers

import (
        "os"
        "github.com/Azure/go-autorest/autorest/azure"
)

// NewToken creates a token using the environmental vars passed during teamcity runs

func NewToken(creds map[string]string, scope string) (*azure.ServicePrincipalToken, error) {
  oauthConfig, err := azure.PublicCloud.OAuthConfigForTenant(creds("AZURE_TENANT_ID"))
  if err != nil {
          panic(err)
  }
  return azure.NewServicePrincipalToken(*oauthConfig, creds("AZURE_CLIENT_ID"), creds("AZURE_CLIENT_SECRET"), scope)
}
