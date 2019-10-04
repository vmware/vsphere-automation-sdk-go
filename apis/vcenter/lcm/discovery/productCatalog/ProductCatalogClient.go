/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ProductCatalog
 * Used by client-side stubs.
 */

package productCatalog

import (
)

// The ``ProductCatalog`` interface provides information which VMware products can be associated with vCenter Server.
type ProductCatalogClient interface {


    // Retrieves a list of all VMware products that can be associated with vCenter Server.
    // @return List of all the VMware products which can be associated with vCenter Server
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Error If there is some unknown internal error. The accompanying error message will give more details about the failure.
    List() ([]Summary, error) 

}
