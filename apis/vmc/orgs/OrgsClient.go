/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Orgs
 * Used by client-side stubs.
 */

package orgs

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type OrgsClient interface {


    // Get details of organization
    //
    // @param orgParam Organization identifier. (required)
    // @return com.vmware.vmc.model.Organization
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Organization doesn't exist
    Get(orgParam string) (model.Organization, error) 


    // Return a list of all organizations the calling user (based on credential) is authorized on.
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    List() ([]model.Organization, error) 

}
