/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Routes
 * Used by client-side stubs.
 */

package routes

import (
)

// ``Routes`` interface provides methods Performs networking routes operations. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type RoutesClient interface {


    // Test connection to a list of gateways. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param gatewaysParam list of gateways.
    // @return connection status
    // @throws Error Generic error
    Test(gatewaysParam []string) (TestStatus, error) 


    // Set static routing rules. A destination of 0.0.0.0 and prefix 0 (for IPv4) or destination of :: and prefix 0 (for IPv6) refers to the default gateway. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param routesParam Static routing rules.
    // @throws Error Generic error
    Set(routesParam []Config) error 


    // Get main routing table. A destination of 0.0.0.0 and prefix 0 (for IPv4) or destination of :: and prefix 0 (for IPv6) refers to the default gateway. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Routing table.
    // @throws Error Generic error
    Get() ([]Info, error) 

}
