/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: CustomizationSpecs
 * Used by client-side stubs.
 */

package customizationSpecs

import (
)

// The ``CustomizationSpecs`` interface provides methods to manage guest customization specifications in the vCenter Server.
type CustomizationSpecsClient interface {


    // Returns information about at most 1000 visible (subject to permission checks) guest customization specifications in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching guest customization specifications for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all guest customization specifications match the filter.
    // @return Commonly used information about the guest customization specifications matching the FilterSpec.
    // @throws InvalidArgument if the FilterSpec#osType property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 1000 guest customization specifications match the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 


    // Creates a customization specification. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam The information i.e. name, description and the settings i.e hostname, ip address etc for the new customization specification that needs to be created.
    // @return The name of the customization specification that is created.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.guest.CustomizationSpec``.
    // @throws AlreadyExists if a customization specification is already present with the same name.
    // @throws InvalidArgument if the specified specification is not a valid one.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Create(specParam CreateSpec) (string, error) 


    // Returns the guest customization specification from vCenter with the specified identifier. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param nameParam The name of the customization specification.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.guest.CustomizationSpec``.
    // @return A customization spec with the specified identifier.
    // @throws NotFound if a customization specification is not found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(nameParam string) (Info, error) 


    // Sets an existing specification, possibly after retrieving (by using CustomizationSpecs#get) and editing it. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param nameParam The name of the customization specification that needs to be set.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.guest.CustomizationSpec``.
    // @param specParam The new specification that will overwrite the existing specification.
    // @throws InvalidArgument If, based on the item's fingerprint value, the set process detects that the specification has changed since its retrieval, then the method throws InvalidArgument exception to warn the client that he might overwrite another client's change.
    // @throws InvalidArgument If the settings in ``spec`` are not valid.
    // @throws NotFound if a customization specification is not found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Set(nameParam string, specParam Spec) error 


    // Deletes a customization specification with the specified identifier. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param nameParam The name of the customization specification that needs to be deleted.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.guest.CustomizationSpec``.
    // @throws NotFound if a customization specification is not found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(nameParam string) error 


    // Returns the content of the customization specification in the specified format. Note that any passwords in the customization specification will be set to blank values during the export method. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param nameParam The name of the customization specification that has to be returned.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.guest.CustomizationSpec``.
    // @param formatParam The format in which the customization specification has to be returned.
    // @return The string representation of the customization specification in the specified format.
    // @throws InvalidArgument If value of ``format`` is not valid.
    // @throws NotFound if the customization specification is not found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Export(nameParam string, formatParam Format) (string, error) 


    // Converts a well formatted string to a CreateSpec. The resulting object can be passed to CustomizationSpecs#create method. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param customizationSpecParam content to be converted to the spec.
    // @return A proper specification of type CreateSpec
    // @throws InvalidArgument if the specified content cannot be properly converted into a proper valid CreateSpec object.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    ImportSpecification(customizationSpecParam string) (CreateSpec, error) 

}
