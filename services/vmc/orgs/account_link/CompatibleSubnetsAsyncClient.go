/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CompatibleSubnetsAsync
 * Used by client-side stubs.
 */

package account_link

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type CompatibleSubnetsAsyncClient interface {

    // Gets a customer's compatible subnets for account linking via a task. The information is returned as a member of the task (found in task.params['subnet_list_result'] when you are notified it is complete), and it's documented under ref /definitions/AwsCompatibleSubnets
    //
    // @param orgParam Organization identifier. (required)
    // @param linkedAccountIdParam The linked connected account identifier (required)
    // @param regionParam The region of the cloud resources to work in (optional)
    // @param sddcParam sddc (optional)
    // @param instanceTypeParam The server instance type to be used. (optional)
    // @param sddcTypeParam The sddc type to be used. (1NODE, SingleAZ, MultiAZ) (optional)
    // @param numOfHostsParam The number of hosts (optional)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (model.Task, error)

    // Sets which subnet to use to link accounts and finishes the linking process via a task
    //
    // @param awsSubnetParam The subnet chosen by the customer (required)
    // @param orgParam Organization identifier. (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Post(awsSubnetParam model.AwsSubnet, orgParam string) (model.Task, error)
}
