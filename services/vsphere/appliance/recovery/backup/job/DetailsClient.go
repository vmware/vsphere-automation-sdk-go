/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Details
 * Used by client-side stubs.
 */

package job


// The ``Details`` interface provides methods to get the details about backup jobs. This interface was added in vSphere API 6.7.
type DetailsClient interface {

    // Returns detailed information about the current and historical backup jobs. This method was added in vSphere API 6.7.
    //
    // @param filterParam Specification of matching backup jobs for which information should be returned.
    // If null, the behavior is equivalent to DetailsFilterSpec with all properties null which means all the backup jobs match the filter.
    // @return Map of backup job identifier to Info Structure.
    // The key in the return value map will be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.job``.
    // @throws Error if any error occurs during the execution of the operation.
	List(filterParam *DetailsFilterSpec) (map[string]DetailsInfo, error)
}
