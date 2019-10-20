package kind

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

type Kind interface {
	isHostType() Kind
	String() string
	Protocol() arch.Type
	SupportedAuthTypes() []scheme.Scheme
	APIEndPointPrefix() string
}
