/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AcqSpecs
 * Used by client-side stubs.
 */

package vstats


// The ``AcqSpecs`` interface provides methods to perform acquisition specification related operations. An acquisition specification defines the statistical data that should be collected at desired sampling rates from the underlying providers. It designates the resources and their counters which should be sampled, and a desired sampling rate. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type AcqSpecsClient interface {

    // Create a new acquisition specification record. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param acqSpecParam Specification for the acquisition of stats data.
    // @return Identifier of the newly created acquisition specification.
    // The return value will be an identifier for the resource type: ``com.vmware.vstats.model.AcqSpec``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``acq_spec`` contain any errors.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Create(acqSpecParam AcqSpecsCreateSpec) (string, error)

    // Delete an acquisition specification. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param idParam Acquisition specification ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.AcqSpec``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``id`` is invalid.
    // @throws NotFound if acquisition specification could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Delete(idParam string) error

    // Returns information about all acquisition specifications. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param filterParam Criteria for selecting records to return.
    // When map with bool value filtering will be applied to the result.
    // @return List of acquisition specification records.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	List(filterParam *AcqSpecsFilterSpec) (AcqSpecsListResult, error)

    // Returns information about a specific acquisition specification. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param idParam Acquisition specification ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.AcqSpec``.
    // @return Information about the desired acquisition specification.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``id`` is invalid.
    // @throws NotFound acquisition specification could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get(idParam string) (AcqSpecsInfo, error)

    // Update an existing acquisition specification. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param idParam Acquisition specification ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.AcqSpec``.
    // @param acqSpecParam Updated acquisition specification.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound acquisition specification could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Update(idParam string, acqSpecParam AcqSpecsUpdateSpec) error
}
