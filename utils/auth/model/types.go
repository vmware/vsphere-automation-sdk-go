/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package model defines the common models for auth package.
package model

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
)

// AuthInfo defines the base interface for all the Authentication Schemes.
type AuthInfo interface {
	// Name gets the name of the Authentication Scheme.
	Name() string
	// AuthInterface gets the interface of the used Authentication Scheme.
	AuthInterface() interface{}
	// SecurityContext creates and return new Security Context from the Authentication.
	SecurityContext() (core.SecurityContext, error)
}
