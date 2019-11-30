package model

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
)

type AuthInfo interface {
	GetName() string
	GetAuthInterface() interface{}
	GetSecurityContext() (core.SecurityContext, error)
}
