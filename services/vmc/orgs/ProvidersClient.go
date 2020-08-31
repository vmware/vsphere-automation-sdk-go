/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Providers
 * Used by client-side stubs.
 */

package orgs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ProvidersClient interface {

    // @param orgParam Organization identifier (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Organization doesn't exist
	List(orgParam string) ([]model.AwsCloudProvider, error)
}
