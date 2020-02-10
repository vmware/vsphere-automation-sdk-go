/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ResourceTypes
 * Used by client-side stubs.
 */

package vstats


// The ``ResourceTypes`` interface provides list of resource types. Resource refers to any item or concept that could have measurable properties. It is a prerequisite that a resource can be identified. 
//
//  Example resource types are virtual machine, virtual disk etc.. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ResourceTypesClient interface {

    // Returns a list of available resource types. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    // @return List of resource types.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	List() ([]ResourceTypesSummary, error)
}
