/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: HostProvisionSpec
 * Used by client-side stubs.
 */

package clusters

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
)

type HostProvisionSpecClient interface {

    // Get host provision spec for a new host in a cluster in an sddc
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param clusterParam cluster identifier (required)
    // @param actionParam If = 'scale', Will get provisioning spec for add host operations for scaling the cluster/sddc. If = 'remediate', Will get provisioning spec for add host operations for all internal operations (remediation, maintenance etc). Default behaviour is 'scale' (optional)
    // @return com.vmware.vmc.model.HostProvisionSpec
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws InternalServerError  Internal server error.
	Get(orgParam string, sddcParam string, clusterParam string, actionParam *string) (model.HostProvisionSpec, error)
}
