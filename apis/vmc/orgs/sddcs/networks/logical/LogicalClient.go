/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Logical
 * Used by client-side stubs.
 */

package logical

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type LogicalClient interface {


    // Create a network in an SDDC.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param sddcNetworkParam (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
    Create(orgParam string, sddcParam string, sddcNetworkParam model.SddcNetwork) error 


    // Delete a network in an SDDC.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param networkIdParam Logical Network Identifier (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
    Delete(orgParam string, sddcParam string, networkIdParam string) error 


    // Retrieve information about a network in an SDDC.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param networkIdParam Logical Network Identifier (required)
    // @return com.vmware.vmc.model.SddcNetwork
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
    Get(orgParam string, sddcParam string, networkIdParam string) (model.SddcNetwork, error) 


    // Retrieve all networks in an SDDC.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param pageSizeParam Page size for pagination. (optional)
    // @param startIndexParam Start index of page. (optional)
    // @param prevSddcNetworkIdParam Previous logical network id. (optional)
    // @param sortOrderAscendingParam Sort order ascending. (optional)
    // @return com.vmware.vmc.model.DataPageSddcNetwork
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
    Get0(orgParam string, sddcParam string, pageSizeParam *int64, startIndexParam *int64, prevSddcNetworkIdParam *string, sortOrderAscendingParam *bool) (model.DataPageSddcNetwork, error) 


    // Modify a network in an SDDC.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param networkIdParam Logical Network Identifier (required)
    // @param sddcNetworkParam (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
    Update(orgParam string, sddcParam string, networkIdParam string, sddcNetworkParam model.SddcNetwork) error 

}
