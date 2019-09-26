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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/metamodel"
)

// The ``Operation`` interface provides methods to retrieve metamodel information of an operation element in the interface definition language.
type OperationClient interface {


    // Returns the identifiers for the operation elements that are defined in the scope of ``service_id``.
    //
    // @param serviceIdParam Identifier of the service element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return The list of identifiers for the operation elements that are defined in the scope of ``service_id``.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.operation``.
    // @throws NotFound if the service element associated with ``service_id`` does not exist in any of the package elements.
    List(serviceIdParam string) ([]string, error) 


    // Retrieves the metamodel information about an operation element corresponding to ``operation_id`` contained in the service element corresponding to ``service_id``.
    //
    // @param serviceIdParam Identifier of the service element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @param operationIdParam Identifier of the operation element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.operation``.
    // @return The metamodel.OperationInfo instance that corresponds to ``operation_id`` defined in scope ``service_id``.
    // @throws NotFound if the service element associated with ``service_id`` does not exist in any of the package elements.
    // @throws NotFound if the operation element associated with ``operation_id`` does not exist in the service element.
    Get(serviceIdParam string, operationIdParam string) (metamodel.OperationInfo, error) 

}
