package utils

import (
	"reflect"

	session "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/session"
	client "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type ArgOptionalIf func() bool

//Arguments is initialized with all the command line Options, Flags & Arguments.
var Arguments Args

type properties struct {
	options map[string]*propValue
	args    []string
}

type propValue struct {
	Value             interface{}
	DefaultValue      interface{}
	Kind              reflect.Kind
	Description       string
	Required          bool
	OptionalCondition ArgOptionalIf
}

type TaskFunc func(connector client.Connector) error

type taskImpl struct {
	task TaskFunc
	name string
}

type sessionManagerImpl struct {
	sessionID     string
	sessionClient session.SessionClient
	connector     client.Connector
	sessionError  error
}
