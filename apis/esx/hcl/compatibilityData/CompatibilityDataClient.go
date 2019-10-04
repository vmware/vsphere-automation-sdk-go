/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: CompatibilityData
 * Used by client-side stubs.
 */

package compatibilityData

import (
)

// This interface provides methods to update the local compatibility data residing on the vCenter Appliance or to get information about the said data. The information in the data is generic VMware compatibility information for servers and devices. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityDataClient interface {


    // Provides information about the compatibility data located on the vCenter Appliance. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Information about the compatibility data.
    // @throws NotAllowedInCurrentState if there is no compatibility data on the vCenter executing the operation.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws ResourceInaccessible if the vCenter this API is executed on is not part of the Customer Experience Improvement Program (CEIP).
    // @throws Error If there is some unknown error. The accompanying error message will give more details about the failure.
    Get() (Status, error) 


    // Replaces the local compatibility data with the latest version found from VMware official source. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws NotAllowedInCurrentState if there is compatibility data update in progress.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws ResourceInaccessible if the vCenter this API is executed on is not part of the Customer Experience Improvement Program (CEIP).
    // @throws Error If there is some unknown error. The accompanying error message will give more details about the failure.
    Update() error 

}
