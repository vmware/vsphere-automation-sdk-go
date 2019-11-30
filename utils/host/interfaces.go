package host

import (
	"net/http"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/host/kind"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/task"
)

type Info interface {
	Name() string
	Server() string
	Kind() kind.Info
	AuthScheme() auth.Scheme
	AuthDetails() auth.Info
	ExecuteTask(t task.TaskFunc) error
	Connector() client.Connector
	getHTTPClient() http.Client
	getSessionManager() (session.Manager, error)
	args.Properties
}
