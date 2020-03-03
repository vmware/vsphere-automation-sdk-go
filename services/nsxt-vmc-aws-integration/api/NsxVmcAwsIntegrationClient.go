/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: NsxVmcAwsIntegration
 * Used by client-side stubs.
 */

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

type NsxVmcAwsIntegrationClient interface {

    // Perform the attach operation in VMC provider to attach the specified VIF (Virtual Interface) to the SDDC.
    //
    // @param vifIdParam (required)
    // @param actionParam Direct Connect VIF (Virtual Interface) action (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	AttachVif(vifIdParam string, actionParam string) error

    // This API is used to create or update a public IP. In creating, the API allocates a new public IP from VMC provider. In updating, only the display name can be modified, the IP is read-only.
    //
    // @param publicIpIdParam (required)
    // @param publicIpParam (required)
    // @return com.vmware.model.PublicIp
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	CreatePublicIp(publicIpIdParam string, publicIpParam model.PublicIp) (model.PublicIp, error)

    // Delete a public IP. The IP will be released in VMC provider.
    //
    // @param publicIpIdParam (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	DeletePublicIp(publicIpIdParam string, forceParam *bool) error

    // Detach and delete a Direct Connect VIF (Virtual Interface) from the SDDC. Detach and delete are one operation in VMC provider, a Direct Connect VIF will be deleted after detached.
    //
    // @param vifIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	DeleteVifById(vifIdParam string) error

    // Get Direct Connect BGP related information, including current Autonomous System Number of the VGW attached to the VPC
    // @return com.vmware.model.DirectConnectBgpInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetDxBgpInfo() (model.DirectConnectBgpInfo, error)

    // Get linked VPC information.
    //
    // @param linkedVpcIdParam (required)
    // @return com.vmware.model.LinkedVpcInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetLinkedVpc(linkedVpcIdParam string) (model.LinkedVpcInfo, error)

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

    // Get the public IP information.
    //
    // @param publicIpIdParam (required)
    // @return com.vmware.model.PublicIp
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetPublicIp(publicIpIdParam string) (model.PublicIp, error)

    // Get the consolidated status of an intent object, specified by path in query parameter. The intent object is indicated by a specific VMC-App API and can contain multiple objects. For example, /infra/direct-connect/bgp can return the consolidated status of ASN update and route preference update.
    //
    // @param intentPathParam String Path of the intent object (required)
    // @return com.vmware.model.VmcConsolidatedRealizedStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetRealizedStateStatus(intentPathParam string) (model.VmcConsolidatedRealizedStatus, error)

    // Get the user-level SDDC configuration parameters
    // @return com.vmware.model.SddcUserConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetSddcUserConfig() (model.SddcUserConfiguration, error)

    // Retrieve the shadow account and linked VPC account information from VMC provider. This API is a live query to VMC provider.
    // @return com.vmware.model.VMCAccounts
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListAccounts() (model.VMCAccounts, error)

    // Retrieve BGP routes that are advertised by Direct Connect from VMC provider to on-premise datacenter. This API is a live query to VMC provider.
    // @return com.vmware.model.BGPAdvertisedRoutes
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListAdvertisedRoutes() (model.BGPAdvertisedRoutes, error)

    // List services connected to this linked vpc, for example, S3. The response consist of all available services along with their status.
    //
    // @param linkedVpcIdParam linked vpc id (required)
    // @return com.vmware.model.ConnectedServiceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListConnectedServices(linkedVpcIdParam string) (model.ConnectedServiceListResult, error)

    // Retrieve BGP routes that are learned by Direct Connect from on-premise datacenter. This API is a live query to VMC provider.
    // @return com.vmware.model.BGPLearnedRoutes
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListLearnedRoutes() (model.BGPLearnedRoutes, error)

    // List linked VPC information.
    // @return com.vmware.model.LinkedVpcsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListLinkedVpcs() (model.LinkedVpcsListResult, error)

    // List Management VM information.
    // @return com.vmware.model.MgmtVmsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListMgmtVms() (model.MgmtVmsListResult, error)

    // List all public IPs obtained in the SDDC.
    // @return com.vmware.model.PublicIpsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListPublicIps() (model.PublicIpsListResult, error)

    // Return all non-connected VIFs (with states \"avalible\", \"down\", \"pending\" and \"confirming\") and connected VIFs that are available to the SDDC.
    // @return com.vmware.model.VifsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListVifs() (model.VifsListResult, error)

    // Connect/Disconnect the service to the given linked vpc. For example, connect S3. The user will know what services are available through the GET call. If the user is trying to connect/disconnect an unknown service, the POST call will throw a 400 Bad Request error.
    //
    // @param linkedVpcIdParam linked vpc id (required)
    // @param serviceNameParam connected service name, e.g. s3 (required)
    // @param connectedServiceStatusParam (required)
    // @return com.vmware.model.ConnectedServiceStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	UpdateConnectedService(linkedVpcIdParam string, serviceNameParam string, connectedServiceStatusParam model.ConnectedServiceStatus) (model.ConnectedServiceStatus, error)

    // This API is used to create or update DX BGP related information. For ASN update, VIFs should be disconnected from the DX VGW before making this API call. The ASN update operation will be synchronous at this point. In the future the user should make use of the Get RealizationStatus call to check update status. While an ASN update call is in progress, any other DX BGP update request will be rejected.
    //
    // @param directConnectBgpInfoParam (required)
    // @return com.vmware.model.DirectConnectBgpInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	UpdateDxBgpInfo(directConnectBgpInfoParam model.DirectConnectBgpInfo) (model.DirectConnectBgpInfo, error)
}
