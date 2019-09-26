/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Instances
 * Used by client-side stubs.
 */

package instances

import (
)

// The ``Instances`` interface provides methods to access namespaces for non-administrative users. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type InstancesClient interface {


    // Returns namespaces that user making the call is authorized to access. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return List of Namespace identifiers together with the API endpoint for each namespace.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    List() ([]Summary, error) 

}
