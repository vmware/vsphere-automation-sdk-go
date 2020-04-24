/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Config
 * Used by client-side stubs.
 */

package backups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type ConfigClient interface {

    // Get a configuration of a file server and timers for automated backup. Fields that contain secrets (password, passphrase) are not returned.
    // @return com.vmware.nsx_global_policy.model.BackupConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.BackupConfiguration, error)

    // Configure file server and timers for automated backup. If secret fields are omitted (password, passphrase) then use the previously set value.
    //
    // @param backupConfigurationParam (required)
    // @param frameTypeParam Frame type (optional, default to LOCAL_LOCAL_MANAGER)
    // @param siteIdParam Site ID (optional, default to localhost)
    // @return com.vmware.nsx_global_policy.model.BackupConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(backupConfigurationParam model.BackupConfiguration, frameTypeParam *string, siteIdParam *string) (model.BackupConfiguration, error)
}
