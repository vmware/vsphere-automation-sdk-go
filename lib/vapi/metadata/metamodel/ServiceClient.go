/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Service
 * Used by client-side stubs.
 */

package metamodel


// The ``Service`` interface provides methods to retrieve metamodel information about a service element in the interface definition language. 
//
//  A service is a logical grouping of operations that operate on some entity. A service element describes a service. It contains operation elements that describe the operations grouped in the service. It also contains structure elements and enumeration elements corresponding to the structures and enumerations defined in the service.
type ServiceClient interface {

    // Returns the identifiers for the service elements that are currently registered with the infrastructure. 
    //
    //  The list of service elements is an aggregate list of all the service elements contained in all the package elements.
    // @return The list of identifiers for the service elements that are currently registered with the infrastructure.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.service``.
	List() ([]string, error)

    // Retrieves information about the service element corresponding to ``service_id``. 
    //
    //  The ServiceInfo contains the metamodel information for the operation elements, structure elements and enumeration elements contained in the service element.
    //
    // @param serviceIdParam Identifier of the service element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return The ServiceInfo instance that corresponds to ``service_id``
    // @throws NotFound if the service element associated with ``service_id`` is not registered with the infrastructure.
	Get(serviceIdParam string) (ServiceInfo, error)
}
