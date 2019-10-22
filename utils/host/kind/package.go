package kind

import (
	"strings"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

var VCENTER Kind = hostTypeImpl{name: "vCenter", protocol: arch.JSONRPC, apiEndPointPrefix: "/api", supportedAuthTypes: []scheme.Scheme{scheme.BasicAuth, scheme.SAMLBearer}}
var VMC Kind = hostTypeImpl{name: "VMC", protocol: arch.REST, apiEndPointPrefix: "", supportedAuthTypes: []scheme.Scheme{scheme.OAuthRefreshToken}}
var NSXT Kind = hostTypeImpl{name: "NSXT", protocol: arch.REST, apiEndPointPrefix: "", supportedAuthTypes: []scheme.Scheme{scheme.OAuthRefreshToken}}

var All = []Kind{
	VCENTER,
	VMC,
	NSXT,
}

func GetHostTypeByName(name string) Kind {
	for _, host := range All {
		if strings.ToLower(host.String()) == strings.ToLower(name) {
			return host
		}
	}
	return nil
}

func GetAllHostKind() []string {
	var allHostKind []string
	for _, hostKind := range All {
		allHostKind = append(allHostKind, hostKind.String())
	}
	return allHostKind
}
