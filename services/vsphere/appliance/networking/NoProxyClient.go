/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: NoProxy
 * Used by client-side stubs.
 */

package networking


// The ``NoProxy`` interface provides methods to configure a connection that does not need a proxy. This interface was added in vSphere API 6.7.
type NoProxyClient interface {

    // Sets servers for which no proxy configuration should be applied. This operation sets environment variables. In order for this operation to take effect, a logout from appliance or a service restart is required. If IPv4 is enabled, "127.0.0.1" and "localhost" will always bypass the proxy (even if they are not explicitly configured). This method was added in vSphere API 6.7.
    //
    // @param serversParam List of strings representing servers to bypass proxy. A server can be a FQDN, IP address, FQDN:port or IP:port combinations.
    // @throws Error Generic error.
	Set(serversParam []string) error

    // Returns servers for which no proxy configuration will be applied. This method was added in vSphere API 6.7.
    // @return List of servers for which no proxy configuration will be applied.
    // @throws Error Generic error.
	Get() ([]string, error)
}
