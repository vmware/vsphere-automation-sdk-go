/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package transport_nodes

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatusClient interface {

    // This API will list all transport node and statuses based on idfw enabled compute collection ID.
    //
    // @param computeCollectionIdParam Compute colelction id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.IdfwTransportNodeStatusListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(computeCollectionIdParam string, enforcementPointPathParam *string) (model.IdfwTransportNodeStatusListResult, error)
}
