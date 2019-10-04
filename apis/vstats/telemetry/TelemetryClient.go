/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Telemetry
 * Used by client-side stubs.
 */

package telemetry

import (
)

// The ``Telemetry`` interface provides methods to retrieve telemetry data.
type TelemetryClient interface {


    // Returns a view to metrics internal to the vStats service instance.
    //
    // @param filterParam Criteria for selecting telemetry data to return.
    // If null returns all the telemetry data.
    // @return Telemetry data.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if prefix does not match any telemetry data.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Get(filterParam *FilterSpec) (Info, error) 

}
