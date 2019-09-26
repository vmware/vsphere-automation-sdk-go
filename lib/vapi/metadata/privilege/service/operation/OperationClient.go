/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Operation
 * Used by client-side stubs.
 */

package operation

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/privilege"
)

// The ``Operation`` interface provides methods to retrieve privilege information of an operation element. 
//
//  An operation element is said to contain privilege information if there are any privileges assigned to the operation element or if one of the parameter elements contained in it has privileges assigned in privilege definition file.
type OperationClient interface {


    // Returns the identifiers for the operation elements contained in the service element corresponding to ``service_id`` that have privilege information.
    //
    // @param serviceIdParam Identifier of the service element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return List of identifiers for the operation elements contained in the service element that have privilege information.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.operation``.
    // @throws NotFound if the service element associated with ``service_id`` does not have any operation elements that have privilege information.
    List(serviceIdParam string) ([]string, error) 


    // Retrieves the privilege information about an operation element corresponding to ``operation_id`` contained in the service element corresponding to ``service_id``.
    //
    // @param serviceIdParam Identifier of the service element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @param operationIdParam Identifier of the operation element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.operation``.
    // @return The privilege.OperationInfo instance that corresponds to ``operation_id``.
    // @throws NotFound if the service element associated with ``service_id`` does not exist.
    // @throws NotFound if the operation element associated with ``operation_id`` does not exist.
    // @throws NotFound if the operation element associated with ``operation_id`` does not have any privilege information.
    Get(serviceIdParam string, operationIdParam string) (privilege.OperationInfo, error) 

}
