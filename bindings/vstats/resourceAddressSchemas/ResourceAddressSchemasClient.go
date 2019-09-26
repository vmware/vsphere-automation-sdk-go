/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ResourceAddressSchemas
 * Used by client-side stubs.
 */

package resourceAddressSchemas

import (
)

// The ``ResourceAddressSchemas`` interface manages inventory of resource addressing schemas used by Counters. Each schema consists of a named list of resource identifiers of specific resource type.
type ResourceAddressSchemasClient interface {


    // Returns information about a specific resource address schema.
    //
    // @param idParam Resource address schema identifier.
    // @return Information about the desired resource address schema.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``id`` is invalid.
    // @throws NotFound if RsrcAddrSchema could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Get(idParam string) (Info, error) 

}
