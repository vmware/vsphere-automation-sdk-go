package host

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/oauth/refreshtoken"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/host/kind"
)

func DefineDefault() Info {
	return Define("")
}

func Define(hostName string) Info {
	var host Info = &info{Properties: args.InitProperties()}
	if len(hostName) == 0 {
		hostName = "default"
	} else {
		host.DefineString(keys.HostNameKey, "Unique name of the host.", true, nil, nil)
	}
	host.AddOption(keys.HostNameKey, hostName)
	host.DefineString(keys.HostAddressKey, "Host server URL.", true, nil, nil)
	host.DefineString(keys.HostTypeKey, "Type of host i.e. VMware Product.", true, nil, kind.GetAllHostKind())
	host.DefineString(keys.AuthSchemeKey, "Authentication Protocol.", true, nil, auth.GetAllAuthScheme())
	host.DefineBool(keys.SkipServerVerifiationKey, "Skip Server Verification. Do not skip server verification for the production servers.")
	// Arguments definitions for Basic Authentication
	host.DefineString(keys.AuthSchemeBasicUsernameKey,
		"User Name for authentication. This Argument is mandatory only if "+keys.AuthSchemeKey+" has value "+string(auth.BasicAuth),
		true, areBasicAuthParamsRequired, nil)
	host.DefineString(keys.AuthSchemeBasicPasswordKey,
		"Password for authentication. This Argument is mandatory only if "+keys.AuthSchemeKey+" has value "+string(auth.BasicAuth),
		true, areBasicAuthParamsRequired, nil)
	// Arguments definitions for SAML Bearer Token Authentication
	host.DefineString(keys.AuthSchemeSAMLBearerTokenKey,
		"SAML Bearer Token for authentication. This Argument is mandatory only if "+keys.AuthSchemeKey+" has value "+string(auth.SAMLBearer),
		true, areSAMLBearerParamsRequired, nil)
	// Arguments definitions for OAuth with Refresh Token
	host.DefineStringWithDefaultValue(keys.AuthSchemeOAuthCSPURLKey, "CSP URL to get Access Token.", refreshtoken.CSPDefaultURL, areOAuthParamsRequired, nil)
	host.DefineString(keys.AuthSchemeOAuthRefreshTokenKey,
		"OAuth Refresh Token for authentication. This Argument is mandatory only if "+keys.AuthSchemeKey+" has value "+string(auth.OAuthRefreshToken),
		true, areOAuthParamsRequired, nil)
	if hostName == "default" {
		host.DefineString(keys.ConfigFileKey, "Full path to config file.", false, nil, nil)
	}
	return host
}

func New(hostProperties map[string]interface{}) (Info, error) {
	hostName, hostNamePresent := hostProperties[keys.HostNameKey]
	if !hostNamePresent || hostName == nil || len(hostName.(string)) == 0 {
		hostName = "default"
	}
	host := Define(hostName.(string))
	err := host.AddProperties(hostProperties)
	if err != nil {
		return nil, err
	}
	return host, nil
}

func areOAuthParamsRequired(properties args.Properties) bool {
	return !auth.IsOAuthRefreshToken(properties)
}

func areBasicAuthParamsRequired(properties args.Properties) bool {
	return !auth.IsBasicAuth(properties)
}

func areSAMLBearerParamsRequired(properties args.Properties) bool {
	return !auth.IsSAMLBearerAuth(properties)
}

// func getPropertiesMapByHostName(hostName string) (map[string]interface{}, error) {
// 	var configFilePath string
// 	for i := 1; i < len(os.Args)-1; i++ {
// 		if os.Args[i] == "-config-file" {
// 			configFilePath = os.Args[i+1]
// 			break
// 		}
// 	}

// 	if len(configFilePath) > 0 {
// 		configFile, err := os.Open(configFilePath)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer configFile.Close()
// 		byteValue, _ := ioutil.ReadAll(configFile)
// 		var configFileMap map[string]interface{}
// 		json.Unmarshal(byteValue, &configFileMap)
// 		hosts, hostsValue := configFileMap["Hosts"]
// 		if hostsValue {
// 			//hosts := configFileMap["Hosts"].([]map[string]interface{})
// 		}
// 		for key, value := range configFileMap {
// 			if len(hostName) > 0 && key != hostName {
// 				continue
// 			}
// 			prop.setValue(key, value)
// 			if key == hostName {
// 				break
// 			}
// 		}
// 	}
// }
