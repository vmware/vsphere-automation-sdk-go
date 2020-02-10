/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: TagUsage
 * Used by client-side stubs.
 */

package policies


// The ``TagUsage`` interface provides methods to query which tags are used by policies in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TagUsageClient interface {

    // Returns information about the tags used by policies available in this vCenter server matching the TagUsageFilterSpec in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param filterParam Specification for matching tags used by policies.
    // If null, the behavior is equivalent to a TagUsageFilterSpec with all properties null, which means all tags used by policies match the filter.
    // @return The list of tags used by policies available on this vCenter server matching the TagUsageFilterSpec.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	List(filterParam *TagUsageFilterSpec) ([]TagUsageSummary, error)
}
