package helpers

import (
        "net/http"
        "os"
        "github.com/Azure/go-autorest/autorest/azure"
)

// rest helpers

func withInspection() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			fmt.Printf("Inspecting Request: %s %s\n", r.Method, r.URL)
			return p.Prepare(r)
		})
	}
}

func byInspecting() autorest.RespondDecorator {
	return func(r autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(resp *http.Response) error {
			fmt.Printf("Inspecting Response: %s for %s %s\n", resp.Status, resp.Request.Method, resp.Request.URL)
			return r.Respond(resp)
		})
	}
}

// NewToken creates a token using the environmental vars passed during teamcity runs

func NewToken(creds map[string]string, scope string) (*azure.ServicePrincipalToken, error) {
  oauthConfig, err := azure.PublicCloud.OAuthConfigForTenant(creds("AZURE_TENANT_ID"))
  if err != nil {
          panic(err)
  }
  return azure.NewServicePrincipalToken(*oauthConfig, creds("AZURE_CLIENT_ID"), creds("AZURE_CLIENT_SECRET"), scope)
}
