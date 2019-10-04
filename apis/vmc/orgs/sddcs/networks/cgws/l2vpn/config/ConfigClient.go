/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Config
 * Used by client-side stubs.
 */

package config

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type ConfigClient interface {


    // Delete SDDC L2 VPN configuration.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Compute Gateway Edge Identifier (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
    Delete(orgParam string, sddcParam string, edgeIdParam string) error 


    // Retrieve SDDC L2 VPN configuration.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Compute Gateway Edge Identifier (required)
    // @param showSensitiveDataParam (optional)
    // @return com.vmware.vmc.model.Nsxl2vpn
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
    Get(orgParam string, sddcParam string, edgeIdParam string, showSensitiveDataParam *bool) (model.Nsxl2vpn, error) 


    // Modify SDDC L2 VPN configuration
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Compute Gateway Edge Identifier (required)
    // @param nsxl2vpnParam (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
    Update(orgParam string, sddcParam string, edgeIdParam string, nsxl2vpnParam model.Nsxl2vpn) error 

}
