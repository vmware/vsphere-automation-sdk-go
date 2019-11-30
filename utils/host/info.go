package host

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
	"path"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/host/kind"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/task"
)

type info struct {
	args.Properties
	sessionManager session.Manager
}

func (hd *info) Name() string {
	name, err := hd.GetStringPropertyValue(keys.HostNameKey)
	if err != nil {
		return ""
	}
	return name
}

func (hd *info) Server() string {
	server, err := hd.GetStringPropertyValue(keys.HostAddressKey)
	if err != nil {
		return ""
	}
	return server
}

func (hd *info) Kind() kind.Info {
	hostType, err := hd.GetStringPropertyValue(keys.HostTypeKey)
	if err != nil {
		return nil
	}
	return kind.GetHostTypeByName(hostType)
}

func (hd *info) AuthScheme() auth.Scheme {
	authDetails := hd.AuthDetails()
	if authDetails == nil {
		return auth.NoAuth
	}
	return authDetails.GetAuthScheme()
}

func (hd *info) AuthDetails() auth.Info {
	authDetails, err := auth.GetAuthSchemeInfo(hd.Properties)
	if err != nil {
		return nil
	}
	return authDetails
}

func (hd *info) getHTTPClient() http.Client {
	var skipServerVerification bool = false
	ssv := hd.GetPropertyValue(keys.SkipServerVerifiationKey)
	if ssv != nil {
		skipServerVerification = ssv.(bool)
	}
	if skipServerVerification {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
			InsecureSkipVerify: skipServerVerification}
	} else {
		serverCert := hd.GetPropertyValue(keys.CertFileKey)
		if serverCert == nil {
			log.Error("Invalid " + keys.CertFileKey)
			os.Exit(1)
		}
		serverCertKey := hd.GetPropertyValue(keys.CertKeyFileKey)
		if serverCertKey == nil {
			log.Error("Invalid " + keys.CertKeyFileKey)
			os.Exit(1)
		}
		cert, err := tls.LoadX509KeyPair(serverCert.(string), serverCertKey.(string))
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
			InsecureSkipVerify: skipServerVerification,
			Certificates:       []tls.Certificate{cert}}
	}
	httpClient := http.Client{}
	return httpClient
}

func (hd *info) getSessionManager() (session.Manager, error) {
	url, err := url.Parse(hd.Server())
	if err != nil {
		return nil, err
	}
	hostKind := hd.Kind()
	url.Path = path.Join(url.Path, hostKind.APIEndPointPrefix())
	sessionManager, err := session.NewSessionManager(url.String(), hostKind.Protocol(), hd.AuthDetails(), hd.getHTTPClient())
	if err != nil {
		return nil, err
	}
	return sessionManager, nil
}

func (hd *info) Connector() client.Connector {
	if hd.sessionManager == nil {
		return nil
	}
	return hd.sessionManager.Connector()
}

func (hd *info) ExecuteTask(t task.TaskFunc) error {
	sessionManager, err := hd.getSessionManager()
	if err != nil {
		return err
	}
	hd.sessionManager = sessionManager
	execTask := task.CreateNewTask(t)
	err = (*execTask).Execute(sessionManager)
	return err
}
