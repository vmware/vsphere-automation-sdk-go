/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SddcConnections
 * Used by client-side stubs.
 */

package account_link

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type SddcConnectionsClient interface {

    // Get a list of SDDC connections currently setup for the customer's organization.
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam sddc (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Get(orgParam string, sddcParam *string) ([]model.AwsSddcConnection, error)
}
