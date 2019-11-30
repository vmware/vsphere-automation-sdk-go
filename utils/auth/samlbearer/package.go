package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
)

func CreateSAMLBearerAuthFromProperties(authMap map[string]interface{}) (Info, error) {
	samlToken := authMap[keys.AuthSchemeSAMLBearerTokenKey]
	if samlToken == nil {
		return nil, &Error{}
	}

	return &info{Token: samlToken.(string)}, nil
}

func Create(samlToken string) (Info, error) {
	if len(samlToken) <= 0 {
		return nil, &Error{}
	}

	return &info{Token: samlToken}, nil
}
