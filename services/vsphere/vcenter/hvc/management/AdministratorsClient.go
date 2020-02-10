/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Administrators
 * Used by client-side stubs.
 */

package management


// The ``Administrators`` provides methods to update, delete, and list groups in the local sso group. This is limited to the Hybrid Linked Mode service. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type AdministratorsClient interface {

    // Add the local sso group with the new group. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param groupNameParam Name of the new group to be added. Ex - xyz\\\\@abc.com where xyz is the group name and abc.com is the domain name
    // @throws Unauthorized If the user is not authorized.
    // @throws Error if the system reports an error while responding to the request.
	Add(groupNameParam string) error

    // Remove the group from the local sso group. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param groupNameParam Name of the group to be removed. Ex - xyz\\\\@abc.com where xyz is the group name and abc.com is the domain name
    // @throws Unauthorized If the user is not authorized.
    // @throws Error if the system reports an error while responding to the request.
	Remove(groupNameParam string) error

    // Sets the groups in the local sso group. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param groupNamesParam Names the groups to be in the CloudAdminGroup Ex - xyz\\\\@abc.com where xyz is the group name and abc.com is the domain name
    // @throws Unauthorized If the user is not authorized.
    // @throws Error if the system reports an error while responding to the request.
	Set(groupNamesParam map[string]bool) error

    // Enumerates the set of all the groups in the local sso group. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    // @return The map with bool value of all the groups.
    // @throws Error if the system reports an error while responding to the request.
	Get() (map[string]bool, error)
}
