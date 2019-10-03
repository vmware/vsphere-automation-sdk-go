package utils

import (
	"log"
	"net/http"
	"strings"

	session "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/session"
	client "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

func (sm sessionManagerImpl) Connector() client.Connector {
	return sm.connector
}

func (sm sessionManagerImpl) SessionID() string {
	return sm.sessionID
}

func (sm *sessionManagerImpl) CreateSession(cl http.Client) {
	hostAPIAddr := Arguments.Get("server").(string)
	if strings.HasSuffix(hostAPIAddr, "/") {
		hostAPIAddr = hostAPIAddr + "api"
	} else {
		hostAPIAddr = hostAPIAddr + "/api"
	}
	sm.connector = client.NewJsonRpcConnector(hostAPIAddr, httpClient)
	username := Arguments.Get("username").(string)
	password := Arguments.Get("passwrod").(string)
	sm.connector.SetSecurityContext(security.NewUserPasswordSecurityContext(username, password))
	sm.sessionClient = session.NewSessionClientImpl(sm.connector)
}

func (sm *sessionManagerImpl) Login() error {
	sm.sessionID, sm.sessionError = sm.sessionClient.Create()
	if sm.sessionError != nil {
		log.Fatal(sm.sessionError)
	} else {
		log.Printf("Successfully created new session with Session ID: %s", sm.sessionID)
		sm.connector.SetSecurityContext(security.NewSessionSecurityContext(sm.sessionID))
	}
	return sm.sessionError
}

func (sm *sessionManagerImpl) Logout() error {
	err := sm.sessionClient.Delete()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Session %s deleted successfully.", sm.sessionID)
	}
	return err
}
