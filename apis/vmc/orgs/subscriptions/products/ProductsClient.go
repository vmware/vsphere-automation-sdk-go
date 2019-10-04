/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Products
 * Used by client-side stubs.
 */

package products

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type ProductsClient interface {


    // List of all the products that are available for purchase.
    //
    // @param orgParam Organization identifier. (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws InternalServerError  Internal server error
    List(orgParam string) ([]model.SubscriptionProducts, error) 

}
