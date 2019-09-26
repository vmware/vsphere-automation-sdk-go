/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Access
 * Used by client-side stubs.
 */

package access

import (
)

// The ``Access`` interface provides methods to manage access control of subjects on namespaces. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AccessClient interface {


    // Set up access control for the subject on given domain on the namespace. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @param domainParam The domain of the subject.
    // @param subjectParam The principal for this operation.
    // @param specParam Information about the access control to be created.
    // @throws AlreadyExists if the specified principal on given domain is already associated with a role on the namespace.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the namespace is marked for deletion or the associated cluster is being disabled.
    // @throws NotFound if ``namespace`` cannot be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Configure privilege.
    Create(namespaceParam string, domainParam string, subjectParam string, specParam CreateSpec) error 


    // Remove access control of the subject on given domain from the namespace. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @param domainParam The domain of the subject.
    // @param subjectParam The principal for this operation.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the namespace is marked for deletion or the associated cluster is being disabled.
    // @throws NotFound if ``namespace`` cannot be located.
    // @throws Unsupported if the specified principal on given domain is not associated with the namespace.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Configure privilege.
    Delete(namespaceParam string, domainParam string, subjectParam string) error 


    // Set new access control on the namespace for the subject on given domain. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @param domainParam The domain of the subject.
    // @param subjectParam The principal for this operation.
    // @param specParam Information about the new access control to be assigned.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the namespace is marked for deletion or the associated cluster is being disabled.
    // @throws NotFound if ``namespace`` cannot be located.
    // @throws Unsupported if the specified principal on given domain is not associated with the namespace.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Configure privilege.
    Set(namespaceParam string, domainParam string, subjectParam string, specParam SetSpec) error 


    // Get the information about the access control of the subject on given domain on the namespace. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @param domainParam The domain of the subject.
    // @param subjectParam The principal for this operation.
    // @return Information about the subject including the type and the role on the namespace.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if ``namespace`` cannot be located.
    // @throws Unsupported if the specified principal on given domain is not associated with the namespace.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    Get(namespaceParam string, domainParam string, subjectParam string) (Info, error) 

}
