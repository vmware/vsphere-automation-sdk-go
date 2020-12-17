/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CloudServiceCommon
 * Used by client-side stubs.
 */

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

type CloudServiceCommonClient interface {

    // This API is used to get external configuration for north-south traffic. For eg., in AWS case, this would refer to Direct Connect config. For Dimension, this would refer to the config at Upstream Intranet interface to Tor.
    // @return com.vmware.model.ExternalConnectivityConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetExternalConnectivityConfig() (model.ExternalConnectivityConfig, error)

    // Get management VM access information.
    //
    // @param vmIdParam (required)
    // @return com.vmware.model.MgmtVmInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetMgmtVmInfo(vmIdParam string) (model.MgmtVmInfo, error)

    // Get the consolidated realized entities given an intent path, specified in the query parameter. The intent object under the intent path is indicated by a specific VMC-App API and can contain multiple objects.
    //
    // @param intentPathParam String Path of the intent object (required)
    // @param sitePathParam String Path of the site (optional)
    // @return com.vmware.model.VmcRealizedEntities
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetRealizedEntities(intentPathParam string, sitePathParam *string) (model.VmcRealizedEntities, error)

    // Get the consolidated status of an intent object, specified by path in query parameter. The intent object is indicated by a specific VMC-App API and can contain multiple objects. For example, /infra/direct-connect/bgp can return the consolidated status of ASN update and route preference update.
    //
    // @param intentPathParam String Path of the intent object (required)
    // @param sitePathParam String Path of the site (optional)
    // @return com.vmware.model.VmcConsolidatedRealizedStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetRealizedStateStatus(intentPathParam string, sitePathParam *string) (model.VmcConsolidatedRealizedStatus, error)

    // Get vmc environment feature flags
    //
    // @param internalNameParam internal feature name (optional)
    // @param nameParam feature name (optional)
    // @return com.vmware.model.VmcFeatureFlags
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetVmcFeatureFlags(internalNameParam *string, nameParam *string) (model.VmcFeatureFlags, error)

    // List Management VM information.
    // @return com.vmware.model.MgmtVmsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListMgmtVms() (model.MgmtVmsListResult, error)

    // This API is used to update intranet configuration for north-south traffic. For eg., in AWS case, this would refer to Direct Connect config. For Dimension, this would refer to the config at Upstream Intranet interface to Tor. Only the intranet MTU can be updated, internet mtu and connected VPC mtu are read-only.
    //
    // @param externalConnectivityConfigParam (required)
    // @return com.vmware.model.ExternalConnectivityConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	UpdateIntranetUplinkMtu(externalConnectivityConfigParam model.ExternalConnectivityConfig) (model.ExternalConnectivityConfig, error)
}
