package session

import (
	"net/http"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

func NewSessionManager(url string, archType arch.Type, authDetails scheme.Details, httpClient http.Client) (SessionManager, error) {
	if len(url) == 0 {
		return nil, &URLError{}
	}
	if archType == nil {
		return nil, &ArchTypeError{archType: "nil"}
	}
	if authDetails == nil {
		return nil, &AuthDetailsError{}
	}

	var sessionManager sessionManagerImpl = sessionManagerImpl{URL: url, ArchType: archType, AuthDetails: authDetails, HTTPClient: httpClient}
	if sessionManager.ArchType == arch.JSONRPC {
		sessionManager.connector = client.NewJsonRpcConnector(sessionManager.URL, sessionManager.HTTPClient)
	} else if sessionManager.ArchType == arch.REST {
		sessionManager.connector = client.NewRestConnector(sessionManager.URL, sessionManager.HTTPClient)
	} else {
		return nil, &ArchTypeError{archType: sessionManager.ArchType.String()}
	}
	return &sessionManager, nil
}
