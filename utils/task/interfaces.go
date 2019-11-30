package task

import "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"

type Task interface {
	Execute(sessionManager session.Manager) error
}
