/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Subscriptions
 * Used by client-side stubs.
 */

package library


// The ``Subscriptions`` interface provides methods for managing the subscription information of the subscribers of a published library. This interface was added in vSphere API 6.7.2.
type SubscriptionsClient interface {

    // Creates a subscription of the published library. This method was added in vSphere API 6.7.2.
    //
    // @param clientTokenParam A unique token generated on the client for each creation request. The token should be a universally unique identifier (UUID), for example: ``b8a2a2e3-2314-43cd-a871-6ede0f429751``. This token can be used to guarantee idempotent creation.
    // If not specified, creation is not idempotent.
    // @param libraryParam Identifier of the published library.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @param specParam Specification for the subscription.
    // @return Subscription identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.Subscriptions``.
    // @throws AlreadyExists  If a subscription of the published library to the specified subscribed library already exists. This is only applicable when ``subscribedLibrary#subscribedLibrary`` is specified.
    // @throws Error  If the system reports an error while responding to the request.
    // @throws NotFound  If the library specified by ``library`` does not exist.
    // @throws NotFound  If the subscribed library specified by ``subscribedLibrary#subscribedLibrary`` does not exist at the vCenter instance specified by ``subscribedLibrary#vcenter``.
    // @throws ResourceInaccessible  If the vCenter instance specified by ``subscribedLibrary#vcenter`` cannot be contacted or found.
    // @throws InvalidArgument  If SubscriptionsCreateSpec contains invalid arguments.
    // @throws InvalidElementType  If the library specified by ``library`` is a subscribed library.
    // @throws NotAllowedInCurrentState  If the library specified by ``library`` is not a published library.
    // @throws Unauthenticated  If the user that requested the method cannot be authenticated.
    // @throws Unauthorized  If the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library`` requires ``ContentLibrary.AddSubscription``.
	Create(clientTokenParam *string, libraryParam string, specParam SubscriptionsCreateSpec) (string, error)

    // Deletes the specified subscription of the published library. The subscribed library associated with the subscription will not be deleted. This method was added in vSphere API 6.7.2.
    //
    // @param libraryParam Identifier of the published library.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @param subscriptionParam Subscription identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Subscriptions``.
    // @throws Error  If the system reports an error while responding to the request.
    // @throws InvalidElementType  If the library specified by ``library`` is a subscribed library.
    // @throws NotAllowedInCurrentState  If the library specified by ``library`` is not a published library.
    // @throws NotFound  If the library specified by ``library`` does not exist.
    // @throws NotFound  If the subscription specified by ``subscription`` does not exist for the library specified by ``library``.
    // @throws Unauthenticated  If the user that requested the method cannot be authenticated.
    // @throws Unauthorized  If the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library`` requires ``ContentLibrary.DeleteSubscription``.
	Delete(libraryParam string, subscriptionParam string) error

    // Lists the subscriptions of the published library. This method was added in vSphere API 6.7.2.
    //
    // @param libraryParam Identifier of the published library.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @return List of commonly used information about subscriptions of the published library.
    // @throws Error  If the system reports an error while responding to the request.
    // @throws InvalidElementType  If the library specified by ``library`` is a subscribed library.
    // @throws NotFound  If the library specified by ``library`` does not exist.
    // @throws NotAllowedInCurrentState  If the library specified by ``library`` is not a published library.
    // @throws Unauthenticated  If the user that requested the method cannot be authenticated.
    // @throws Unauthorized  If the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library`` requires ``System.Read``.
	List(libraryParam string) ([]SubscriptionsSummary, error)

    // Updates the specified subscription of the published library. 
    //
    //  This is an incremental update to the subscription. Except for the SubscriptionsUpdateSpecPlacement class, properties that are null in the update specification will be left unchanged. If ``spec#subscribedLibraryPlacement`` is specified, all properties of the current subscribed library placement will be replaced by this placement.. This method was added in vSphere API 6.7.2.
    //
    // @param libraryParam Identifier of the published library.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @param subscriptionParam subscription identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Subscriptions``.
    // @param specParam Specification of the new property values to set on the subscription.
    // @throws Error  If the system reports an error while responding to the request.
    // @throws NotFound  If the library specified by ``library`` does not exist.
    // @throws NotFound  If the subscription specified by ``subscription`` does not exist for the library specified by ``library``.
    // @throws ResourceInaccessible  If the subscribed library cannot be contacted or found.
    // @throws InvalidArgument  If SubscriptionsUpdateSpec contains invalid arguments.
    // @throws InvalidElementType  If the library specified by ``library`` is a subscribed library.
    // @throws NotAllowedInCurrentState  If the library specified by ``library`` is not a published library.
    // @throws Unauthenticated  If the user that requested the method cannot be authenticated.
    // @throws Unauthorized  If the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library`` requires ``ContentLibrary.UpdateSubscription``.
	Update(libraryParam string, subscriptionParam string, specParam SubscriptionsUpdateSpec) error

    // Returns information about the specified subscription of the published library. This method was added in vSphere API 6.7.2.
    //
    // @param libraryParam Identifier of the published library.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @param subscriptionParam Identifier of the subscription.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Subscriptions``.
    // @return Information about the subscription.
    // @throws Error  If the system reports an error while responding to the request.
    // @throws NotFound  If the library specified by ``library`` does not exist.
    // @throws InvalidArgument  If the ``subscription`` is not valid.
    // @throws InvalidElementType  If the library specified by ``library`` is a subscribed library.
    // @throws NotAllowedInCurrentState  If the library specified by ``library`` is not a published library.
    // @throws Unauthenticated  If the user that requested the method cannot be authenticated.
    // @throws Unauthorized  If the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library`` requires ``System.Read``.
	Get(libraryParam string, subscriptionParam string) (SubscriptionsInfo, error)
}
