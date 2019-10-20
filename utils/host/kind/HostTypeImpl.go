package kind

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

type hostTypeImpl struct {
	name               string
	protocol           arch.Type
	supportedAuthTypes []scheme.Scheme
	apiEndPointPrefix  string
}

func (ht hostTypeImpl) isHostType() Kind {
	return ht
}

func (ht hostTypeImpl) String() string {
	return ht.name
}

func (ht hostTypeImpl) Protocol() arch.Type {
	return ht.protocol
}

func (ht hostTypeImpl) APIEndPointPrefix() string {
	return ht.apiEndPointPrefix
}

func (ht hostTypeImpl) SupportedAuthTypes() []scheme.Scheme {
	return ht.supportedAuthTypes
}
