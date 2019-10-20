package host

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
	"path"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/host/kind"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/task"
)

type hostDetails struct {
	args.Properties
}

func (hd *hostDetails) Name() string {
	name, err := hd.GetStringPropertyValue(keys.HostNameKey)
	if err != nil {
		return ""
	}
	return name
}

func (hd *hostDetails) Server() string {
	server, err := hd.GetStringPropertyValue(keys.HostAddressKey)
	if err != nil {
		return ""
	}
	return server
}

func (hd *hostDetails) Kind() kind.Kind {
	hostType, err := hd.GetStringPropertyValue(keys.HostTypeKey)
	if err != nil {
		return nil
	}
	return kind.GetHostTypeByName(hostType)
}

func (hd *hostDetails) AuthScheme() scheme.Scheme {
	authDetails := hd.AuthDetails()
	if authDetails == nil {
		return scheme.NoAuth
	}
	return authDetails.GetAuthScheme()
}

func (hd *hostDetails) AuthDetails() scheme.Details {
	authDetails, err := auth.GetAuthSchemeDetails(hd.Properties)
	if err != nil {
		return nil
	}
	return authDetails
}

func (hd *hostDetails) getHTTPClient() http.Client {
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

func (hd *hostDetails) getSessionManager() (session.SessionManager, error) {
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

func (hd *hostDetails) ExecuteTask(t task.TaskFunc) error {
	sessionManager, err := hd.getSessionManager()
	if err != nil {
		return err
	}
	execTask := task.CreateNewTask(t)
	err = (*execTask).Execute(sessionManager)
	return err
}
