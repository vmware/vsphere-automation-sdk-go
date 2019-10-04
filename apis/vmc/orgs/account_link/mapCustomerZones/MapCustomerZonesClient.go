/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: MapCustomerZones
 * Used by client-side stubs.
 */

package mapCustomerZones

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type MapCustomerZonesClient interface {


    // Creates a task to re-map customer's datacenters across zones.
    //
    // @param orgParam Organization identifier. (required)
    // @param mapZonesRequestParam The zones request information about who to map and what to map. (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    Post(orgParam string, mapZonesRequestParam model.MapZonesRequest) (model.Task, error) 

}
