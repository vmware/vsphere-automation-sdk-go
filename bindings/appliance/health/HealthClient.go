/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Health
 * Used by client-side stubs.
 */

package health

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance"
)

// The ``Health`` interface provides methods to retrieve the appliance health information.
type HealthClient interface {


    // Get health messages.
    //
    // @param itemParam ID of the data item
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.health``.
    // @return List of the health messages
    // @throws NotFound Unknown health item
    // @throws Error Generic error
    Messages(itemParam string) ([]appliance.Notification, error) 

}
