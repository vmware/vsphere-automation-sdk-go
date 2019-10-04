/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SddcTemplate
 * Used by client-side stubs.
 */

package sddcTemplate

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type SddcTemplateClient interface {


    // Get configuration template for an SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @return com.vmware.vmc.model.SddcTemplate
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  SDDC is in a state that cannot be use for generating configuration template
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the SDDC with given identifier
    Get(orgParam string, sddcParam string) (model.SddcTemplate, error) 

}
