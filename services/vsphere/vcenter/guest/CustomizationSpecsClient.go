/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CustomizationSpecs
 * Used by client-side stubs.
 */

package guest


// The ``CustomizationSpecs`` interface provides methods to manage guest customization specifications in the vCenter Server. This interface was added in vSphere API 6.7.1.
type CustomizationSpecsClient interface {

    // Returns information about at most 1000 visible (subject to permission checks) guest customization specifications in vCenter matching the CustomizationSpecsFilterSpec. This method was added in vSphere API 6.7.1.
    //
    // @param filterParam Specification of matching guest customization specifications for which information should be returned.
    // If null, the behavior is equivalent to a CustomizationSpecsFilterSpec with all properties null which means all guest customization specifications match the filter.
    // @return Commonly used information about the guest customization specifications matching the CustomizationSpecsFilterSpec.
    // @throws InvalidArgument if the CustomizationSpecsFilterSpec#osType property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 1000 guest customization specifications match the CustomizationSpecsFilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(filterParam *CustomizationSpecsFilterSpec) ([]CustomizationSpecsSummary, error)

    // Creates a customization specification. This method was added in vSphere API 7.0.0.
    //
    // @param specParam The information i.e. name, description and the settings i.e hostname, ip address etc for the new customization specification that needs to be created.
    // @return The name of the customization specification that is created.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.guest.CustomizationSpec``.
    // @throws AlreadyExists if a customization specification is already present with the same name.
    // @throws InvalidArgument if the specified specification is not a valid one.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Create(specParam CustomizationSpecsCreateSpec) (string, error)

    // Returns the guest customization specification from vCenter with the specified identifier. This method was added in vSphere API 7.0.0.
    //
    // @param nameParam The name of the customization specification.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.guest.CustomizationSpec``.
    // @return A customization spec with the specified identifier.
    // @throws NotFound if a customization specification is not found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(nameParam string) (CustomizationSpecsInfo, error)

    // Sets an existing specification, possibly after retrieving (by using CustomizationSpecs#get) and editing it. This method was added in vSphere API 7.0.0.
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
	Set(nameParam string, specParam CustomizationSpecsSpec) error

    // Deletes a customization specification with the specified identifier. This method was added in vSphere API 7.0.0.
    //
    // @param nameParam The name of the customization specification that needs to be deleted.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.guest.CustomizationSpec``.
    // @throws NotFound if a customization specification is not found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Delete(nameParam string) error

    // Returns the content of the customization specification in the specified format. Note that any passwords in the customization specification will be set to blank values during the export method. This method was added in vSphere API 7.0.0.
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
	Export(nameParam string, formatParam CustomizationSpecsFormat) (string, error)

    // Converts a well formatted string to a CustomizationSpecsCreateSpec. The resulting object can be passed to CustomizationSpecs#create method. This method was added in vSphere API 7.0.0.
    //
    // @param customizationSpecParam content to be converted to the spec.
    // @return A proper specification of type CustomizationSpecsCreateSpec
    // @throws InvalidArgument if the specified content cannot be properly converted into a proper valid CustomizationSpecsCreateSpec object.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	ImportSpecification(customizationSpecParam string) (CustomizationSpecsCreateSpec, error)
}
