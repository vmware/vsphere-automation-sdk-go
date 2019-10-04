/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ConnectedAccounts
 * Used by client-side stubs.
 */

package connectedAccounts

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
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
