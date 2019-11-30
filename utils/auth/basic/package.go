package basic

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
)

func CreateBasicAuthFromPropMap(authMap map[string]interface{}) (Info, error) {
	user := authMap[keys.AuthSchemeBasicUsernameKey]
	password := authMap[keys.AuthSchemeBasicPasswordKey]
	if user == nil && password == nil {
		return nil, &Error{}
	} else if user == nil {
		return nil, &UsernameError{}
	} else if password == nil {
		return nil, &PasswordError{}
	}

	return &info{userName: user.(string), password: password.(string)}, nil
}

func Create(user string, password string) (Info, error) {
	if len(user) <= 0 && len(password) <= 0 {
		return nil, &Error{}
	} else if len(user) <= 0 {
		return nil, &UsernameError{}
	} else if len(password) <= 0 {
		return nil, &PasswordError{}
	}

	return &info{userName: user, password: password}, nil
}
