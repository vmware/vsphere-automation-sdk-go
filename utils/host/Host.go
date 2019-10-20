package host

import (
	"net/http"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/host/kind"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/task"
)

type Host interface {
	Name() string
	Server() string
	Kind() kind.Kind
	AuthScheme() scheme.Scheme
	AuthDetails() scheme.Details
	ExecuteTask(t task.TaskFunc) error
	getHTTPClient() http.Client
	getSessionManager() (session.SessionManager, error)
	args.Properties
}
