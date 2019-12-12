/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Esxs
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type EsxsClient interface {

    // Add/Remove one or more ESX hosts in the target cloud
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param esxConfigParam esxConfig (required)
    // @param actionParam If = 'add', will add the esx. If = 'remove', will delete the esx/esxs bound to a single cluster (Cluster Id is mandatory for non cluster 1 esx remove). If = 'force-remove', will delete the esx even if it can lead to data loss (This is an privileged operation). If = 'addToAll', will add esxs to all clusters in the SDDC (This is an privileged operation). If = 'removeFromAll', will delete the esxs from all clusters in the SDDC (This is an privileged operation). If = 'attach-diskgroup', will attach the provided diskgroups to a given host (privileged). If = 'detach-diskgroup', will detach the diskgroups of a given host (privileged). Default behaviour is 'add' (optional)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the SDDC with the given identifier
	Create(orgParam string, sddcParam string, esxConfigParam model.EsxConfig, actionParam *string) (model.Task, error)
}
