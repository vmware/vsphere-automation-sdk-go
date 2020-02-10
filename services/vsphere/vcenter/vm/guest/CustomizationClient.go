/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Customization
 * Used by client-side stubs.
 */

package guest


// The ``Customization`` interface provides methods to apply a customization specification to a virtual machine. This interface was added in vSphere API 7.0.0.
type CustomizationClient interface {

    // Applies a customization specification in ``spec`` on the virtual machine in ``vm``. This method only sets the specification settings for the virtual machine. The actual customization happens inside the guest when the virtual machine is powered on. If ``spec`` has null values, then any pending customization settings for the virtual machine are cleared. If there is a pending customization for the virtual machine and ``spec`` has valid content, then the existing customization setting will be overwritten with the new settings. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam The unique identifier of the virtual machine that needs to be customized.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam The customization settings to be applied to the guest operating system.
    // @throws InvalidArgument if the customization settings in ``spec`` are not valid.
    // @throws NotAllowedInCurrentState if the virtual machine ``vm`` is not in a powered off state.
    // @throws NotFound if a customization specification is not found with the unique name in ``spec``.
    // @throws NotFound If the virtual machine ``vm`` is not found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Set(vmParam string, specParam CustomizationSetSpec) error
}
