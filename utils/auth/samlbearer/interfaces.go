package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
)

type Info interface {
	GetSAMLBearerToken() string
	model.AuthInfo
}
