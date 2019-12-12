/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Publicips
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type PublicipsClient interface {

    // Allocate public IPs for a SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param specParam allocation spec (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the SDDC with given identifier
	Create(orgParam string, sddcParam string, specParam model.SddcAllocatePublicIpSpec) (model.Task, error)

    // Free one public IP for a SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param idParam ip allocation id (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the public IP with given IP address
	Delete(orgParam string, sddcParam string, idParam string) (model.Task, error)

    // Get one public IP for a SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param idParam ip allocation id (required)
    // @return com.vmware.vmc.model.SddcPublicIp
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the public IP with given IP address
	Get(orgParam string, sddcParam string, idParam string) (model.SddcPublicIp, error)

    // list all public IPs for a SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the SDDC with given identifier
	List(orgParam string, sddcParam string) ([]model.SddcPublicIp, error)

    // Attach or detach a public IP to workload VM for a SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param idParam ip allocation id (required)
    // @param actionParam Type of action as 'attach', 'detach', 'reattach', or 'rename'. For 'attch', the public IP must not be attached and 'associated_private_ip' in the payload needs to be set with a workload VM private IP. For 'detach', the public IP must be attached and 'associated_private_ip' in the payload should not be set with any value. For 'reattach', the public IP must be attached and 'associated_private_ip' in the payload needs to be set with a new workload VM private IP. For 'rename', the 'name' in the payload needs to have a new name string. (required)
    // @param sddcPublicIpObjectParam SddcPublicIp object to update (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the public IP with given IP address
	Update(orgParam string, sddcParam string, idParam string, actionParam string, sddcPublicIpObjectParam model.SddcPublicIp) (model.Task, error)
}
