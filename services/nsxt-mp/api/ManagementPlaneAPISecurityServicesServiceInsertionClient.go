// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPISecurityServicesServiceInsertion
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type ManagementPlaneAPISecurityServicesServiceInsertionClient interface {

	// Adds a new instance endpoint. It belongs to one service instance and is attached to one service attachment. It represents a redirection target for a Rule.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param instanceEndpointParam (required)
	// @return com.vmware.model.InstanceEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddInstanceEndpoint(serviceIdParam string, serviceInstanceIdParam string, instanceEndpointParam model.InstanceEndpoint) (model.InstanceEndpoint, error)

	// Adds a new service profile.
	//
	// @param serviceIdParam (required)
	// @param baseServiceProfileParam (required)
	// The parameter must contain all the properties defined in model.BaseServiceProfile.
	// @return com.vmware.model.BaseServiceProfile
	// The return value will contain all the properties defined in model.BaseServiceProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddSIServiceProfile(serviceIdParam string, baseServiceProfileParam *data.StructValue) (*data.StructValue, error)

	// Adds a new Service attachment. A service attachment represents a point on NSX entity (Example: Logical Router) to which service instance can be connected through an InstanceEndpoint.
	//
	// @param serviceAttachmentParam (required)
	// @return com.vmware.model.ServiceAttachment
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceAttachment(serviceAttachmentParam model.ServiceAttachment) (model.ServiceAttachment, error)

	// Adds a new service chain. Service Chains is can contain profile belonging to same or different Service(s). It represents a redirection target for a Rule.
	//
	// @param serviceChainParam (required)
	// @return com.vmware.model.ServiceChain
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceChain(serviceChainParam model.ServiceChain) (model.ServiceChain, error)

	// Note- POST serviceinsertion excludelist API is deprecated. Please use the policy serviceinsertion excludelist API instead.
	//
	// @param resourceReferenceParam (required)
	// @return com.vmware.model.ResourceReference
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceInsertionExcludeListMemberAddMember(resourceReferenceParam model.ResourceReference) (model.ResourceReference, error)

	// Adds a new serviceinsertion rule in existing serviceinsertion section. Note- POST service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionRuleParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.model.ServiceInsertionRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceInsertionRuleInSection(sectionIdParam string, serviceInsertionRuleParam model.ServiceInsertionRule, idParam *string, operationParam *string) (model.ServiceInsertionRule, error)

	// Create multiple serviceinsertion rules in existing serviceinsertion section bounded by limit of 1000 serviceinsertion rules per section. Note- POST service insertion rules API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.model.ServiceInsertionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceInsertionRulesInSectionCreateMultiple(sectionIdParam string, serviceInsertionRuleListParam model.ServiceInsertionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionRuleList, error)

	// Creates new empty Service Insertion section in the system. Note- POST service insertion section API is deprecated. Please use the policy redirection-policy API.
	//
	// @param serviceInsertionSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.model.ServiceInsertionSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceInsertionSection(serviceInsertionSectionParam model.ServiceInsertionSection, idParam *string, operationParam *string) (model.ServiceInsertionSection, error)

	// Creates a new serviceinsertion section with rules. The limit on the number of rules is defined by maxItems in collection types for ServiceInsertionRule (ServiceInsertionRuleXXXList types). When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules are not supported. Instead, to create sections, use: POST /api/v1/serviceinsertion/sections To create rules, use: POST /api/v1/serviceinsertion/sections/<section-id>/rules Note- POST service insertion section creation with rules API is deprecated. Please use policy redirection-policy API.
	//
	// @param serviceInsertionSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.model.ServiceInsertionSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceInsertionSectionWithRulesCreateWithRules(serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionSectionRuleList, error)

	// Creates new Service-Insertion Service in the system.
	//
	// @param serviceDefinitionParam (required)
	// @return com.vmware.model.ServiceDefinition
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceInsertionService(serviceDefinitionParam model.ServiceDefinition) (model.ServiceDefinition, error)

	// Adds a new Service-Instance under the specified Service.
	//
	// @param serviceIdParam (required)
	// @param baseServiceInstanceParam (required)
	// The parameter must contain all the properties defined in model.BaseServiceInstance.
	// @return com.vmware.model.BaseServiceInstance
	// The return value will contain all the properties defined in model.BaseServiceInstance.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddServiceInstance(serviceIdParam string, baseServiceInstanceParam *data.StructValue) (*data.StructValue, error)

	// Adds a new vendor template. Vendor templates are service level objects, registered to be used in Service Profiles.
	//
	// @param serviceIdParam (required)
	// @param vendorTemplateParam (required)
	// @return com.vmware.model.VendorTemplate
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddVendorTemplate(serviceIdParam string, vendorTemplateParam model.VendorTemplate) (model.VendorTemplate, error)

	// Adds a solution config. Solution Config are service level objects, required for configuring the NXGI partner Service after deployment.
	//
	// @param serviceIdParam (required)
	// @param solutionConfigParam (required)
	// @return com.vmware.model.SolutionConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateSolutionConfig(serviceIdParam string, solutionConfigParam model.SolutionConfig) (model.SolutionConfig, error)

	// Delete instance endpoint information for a given instace endpoint. Please make sure to delete all the Service Insertion Rules, which refer to this Endpoint as 'redirect_tos' target.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param instanceEndpointIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteInstanceEndpoint(serviceIdParam string, serviceInstanceIdParam string, instanceEndpointIdParam string) error

	// Delete service profile for a given service.
	//
	// @param serviceIdParam (required)
	// @param serviceProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteSIServiceProfile(serviceIdParam string, serviceProfileIdParam string) error

	// Delete existing service attachment from system. Before deletion, please make sure that, no instance endpoints are connected to this attachment. In turn no appliance should be connected to this attachment.
	//
	// @param serviceAttachmentIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServiceAttachment(serviceAttachmentIdParam string) error

	// Delete a particular service chain.
	//
	// @param serviceChainIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServiceChain(serviceChainIdParam string) error

	// Remove the service deployment. Will remove all the Service VMs that were created as part of this deployment. User can send optional force delete option which will force remove the deployment, but should be used only when the regular delete is not working. Regular delete will ensure proper cleanup of Service VMs and related objects. Directly calling this API without trying regular undeploy will result in unexpected results, and orphan objects.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServiceDeployment(serviceIdParam string, serviceDeploymentIdParam string, forceParam *bool) error

	// Delete existing serviceinsertion rule in a serviceinsertion section. Note- DELETE service insertion rule API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServiceInsertionRule(sectionIdParam string, ruleIdParam string) error

	// Removes serviceinsertion section from the system. ServiceInsertion section with rules can only be deleted by passing \"cascade=true\" parameter. Note- DELETE service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param cascadeParam Flag to cascade delete of this object to all it's child objects. (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServiceInsertionSection(sectionIdParam string, cascadeParam *bool) error

	// Removes Service-Insertion Service from the system. A Service with Service-Instances can only be deleted by passing \"cascade=true\" parameter.
	//
	// @param serviceIdParam (required)
	// @param cascadeParam Flag to cascade delete all the child objects, associated with it. (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServiceInsertionService(serviceIdParam string, cascadeParam *bool) error

	// Delete existing Service-Instance for a given Service-Insertion Service.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServiceInstance(serviceIdParam string, serviceInstanceIdParam string) error

	// Delete service-manager which is registered with NSX with basic details like name, username, password.
	//
	// @param serviceManagerIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServiceManager(serviceManagerIdParam string) error

	// Undeploy one service VM as standalone or two service VMs as HA. Associated deployment information and instance runtime will also be deleted once service VMs have been un-deployed successfully.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteServicevMsDelete(serviceIdParam string, serviceInstanceIdParam string) error

	// Deletes solution config information for a given service.
	//
	// @param serviceIdParam (required)
	// @param solutionConfigIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteSolutionConfig(serviceIdParam string, solutionConfigIdParam string) error

	// Delete vendor template information for a given service. Please make sure to delete all the Service Profile(s), which refer to this vendor tempalte before deleting the template itself.
	//
	// @param serviceIdParam (required)
	// @param vendorTemplateIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteVendorTemplate(serviceIdParam string, vendorTemplateIdParam string) error

	// This will deploy a particular service on a given cluster / host. Internally multiple service instance can be created during the deployment. If there are no issues in the parameters, the call returns immediately, and the service VMs will be deployed asynchronously. To get the overall status of the deployment or to get the status of individual service vm, please use the deployment status APIs.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentParam (required)
	// @return com.vmware.model.ServiceDeployment
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeployService(serviceIdParam string, serviceDeploymentParam model.ServiceDeployment) (model.ServiceDeployment, error)

	// Deploys one service VM as standalone, or two service VMs as HA where one VM is active and another one is standby. During the deployment of service VMs, service will be set up based on deployment events using callbacks.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeployServicevMsDeploy(serviceIdParam string, serviceInstanceIdParam string) error

	// Returns detailed Endpoint information for a given InstanceEndpoint.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param instanceEndpointIdParam (required)
	// @return com.vmware.model.InstanceEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetInstanceEndpoint(serviceIdParam string, serviceInstanceIdParam string, instanceEndpointIdParam string) (model.InstanceEndpoint, error)

	// Returns operational status of a specified interface
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param instanceRuntimeIdParam (required)
	// @param interfaceIndexParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.RuntimeInterfaceOperationalStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetRuntimeInterfaceOperationalStatus(serviceIdParam string, serviceInstanceIdParam string, instanceRuntimeIdParam string, interfaceIndexParam string, sourceParam *string) (model.RuntimeInterfaceOperationalStatus, error)

	// Returns statistics of a specified interface via associated logical port. If the logical port is attached to a logical router port, query parameter \"source=realtime\" is not supported.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param instanceRuntimeIdParam (required)
	// @param interfaceIndexParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.RuntimeInterfaceStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetRuntimeInterfaceStatistics(serviceIdParam string, serviceInstanceIdParam string, instanceRuntimeIdParam string, interfaceIndexParam string, sourceParam *string) (model.RuntimeInterfaceStatistics, error)

	// Returns detailed service profile information for a given Service.
	//
	// @param serviceIdParam (required)
	// @param serviceProfileIdParam (required)
	// @return com.vmware.model.BaseServiceProfile
	// The return value will contain all the properties defined in model.BaseServiceProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSIServiceProfile(serviceIdParam string, serviceProfileIdParam string) (*data.StructValue, error)

	// Returns detailed Attachment information for a given service attachment.
	//
	// @param serviceAttachmentIdParam (required)
	// @return com.vmware.model.ServiceAttachment
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceAttachment(serviceAttachmentIdParam string) (model.ServiceAttachment, error)

	// Returns detailed service chain information.
	//
	// @param serviceChainIdParam (required)
	// @return com.vmware.model.ServiceChain
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceChain(serviceChainIdParam string) (model.ServiceChain, error)

	// Returns detail of service deployment.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @return com.vmware.model.ServiceDeployment
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceDeployment(serviceIdParam string, serviceDeploymentIdParam string) (model.ServiceDeployment, error)

	// Returns configuration state of deployed partner service using service insertion framework.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @return com.vmware.model.ConfigurationState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceDeploymentState(serviceIdParam string, serviceDeploymentIdParam string) (model.ConfigurationState, error)

	// Returns current status of the deployment of partner service. Available only for EPP Services. By default this API would return cached status. Caching happens every 3 minutes. For realtime status, query parameter \"source=realtime\" needs to be passed.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.ServiceDeploymentStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceDeploymentStatus(serviceIdParam string, serviceDeploymentIdParam string, sourceParam *string) (model.ServiceDeploymentStatus, error)

	// Returns the list of deployments for the given service
	//
	// @param serviceIdParam (required)
	// @return com.vmware.model.ServiceDeploymentListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceDeployments(serviceIdParam string) (model.ServiceDeploymentListResult, error)

	// Note- GET serviceinsertion excludelist API is deprecated. Please use the policy serviceinsertion excludelist API instead.
	// @return com.vmware.model.SIExcludeList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInsertionExcludeList() (model.SIExcludeList, error)

	// Return existing serviceinsertion rule information in a serviceinsertion section. Note- GET service insertion rule API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @return com.vmware.model.ServiceInsertionRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInsertionRule(sectionIdParam string, ruleIdParam string) (model.ServiceInsertionRule, error)

	// Return all serviceinsertion rule(s) information for a given serviceinsertion section. Note- GET service insertion rules API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param appliedTosParam AppliedTo's referenced by this section or section's Distributed Service Rules . (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param destinationsParam Destinations referenced by this section's Distributed Service Rules . (optional)
	// @param filterTypeParam Filter type (optional, default to FILTER)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param servicesParam NSService referenced by this section's Distributed Service Rules . (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourcesParam Sources referenced by this section's Distributed Service Rules . (optional)
	// @return com.vmware.model.ServiceInsertionRuleListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInsertionRules(sectionIdParam string, appliedTosParam *string, cursorParam *string, destinationsParam *string, filterTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string) (model.ServiceInsertionRuleListResult, error)

	// Returns information about serviceinsertion section for the identifier. Note- GET service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.model.ServiceInsertionSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInsertionSection(sectionIdParam string) (model.ServiceInsertionSection, error)

	// Returns serviceinsertion section information with rules for a section identifier. When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules are not supported. Instead, to read serviceinsertion rules, use: GET /api/v1/serviceinsertion/sections/<section-id>/rules with the appropriate page_size. Note- GET service insertion section with rules API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.model.ServiceInsertionSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInsertionSectionWithRulesListWithRules(sectionIdParam string) (model.ServiceInsertionSectionRuleList, error)

	// Returns information about Service-Insertion Service with the given identifier.
	//
	// @param serviceIdParam (required)
	// @return com.vmware.model.ServiceDefinition
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInsertionService(serviceIdParam string) (model.ServiceDefinition, error)

	// Get ServiceInsertion global status for a context
	//
	// @param contextTypeParam (required)
	// @return com.vmware.model.ServiceInsertionStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInsertionStatus(contextTypeParam string) (model.ServiceInsertionStatus, error)

	// Returns Service-Instance information for a given Service-Insertion Service.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @return com.vmware.model.BaseServiceInstance
	// The return value will contain all the properties defined in model.BaseServiceInstance.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInstance(serviceIdParam string, serviceInstanceIdParam string) (*data.StructValue, error)

	// Returns list of NSGroups used in Service Insertion North-South rules for a given Service Instance.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @return com.vmware.model.ServiceInstanceNSGroups
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInstanceNSGroups(serviceIdParam string, serviceInstanceIdParam string) (model.ServiceInstanceNSGroups, error)

	// Returns configuration state of one instance of a deployed partner service using service insertion framework.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @return com.vmware.model.ConfigurationState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInstanceState(serviceIdParam string, serviceInstanceIdParam string) (model.ConfigurationState, error)

	// Returns status of one instance of a deployed partner service using service insertion framework. By default this API would return cached status. Caching happens every 3 minutes. For realtime status, query parameter \"source=realtime\" needs to be passed.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.ServiceInstanceStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceInstanceStatus(serviceIdParam string, serviceInstanceIdParam string, sourceParam *string) (model.ServiceInstanceStatus, error)

	// Retrieve service-manager details like name, username, password, vendor ID, thumbprint for a given ID.
	//
	// @param serviceManagerIdParam (required)
	// @return com.vmware.model.ServiceManager
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceManager(serviceManagerIdParam string) (model.ServiceManager, error)

	// Returns list of NSGroups used in Service Insertion rules for a given Service Profile.
	//
	// @param serviceIdParam (required)
	// @param serviceProfileIdParam (required)
	// @return com.vmware.model.ServiceProfileNSGroups
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetServiceProfileNSGroups(serviceIdParam string, serviceProfileIdParam string) (model.ServiceProfileNSGroups, error)

	// Returns Solution Config information for a given solution config id.
	//
	// @param serviceIdParam (required)
	// @param solutionConfigIdParam (required)
	// @return com.vmware.model.SolutionConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSolutionConfig(serviceIdParam string, solutionConfigIdParam string) (model.SolutionConfig, error)

	// Returns detailed vendor template information for a given service.
	//
	// @param serviceIdParam (required)
	// @param vendorTemplateIdParam (required)
	// @return com.vmware.model.VendorTemplate
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetVendorTemplate(serviceIdParam string, vendorTemplateIdParam string) (model.VendorTemplate, error)

	// List all InstanceEndpoints of a service instance.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @return com.vmware.model.InstanceEndpointListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListInstanceEndpoints(serviceIdParam string, serviceInstanceIdParam string) (model.InstanceEndpointListResult, error)

	// Returns list of instance runtimes of service VMs being deployed for a given service instance id
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @return com.vmware.model.InstanceRuntimeListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListInstanceRuntimes(serviceIdParam string, serviceInstanceIdParam string) (model.InstanceRuntimeListResult, error)

	// List all service profiles of a service.
	//
	// @param serviceIdParam (required)
	// @return com.vmware.model.SIServiceProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListSIServiceProfiles(serviceIdParam string) (model.SIServiceProfileListResult, error)

	// Returns all Service-Attachement(s) present in the system.
	// @return com.vmware.model.ServiceAttachmentListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceAttachments() (model.ServiceAttachmentListResult, error)

	// List all service chain mappings in the system for the given service profile.
	//
	// @param serviceIdParam (required)
	// @param serviceProfileIdParam (required)
	// @return com.vmware.model.ServiceChainMappingListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceChainMappings(serviceIdParam string, serviceProfileIdParam string) (model.ServiceChainMappingListResult, error)

	// List all service chains in the system.
	// @return com.vmware.model.ServiceChainListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceChains() (model.ServiceChainListResult, error)

	// List all Service Insertion section in paginated form. A default page size is limited to 1000 sections. By default, the list of section is filtered by L3REDIRECT type. Note- GET service insertion sections API is deprecated. Please use the policy redirection-policy API.
	//
	// @param appliedTosParam AppliedTo's referenced by this section or section's Distributed Service Rules . (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param destinationsParam Destinations referenced by this section's Distributed Service Rules . (optional)
	// @param excludeAppliedToTypeParam Resource type valid for use as AppliedTo filter in section API (optional)
	// @param filterTypeParam Filter type (optional, default to FILTER)
	// @param includeAppliedToTypeParam Resource type valid for use as AppliedTo filter in section API (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param servicesParam NSService referenced by this section's Distributed Service Rules . (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourcesParam Sources referenced by this section's Distributed Service Rules . (optional)
	// @param type_Param Section Type (optional, default to L3REDIRECT)
	// @return com.vmware.model.ServiceInsertionSectionListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceInsertionSections(appliedTosParam *string, cursorParam *string, destinationsParam *string, excludeAppliedToTypeParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (model.ServiceInsertionSectionListResult, error)

	// List all Service-Insertion Service Definitions.
	// @return com.vmware.model.ServiceInsertionServiceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceInsertionServices() (model.ServiceInsertionServiceListResult, error)

	// List all service insertion status for supported contexts
	// @return com.vmware.model.ServiceInsertionStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceInsertionStatus() (model.ServiceInsertionStatusListResult, error)

	// Returns all Service-Instance(s) of all Services present in system. When request parameter (deployed_to or service_deployment_id) is provided as a part of request, it will filter out Service-Instances accordingly.
	//
	// @param deployedToParam Deployed_to referenced by service instances present in system (optional)
	// @param serviceDeploymentIdParam Service Deployment Id using which the instances were deployed (optional)
	// @return com.vmware.model.ServiceInstanceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceInstances(deployedToParam *string, serviceDeploymentIdParam *string) (model.ServiceInstanceListResult, error)

	// Returns all Service-Instance(s) for a given Service-Insertion Service.
	//
	// @param serviceIdParam (required)
	// @return com.vmware.model.ServiceInstanceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceInstancesForService(serviceIdParam string) (model.ServiceInstanceListResult, error)

	// List all service managers.
	// @return com.vmware.model.ServiceManagerListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServiceManagers() (model.ServiceManagerListResult, error)

	// List all service paths for the given service chain for the given service chain id NOTE: GET service paths api is deprecated, please use the policy API.
	//
	// @param serviceChainIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.ServicePathListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListServicePaths(serviceChainIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ServicePathListResult, error)

	// Returns Solution Config information for a given service.
	//
	// @param serviceIdParam (required)
	// @return com.vmware.model.SolutionConfigListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListSolutionConfigs(serviceIdParam string) (model.SolutionConfigListResult, error)

	// List all vendor templates of a service.
	//
	// @param serviceIdParam (required)
	// @param vendorTemplateNameParam Name of vendor template (optional)
	// @return com.vmware.model.VendorTemplateListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListVendorTemplates(serviceIdParam string, vendorTemplateNameParam *string) (model.VendorTemplateListResult, error)

	// Register service-manager with NSX with basic details like name, username, password.
	//
	// @param serviceManagerParam (required)
	// @return com.vmware.model.ServiceManager
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RegisterServiceManager(serviceManagerParam model.ServiceManager) (model.ServiceManager, error)

	// Note- POST serviceinsertion excludelist API is deprecated. Please use the policy serviceinsertion excludelist API instead.
	//
	// @param objectIdParam Identifier of the object (required)
	// @return com.vmware.model.ResourceReference
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RemoveServiceInsertionExcludeListMemberRemoveMember(objectIdParam string) (model.ResourceReference, error)

	// Service insertion data path inserts unique 'source node id' value into each packet. This API can be used to identify the source of the packet using this value. It can be resolved to multiple source entities.
	//
	// @param sourceNodeValueParam value (required)
	// @return com.vmware.model.SourceEntityResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ResolveSourceEntities(sourceNodeValueParam string) (model.SourceEntityResult, error)

	// Modifies existing serviceinsertion rule along with relative position among other serviceinsertion rules inside a serviceinsertion section. Note- POST service insertion rule API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @param serviceInsertionRuleParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.model.ServiceInsertionRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReviseServiceInsertionRuleRevise(sectionIdParam string, ruleIdParam string, serviceInsertionRuleParam model.ServiceInsertionRule, idParam *string, operationParam *string) (model.ServiceInsertionRule, error)

	// Modifies an existing serviceinsertion section along with its relative position among other serviceinsertion sections in the system. Note- POST service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.model.ServiceInsertionSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReviseServiceInsertionSectionRevise(sectionIdParam string, serviceInsertionSectionParam model.ServiceInsertionSection, idParam *string, operationParam *string) (model.ServiceInsertionSection, error)

	// Modifies an existing serviceinsertion section along with its relative position among other serviceinsertion sections with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to move a section above or below another section, use: POST /api/v1/serviceinsertion/sections/<section-id>?action=revise To modify rules, use: PUT /api/v1/serviceinsertion/sections/<section-id>/rules/<rule-id> Note- POST service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.model.ServiceInsertionSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReviseServiceInsertionSectionWithRulesReviseWithRules(sectionIdParam string, serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionSectionRuleList, error)

	// Modifies an existing service attachment. Updates to name, description and Logical Router list only supported.
	//
	// @param serviceAttachmentIdParam (required)
	// @param serviceAttachmentParam (required)
	// @return com.vmware.model.ServiceAttachment
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceAttachment(serviceAttachmentIdParam string, serviceAttachmentParam model.ServiceAttachment) (model.ServiceAttachment, error)

	// This API is deprecated since only property we can change on service deployment is display name, which is used for the SVM name. Changing the name will cause the name of the deployment to go out of sync with the deployed VM.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param serviceDeploymentParam (required)
	// @return com.vmware.model.ServiceDeployment
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceDeployment(serviceIdParam string, serviceDeploymentIdParam string, serviceDeploymentParam model.ServiceDeployment) (model.ServiceDeployment, error)

	// Modify exclude list. This includes adding/removing members in the list. Note- PUT serviceinsertion excludelist API is deprecated. Please use the policy serviceinsertion excludelist API instead.
	//
	// @param siExcludeListParam (required)
	// @return com.vmware.model.SIExcludeList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceInsertionExcludeList(siExcludeListParam model.SIExcludeList) (model.SIExcludeList, error)

	// Modifies existing serviceinsertion rule in a serviceinsertion section. Note- PUT service insertion rule API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @param serviceInsertionRuleParam (required)
	// @return com.vmware.model.ServiceInsertionRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceInsertionRule(sectionIdParam string, ruleIdParam string, serviceInsertionRuleParam model.ServiceInsertionRule) (model.ServiceInsertionRule, error)

	// Modifies the specified section, but does not modify the section's associated rules. Note- PUT service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionParam (required)
	// @return com.vmware.model.ServiceInsertionSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceInsertionSection(sectionIdParam string, serviceInsertionSectionParam model.ServiceInsertionSection) (model.ServiceInsertionSection, error)

	// Modifies existing serviceinsertion section along with its association with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to update rule content, use: PUT /api/v1/serviceinsertion/sections/<section-id>/rules/<rule-id> Note- POST service insertion section with rules API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionRuleListParam (required)
	// @return com.vmware.model.ServiceInsertionSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceInsertionSectionWithRulesUpdateWithRules(sectionIdParam string, serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList) (model.ServiceInsertionSectionRuleList, error)

	// Modifies the specified Service.
	//
	// @param serviceIdParam (required)
	// @param serviceDefinitionParam (required)
	// @return com.vmware.model.ServiceDefinition
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceInsertionService(serviceIdParam string, serviceDefinitionParam model.ServiceDefinition) (model.ServiceDefinition, error)

	// Update global ServiceInsertion status for a context
	//
	// @param contextTypeParam (required)
	// @param serviceInsertionStatusParam (required)
	// @return com.vmware.model.ServiceInsertionStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceInsertionStatus(contextTypeParam string, serviceInsertionStatusParam model.ServiceInsertionStatus) (model.ServiceInsertionStatus, error)

	// Modifies an existing Service-Instance for a given Service-Insertion Service.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param baseServiceInstanceParam (required)
	// The parameter must contain all the properties defined in model.BaseServiceInstance.
	// @return com.vmware.model.BaseServiceInstance
	// The return value will contain all the properties defined in model.BaseServiceInstance.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceInstance(serviceIdParam string, serviceInstanceIdParam string, baseServiceInstanceParam *data.StructValue) (*data.StructValue, error)

	// Update service-manager which is registered with NSX with basic details like name, username, password.
	//
	// @param serviceManagerIdParam (required)
	// @param serviceManagerParam (required)
	// @return com.vmware.model.ServiceManager
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceManager(serviceManagerIdParam string, serviceManagerParam model.ServiceManager) (model.ServiceManager, error)

	// Set service VM either in or out of maintenance mode for maintenance mode, or in service or out of service for runtime state. Only one value can be set at one time.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param instanceRuntimeIdParam (required)
	// @param actionParam (optional)
	// @param unhealthyReasonParam Reason for the unhealthy state (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateServiceVmState(serviceIdParam string, serviceInstanceIdParam string, instanceRuntimeIdParam string, actionParam *string, unhealthyReasonParam *string) error

	// Updates a solution config. Solution Config are service level objects, required for configuring the NXGI partner Service after deployment.
	//
	// @param serviceIdParam (required)
	// @param solutionConfigIdParam (required)
	// @param solutionConfigParam (required)
	// @return com.vmware.model.SolutionConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateSolutionConfig(serviceIdParam string, solutionConfigIdParam string, solutionConfigParam model.SolutionConfig) (model.SolutionConfig, error)

	// If new deployment spec is provided, the deployment will be moved to the provided spec provided that current deployment state is either UPGRADE_FAILED or DEPLOYMENT_SUCCESSFUL If same deployment spec is provided, upgrade will be done only if current deployment state is UPGRADE_FAILED
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param deploymentSpecNameParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpgradeServiceDeploymentUpgrade(serviceIdParam string, serviceDeploymentIdParam string, deploymentSpecNameParam model.DeploymentSpecName) error

	// Upgrade service VMs using newer version of OVF. Upgrade is a 2 step process. Update the 'deployment_spec_name' in the ServiceInstance to the new DeploymentSpec to which the service VMs will be upgraded, folowed by this 'upgrade' api. In case of HA, the stand-by service VM will be upgrade first. Once it has been upgraded, it switches to be the Active one and then the other VM will be upgrade.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpgradeServicevMsUpgrade(serviceIdParam string, serviceInstanceIdParam string) error
}

