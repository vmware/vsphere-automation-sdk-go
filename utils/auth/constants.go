package auth

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/basic"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/oauth/refreshtoken"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/samlbearer"
)

const (
	NoAuth            = Scheme("NoAuth")
	BasicAuth         = Scheme(basic.Name)
	SAMLBearer        = Scheme(samlbearer.Name)
	OAuthRefreshToken = Scheme(refreshtoken.Name)
)

var All []Scheme = []Scheme{
	NoAuth,
	BasicAuth,
	SAMLBearer,
	OAuthRefreshToken,
}
