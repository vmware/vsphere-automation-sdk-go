/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Operation
 * Used by client-side stubs.
 */

package introspection


// The Operation service provides operations to retrieve information about the operations present in a vAPI service.
type OperationClient interface {

    // Returns the set of operation identifiers for a given vAPI service.
    //
    // @param serviceIdParam service identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return set of operation identifiers for a given vAPI service.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.operation``.
    // @throws NotFound If the service identifier does not exist.
	List(serviceIdParam string) (map[string]bool, error)

    // Returns the OperationInfo for a given vAPI operation.
    //
    // @param serviceIdParam service identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @param operationIdParam operation identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.operation``.
    // @return OperationInfo for a given vAPI operation.
    // @throws NotFound If the operation identifier does not exist.
    // @throws NotFound If the service identifier does not exist.
	Get(serviceIdParam string, operationIdParam string) (OperationInfo, error)
}
