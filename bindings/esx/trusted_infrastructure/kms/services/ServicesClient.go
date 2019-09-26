/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Services
 * Used by client-side stubs.
 */

package services

import (
)

// The ``Services`` interface manages endpoint configuration about the Key Management Server (KMS) service and its relation to the attestation services.
type ServicesClient interface {


    // Returns the list of all KMS service instances.
    //
    // @param filterParam Return only services matching the specified filters. If the filter doesn't match, return an empty list.
    // If {\\\\@term.unset} return all services.
    // @return List of all KMS service instances.
    // @throws Error if there is a problem accessing the stored data.
    // @throws Unauthenticated if the user can not be authenticated.
    List(filterParam *FilterSpec) ([]Summary, error) 


    // Returns the detailed information about a KMS service instance.
    //
    // @param serviceParam the attestation service instance unique identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.trusted_infrastructure.kms.service``.
    // @return Detailed information about the specified KMS service instance.
    // @throws NotFound if there is no KMS service instance with the specified identifier.
    // @throws Error if there is a problem accessing the stored data.
    // @throws Unauthenticated if the user can not be authenticated.
    Get(serviceParam string) (Info, error) 


    // Adds a new KMS service instance.
    //
    // @param specParam The CreateSpec for the new service.
    // @return Identifier of the newly registered KMS service instance.
    // The return value will be an identifier for the resource type: ``com.vmware.esx.trusted_infrastructure.kms.service``.
    // @throws AlreadyExists if there is already a KMS service instance with the same trusted_infrastructure.NetworkAddress.
    // @throws InvalidArgument if the CreateSpec contains invalid data.
    // @throws UnableToAllocateResource if the host does not have a required license.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Create(specParam CreateSpec) (string, error) 


    // Removes a currently configured KMS service instance.
    //
    // @param serviceParam the attestation service instance unique identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.trusted_infrastructure.kms.service``.
    // @throws NotFound if the service is not found.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Delete(serviceParam string) error 

}
