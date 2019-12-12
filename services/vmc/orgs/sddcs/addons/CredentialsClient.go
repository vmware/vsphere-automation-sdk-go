/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Credentials
 * Used by client-side stubs.
 */

package addons

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type CredentialsClient interface {

    // Associated a new add on credentials with the SDDC such as HCX
    //
    // @param orgParam Org id of the associated SDDC (required)
    // @param sddcIdParam Id of the SDDC (required)
    // @param addonTypeParam Add on type (required)
    // @param credentialsParam Credentials creation payload (required)
    // @return com.vmware.vmc.model.NewCredentials
    // @throws ConcurrentChange  Credentials with same name exists with in the scope of addOnType
    // @throws InvalidRequest  Invalid input
    // @throws Unauthorized  Forbidden
	Create(orgParam string, sddcIdParam string, addonTypeParam string, credentialsParam model.NewCredentials) (model.NewCredentials, error)

    // Get credential details by name
    //
    // @param orgParam Org id of the associated SDDC (required)
    // @param sddcIdParam Id of the SDDC (required)
    // @param addonTypeParam Add on type (required)
    // @param nameParam name of the credentials (required)
    // @return com.vmware.vmc.model.NewCredentials
    // @throws Unauthorized  Forbidden
	Get(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string) (model.NewCredentials, error)

    // List all the credentials assoicated with an addon type with in a SDDC
    //
    // @param orgParam Org id of the associated SDDC (required)
    // @param sddcIdParam Id of the SDDC (required)
    // @param addonTypeParam Add on type (required)
    // @throws Unauthorized  Forbidden
	List(orgParam string, sddcIdParam string, addonTypeParam string) ([]model.NewCredentials, error)

    // Update credential details
    //
    // @param orgParam Org id of the associated SDDC (required)
    // @param sddcIdParam Id of the SDDC (required)
    // @param addonTypeParam Add on type (required)
    // @param nameParam name of the credentials (required)
    // @param credentialsParam Credentials update payload (required)
    // @return com.vmware.vmc.model.NewCredentials
    // @throws InvalidRequest  Bad request
    // @throws Unauthorized  Forbidden
	Update(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string, credentialsParam model.UpdateCredentials) (model.NewCredentials, error)
}
