/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TechPreview
 * Used by client-side stubs.
 */

package techPreview

import (
)

// The ``TechPreview`` interface provides methods to get and set the status of Tech Preview feature state switches on a host managed by vcenter. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TechPreviewClient interface {


    // Returns the current status of the feature switches in Tech Preview mode for the feature names in the input set. If the input set is empty, returns the status of all the features. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostParam Identifier for the host from which Tech Preview information will be retrieved.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @param featuresParam Features for which status is to be retrieved.
    // If null, the status of all Tech Preview feature switches will be returned.
    // @return Mapping of feature names to their switch status.
    // @throws Error if the feature names list is not accessible.
    // @throws NotFound if a feature state switch name is not found.
    // @throws NotFound if the host is not registered on this vCenter server.
    // @throws Unauthenticated if the caller is not authenticated.
    Get(hostParam string, featuresParam map[string]bool) (map[string]Status, error) 


    // Sets the status to enabled/disabled for the Tech Preview feature switches passed in the input map. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostParam Identifier for the host on which Tech Preview information will be set.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @param featureStatusParam Map of feature names and their desired status. If a feature name is not present in the map the status of that feature will not be changed. If an unknown feature name is specified a {\\\\@link com.vmware.vapi.std.errors.NotFound exception} will be reported.
    // @throws Error if the feature names list is not accessible.
    // @throws NotFound if a feature state switch name is not found.
    // @throws NotFound if the host is not registered on this vCenter server.
    // @throws Unauthenticated if the caller is not authenticated.
    Update(hostParam string, featureStatusParam map[string]Status) error 

}
