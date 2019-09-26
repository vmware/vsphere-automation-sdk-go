/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: BaseImages
 * Used by client-side stubs.
 */

package baseImages

import (
)

// The ``BaseImages`` interface provides methods to manage trusted instances of ESX software.
type BaseImagesClient interface {


    // Import ESX metadata as a new trusted base image. 
    //
    //  Import a boot_imgdb.tgz file which contains metadata that describes a trusted ESX base image. A boot_imgdb.tgz file can be downloaded from a representative host over HTTP using: https://[hostname]/boot_imgdb.tgz
    //
    // @param imgdbParam The ESX metadata file.
    // @return The imported imgdb version identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.esx.attestation.os.esx.base_images``.
    // @throws AlreadyExists if ESX metadata for the same version has already been added.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the imgdb is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    ImportFromImgdb(imgdbParam []byte) (string, error) 


    // Return a list of trusted ESX base images.
    // @return A list of configured trusted ESX base images.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    List() ([]Summary, error) 


    // Remove a trusted ESX base image.
    //
    // @param versionParam The ESX base image version.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.attestation.os.esx.base_images``.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the version is invalid.
    // @throws NotFound if the version is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    // * The resource ``com.vmware.esx.attestation.os.esx.base_images`` referenced by the parameter ``version`` requires ``SECURITY_MGMT``.
    Delete(versionParam string) error 


    // Get the trusted ESX base version details.
    //
    // @param versionParam The ESX base image version.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.attestation.os.esx.base_images``.
    // @return The version info.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the version is invalid.
    // @throws NotFound if ESX version is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    // * The resource ``com.vmware.esx.attestation.os.esx.base_images`` referenced by the parameter ``version`` requires ``SECURITY_MGMT``.
    Get(versionParam string) (Info, error) 

}
