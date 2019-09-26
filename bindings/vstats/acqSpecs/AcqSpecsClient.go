/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: AcqSpecs
 * Used by client-side stubs.
 */

package acqSpecs

import (
)

// The ``AcqSpecs`` interface provides methods to perform acquisition specification related operations. An acquisition specification defines the statistical data that should be collected at desired sampling rates from the underlying providers. It designates the resources and their counters which should be sampled, and a desired sampling rate.
type AcqSpecsClient interface {


    // Create a new acquisition specification record.
    //
    // @param acqSpecParam Specification for the acquisition of stats data.
    // @return Identifier of the newly created acquisition specification.
    // The return value will be an identifier for the resource type: ``com.vmware.vstats.model.AcqSpec``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``acq_spec`` contain any errors.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Create(acqSpecParam CreateSpec) (string, error) 


    // Delete an acquisition specification.
    //
    // @param idParam Acquisition specification ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.AcqSpec``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``id`` is invalid.
    // @throws NotFound if acquisition specification could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Delete(idParam string) error 


    // Returns information about all acquisition specifications.
    //
    // @param filterParam Criteria for selecting records to return.
    // When map with bool value filtering will be applied to the result.
    // @return List of acquisition specification records.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    List(filterParam *FilterSpec) (ListResult, error) 


    // Returns information about a specific acquisition specification.
    //
    // @param idParam Acquisition specification ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.AcqSpec``.
    // @return Information about the desired acquisition specification.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``id`` is invalid.
    // @throws NotFound acquisition specification could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Get(idParam string) (Info, error) 


    // Update an existing acquisition specification.
    //
    // @param idParam Acquisition specification ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.AcqSpec``.
    // @param acqSpecParam Updated acquisition specification.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound acquisition specification could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Update(idParam string, acqSpecParam UpdateSpec) error 

}
