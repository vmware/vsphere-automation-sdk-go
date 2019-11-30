package refreshtoken

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
)

type Info interface {
	CSPURL() string
	RefreshToken() string
	model.AuthInfo
}
