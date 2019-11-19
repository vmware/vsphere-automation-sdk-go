/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Peerconfig
 * Used by client-side stubs.
 */

package edges

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

type PeerconfigClient interface {

    // Retrieve IPsec VPN peer configuration for a management or compute gateway (NSX Edge). The response output is free form text generated as per the template specified as request parameter input.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param objecttypeParam Specify object type identifier. Valid value is 'ipsecSiteConfig'. Required. (required)
    // @param objectidParam Specify object identifier, for example 'ipsecsite-1' for IPsec Site configuration. Required. (required)
    // @param templateidParam Specify template identifier and response format. Valid values are 'text', 'json' and 'xml'. Default is 'text'. Optional. (optional)
    // @return DynamicStructure
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
	Get(orgParam string, sddcParam string, edgeIdParam string, objecttypeParam string, objectidParam string, templateidParam *string) (*data.StructValue, error)
}
