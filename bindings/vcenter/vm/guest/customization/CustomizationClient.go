/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Customization
 * Used by client-side stubs.
 */

package customization

import (
)

// The ``Customization`` interface provides methods to apply a customization specification to a virtual machine. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CustomizationClient interface {


    // Applies a customization specification in ``spec`` on the virtual machine in ``vm``. This method only sets the specification settings for the virtual machine. The actual customization happens inside the guest when the virtual machine is powered on. If ``spec`` has null values, then any pending customization settings for the virtual machine are cleared. If there is a pending customization for the virtual machine and ``spec`` has valid content, then the existing customization setting will be overwritten with the new settings. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
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
    Set(vmParam string, specParam SetSpec) error 


    // Returns the status of the customization operation that has been applied for the virtual machine in ``vm``. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam The unique identifier of the virtual machine that needs to be queried.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return The status of the customization operation applied for the virtual machine.
    // @throws NotFound If the virtual machine ``vm`` is not found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string) (Info, error) 

}