type managementPlaneAPISecurityServicesServiceInsertionClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPISecurityServicesServiceInsertionClient(connector client.Connector) *managementPlaneAPISecurityServicesServiceInsertionClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_security_services_service_insertion")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_instance_endpoint":                                         core.NewMethodIdentifier(interfaceIdentifier, "add_instance_endpoint"),
		"add_SI_service_profile":                                        core.NewMethodIdentifier(interfaceIdentifier, "add_SI_service_profile"),
		"add_service_attachment":                                        core.NewMethodIdentifier(interfaceIdentifier, "add_service_attachment"),
		"add_service_chain":                                             core.NewMethodIdentifier(interfaceIdentifier, "add_service_chain"),
		"add_service_insertion_exclude_list_member_add_member":          core.NewMethodIdentifier(interfaceIdentifier, "add_service_insertion_exclude_list_member_add_member"),
		"add_service_insertion_rule_in_section":                         core.NewMethodIdentifier(interfaceIdentifier, "add_service_insertion_rule_in_section"),
		"add_service_insertion_rules_in_section_create_multiple":        core.NewMethodIdentifier(interfaceIdentifier, "add_service_insertion_rules_in_section_create_multiple"),
		"add_service_insertion_section":                                 core.NewMethodIdentifier(interfaceIdentifier, "add_service_insertion_section"),
		"add_service_insertion_section_with_rules_create_with_rules":    core.NewMethodIdentifier(interfaceIdentifier, "add_service_insertion_section_with_rules_create_with_rules"),
		"add_service_insertion_service":                                 core.NewMethodIdentifier(interfaceIdentifier, "add_service_insertion_service"),
		"add_service_instance":                                          core.NewMethodIdentifier(interfaceIdentifier, "add_service_instance"),
		"add_vendor_template":                                           core.NewMethodIdentifier(interfaceIdentifier, "add_vendor_template"),
		"create_solution_config":                                        core.NewMethodIdentifier(interfaceIdentifier, "create_solution_config"),
		"delete_instance_endpoint":                                      core.NewMethodIdentifier(interfaceIdentifier, "delete_instance_endpoint"),
		"delete_SI_service_profile":                                     core.NewMethodIdentifier(interfaceIdentifier, "delete_SI_service_profile"),
		"delete_service_attachment":                                     core.NewMethodIdentifier(interfaceIdentifier, "delete_service_attachment"),
		"delete_service_chain":                                          core.NewMethodIdentifier(interfaceIdentifier, "delete_service_chain"),
		"delete_service_deployment":                                     core.NewMethodIdentifier(interfaceIdentifier, "delete_service_deployment"),
		"delete_service_insertion_rule":                                 core.NewMethodIdentifier(interfaceIdentifier, "delete_service_insertion_rule"),
		"delete_service_insertion_section":                              core.NewMethodIdentifier(interfaceIdentifier, "delete_service_insertion_section"),
		"delete_service_insertion_service":                              core.NewMethodIdentifier(interfaceIdentifier, "delete_service_insertion_service"),
		"delete_service_instance":                                       core.NewMethodIdentifier(interfaceIdentifier, "delete_service_instance"),
		"delete_service_manager":                                        core.NewMethodIdentifier(interfaceIdentifier, "delete_service_manager"),
		"delete_servicev_ms_delete":                                     core.NewMethodIdentifier(interfaceIdentifier, "delete_servicev_ms_delete"),
		"delete_solution_config":                                        core.NewMethodIdentifier(interfaceIdentifier, "delete_solution_config"),
		"delete_vendor_template":                                        core.NewMethodIdentifier(interfaceIdentifier, "delete_vendor_template"),
		"deploy_service":                                                core.NewMethodIdentifier(interfaceIdentifier, "deploy_service"),
		"deploy_servicev_ms_deploy":                                     core.NewMethodIdentifier(interfaceIdentifier, "deploy_servicev_ms_deploy"),
		"get_instance_endpoint":                                         core.NewMethodIdentifier(interfaceIdentifier, "get_instance_endpoint"),
		"get_runtime_interface_operational_status":                      core.NewMethodIdentifier(interfaceIdentifier, "get_runtime_interface_operational_status"),
		"get_runtime_interface_statistics":                              core.NewMethodIdentifier(interfaceIdentifier, "get_runtime_interface_statistics"),
		"get_SI_service_profile":                                        core.NewMethodIdentifier(interfaceIdentifier, "get_SI_service_profile"),
		"get_service_attachment":                                        core.NewMethodIdentifier(interfaceIdentifier, "get_service_attachment"),
		"get_service_chain":                                             core.NewMethodIdentifier(interfaceIdentifier, "get_service_chain"),
		"get_service_deployment":                                        core.NewMethodIdentifier(interfaceIdentifier, "get_service_deployment"),
		"get_service_deployment_state":                                  core.NewMethodIdentifier(interfaceIdentifier, "get_service_deployment_state"),
		"get_service_deployment_status":                                 core.NewMethodIdentifier(interfaceIdentifier, "get_service_deployment_status"),
		"get_service_deployments":                                       core.NewMethodIdentifier(interfaceIdentifier, "get_service_deployments"),
		"get_service_insertion_exclude_list":                            core.NewMethodIdentifier(interfaceIdentifier, "get_service_insertion_exclude_list"),
		"get_service_insertion_rule":                                    core.NewMethodIdentifier(interfaceIdentifier, "get_service_insertion_rule"),
		"get_service_insertion_rules":                                   core.NewMethodIdentifier(interfaceIdentifier, "get_service_insertion_rules"),
		"get_service_insertion_section":                                 core.NewMethodIdentifier(interfaceIdentifier, "get_service_insertion_section"),
		"get_service_insertion_section_with_rules_list_with_rules":      core.NewMethodIdentifier(interfaceIdentifier, "get_service_insertion_section_with_rules_list_with_rules"),
		"get_service_insertion_service":                                 core.NewMethodIdentifier(interfaceIdentifier, "get_service_insertion_service"),
		"get_service_insertion_status":                                  core.NewMethodIdentifier(interfaceIdentifier, "get_service_insertion_status"),
		"get_service_instance":                                          core.NewMethodIdentifier(interfaceIdentifier, "get_service_instance"),
		"get_service_instance_NS_groups":                                core.NewMethodIdentifier(interfaceIdentifier, "get_service_instance_NS_groups"),
		"get_service_instance_state":                                    core.NewMethodIdentifier(interfaceIdentifier, "get_service_instance_state"),
		"get_service_instance_status":                                   core.NewMethodIdentifier(interfaceIdentifier, "get_service_instance_status"),
		"get_service_manager":                                           core.NewMethodIdentifier(interfaceIdentifier, "get_service_manager"),
		"get_service_profile_NS_groups":                                 core.NewMethodIdentifier(interfaceIdentifier, "get_service_profile_NS_groups"),
		"get_solution_config":                                           core.NewMethodIdentifier(interfaceIdentifier, "get_solution_config"),
		"get_vendor_template":                                           core.NewMethodIdentifier(interfaceIdentifier, "get_vendor_template"),
		"list_instance_endpoints":                                       core.NewMethodIdentifier(interfaceIdentifier, "list_instance_endpoints"),
		"list_instance_runtimes":                                        core.NewMethodIdentifier(interfaceIdentifier, "list_instance_runtimes"),
		"list_SI_service_profiles":                                      core.NewMethodIdentifier(interfaceIdentifier, "list_SI_service_profiles"),
		"list_service_attachments":                                      core.NewMethodIdentifier(interfaceIdentifier, "list_service_attachments"),
		"list_service_chain_mappings":                                   core.NewMethodIdentifier(interfaceIdentifier, "list_service_chain_mappings"),
		"list_service_chains":                                           core.NewMethodIdentifier(interfaceIdentifier, "list_service_chains"),
		"list_service_insertion_sections":                               core.NewMethodIdentifier(interfaceIdentifier, "list_service_insertion_sections"),
		"list_service_insertion_services":                               core.NewMethodIdentifier(interfaceIdentifier, "list_service_insertion_services"),
		"list_service_insertion_status":                                 core.NewMethodIdentifier(interfaceIdentifier, "list_service_insertion_status"),
		"list_service_instances":                                        core.NewMethodIdentifier(interfaceIdentifier, "list_service_instances"),
		"list_service_instances_for_service":                            core.NewMethodIdentifier(interfaceIdentifier, "list_service_instances_for_service"),
		"list_service_managers":                                         core.NewMethodIdentifier(interfaceIdentifier, "list_service_managers"),
		"list_service_paths":                                            core.NewMethodIdentifier(interfaceIdentifier, "list_service_paths"),
		"list_solution_configs":                                         core.NewMethodIdentifier(interfaceIdentifier, "list_solution_configs"),
		"list_vendor_templates":                                         core.NewMethodIdentifier(interfaceIdentifier, "list_vendor_templates"),
		"register_service_manager":                                      core.NewMethodIdentifier(interfaceIdentifier, "register_service_manager"),
		"remove_service_insertion_exclude_list_member_remove_member":    core.NewMethodIdentifier(interfaceIdentifier, "remove_service_insertion_exclude_list_member_remove_member"),
		"resolve_source_entities":                                       core.NewMethodIdentifier(interfaceIdentifier, "resolve_source_entities"),
		"revise_service_insertion_rule_revise":                          core.NewMethodIdentifier(interfaceIdentifier, "revise_service_insertion_rule_revise"),
		"revise_service_insertion_section_revise":                       core.NewMethodIdentifier(interfaceIdentifier, "revise_service_insertion_section_revise"),
		"revise_service_insertion_section_with_rules_revise_with_rules": core.NewMethodIdentifier(interfaceIdentifier, "revise_service_insertion_section_with_rules_revise_with_rules"),
		"update_service_attachment":                                     core.NewMethodIdentifier(interfaceIdentifier, "update_service_attachment"),
		"update_service_deployment":                                     core.NewMethodIdentifier(interfaceIdentifier, "update_service_deployment"),
		"update_service_insertion_exclude_list":                         core.NewMethodIdentifier(interfaceIdentifier, "update_service_insertion_exclude_list"),
		"update_service_insertion_rule":                                 core.NewMethodIdentifier(interfaceIdentifier, "update_service_insertion_rule"),
		"update_service_insertion_section":                              core.NewMethodIdentifier(interfaceIdentifier, "update_service_insertion_section"),
		"update_service_insertion_section_with_rules_update_with_rules": core.NewMethodIdentifier(interfaceIdentifier, "update_service_insertion_section_with_rules_update_with_rules"),
		"update_service_insertion_service":                              core.NewMethodIdentifier(interfaceIdentifier, "update_service_insertion_service"),
		"update_service_insertion_status":                               core.NewMethodIdentifier(interfaceIdentifier, "update_service_insertion_status"),
		"update_service_instance":                                       core.NewMethodIdentifier(interfaceIdentifier, "update_service_instance"),
		"update_service_manager":                                        core.NewMethodIdentifier(interfaceIdentifier, "update_service_manager"),
		"update_service_vm_state":                                       core.NewMethodIdentifier(interfaceIdentifier, "update_service_vm_state"),
		"update_solution_config":                                        core.NewMethodIdentifier(interfaceIdentifier, "update_solution_config"),
		"upgrade_service_deployment_upgrade":                            core.NewMethodIdentifier(interfaceIdentifier, "upgrade_service_deployment_upgrade"),
		"upgrade_servicev_ms_upgrade":                                   core.NewMethodIdentifier(interfaceIdentifier, "upgrade_servicev_ms_upgrade"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPISecurityServicesServiceInsertionClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddInstanceEndpoint(serviceIdParam string, serviceInstanceIdParam string, instanceEndpointParam model.InstanceEndpoint) (model.InstanceEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddInstanceEndpointInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("InstanceEndpoint", instanceEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.InstanceEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddInstanceEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_instance_endpoint", inputDataValue, executionContext)
	var emptyOutput model.InstanceEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddInstanceEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.InstanceEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddSIServiceProfile(serviceIdParam string, baseServiceProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddSIServiceProfileInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("BaseServiceProfile", baseServiceProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddSIServiceProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_SI_service_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddSIServiceProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceAttachment(serviceAttachmentParam model.ServiceAttachment) (model.ServiceAttachment, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceAttachmentInputType(), typeConverter)
	sv.AddStructField("ServiceAttachment", serviceAttachmentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceAttachment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceAttachmentRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_attachment", inputDataValue, executionContext)
	var emptyOutput model.ServiceAttachment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceAttachmentOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceAttachment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceChain(serviceChainParam model.ServiceChain) (model.ServiceChain, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceChainInputType(), typeConverter)
	sv.AddStructField("ServiceChain", serviceChainParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceChain
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceChainRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_chain", inputDataValue, executionContext)
	var emptyOutput model.ServiceChain
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceChainOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceChain), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceInsertionExcludeListMemberAddMember(resourceReferenceParam model.ResourceReference) (model.ResourceReference, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionExcludeListMemberAddMemberInputType(), typeConverter)
	sv.AddStructField("ResourceReference", resourceReferenceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ResourceReference
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionExcludeListMemberAddMemberRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_insertion_exclude_list_member_add_member", inputDataValue, executionContext)
	var emptyOutput model.ResourceReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionExcludeListMemberAddMemberOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ResourceReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceInsertionRuleInSection(sectionIdParam string, serviceInsertionRuleParam model.ServiceInsertionRule, idParam *string, operationParam *string) (model.ServiceInsertionRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRuleInSectionInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionRule", serviceInsertionRuleParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRuleInSectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_insertion_rule_in_section", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRuleInSectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceInsertionRulesInSectionCreateMultiple(sectionIdParam string, serviceInsertionRuleListParam model.ServiceInsertionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionRuleList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRulesInSectionCreateMultipleInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionRuleList", serviceInsertionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRulesInSectionCreateMultipleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_insertion_rules_in_section_create_multiple", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRulesInSectionCreateMultipleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceInsertionSection(serviceInsertionSectionParam model.ServiceInsertionSection, idParam *string, operationParam *string) (model.ServiceInsertionSection, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionInputType(), typeConverter)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_insertion_section", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceInsertionSectionWithRulesCreateWithRules(serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionSectionRuleList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionWithRulesCreateWithRulesInputType(), typeConverter)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionWithRulesCreateWithRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_insertion_section_with_rules_create_with_rules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionWithRulesCreateWithRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceInsertionService(serviceDefinitionParam model.ServiceDefinition) (model.ServiceDefinition, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionServiceInputType(), typeConverter)
	sv.AddStructField("ServiceDefinition", serviceDefinitionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDefinition
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_insertion_service", inputDataValue, executionContext)
	var emptyOutput model.ServiceDefinition
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDefinition), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddServiceInstance(serviceIdParam string, baseServiceInstanceParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddServiceInstanceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("BaseServiceInstance", baseServiceInstanceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddServiceInstanceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_service_instance", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddServiceInstanceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) AddVendorTemplate(serviceIdParam string, vendorTemplateParam model.VendorTemplate) (model.VendorTemplate, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionAddVendorTemplateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("VendorTemplate", vendorTemplateParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VendorTemplate
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionAddVendorTemplateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "add_vendor_template", inputDataValue, executionContext)
	var emptyOutput model.VendorTemplate
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionAddVendorTemplateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VendorTemplate), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) CreateSolutionConfig(serviceIdParam string, solutionConfigParam model.SolutionConfig) (model.SolutionConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionCreateSolutionConfigInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("SolutionConfig", solutionConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SolutionConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionCreateSolutionConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "create_solution_config", inputDataValue, executionContext)
	var emptyOutput model.SolutionConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionCreateSolutionConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SolutionConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteInstanceEndpoint(serviceIdParam string, serviceInstanceIdParam string, instanceEndpointIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteInstanceEndpointInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("InstanceEndpointId", instanceEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteInstanceEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_instance_endpoint", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteSIServiceProfile(serviceIdParam string, serviceProfileIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteSIServiceProfileInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceProfileId", serviceProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteSIServiceProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_SI_service_profile", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServiceAttachment(serviceAttachmentIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServiceAttachmentInputType(), typeConverter)
	sv.AddStructField("ServiceAttachmentId", serviceAttachmentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServiceAttachmentRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_service_attachment", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServiceChain(serviceChainIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServiceChainInputType(), typeConverter)
	sv.AddStructField("ServiceChainId", serviceChainIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServiceChainRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_service_chain", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServiceDeployment(serviceIdParam string, serviceDeploymentIdParam string, forceParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServiceDeploymentInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServiceDeploymentRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_service_deployment", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServiceInsertionRule(sectionIdParam string, ruleIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionRuleInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_service_insertion_rule", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServiceInsertionSection(sectionIdParam string, cascadeParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionSectionInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("Cascade", cascadeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionSectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_service_insertion_section", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServiceInsertionService(serviceIdParam string, cascadeParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionServiceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("Cascade", cascadeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_service_insertion_service", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServiceInstance(serviceIdParam string, serviceInstanceIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInstanceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInstanceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_service_instance", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServiceManager(serviceManagerIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServiceManagerInputType(), typeConverter)
	sv.AddStructField("ServiceManagerId", serviceManagerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServiceManagerRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_service_manager", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteServicevMsDelete(serviceIdParam string, serviceInstanceIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteServicevMsDeleteInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteServicevMsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_servicev_ms_delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteSolutionConfig(serviceIdParam string, solutionConfigIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteSolutionConfigInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("SolutionConfigId", solutionConfigIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteSolutionConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_solution_config", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeleteVendorTemplate(serviceIdParam string, vendorTemplateIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeleteVendorTemplateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("VendorTemplateId", vendorTemplateIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeleteVendorTemplateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "delete_vendor_template", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeployService(serviceIdParam string, serviceDeploymentParam model.ServiceDeployment) (model.ServiceDeployment, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeployServiceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeployment", serviceDeploymentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeployment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeployServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "deploy_service", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionDeployServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) DeployServicevMsDeploy(serviceIdParam string, serviceInstanceIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionDeployServicevMsDeployInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionDeployServicevMsDeployRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "deploy_servicev_ms_deploy", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetInstanceEndpoint(serviceIdParam string, serviceInstanceIdParam string, instanceEndpointIdParam string) (model.InstanceEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetInstanceEndpointInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("InstanceEndpointId", instanceEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.InstanceEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetInstanceEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_instance_endpoint", inputDataValue, executionContext)
	var emptyOutput model.InstanceEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetInstanceEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.InstanceEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetRuntimeInterfaceOperationalStatus(serviceIdParam string, serviceInstanceIdParam string, instanceRuntimeIdParam string, interfaceIndexParam string, sourceParam *string) (model.RuntimeInterfaceOperationalStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceOperationalStatusInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("InstanceRuntimeId", instanceRuntimeIdParam)
	sv.AddStructField("InterfaceIndex", interfaceIndexParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RuntimeInterfaceOperationalStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceOperationalStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_runtime_interface_operational_status", inputDataValue, executionContext)
	var emptyOutput model.RuntimeInterfaceOperationalStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceOperationalStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RuntimeInterfaceOperationalStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetRuntimeInterfaceStatistics(serviceIdParam string, serviceInstanceIdParam string, instanceRuntimeIdParam string, interfaceIndexParam string, sourceParam *string) (model.RuntimeInterfaceStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceStatisticsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("InstanceRuntimeId", instanceRuntimeIdParam)
	sv.AddStructField("InterfaceIndex", interfaceIndexParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RuntimeInterfaceStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_runtime_interface_statistics", inputDataValue, executionContext)
	var emptyOutput model.RuntimeInterfaceStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RuntimeInterfaceStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetSIServiceProfile(serviceIdParam string, serviceProfileIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetSIServiceProfileInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceProfileId", serviceProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetSIServiceProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_SI_service_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetSIServiceProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceAttachment(serviceAttachmentIdParam string) (model.ServiceAttachment, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceAttachmentInputType(), typeConverter)
	sv.AddStructField("ServiceAttachmentId", serviceAttachmentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceAttachment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceAttachmentRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_attachment", inputDataValue, executionContext)
	var emptyOutput model.ServiceAttachment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceAttachmentOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceAttachment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceChain(serviceChainIdParam string) (model.ServiceChain, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceChainInputType(), typeConverter)
	sv.AddStructField("ServiceChainId", serviceChainIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceChain
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceChainRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_chain", inputDataValue, executionContext)
	var emptyOutput model.ServiceChain
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceChainOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceChain), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceDeployment(serviceIdParam string, serviceDeploymentIdParam string) (model.ServiceDeployment, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeployment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_deployment", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceDeploymentState(serviceIdParam string, serviceDeploymentIdParam string) (model.ConfigurationState, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConfigurationState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_deployment_state", inputDataValue, executionContext)
	var emptyOutput model.ConfigurationState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConfigurationState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceDeploymentStatus(serviceIdParam string, serviceDeploymentIdParam string, sourceParam *string) (model.ServiceDeploymentStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStatusInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeploymentStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_deployment_status", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeploymentStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeploymentStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceDeployments(serviceIdParam string) (model.ServiceDeploymentListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeploymentListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_deployments", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeploymentListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeploymentListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInsertionExcludeList() (model.SIExcludeList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionExcludeListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SIExcludeList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionExcludeListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_insertion_exclude_list", inputDataValue, executionContext)
	var emptyOutput model.SIExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionExcludeListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SIExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInsertionRule(sectionIdParam string, ruleIdParam string) (model.ServiceInsertionRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRuleInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_insertion_rule", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInsertionRules(sectionIdParam string, appliedTosParam *string, cursorParam *string, destinationsParam *string, filterTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string) (model.ServiceInsertionRuleListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("AppliedTos", appliedTosParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Destinations", destinationsParam)
	sv.AddStructField("FilterType", filterTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Services", servicesParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Sources", sourcesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionRuleListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_insertion_rules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionRuleListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionRuleListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInsertionSection(sectionIdParam string) (model.ServiceInsertionSection, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_insertion_section", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInsertionSectionWithRulesListWithRules(sectionIdParam string) (model.ServiceInsertionSectionRuleList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionWithRulesListWithRulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionWithRulesListWithRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_insertion_section_with_rules_list_with_rules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionWithRulesListWithRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInsertionService(serviceIdParam string) (model.ServiceDefinition, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionServiceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDefinition
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_insertion_service", inputDataValue, executionContext)
	var emptyOutput model.ServiceDefinition
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDefinition), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInsertionStatus(contextTypeParam string) (model.ServiceInsertionStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionStatusInputType(), typeConverter)
	sv.AddStructField("ContextType", contextTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_insertion_status", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInstance(serviceIdParam string, serviceInstanceIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_instance", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInstanceNSGroups(serviceIdParam string, serviceInstanceIdParam string) (model.ServiceInstanceNSGroups, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceNSGroupsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInstanceNSGroups
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceNSGroupsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_instance_NS_groups", inputDataValue, executionContext)
	var emptyOutput model.ServiceInstanceNSGroups
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceNSGroupsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInstanceNSGroups), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInstanceState(serviceIdParam string, serviceInstanceIdParam string) (model.ConfigurationState, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConfigurationState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_instance_state", inputDataValue, executionContext)
	var emptyOutput model.ConfigurationState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConfigurationState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceInstanceStatus(serviceIdParam string, serviceInstanceIdParam string, sourceParam *string) (model.ServiceInstanceStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStatusInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInstanceStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_instance_status", inputDataValue, executionContext)
	var emptyOutput model.ServiceInstanceStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInstanceStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceManager(serviceManagerIdParam string) (model.ServiceManager, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceManagerInputType(), typeConverter)
	sv.AddStructField("ServiceManagerId", serviceManagerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceManager
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceManagerRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_manager", inputDataValue, executionContext)
	var emptyOutput model.ServiceManager
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceManagerOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceManager), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetServiceProfileNSGroups(serviceIdParam string, serviceProfileIdParam string) (model.ServiceProfileNSGroups, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetServiceProfileNSGroupsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceProfileId", serviceProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceProfileNSGroups
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetServiceProfileNSGroupsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_service_profile_NS_groups", inputDataValue, executionContext)
	var emptyOutput model.ServiceProfileNSGroups
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetServiceProfileNSGroupsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceProfileNSGroups), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetSolutionConfig(serviceIdParam string, solutionConfigIdParam string) (model.SolutionConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetSolutionConfigInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("SolutionConfigId", solutionConfigIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SolutionConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetSolutionConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_solution_config", inputDataValue, executionContext)
	var emptyOutput model.SolutionConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetSolutionConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SolutionConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) GetVendorTemplate(serviceIdParam string, vendorTemplateIdParam string) (model.VendorTemplate, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionGetVendorTemplateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("VendorTemplateId", vendorTemplateIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VendorTemplate
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionGetVendorTemplateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "get_vendor_template", inputDataValue, executionContext)
	var emptyOutput model.VendorTemplate
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionGetVendorTemplateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VendorTemplate), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListInstanceEndpoints(serviceIdParam string, serviceInstanceIdParam string) (model.InstanceEndpointListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListInstanceEndpointsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.InstanceEndpointListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListInstanceEndpointsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_instance_endpoints", inputDataValue, executionContext)
	var emptyOutput model.InstanceEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListInstanceEndpointsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.InstanceEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListInstanceRuntimes(serviceIdParam string, serviceInstanceIdParam string) (model.InstanceRuntimeListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListInstanceRuntimesInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.InstanceRuntimeListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListInstanceRuntimesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_instance_runtimes", inputDataValue, executionContext)
	var emptyOutput model.InstanceRuntimeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListInstanceRuntimesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.InstanceRuntimeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListSIServiceProfiles(serviceIdParam string) (model.SIServiceProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListSIServiceProfilesInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SIServiceProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListSIServiceProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_SI_service_profiles", inputDataValue, executionContext)
	var emptyOutput model.SIServiceProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListSIServiceProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SIServiceProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceAttachments() (model.ServiceAttachmentListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceAttachmentsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceAttachmentListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceAttachmentsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_attachments", inputDataValue, executionContext)
	var emptyOutput model.ServiceAttachmentListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceAttachmentsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceAttachmentListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceChainMappings(serviceIdParam string, serviceProfileIdParam string) (model.ServiceChainMappingListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceChainMappingsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceProfileId", serviceProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceChainMappingListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceChainMappingsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_chain_mappings", inputDataValue, executionContext)
	var emptyOutput model.ServiceChainMappingListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceChainMappingsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceChainMappingListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceChains() (model.ServiceChainListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceChainsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceChainListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceChainsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_chains", inputDataValue, executionContext)
	var emptyOutput model.ServiceChainListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceChainsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceChainListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceInsertionSections(appliedTosParam *string, cursorParam *string, destinationsParam *string, excludeAppliedToTypeParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (model.ServiceInsertionSectionListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionSectionsInputType(), typeConverter)
	sv.AddStructField("AppliedTos", appliedTosParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Destinations", destinationsParam)
	sv.AddStructField("ExcludeAppliedToType", excludeAppliedToTypeParam)
	sv.AddStructField("FilterType", filterTypeParam)
	sv.AddStructField("IncludeAppliedToType", includeAppliedToTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Services", servicesParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Sources", sourcesParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionSectionsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_insertion_sections", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionSectionsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceInsertionServices() (model.ServiceInsertionServiceListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionServicesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionServiceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionServicesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_insertion_services", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionServiceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionServicesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionServiceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceInsertionStatus() (model.ServiceInsertionStatusListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionStatusInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_insertion_status", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceInstances(deployedToParam *string, serviceDeploymentIdParam *string) (model.ServiceInstanceListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesInputType(), typeConverter)
	sv.AddStructField("DeployedTo", deployedToParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInstanceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_instances", inputDataValue, executionContext)
	var emptyOutput model.ServiceInstanceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInstanceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceInstancesForService(serviceIdParam string) (model.ServiceInstanceListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesForServiceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInstanceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesForServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_instances_for_service", inputDataValue, executionContext)
	var emptyOutput model.ServiceInstanceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesForServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInstanceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServiceManagers() (model.ServiceManagerListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServiceManagersInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceManagerListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServiceManagersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_managers", inputDataValue, executionContext)
	var emptyOutput model.ServiceManagerListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServiceManagersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceManagerListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListServicePaths(serviceChainIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ServicePathListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListServicePathsInputType(), typeConverter)
	sv.AddStructField("ServiceChainId", serviceChainIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServicePathListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListServicePathsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_service_paths", inputDataValue, executionContext)
	var emptyOutput model.ServicePathListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListServicePathsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServicePathListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListSolutionConfigs(serviceIdParam string) (model.SolutionConfigListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListSolutionConfigsInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SolutionConfigListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListSolutionConfigsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_solution_configs", inputDataValue, executionContext)
	var emptyOutput model.SolutionConfigListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListSolutionConfigsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SolutionConfigListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ListVendorTemplates(serviceIdParam string, vendorTemplateNameParam *string) (model.VendorTemplateListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionListVendorTemplatesInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("VendorTemplateName", vendorTemplateNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VendorTemplateListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionListVendorTemplatesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "list_vendor_templates", inputDataValue, executionContext)
	var emptyOutput model.VendorTemplateListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionListVendorTemplatesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VendorTemplateListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) RegisterServiceManager(serviceManagerParam model.ServiceManager) (model.ServiceManager, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionRegisterServiceManagerInputType(), typeConverter)
	sv.AddStructField("ServiceManager", serviceManagerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceManager
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionRegisterServiceManagerRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "register_service_manager", inputDataValue, executionContext)
	var emptyOutput model.ServiceManager
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionRegisterServiceManagerOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceManager), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) RemoveServiceInsertionExcludeListMemberRemoveMember(objectIdParam string) (model.ResourceReference, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionRemoveServiceInsertionExcludeListMemberRemoveMemberInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ResourceReference
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionRemoveServiceInsertionExcludeListMemberRemoveMemberRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "remove_service_insertion_exclude_list_member_remove_member", inputDataValue, executionContext)
	var emptyOutput model.ResourceReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionRemoveServiceInsertionExcludeListMemberRemoveMemberOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ResourceReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ResolveSourceEntities(sourceNodeValueParam string) (model.SourceEntityResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionResolveSourceEntitiesInputType(), typeConverter)
	sv.AddStructField("SourceNodeValue", sourceNodeValueParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SourceEntityResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionResolveSourceEntitiesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "resolve_source_entities", inputDataValue, executionContext)
	var emptyOutput model.SourceEntityResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionResolveSourceEntitiesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SourceEntityResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ReviseServiceInsertionRuleRevise(sectionIdParam string, ruleIdParam string, serviceInsertionRuleParam model.ServiceInsertionRule, idParam *string, operationParam *string) (model.ServiceInsertionRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionRuleReviseInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("ServiceInsertionRule", serviceInsertionRuleParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionRuleReviseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "revise_service_insertion_rule_revise", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionRuleReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ReviseServiceInsertionSectionRevise(sectionIdParam string, serviceInsertionSectionParam model.ServiceInsertionSection, idParam *string, operationParam *string) (model.ServiceInsertionSection, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionReviseInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionReviseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "revise_service_insertion_section_revise", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) ReviseServiceInsertionSectionWithRulesReviseWithRules(sectionIdParam string, serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionSectionRuleList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionWithRulesReviseWithRulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionWithRulesReviseWithRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "revise_service_insertion_section_with_rules_revise_with_rules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionWithRulesReviseWithRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceAttachment(serviceAttachmentIdParam string, serviceAttachmentParam model.ServiceAttachment) (model.ServiceAttachment, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceAttachmentInputType(), typeConverter)
	sv.AddStructField("ServiceAttachmentId", serviceAttachmentIdParam)
	sv.AddStructField("ServiceAttachment", serviceAttachmentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceAttachment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceAttachmentRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_attachment", inputDataValue, executionContext)
	var emptyOutput model.ServiceAttachment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceAttachmentOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceAttachment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceDeployment(serviceIdParam string, serviceDeploymentIdParam string, serviceDeploymentParam model.ServiceDeployment) (model.ServiceDeployment, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceDeploymentInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("ServiceDeployment", serviceDeploymentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeployment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceDeploymentRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_deployment", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceDeploymentOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceInsertionExcludeList(siExcludeListParam model.SIExcludeList) (model.SIExcludeList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionExcludeListInputType(), typeConverter)
	sv.AddStructField("SiExcludeList", siExcludeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SIExcludeList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionExcludeListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_insertion_exclude_list", inputDataValue, executionContext)
	var emptyOutput model.SIExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionExcludeListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SIExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceInsertionRule(sectionIdParam string, ruleIdParam string, serviceInsertionRuleParam model.ServiceInsertionRule) (model.ServiceInsertionRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionRuleInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("ServiceInsertionRule", serviceInsertionRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_insertion_rule", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceInsertionSection(sectionIdParam string, serviceInsertionSectionParam model.ServiceInsertionSection) (model.ServiceInsertionSection, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_insertion_section", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceInsertionSectionWithRulesUpdateWithRules(sectionIdParam string, serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList) (model.ServiceInsertionSectionRuleList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionWithRulesUpdateWithRulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionWithRulesUpdateWithRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_insertion_section_with_rules_update_with_rules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionWithRulesUpdateWithRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceInsertionService(serviceIdParam string, serviceDefinitionParam model.ServiceDefinition) (model.ServiceDefinition, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionServiceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDefinition", serviceDefinitionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDefinition
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_insertion_service", inputDataValue, executionContext)
	var emptyOutput model.ServiceDefinition
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDefinition), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceInsertionStatus(contextTypeParam string, serviceInsertionStatusParam model.ServiceInsertionStatus) (model.ServiceInsertionStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionStatusInputType(), typeConverter)
	sv.AddStructField("ContextType", contextTypeParam)
	sv.AddStructField("ServiceInsertionStatus", serviceInsertionStatusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_insertion_status", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceInstance(serviceIdParam string, serviceInstanceIdParam string, baseServiceInstanceParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInstanceInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("BaseServiceInstance", baseServiceInstanceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInstanceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_instance", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInstanceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceManager(serviceManagerIdParam string, serviceManagerParam model.ServiceManager) (model.ServiceManager, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceManagerInputType(), typeConverter)
	sv.AddStructField("ServiceManagerId", serviceManagerIdParam)
	sv.AddStructField("ServiceManager", serviceManagerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceManager
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceManagerRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_manager", inputDataValue, executionContext)
	var emptyOutput model.ServiceManager
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateServiceManagerOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceManager), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateServiceVmState(serviceIdParam string, serviceInstanceIdParam string, instanceRuntimeIdParam string, actionParam *string, unhealthyReasonParam *string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateServiceVmStateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("InstanceRuntimeId", instanceRuntimeIdParam)
	sv.AddStructField("Action", actionParam)
	sv.AddStructField("UnhealthyReason", unhealthyReasonParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateServiceVmStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_service_vm_state", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpdateSolutionConfig(serviceIdParam string, solutionConfigIdParam string, solutionConfigParam model.SolutionConfig) (model.SolutionConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpdateSolutionConfigInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("SolutionConfigId", solutionConfigIdParam)
	sv.AddStructField("SolutionConfig", solutionConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SolutionConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpdateSolutionConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "update_solution_config", inputDataValue, executionContext)
	var emptyOutput model.SolutionConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityServicesServiceInsertionUpdateSolutionConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SolutionConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpgradeServiceDeploymentUpgrade(serviceIdParam string, serviceDeploymentIdParam string, deploymentSpecNameParam model.DeploymentSpecName) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpgradeServiceDeploymentUpgradeInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("DeploymentSpecName", deploymentSpecNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpgradeServiceDeploymentUpgradeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "upgrade_service_deployment_upgrade", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityServicesServiceInsertionClient) UpgradeServicevMsUpgrade(serviceIdParam string, serviceInstanceIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityServicesServiceInsertionUpgradeServicevMsUpgradeInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityServicesServiceInsertionUpgradeServicevMsUpgradeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_services_service_insertion", "upgrade_servicev_ms_upgrade", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
