/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: StandaloneHostSwitchSetting
 * Used by client-side stubs.
 */

package idfw

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type StandaloneHostSwitchSettingClient interface {

    // Read identity firewall configuration for standalone host
    // @return com.vmware.nsx_policy.model.StandaloneHostIdfwConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.StandaloneHostIdfwConfiguration, error)

    // Patch identity firewall configuration for standalone host
    //
    // @param standaloneHostIdfwConfigurationParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(standaloneHostIdfwConfigurationParam model.StandaloneHostIdfwConfiguration) error

    // Update the idfw configuration for standalone host
    //
    // @param standaloneHostIdfwConfigurationParam (required)
    // @return com.vmware.nsx_policy.model.StandaloneHostIdfwConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(standaloneHostIdfwConfigurationParam model.StandaloneHostIdfwConfiguration) (model.StandaloneHostIdfwConfiguration, error)
}
