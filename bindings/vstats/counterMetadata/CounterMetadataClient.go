/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: CounterMetadata
 * Used by client-side stubs.
 */

package counterMetadata

import (
)

// The ``CounterMetadata`` interface provides access to the different historical editions of counters. As computing platforms evolve over time the measurement units for different characteristics of the systems change. As such changes occur, counters will receive different editions reflected in a new metadata record. For example computer memory had changes from kilobytes through megabytes into gigabytes.
type CounterMetadataClient interface {


    // Returns information about all counter metadata records for a specific Counter.
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
    List(cidParam string, filterParam *FilterSpec) ([]Info, error) 


    // This method returns the "default" CounterMetadata. A Counter has at least one related metadata object. Usually, stats providers (for example hosts) are in agreement about the default metadata. In this case the returned list has a single metadata object. 
    //
    //  Sometimes, when providers are in "disagreement" about the default, then the returned list would include all the possible "defaults". This potential ambiguity isn't a real issue because consumers of the vStats API rarely need to specify the "mid" of metadata. Therefore, this API is used primarily for informational purposes.
    //
    // @param cidParam Counter ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.Counter``.
    // @return List of counter metadata records.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``cid`` is invalid.
    // @throws NotFound if Counter could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    GetDefault(cidParam string) ([]Info, error) 


    // Returns information about a specific CounterMetadata.
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
    Get(cidParam string, midParam string) (Info, error) 

}
