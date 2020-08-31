/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SddcTemplate
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type SddcTemplateClient interface {

    // Get configuration template for an SDDC
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.vmc.model.SddcTemplate
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  SDDC is in a state that cannot be use for generating configuration template
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the SDDC with given identifier
	Get(orgParam string, sddcParam string) (model.SddcTemplate, error)
}
