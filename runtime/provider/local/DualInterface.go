/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */
package local

import (
	"github.com/gorilla/mux"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/server/rest/contextbuilder"
)

type (
	// DualInterface combines the REST and JSON-RPC natures of bindings. This is
	// useful to provide common handling of REST and JSON-RPC natures.
	DualInterface interface {
		core.ApiInterface
		APIRESTInterface
	}

	// APIRESTInterface defines the methods specific to REST enabled server side
	// bindings
	APIRESTInterface interface {
		// RegisterRoutesHandlers registers the operations in a given binding type
		// to a router and links the binding to APIProvider and parsers for security
		// and headers.
		RegisterRoutesHandlers(router *mux.Router, provider core.APIProvider,
			appCtxBuilders []contextbuilder.ApplicationContextBuilder,
			secCtxBuilders []contextbuilder.SecurityContextBuilder)
	}
)

// RegisterRESTInterfaces registers interfaces in the router
func RegisterRESTInterfaces(router *mux.Router,
	localProvider *LocalProvider,
	appCtxBuilders []contextbuilder.ApplicationContextBuilder,
	secCtxBuilders []contextbuilder.SecurityContextBuilder,
	interfaces ...DualInterface) {
	for _, i := range interfaces {
		i.RegisterRoutesHandlers(router, localProvider, appCtxBuilders, secCtxBuilders)
	}
}
