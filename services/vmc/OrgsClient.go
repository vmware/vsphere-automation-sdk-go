/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Orgs
 * Used by client-side stubs.
 */

package vmc

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type OrgsClient interface {

    // Get details of organization
    //
    // @param orgParam Organization identifier (required)
    // @return com.vmware.vmc.model.Organization
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Organization doesn't exist
	Get(orgParam string) (model.Organization, error)

    // Return a list of all organizations the calling user (based on credential) is authorized on.
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	List() ([]model.Organization, error)
}
