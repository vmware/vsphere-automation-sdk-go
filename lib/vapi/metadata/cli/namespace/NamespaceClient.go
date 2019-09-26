/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Namespace
 * Used by client-side stubs.
 */

package namespace

import (
)

// The ``Namespace`` interface provides methods to get information about command line interface (CLI) namespaces.
type NamespaceClient interface {


    // Returns the identifiers of all namespaces registered with the infrastructure.
    // @return Identifiers of all the namespaces.
    List() ([]Identity, error) 


    // Retreives information about a namespace including information about children of that namespace.
    //
    // @param identityParam Identifier of the namespace for which to retreive information.
    // @return Information about the namespace including information about child of that namespace.
    // @throws NotFound if a namespace corresponding to ``identity`` doesn't exist.
    Get(identityParam Identity) (Info, error) 


    // Returns the aggregate fingerprint of all the namespace metadata from all the metadata sources. 
    //
    //  The fingerprint provides clients an efficient way to check if the metadata for namespaces has been modified on the server.
    // @return Fingerprint of all the namespace metadata present on the server.
    Fingerprint() (string, error) 

}
