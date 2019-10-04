/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: FeatureState
 * Used by client-side stubs.
 */

package featureState

import (
)

// The ``FeatureState`` interface provides methods to get the status of feature state switches.
type FeatureStateClient interface {


    // Returns the current status of feature state switches.
    //
    // @param featuresParam List of features for which status is to be retrieved.
    // If null, return the status for all feature switch states.
    // @return Mapping of feature names to their switch status.
    // @throws Error if feature names list is not accessible.
    // @throws NotFound if feature state switch name is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    Get(featuresParam []string) (map[string]Status, error) 

}
