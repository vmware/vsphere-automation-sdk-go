/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Operation
 * Used by client-side stubs.
 */

package service

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/metadata/authentication"
)

// The ``Operation`` interface provides methods to retrieve authentication information of an operation element. 
//
//  An operation element is said to contain authentication information if authentication schemes are specified in the authentication definition file.
type OperationClient interface {

    // Returns the identifiers for the operation elements contained in the service element corresponding to ``service_id`` that have authentication information.
    //
    // @param serviceIdParam Identifier of the service element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return List of identifiers for the operation elements contained in the service element that have authentication information.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.operation``.
    // @throws NotFound if the service element associated with ``service_id`` does not have any operation elements that have authentication information.
	List(serviceIdParam string) ([]string, error)

    // Retrieves the authentication information about an operation element corresponding to ``operation_id`` contained in the service element corresponding to ``service_id``.
    //
    // @param serviceIdParam Identifier of the service element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @param operationIdParam Identifier of the operation element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.operation``.
    // @return The authentication.OperationInfo instance that corresponds to ``operation_id``.
    // @throws NotFound if the service element associated with ``service_id`` does not exist.
    // @throws NotFound if the operation element associated with ``operation_id`` does not exist.
    // @throws NotFound if the operation element associated with ``operation_id`` does not have any authentication information.
	Get(serviceIdParam string, operationIdParam string) (authentication.OperationInfo, error)
}
