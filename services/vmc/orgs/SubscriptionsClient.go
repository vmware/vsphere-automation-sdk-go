/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Subscriptions
 * Used by client-side stubs.
 */

package orgs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type SubscriptionsClient interface {

    // Initiates the creation of a subscription
    //
    // @param orgParam Organization identifier (required)
    // @param subscriptionRequestParam subscriptionRequest (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws InternalServerError  Server error. Check retryable flag to see if request should be retried.
	Create(orgParam string, subscriptionRequestParam model.SubscriptionRequest) (model.Task, error)

    // Get subscription details for a given subscription id
    //
    // @param orgParam Organization identifier (required)
    // @param subscriptionParam SubscriptionId for an sddc. (required)
    // @return com.vmware.vmc.model.SubscriptionDetails
    // @throws InternalServerError  Server error. Check retryable flag to see if request should be retried.
    // @throws NotFound  Not Found
	Get(orgParam string, subscriptionParam string) (model.SubscriptionDetails, error)

    // Returns all subscriptions for a given org id
    //
    // @param orgParam Organization identifier (required)
    // @param offerTypeParam Offer Type \* \\\\`ON_DEMAND\\\\` - on-demand subscription \* \\\\`TERM\\\\` - term subscription \* All subscriptions if not specified (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws InternalServerError  Server error. Check retryable flag to see if request should be retried.
    // @throws NotFound  Not Found
	Get0(orgParam string, offerTypeParam *string) ([]model.SubscriptionDetails, error)
}
