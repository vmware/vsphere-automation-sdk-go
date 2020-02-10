/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Services.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package attestation

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/trusted_infrastructure"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Summary`` class contains a summary of a ``Services``. This class was added in vSphere API 7.0.0.
type ServicesSummary struct {
    // The service's unique identifier. This property was added in vSphere API 7.0.0.
	Service string
    // The service's address. This property was added in vSphere API 7.0.0.
	Address trusted_infrastructure.NetworkAddress
    // The group determines which ``Services`` instances can accept reports issued by this ``Services`` instance. This property was added in vSphere API 7.0.0.
	Group string
    // The cluster determines which Trust Authority Cluster this ``Services`` belongs to. This property was added in vSphere API 7.0.0.
	TrustAuthorityCluster string
}

// The ``Info`` class contains all the stored information about a ``Services``. This class was added in vSphere API 7.0.0.
type ServicesInfo struct {
    // The service's address. This property was added in vSphere API 7.0.0.
	Address trusted_infrastructure.NetworkAddress
    // The service's TLS certificate chain. This property was added in vSphere API 7.0.0.
	TrustedCA trusted_infrastructure.X509CertChain
    // The group ID determines which ``Services`` instances this ``Services`` can communicate with. This property was added in vSphere API 7.0.0.
	Group string
    // The cluster determines which Trust Authority Cluster this ``Services`` belongs to. This property was added in vSphere API 7.0.0.
	TrustAuthorityCluster string
}

// The ``CreateSpec`` class contains the data necessary for adding a ``Services`` to the environment. This class was added in vSphere API 7.0.0.
type ServicesCreateSpec struct {
    // Source of truth for the configuration of the Attestation Service. This property was added in vSphere API 7.0.0.
	Type_ ServicesCreateSpecSourceType
    // The service's unique ID. This property was added in vSphere API 7.0.0.
	Service *string
    // The attestation cluster's unique ID. This property was added in vSphere API 7.0.0.
	TrustAuthorityCluster *string
}

// The ``SourceType`` enumeration class lists options which source the the Attestation Service to use for its configuration. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ServicesCreateSpecSourceType string

const (
    // The Attestation Service will be configured based on an ID of an specific Attestation Service. This constant field was added in vSphere API 7.0.0.
	ServicesCreateSpecSourceType_SERVICE ServicesCreateSpecSourceType = "SERVICE"
    // The Attestation Service will be configured based on an ID of a whole attestation cluster. This constant field was added in vSphere API 7.0.0.
	ServicesCreateSpecSourceType_CLUSTER ServicesCreateSpecSourceType = "CLUSTER"
)

func (s ServicesCreateSpecSourceType) ServicesCreateSpecSourceType() bool {
	switch s {
	case ServicesCreateSpecSourceType_SERVICE:
		return true
	case ServicesCreateSpecSourceType_CLUSTER:
		return true
	default:
		return false
	}
}


// The ``FilterSpec`` class contains the data necessary for identifying a ``Services``. This class was added in vSphere API 7.0.0.
type ServicesFilterSpec struct {
    // A set of IDs by which to filter the services. This property was added in vSphere API 7.0.0.
	Services map[string]bool
    // The service's address. This property was added in vSphere API 7.0.0.
	Address []trusted_infrastructure.NetworkAddress
    // The group determines which ``Services`` instances can accept reports issued by this ``Services`` instance. This property was added in vSphere API 7.0.0.
	Group map[string]bool
    // The cluster determines which Trust Authority Cluster this ``Services`` belongs to. This property was added in vSphere API 7.0.0.
	TrustAuthorityCluster map[string]bool
}



func servicesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ServicesFilterSpecBindingType))
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ServicesSummaryBindingType), reflect.TypeOf([]ServicesSummary{}))
}

func servicesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ServicesFilterSpecBindingType))
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ServicesFilterSpecBindingType))
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	pathParams["cluster"] = "cluster"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"action=query",
		"spec",
		"POST",
		"/vcenter/trusted-infrastructure/trusted-clusters/{cluster}/attestation/services",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}

func servicesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ServicesInfoBindingType)
}

func servicesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["service"] = "Service"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["service"] = "service"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/vcenter/trusted-infrastructure/trusted-clusters/{cluster}/attestation/services/{service}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}

func servicesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(ServicesCreateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
}

func servicesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(ServicesCreateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(ServicesCreateSpecBindingType)
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	pathParams["cluster"] = "cluster"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"spec",
		"POST",
		"/vcenter/trusted-infrastructure/trusted-clusters/{cluster}/attestation/services",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"UnableToAllocateResource": 500,"Unauthenticated": 401})
}

func servicesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func servicesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["service"] = "Service"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["service"] = "service"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"DELETE",
		"/vcenter/trusted-infrastructure/trusted-clusters/{cluster}/attestation/services/{service}",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func ServicesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, "")
	fieldNameMap["service"] = "Service"
	fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
	fieldNameMap["address"] = "Address"
	fields["group"] = bindings.NewStringType()
	fieldNameMap["group"] = "Group"
	fields["trust_authority_cluster"] = bindings.NewStringType()
	fieldNameMap["trust_authority_cluster"] = "TrustAuthorityCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trusted_clusters.attestation.services.summary", fields, reflect.TypeOf(ServicesSummary{}), fieldNameMap, validators)
}

func ServicesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
	fieldNameMap["address"] = "Address"
	fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
	fieldNameMap["trusted_CA"] = "TrustedCA"
	fields["group"] = bindings.NewStringType()
	fieldNameMap["group"] = "Group"
	fields["trust_authority_cluster"] = bindings.NewStringType()
	fieldNameMap["trust_authority_cluster"] = "TrustAuthorityCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trusted_clusters.attestation.services.info", fields, reflect.TypeOf(ServicesInfo{}), fieldNameMap, validators)
}

func ServicesCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trusted_clusters.attestation.services.create_spec.source_type", reflect.TypeOf(ServicesCreateSpecSourceType(ServicesCreateSpecSourceType_SERVICE)))
	fieldNameMap["type"] = "Type_"
	fields["service"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, ""))
	fieldNameMap["service"] = "Service"
	fields["trust_authority_cluster"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["trust_authority_cluster"] = "TrustAuthorityCluster"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"SERVICE": []bindings.FieldData{
				bindings.NewFieldData("service", true),
			},
			"CLUSTER": []bindings.FieldData{
				bindings.NewFieldData("trust_authority_cluster", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trusted_clusters.attestation.services.create_spec", fields, reflect.TypeOf(ServicesCreateSpec{}), fieldNameMap, validators)
}

func ServicesFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["services"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.attestation.Service"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["services"] = "Services"
	fields["address"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType), reflect.TypeOf([]trusted_infrastructure.NetworkAddress{})))
	fieldNameMap["address"] = "Address"
	fields["group"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["group"] = "Group"
	fields["trust_authority_cluster"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["trust_authority_cluster"] = "TrustAuthorityCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trusted_clusters.attestation.services.filter_spec", fields, reflect.TypeOf(ServicesFilterSpec{}), fieldNameMap, validators)
}

