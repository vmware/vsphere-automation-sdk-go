package session

import (
	"log"
	"net/http"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

type sessionManagerImpl struct {
	sessionID    string
	client       *cis.DefaultSessionClient
	connector    client.Connector
	sessionError error
	ArchType     arch.Type
	AuthDetails  scheme.Details
	URL          string
	HTTPClient   http.Client
}

func (sm sessionManagerImpl) Connector() client.Connector {
	return sm.connector
}

func (sm sessionManagerImpl) SessionID() string {
	return sm.sessionID
}

func (sm *sessionManagerImpl) Login() error {
	sm.sessionError = nil
	secCtx, err := sm.AuthDetails.GetSecurityContext()
	sm.sessionError = err
	if sm.sessionError != nil {
		return sm.sessionError
	}
	sm.connector.SetSecurityContext(secCtx)

	if sm.ArchType == arch.JSONRPC {
		sm.client = cis.NewDefaultSessionClient(sm.connector)
		sm.sessionID, sm.sessionError = sm.client.Create()
		if sm.sessionError == nil && sm.AuthDetails.GetAuthScheme() == scheme.BasicAuth {
			sm.connector.SetSecurityContext(security.NewSessionSecurityContext(sm.sessionID))
		}
	}
	if sm.sessionError != nil {
		log.Fatal(sm.sessionError)
	} else {
		log.Printf("Successfully created new session with Session ID: %s", sm.sessionID)
	}
	return sm.sessionError
}

func (sm *sessionManagerImpl) Logout() error {
	sm.sessionError = nil
	if sm.ArchType == arch.JSONRPC {
		sm.sessionError = sm.client.Delete()
	}
	if sm.sessionError != nil {
		log.Fatal(sm.sessionError)
	} else {
		sm.client = nil
		sm.connector = nil
		sm.sessionID = ""
		log.Printf("Session %s deleted successfully.", sm.sessionID)
	}
	return sm.sessionError
}
