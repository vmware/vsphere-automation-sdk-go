package auth

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
)

type Info interface {
	GetAuthScheme() Scheme
	GetAuthInterface() interface{}
	GetSecurityContext() (core.SecurityContext, error)
}
