/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Public
 * Used by client-side stubs.
 */

package public

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type PublicClient interface {


    // Update the DNS records of management VMs to use public IP addresses
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates or invalid body
    // @throws Unauthorized  Forbidden
    Update(orgParam string, sddcParam string) (model.Task, error) 

}
