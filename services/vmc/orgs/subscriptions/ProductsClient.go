/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Products
 * Used by client-side stubs.
 */

package subscriptions

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ProductsClient interface {

    // List of all the products that are available for purchase.
    //
    // @param orgParam Organization identifier (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws InternalServerError  Internal server error
	List(orgParam string) ([]model.SubscriptionProducts, error)
}
