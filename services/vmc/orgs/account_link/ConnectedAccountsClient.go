/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ConnectedAccounts
 * Used by client-side stubs.
 */

package account_link

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ConnectedAccountsClient interface {

    // Delete a particular connected (linked) account.
    //
    // @param orgParam Organization identifier. (required)
    // @param linkedAccountPathIdParam The linked connected account identifier (required)
    // @param forceEvenWhenSddcPresentParam When true, forcibly removes a connected account even when SDDC's are still linked to it. (optional)
    // @return com.vmware.vmc.model.AwsCustomerConnectedAccount
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  An invalid connected account ID was specified, or the connection still has SDDCs active on it.
    // @throws Unauthorized  Forbidden
	Delete(orgParam string, linkedAccountPathIdParam string, forceEvenWhenSddcPresentParam *bool) (model.AwsCustomerConnectedAccount, error)

    // Get a list of connected accounts
    //
    // @param orgParam Organization identifier. (required)
    // @param providerParam The cloud provider of the SDDC (AWS or ZeroCloud). Default value is AWS. (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Get(orgParam string, providerParam *string) ([]model.AwsCustomerConnectedAccount, error)
}
