package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
)

type samlBearerDetails struct {
	Token string
}

func CreateSAMLBearerAuthDetails(authMap map[string]interface{}) (Auth, error) {
	samlToken := authMap[keys.AuthSchemeSAMLBearerTokenKey]
	if samlToken == nil {
		return nil, &Error{}
	}

	return Auth(samlBearerDetails{Token: samlToken.(string)}), nil
}
