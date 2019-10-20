package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
)

type Auth interface {
	scheme.Details
	GetSAMLBearerToken() string
}
