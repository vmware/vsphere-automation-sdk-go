package refreshtoken

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
)

type Auth interface {
	scheme.Details
	CSPURL() string
	RefreshToken() string
}
