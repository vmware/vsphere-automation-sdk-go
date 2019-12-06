/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package introspection

import "github.com/vmware/vsphere-automation-sdk-go/runtime/core"

//should be called ApiIntrospector instead.
// the act is introspection. Actor is introspector. like stringer
type APIIntrospection interface {
	GetServices() []string
	GetServiceInfo(serviceID string) core.MethodResult
	GetCheckSum() string
	GetIntrospectionServices() []core.ApiInterface
	GetOperations(serviceID string) core.MethodResult
	GetOperationInfo(serviceID string, operationID string) core.MethodResult
}
