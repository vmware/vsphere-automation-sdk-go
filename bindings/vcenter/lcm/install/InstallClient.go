/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Install
 * Used by client-side stubs.
 */

package install

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/lcm"
)

// The service to install Embedded VCSA, PSC, Management VCSA, VMC gateway. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type InstallClient interface {


    // Performs a precheck for the given specification. The result of this operation can be queried by calling the cis/tasks/{task-id} with the task-id in the response of this call. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam  The specification of the deployment.
    // @param optionsParam  The deployment precheck options.
    // @return  The identifier of the operation.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.lcm.task``.
    // @throws InvalidArgument  If the given spec and/or option contains error.
    Check(specParam lcm.InstallSpec, optionsParam *lcm.DeploymentOption) (string, error) 


    // Deploys the appliance for the given specification. The result of this operation can be queried by calling the cis/tasks/{task-id} with the task-id in the response of this call. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam  The specification of the deployment.
    // @return  The identifier of the operation.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.lcm.task``.
    // @throws InvalidArgument  If the given specification contains error.
    Start(specParam lcm.InstallSpec) (string, error) 

}
