/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ClusterConstraints
 * Used by client-side stubs.
 */

package clusterConstraints

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type ClusterConstraintsClient interface {


    // Get constraints on cluster storage size for EBS-backed clusters.
    //
    // @param orgParam Organization identifier. (required)
    // @param providerParam Cloud storage provider ID (example AWS) (required)
    // @param numHostsParam Number of hosts in cluster (required)
    // @param oneNodeReducedCapacityParam Whether this sddc is reduced capacity 1NODE. (optional)
    // @return com.vmware.vmc.model.VsanConfigConstraints
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid or missing parameters
    // @throws Unauthorized  Forbidden
    Get(orgParam string, providerParam string, numHostsParam int64, oneNodeReducedCapacityParam *bool) (model.VsanConfigConstraints, error) 

}
