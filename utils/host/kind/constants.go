package kind

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

var VCENTER Info = info{name: "vCenter", protocol: arch.JSONRPC, apiEndPointPrefix: "/api", supportedAuthTypes: []auth.Scheme{auth.BasicAuth, auth.SAMLBearer}}
var VMC Info = info{name: "VMC", protocol: arch.REST, apiEndPointPrefix: "", supportedAuthTypes: []auth.Scheme{auth.OAuthRefreshToken}}
var NSXT Info = info{name: "NSXT", protocol: arch.REST, apiEndPointPrefix: "", supportedAuthTypes: []auth.Scheme{auth.OAuthRefreshToken}}

var All = []Info{
	VCENTER,
	VMC,
	NSXT,
}
