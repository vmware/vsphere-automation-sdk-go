package basic

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
)

type basicAuthDetails struct {
	UserName string
	Password string
}

func CreateBasicAuthDetails(authMap map[string]interface{}) (Auth, error) {
	user := authMap[keys.AuthSchemeBasicUsernameKey]
	password := authMap[keys.AuthSchemeBasicPasswordKey]
	if user == nil && password == nil {
		return nil, &Error{}
	} else if user == nil {
		return nil, &UsernameError{}
	} else if password == nil {
		return nil, &PasswordError{}
	}

	return Auth(basicAuthDetails{UserName: user.(string), Password: password.(string)}), nil
}
