/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Operation
 * Used by client-side stubs.
 */

package service

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/routing"
)

// Operations to retrieve information about routing in a vAPI operation
type OperationClient interface {

    // Get the IDs of all the vAPI operations in the given service that contain routing information
    //
    // @param serviceIdParam Identifier of the service
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return list of operation identifiers
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.operation``.
    // @throws NotFound If the service identifier does not exist.
	List(serviceIdParam string) ([]string, error)

    // Get information about a vAPI operation that contains routing information
    //
    // @param serviceIdParam Identifier of the service
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @param operationIdParam Identifier of the operation
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.operation``.
    // @return Operation info for the vAPI operation that contains routing information.
    // @throws NotFound If the service identifier does not exist or if the specified operation identifier does not exist in the service.
	Get(serviceIdParam string, operationIdParam string) (routing.OperationInfo, error)
}
