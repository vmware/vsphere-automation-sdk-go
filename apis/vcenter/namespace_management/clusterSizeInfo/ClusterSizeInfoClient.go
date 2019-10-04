/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ClusterSizeInfo
 * Used by client-side stubs.
 */

package clusterSizeInfo

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/namespace_management"
)

// The {\\\\@name cluster-size-info} interface provides methods to retrieve various sizes available for enabling Namespaces and information about each size. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClusterSizeInfoClient interface {


    // Get information about the default values associated with various sizes. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Information for each size.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    Get() (map[namespace_management.SizingHint]Info, error) 

}
