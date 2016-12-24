package helpers

import (
        "fmt"
        "net/http"

	      "github.com/Azure/go-autorest/autorest"
        "github.com/Azure/go-autorest/autorest/azure"
)

// rest helpers

func WithInspection() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			fmt.Printf("Inspecting Request: %s %s\n", r.Method, r.URL)
			return p.Prepare(r)
		})
	}
}

func ByInspecting() autorest.RespondDecorator {
	return func(r autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(resp *http.Response) error {
			fmt.Printf("Inspecting Response: %s for %s %s\n", resp.Status, resp.Request.Method, resp.Request.URL)
			return r.Respond(resp)
		})
	}
}

// NewToken creates a token using the environmental vars passed during teamcity runs

func NewToken(creds map[string]string, scope string) (*azure.ServicePrincipalToken, error) {
  oauthConfig, err := azure.PublicCloud.OAuthConfigForTenant(creds["AZURE_TENANT_ID"])
  if err != nil {
          panic(err)
  }
  return azure.NewServicePrincipalToken(*oauthConfig, creds["AZURE_CLIENT_ID"], creds["AZURE_CLIENT_SECRET"], scope)
}

func ensureValueStrings(mapOfInterface map[string]interface{}) map[string]string {
	mapOfStrings := make(map[string]string)
	for key, value := range mapOfInterface {
		mapOfStrings[key] = ensureValueString(value)
	}
	return mapOfStrings
}

func ensureValueString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
