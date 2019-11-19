/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: OfferInstances
 * Used by client-side stubs.
 */

package subscriptions

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
)

type OfferInstancesClient interface {

    // List all offers available for the specific product type in the specific region
    //
    // @param orgParam Organization identifier. (required)
    // @param regionParam Region for the offer (required)
    // @param productTypeParam Type of the product in offers. \*This has been deprecated\*. Please use product & type parameters (required)
    // @param productParam The product that you are trying to purchase, eg. host. This needs to be accompanied by the type parameter (optional)
    // @param type_Param The type/flavor of the product you are trying it purchase,eg. an \\\\`r5.metal\\\\` host. This needs to be accompanied by the product parameter. (optional)
    // @return com.vmware.vmc.model.OfferInstancesHolder
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Bad Request. Type of the product not supported.
    // @throws Unauthorized  Forbidden
	List(orgParam string, regionParam string, productTypeParam string, productParam *string, type_Param *string) (model.OfferInstancesHolder, error)
}
