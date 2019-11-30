package session

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type Manager interface {
	Login() error
	Logout() error
	Connector() client.Connector
}
