package kind

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

type Info interface {
	isHostType() Info
	String() string
	Protocol() arch.Type
	SupportedAuthTypes() []auth.Scheme
	APIEndPointPrefix() string
}
