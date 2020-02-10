/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ClientCertificate.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Info`` class contains the client certificate used by the hosts in a cluster for authenticating with the Provider. This class was added in vSphere API 7.0.0.
type ClientCertificateInfo struct {
    // Public certificate. This property was added in vSphere API 7.0.0.
	Certificate string
}

// The ``UpdateSpec`` class contains properties that describe the client certificate update for a Key Provider. This class was added in vSphere API 7.0.0.
type ClientCertificateUpdateSpec struct {
    // Public certificate used by every host in the cluster. This property was added in vSphere API 7.0.0.
	Certificate string
    // Private part of the certificate. This property was added in vSphere API 7.0.0.
	PrivateKey *string
}



func clientCertificateCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientCertificateCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clientCertificateCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
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
		"POST",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}/client-certificate",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}

func clientCertificateGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientCertificateGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ClientCertificateInfoBindingType)
}

func clientCertificateGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}/client-certificate",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}

func clientCertificateUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fields["spec"] = bindings.NewReferenceType(ClientCertificateUpdateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientCertificateUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clientCertificateUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fields["spec"] = bindings.NewReferenceType(ClientCertificateUpdateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(ClientCertificateUpdateSpecBindingType)
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
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
		"PATCH",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}/client-certificate",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}


func ClientCertificateInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificate"] = bindings.NewStringType()
	fieldNameMap["certificate"] = "Certificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.client_certificate.info", fields, reflect.TypeOf(ClientCertificateInfo{}), fieldNameMap, validators)
}

func ClientCertificateUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificate"] = bindings.NewStringType()
	fieldNameMap["certificate"] = "Certificate"
	fields["private_key"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["private_key"] = "PrivateKey"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.client_certificate.update_spec", fields, reflect.TypeOf(ClientCertificateUpdateSpec{}), fieldNameMap, validators)
}

