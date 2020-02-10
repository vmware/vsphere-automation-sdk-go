/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ResourceAddressSchemas
 * Used by client-side stubs.
 */

package vstats


// The ``ResourceAddressSchemas`` interface manages inventory of resource addressing schemas used by Counters. Each schema consists of a named list of resource identifiers of specific resource type. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ResourceAddressSchemasClient interface {

    // Returns information about a specific resource address schema. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param idParam Resource address schema identifier.
    // @return Information about the desired resource address schema.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``id`` is invalid.
    // @throws NotFound if RsrcAddrSchema could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get(idParam string) (ResourceAddressSchemasInfo, error)
}
