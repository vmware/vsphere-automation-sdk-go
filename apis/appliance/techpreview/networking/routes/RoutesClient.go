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

// ``Routes`` interface provides methods Performs networking routes operations.
type RoutesClient interface {


    // Test connection to a list of gateways
    //
    // @param gatewaysParam list of gateways.
    // @return connection status
    // @throws Error Generic error
    Test(gatewaysParam []string) (TestStatusInfo, error) 


    // Set static routing rules. A destination of 0.0.0.0 and prefix 0 (for IPv4) or destination of :: and prefix 0 (for IPv6) refers to the default gateway.
    //
    // @param routeParam Static routing rule.
    // @throws Error Generic error
    Add(routeParam Route) error 


    // Set static routing rules. A destination of 0.0.0.0 and prefix 0 (for IPv4) or destination of :: and prefix 0 (for IPv6) refers to the default gateway.
    //
    // @param routesParam Static routing rules.
    // @throws Error Generic error
    Set(routesParam []Route) error 


    // Get main routing table. A destination of 0.0.0.0 and prefix 0 (for IPv4) or destination of :: and prefix 0 (for IPv6) refers to the default gateway.
    // @return Routing table.
    // @throws Error Generic error
    List() ([]RouteReadOnly, error) 


    // Delete static routing rules.
    //
    // @param routeParam Static routing rule.
    // @throws Error Generic error
    Delete(routeParam Route) error 

}
