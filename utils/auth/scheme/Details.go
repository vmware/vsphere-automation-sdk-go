package scheme

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
)

type Details interface {
	GetAuthScheme() Scheme
	GetAuthInterface() interface{}
	GetSecurityContext() (core.SecurityContext, error)
}
