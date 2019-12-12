/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SddcTemplates
 * Used by client-side stubs.
 */

package orgs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type SddcTemplatesClient interface {

    // Delete SDDC template identified by given id.
    //
    // @param orgParam Organization identifier. (required)
    // @param templateIdParam SDDC Template identifier (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Delete(orgParam string, templateIdParam string) (model.Task, error)

    // Get configuration template by given template id.
    //
    // @param orgParam Organization identifier. (required)
    // @param templateIdParam SDDC Template identifier (required)
    // @return com.vmware.vmc.model.SddcTemplate
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the SDDC Template with given identifier
	Get(orgParam string, templateIdParam string) (model.SddcTemplate, error)

    // List all available SDDC configuration templates in an organization
    //
    // @param orgParam Organization identifier. (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	List(orgParam string) ([]model.SddcTemplate, error)
}
