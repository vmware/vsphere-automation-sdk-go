/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Provider
 * Used by client-side stubs.
 */

package introspection


// The Provider service provides operations to retrieve information about a vAPI Provider. A provider is a container that exposes one or more vAPI services.
type ProviderClient interface {

    // Returns a ProviderInfo describing the vAPI provider on which the operation is invoked
    // @return ProviderInfo describing the vAPI provider on which the operation is invoked
	Get() (ProviderInfo, error)
}
