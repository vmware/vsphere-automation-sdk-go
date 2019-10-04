/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: NoProxy
 * Used by client-side stubs.
 */

package noProxy

import (
)

// The ``NoProxy`` interface provides methods to configure a connection that does not need a proxy.
type NoProxyClient interface {


    // Sets servers for which no proxy configuration should be applied. This operation sets environment variables. In order for this operation to take effect, a logout from appliance or a service restart is required. If IPv4 is enabled, "127.0.0.1" and "localhost" will always bypass the proxy (even if they are not explicitly configured).
    //
    // @param serversParam List of strings representing servers to bypass proxy. A server can be a FQDN, IP address, FQDN:port or IP:port combinations.
    // @throws Error Generic error.
    Set(serversParam []string) error 


    // Returns servers for which no proxy configuration will be applied.
    // @return List of servers for which no proxy configuration will be applied.
    // @throws Error Generic error.
    Get() ([]string, error) 

}
