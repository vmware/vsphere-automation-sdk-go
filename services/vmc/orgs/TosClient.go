/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Tos
 * Used by client-side stubs.
 */

package orgs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type TosClient interface {

    // Queries for the terms of service of a given org.
    //
    // @param orgParam Organization identifier (required)
    // @param termsIdParam The terms of service reference ID to check on. (required)
    // @return com.vmware.vmc.model.TermsOfServiceResult
    // @throws Unauthorized  Forbidden
	Get(orgParam string, termsIdParam string) (model.TermsOfServiceResult, error)
}
