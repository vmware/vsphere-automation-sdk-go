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

// The ``TechPreview`` interface provides methods to get and set the status of Tech Preview feature state switches. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TechPreviewClient interface {


    // Returns the current status of all the feature switches in Tech Preview. mode. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param featuresParam Features for which status is to be retrieved.
    // If null, the status of all Tech Preview feature switches will be returned.
    // @return Mapping of feature names to their switch status.
    // @throws Error if the feature names list is not accessible.
    // @throws NotFound if a feature state switch name is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    Get(featuresParam map[string]bool) (map[string]Status, error) 


    // Sets the status to enabled/disabled for the Tech Preview feature switches present in the input map. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param featureStatusParam map of feature names and their desired status. If a feature name is not present in the map the status of that feature will not be changed. If an unknown feature name is specified a {\\\\@link com.vmware.vapi.std.errors.NotFound exception} will be reported.
    // @throws Error if the feature names list is not accessible.
    // @throws NotFound if a feature state switch name is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    Update(featureStatusParam map[string]Status) error 

}
