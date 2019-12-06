/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Source
 * Used by client-side stubs.
 */

package privilege


// The ``Source`` interface provides methods to manage the sources of privilege metadata information. 
//
//  The interface definition language infrastructure provides tools to generate various kinds of metadata in JSON format from the interface definition files and additional properties files. One of the generated files contains privilege information. The generated file can be registered as a source of metadata. 
//
//  The privilege file contains all the data present in the interface definition files. Each privilege file contains data about one component element. When a privilege file is added as a source, each source contributes only one component element's metadata. 
//
//  Privilege metadata can also be discovered from a remote server that supports the privilege metadata interfaces (see com.vmware.vapi.metadata.privilege). Since multiple components can be registered with a single metadata server, when a remote server is registered as a source, that source can contribute more than one component.
type SourceClient interface {

    // Creates a new metadata source. Once the server validates the registration information of the metadata source, the privilege metadata is retrieved from the source. This populates elements in all the interfaces defined in com.vmware.vapi.metadata.privilege package.
    //
    // @param sourceIdParam metadata source identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
    // @param specParam create specification.
    // @throws AlreadyExists if the metadata source identifier is already registered with the infrastructure.
    // @throws InvalidArgument if the type of the source specified in null is invalid.
    // @throws InvalidArgument if the file specified in null is not a valid JSON file or if the format of the privilege metadata in the JSON file is invalid.
    // @throws InvalidArgument if the URI specified in null is unreachable or if there is a transport protocol or message protocol mismatch between the client and the server or if the remote server do not have interfaces present in com.vmware.vapi.metadata.privilege package.
    // @throws NotFound if the file specified in null does not exist.
	Create(sourceIdParam string, specParam SourceCreateSpec) error

    // Deletes an existing privilege metadata source from the infrastructure.
    //
    // @param sourceIdParam Identifier of the metadata source.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
    // @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Delete(sourceIdParam string) error

    // Retrieves information about the metadata source corresponding to ``source_id``.
    //
    // @param sourceIdParam Identifier of the metadata source.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
    // @return The SourceInfo instance that corresponds to ``source_id``
    // @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Get(sourceIdParam string) (SourceInfo, error)

    // Returns the identifiers of the metadata sources currently registered with the infrastructure.
    // @return The list of identifiers for metadata sources currently registered.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	List() ([]string, error)

    // Reloads the privilege metadata from all the metadata sources or of a particular metadata source if ``source_id`` is specified.
    //
    // @param sourceIdParam Identifier of the metadata source.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
    // If unspecified, all the metadata sources are reloaded.
    // @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Reload(sourceIdParam *string) error

    // Returns the aggregate fingerprint of metadata from all the metadata sources or from a particular metadata source if ``source_id`` is specified.
    //
    // @param sourceIdParam Identifier of the metadata source.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
    // If unspecified, the fingerprint of all the metadata sources is returned.
    // @return Aggregate fingerprint of all the metadata sources or of a particular metadata source.
    // @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Fingerprint(sourceIdParam *string) (string, error)
}
