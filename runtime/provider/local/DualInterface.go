/* Copyright © 2020-2021 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package local

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rest"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rest/contextbuilder"
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
		RegisterRoutesHandlers(router rest.Router, provider core.APIProvider,
			appCtxBuilders []contextbuilder.ApplicationContextBuilder,
			secCtxBuilders []contextbuilder.SecurityContextBuilder,
			opts ...rest.RequestHandlerOption)
	}
)

// RegisterRESTInterfaces registers interfaces in the router
// Deprecated: This method does not allow configuring opentracing Tracer. Please use the bindings *ApiInterface
// RegisterRoutesHandlers
func RegisterRESTInterfaces(router rest.Router,
	localProvider *LocalProvider,
	appCtxBuilders []contextbuilder.ApplicationContextBuilder,
	secCtxBuilders []contextbuilder.SecurityContextBuilder,
	interfaces ...DualInterface) {
	for _, i := range interfaces {
		i.RegisterRoutesHandlers(router, localProvider, appCtxBuilders, secCtxBuilders)
	}
}
