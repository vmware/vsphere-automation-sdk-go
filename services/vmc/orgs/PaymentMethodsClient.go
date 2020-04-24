/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PaymentMethods
 * Used by client-side stubs.
 */

package orgs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type PaymentMethodsClient interface {

    // Get payment methods of organization
    //
    // @param orgParam Organization identifier. (required)
    // @param defaultFlagParam When true, will only return default payment methods. (optional, default to false)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Organization doesn't exist
	GetOrgPaymentMethods(orgParam string, defaultFlagParam *bool) ([]model.PaymentMethodInfo, error)
}
