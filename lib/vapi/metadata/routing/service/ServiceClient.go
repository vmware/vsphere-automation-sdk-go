/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Service
 * Used by client-side stubs.
 */

package service

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/routing"
)

// Operations to retrieve information about routing information of a vAPI service
type ServiceClient interface {


    // Get list of all vAPI services that have operations with routing information
    // @return list of fully qualified service names
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.service``.
    List() ([]string, error) 


    // Get the routing information for a vAPI service
    //
    // @param serviceIdParam fully qualified service name
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return identifier information for the vAPI service
    // @throws NotFound If the service name does not exist
    Get(serviceIdParam string) (routing.ServiceInfo, error) 

}
