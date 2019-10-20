package scheme

import (
	"fmt"
	"strings"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args"
)

func GetAuthSchemeByName(name string) Scheme {
	for _, authScheme := range All {
		if strings.ToLower(string(authScheme)) == strings.ToLower(name) {
			return authScheme
		}
	}
	return NoAuth
}

func GetAuthSchemeFromProperties(properties args.Properties) (Scheme, error) {
	if properties == nil {
		return NoAuth, fmt.Errorf("No Authentication Scheme set")
	}
	schemeName, err := properties.GetStringPropertyValue(keys.AuthSchemeKey)
	if err != nil {
		return NoAuth, err
	}
	return GetAuthSchemeByName(schemeName), nil
}

func GetAllAuthScheme() []string {
	var allAuthSchemes []string
	for _, authScheme := range All {
		allAuthSchemes = append(allAuthSchemes, string(authScheme))
	}
	return allAuthSchemes
}
