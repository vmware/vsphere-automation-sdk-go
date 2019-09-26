/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Source
 * Used by client-side stubs.
 */

package source

import (
)

// Operations to manage the metadata sources for routing information
type SourceClient interface {


    // Create a new metadata source.
    //
    // @param sourceIdParam  metadata source identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
    // @param specParam  create specification.
    // @throws AlreadyExists  If the metadata source identifier is already present.
    // @throws InvalidArgument  If type of the source specified in \\\\@{link CreateSpec#type} is invalid.
    // @throws InvalidArgument  If the file specified in \\\\@{link CreateSpec#filepath} is not a valid json file.
    // @throws InvalidArgument  If the URI specified in \\\\@{link CreateSpec#address} is unreachable or not a vAPI compatible server.
    // @throws NotFound  If the file specified in \\\\@{link CreateSpec#filepath} does not exist.
    Create(sourceIdParam string, specParam CreateSpec) error 


    // Delete a metadata source.
    //
    // @param sourceIdParam  Metadata source identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
    // @throws NotFound  If the metadata source identifier is not found.
    Delete(sourceIdParam string) error 


    // Get the details about a metadata source.
    //
    // @param sourceIdParam  Metadata source identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
    // @return Metadata source info.
    // @throws NotFound  If the metadata source identifier is not found.
    Get(sourceIdParam string) (Info, error) 


    // List all the metadata sources.
    // @return List of all metadata sources.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.metadata.source``.
    List() ([]string, error) 


    // Reload metadata from all the sources or of a particular source.
    //
    // @param sourceIdParam  Metadata source identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
    //  If unspecified, all the sources are reloaded
    // @throws NotFound  If the metadata source identifier is not found.
    Reload(sourceIdParam *string) error 


    // Returns the fingerprint of all the sources or of a particular source.
    //
    // @param sourceIdParam  Metadata source identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
    //  If unspecified, fingerprint of all the sources is returned
    // @return fingerprint of all the sources or of a particular source.
    // @throws NotFound  If the metadata source identifier is not found.
    Fingerprint(sourceIdParam *string) (string, error) 

}
