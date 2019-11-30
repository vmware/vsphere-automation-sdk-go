package basic

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
)

type Info interface {
	UserName() string
	Password() string
	model.AuthInfo
}
