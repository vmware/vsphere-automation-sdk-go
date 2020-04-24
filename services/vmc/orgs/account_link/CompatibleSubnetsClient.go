/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CompatibleSubnets
 * Used by client-side stubs.
 */

package account_link

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type CompatibleSubnetsClient interface {

    // Gets a customer's compatible subnets for account linking
    //
    // @param orgParam Organization identifier. (required)
    // @param linkedAccountIdParam The linked connected account identifier (required)
    // @param regionParam The region of the cloud resources to work in (optional)
    // @param sddcParam sddc (optional)
    // @param forceRefreshParam When true, forces the mappings for datacenters to be refreshed for the connected account. (optional)
    // @param instanceTypeParam The server instance type to be used. (optional)
    // @param sddcTypeParam The sddc type to be used. (1NODE, SingleAZ, MultiAZ) (optional)
    // @param numOfHostsParam The number of hosts (optional)
    // @return com.vmware.vmc.model.AwsCompatibleSubnets
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, forceRefreshParam *bool, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (model.AwsCompatibleSubnets, error)

    // Sets which subnet to use to link accounts and finishes the linking process
    //
    // @param orgParam Organization identifier. (required)
    // @return com.vmware.vmc.model.AwsSubnet
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Post(orgParam string) (model.AwsSubnet, error)
}
