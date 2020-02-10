/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Proxy
 * Used by client-side stubs.
 */

package networking


// The ``Proxy`` interface provides methods Proxy configuration. This interface was added in vSphere API 6.7.
type ProxyClient interface {

    // Tests a proxy configuration by testing the connection to the proxy server and test host. This method was added in vSphere API 6.7.
    //
    // @param hostParam A hostname, IPv4 or Ipv6 address.
    // @param protocolParam Protocol whose proxy is to be tested.
    // @param configParam Proxy configuration to test.
    // @return Status of proxy settings.
    // @throws Error Generic error.
	Test(hostParam string, protocolParam string, configParam ProxyConfig) (ProxyTestResult, error)

    // Configures which proxy server to use for the specified protocol. This operation sets environment variables for using proxy. In order for this configuration to take effect a logout / service restart is required. This method was added in vSphere API 6.7.
    //
    // @param protocolParam The protocol for which proxy should be set.
    // @param configParam Proxy configuration for the specific protocol.
    // @throws Error Generic error.
	Set(protocolParam string, configParam ProxyConfig) error

    // Deletes a proxy configuration for a specific protocol. This method was added in vSphere API 6.7.
    //
    // @param protocolParam ID whose proxy is to be deleted.
    // @throws Error Generic error.
	Delete(protocolParam string) error

    // Gets proxy configuration for all configured protocols. This method was added in vSphere API 6.7.
    // @return Proxy configuration for all configured protocols.
    // @throws Error Generic error.
	List() (map[ProxyProtocol]ProxyConfig, error)

    // Gets the proxy configuration for a specific protocol. This method was added in vSphere API 6.7.
    //
    // @param protocolParam The protocol whose proxy configuration is requested.
    // @return Proxy configuration for a specific protocol.
    // @throws Error Generic error.
	Get(protocolParam string) (ProxyConfig, error)
}
