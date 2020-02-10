/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CounterMetadata
 * Used by client-side stubs.
 */

package vstats


// The ``CounterMetadata`` interface provides access to the different historical editions of counters. As computing platforms evolve over time the measurement units for different characteristics of the systems change. As such changes occur, counters will receive different editions reflected in a new metadata record. For example computer memory had changes from kilobytes through megabytes into gigabytes. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CounterMetadataClient interface {

    // Returns information about all counter metadata records for a specific Counter. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param cidParam Counter ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.Counter``.
    // @param filterParam Filter specification.
    // When null no filtering will be performed.
    // @return List of counter metadata for the specified counter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if Counter could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	List(cidParam string, filterParam *CounterMetadataFilterSpec) ([]CounterMetadataInfo, error)

    // This method returns the "default" CounterMetadata. A Counter has at least one related metadata object. Usually, stats providers (for example hosts) are in agreement about the default metadata. In this case the returned list has a single metadata object. 
    //
    //  Sometimes, when providers are in "disagreement" about the default, then the returned list would include all the possible "defaults". This potential ambiguity isn't a real issue because consumers of the vStats API rarely need to specify the "mid" of metadata. Therefore, this API is used primarily for informational purposes.. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param cidParam Counter ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.Counter``.
    // @return List of counter metadata records.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``cid`` is invalid.
    // @throws NotFound if Counter could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	GetDefault(cidParam string) ([]CounterMetadataInfo, error)

    // Returns information about a specific CounterMetadata. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param cidParam Counter ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.Counter``.
    // @param midParam CounterMetadata ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.CounterMetadata``.
    // @return Information about the desired CounterMetadata.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if Counter or CounterMetadata could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get(cidParam string, midParam string) (CounterMetadataInfo, error)
}
