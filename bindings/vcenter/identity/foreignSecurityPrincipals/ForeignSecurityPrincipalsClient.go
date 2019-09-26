/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ForeignSecurityPrincipals
 * Used by client-side stubs.
 */

package foreignSecurityPrincipals

import (
)

// The ``ForeignSecurityPrincipals`` interface provides methods to read and modify foreign security principals. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ForeignSecurityPrincipalsClient interface {


    // Retrieve foreign security principal based on the identifier. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param principalParam the principal identifier or name.
    // @return foreign security principal information.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if the principal is not found.
    // @throws InvalidArgument if id contains invalid information
    Get(principalParam string) (Info, error) 


    // Create a foreign security principal. If the same Id already exists, update the record. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param idParam the principal identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.identity.ForeignSecurityPrincipals``.
    // @param specParam the information to create the principal.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws InvalidArgument if id or spec contains invalid information
    Create(idParam string, specParam CreateSpec) error 


    // Update a foreign security principal. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param idParam the principal identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.identity.ForeignSecurityPrincipals``.
    // @param specParam the information to update the principal.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if the principal is not found.
    // @throws InvalidArgument if the spec contains invalid information
    Update(idParam string, specParam UpdateSpec) error 


    // Retrieve all domain names associated with foreign security principals. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return a set of domain names. Returns an empty set if no domains.
    // @throws Unauthorized if authorization is not given to caller.
    ListDomains() (map[string]bool, error) 

}
