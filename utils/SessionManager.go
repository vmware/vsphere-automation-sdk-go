package utils

import (
	"net/http"

	client "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type SessionManager interface {
	CreateSession(cl http.Client)
	Login() error
	Logout() error
	Connector() client.Connector
	SessionID() string
}
