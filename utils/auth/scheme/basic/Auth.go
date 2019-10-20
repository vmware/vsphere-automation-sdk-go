package basic

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
)

type Auth interface {
	scheme.Details
	GetUserName() string
	GetPassword() string
}
