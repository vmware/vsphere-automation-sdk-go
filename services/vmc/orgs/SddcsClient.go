/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Sddcs
 * Used by client-side stubs.
 */

package orgs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type SddcsClient interface {

    // Provision an SDDC in target cloud
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcConfigParam sddcConfig (required)
    // @param validateOnlyParam When true, only validates the given sddc configuration without provisioning. (optional)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Create(orgParam string, sddcConfigParam model.AwsSddcConfig, validateOnlyParam *bool) (model.Task, error)

    // Delete SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param retainConfigurationParam If = 'true', the SDDC's configuration is retained as a template for later use. This flag is applicable only to SDDCs in ACTIVE state. (optional)
    // @param templateNameParam Only applicable when retainConfiguration is also set to 'true'. When set, this value will be used as the name of the SDDC configuration template generated. (optional)
    // @param forceParam If = true, will delete forcefully. Beware: do not use the force flag if there is a chance an active provisioning or deleting task is running against this SDDC. This option is restricted. (optional)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for deletion
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the SDDC with given identifier
	Delete(orgParam string, sddcParam string, retainConfigurationParam *bool, templateNameParam *string, forceParam *bool) (model.Task, error)

    // Get SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @return com.vmware.vmc.model.Sddc
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the SDDC with given identifier
	Get(orgParam string, sddcParam string) (model.Sddc, error)

    // List all the SDDCs of an organization
    //
    // @param orgParam Organization identifier. (required)
    // @param includeDeletedParam When true, forces the result to also include deleted SDDCs. (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	List(orgParam string, includeDeletedParam *bool) ([]model.Sddc, error)

    // Patch SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param sddcPatchRequestParam Patch request for the SDDC (required)
    // @return com.vmware.vmc.model.Sddc
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  SDDC cannot be patched
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the SDDC with given identifier
	Patch(orgParam string, sddcParam string, sddcPatchRequestParam model.SddcPatchRequest) (model.Sddc, error)
}
