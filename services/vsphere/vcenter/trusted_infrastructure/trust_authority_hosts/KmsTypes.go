/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Kms.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package trust_authority_hosts

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/trusted_infrastructure"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The connection information could include the certificates or be a shorter summary. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type KmsSummaryType string

const (
    // The full connection information, including certificates. This constant field was added in vSphere API 7.0.0.
	KmsSummaryType_FULL KmsSummaryType = "FULL"
    // A summary containing only the hostname, port, and the group which determines the Attestation Services this Key Provider Service can communicate with. This constant field was added in vSphere API 7.0.0.
	KmsSummaryType_NORMAL KmsSummaryType = "NORMAL"
    // A brief summary, containing only the hostname for the Key Provider Service. This constant field was added in vSphere API 7.0.0.
	KmsSummaryType_BRIEF KmsSummaryType = "BRIEF"
)

func (s KmsSummaryType) KmsSummaryType() bool {
	switch s {
	case KmsSummaryType_FULL:
		return true
	case KmsSummaryType_NORMAL:
		return true
	case KmsSummaryType_BRIEF:
		return true
	default:
		return false
	}
}


// The ``Summary`` class contains all the stored information about a Key Provider Service. This class was added in vSphere API 7.0.0.
type KmsSummary struct {
    // Defines the verbosity of the summary. This property was added in vSphere API 7.0.0.
	SummaryType KmsSummaryType
    // The trusted ESX on which the service runs. This property was added in vSphere API 7.0.0.
	Host *string
    // The service's address. This property was added in vSphere API 7.0.0.
	Address *trusted_infrastructure.NetworkAddress
    // The group ID determines which Attestation Service instances this Key Provider Service can communicate with. This property was added in vSphere API 7.0.0.
	Group *string
    // The opaque string identifier of the cluster in which the Key Provider Service is part of. This property was added in vSphere API 7.0.0.
	Cluster *string
    // The service's TLS certificate chain. This property was added in vSphere API 7.0.0.
	TrustedCA *trusted_infrastructure.X509CertChain
}

// The ``Info`` class contains all the stored information about a Key Provider Service. This class was added in vSphere API 7.0.0.
type KmsInfo struct {
    // The trusted ESX on which the service runs. This property was added in vSphere API 7.0.0.
	Host string
    // The service's address. This property was added in vSphere API 7.0.0.
	Address trusted_infrastructure.NetworkAddress
    // The group ID determines which Attestation Service instances this Key Provider Service can communicate with. This property was added in vSphere API 7.0.0.
	Group string
    // The opaque string identifier of the cluster in which the Key Provider Service is part of. This property was added in vSphere API 7.0.0.
	Cluster string
    // The service's TLS certificate chain. This property was added in vSphere API 7.0.0.
	TrustedCA trusted_infrastructure.X509CertChain
}

// The ``FilterSpec`` class contains the data necessary for identifying a Key Provider Service. This class was added in vSphere API 7.0.0.
type KmsFilterSpec struct {
    // A set of host IDs by which to filter the services. This property was added in vSphere API 7.0.0.
	Hosts map[string]bool
    // A set of cluster IDs by which to filter the services. This property was added in vSphere API 7.0.0.
	Clusters map[string]bool
    // The service's address. This property was added in vSphere API 7.0.0.
	Address []trusted_infrastructure.NetworkAddress
    // The group determines reports issued by which Attestation Service instances this Key Provider Service can accept. This property was added in vSphere API 7.0.0.
	Groups map[string]bool
}



func kmsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func kmsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(KmsInfoBindingType)
}

func kmsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	paramsTypeMap["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	paramsTypeMap["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	pathParams["host"] = "host"
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
		"/vcenter/trusted-infrastructure/trust-authority-hosts/{host}/kms/",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"ResourceInaccessible": 500})
}

func kmsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(KmsFilterSpecBindingType))
	fields["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.kms.summary_type", reflect.TypeOf(KmsSummaryType(KmsSummaryType_FULL))))
	fieldNameMap["spec"] = "Spec"
	fieldNameMap["projection"] = "Projection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func kmsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(KmsSummaryBindingType), reflect.TypeOf([]KmsSummary{}))
}

func kmsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(KmsFilterSpecBindingType))
	fields["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.kms.summary_type", reflect.TypeOf(KmsSummaryType(KmsSummaryType_FULL))))
	fieldNameMap["spec"] = "Spec"
	fieldNameMap["projection"] = "Projection"
	paramsTypeMap["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.kms.summary_type", reflect.TypeOf(KmsSummaryType(KmsSummaryType_FULL))))
	paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(KmsFilterSpecBindingType))
	queryParams["projection"] = "projection"
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
		"/vcenter/trusted-infrastructure/trust-authority-hosts/kms",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401})
}


func KmsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["summary_type"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.kms.summary_type", reflect.TypeOf(KmsSummaryType(KmsSummaryType_FULL)))
	fieldNameMap["summary_type"] = "SummaryType"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem"}, ""))
	fieldNameMap["host"] = "Host"
	fields["address"] = bindings.NewOptionalType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType))
	fieldNameMap["address"] = "Address"
	fields["group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["group"] = "Group"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster"] = "Cluster"
	fields["trusted_CA"] = bindings.NewOptionalType(bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType))
	fieldNameMap["trusted_CA"] = "TrustedCA"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("summary_type",
		map[string][]bindings.FieldData{
			"BRIEF": []bindings.FieldData{
				bindings.NewFieldData("host", true),
				bindings.NewFieldData("address", true),
			},
			"NORMAL": []bindings.FieldData{
				bindings.NewFieldData("host", true),
				bindings.NewFieldData("address", true),
				bindings.NewFieldData("group", true),
				bindings.NewFieldData("cluster", true),
			},
			"FULL": []bindings.FieldData{
				bindings.NewFieldData("host", true),
				bindings.NewFieldData("address", true),
				bindings.NewFieldData("group", true),
				bindings.NewFieldData("cluster", true),
				bindings.NewFieldData("trusted_CA", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.kms.summary", fields, reflect.TypeOf(KmsSummary{}), fieldNameMap, validators)
}

func KmsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
	fieldNameMap["address"] = "Address"
	fields["group"] = bindings.NewStringType()
	fieldNameMap["group"] = "Group"
	fields["cluster"] = bindings.NewStringType()
	fieldNameMap["cluster"] = "Cluster"
	fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
	fieldNameMap["trusted_CA"] = "TrustedCA"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.kms.info", fields, reflect.TypeOf(KmsInfo{}), fieldNameMap, validators)
}

func KmsFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["address"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType), reflect.TypeOf([]trusted_infrastructure.NetworkAddress{})))
	fieldNameMap["address"] = "Address"
	fields["groups"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["groups"] = "Groups"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.kms.filter_spec", fields, reflect.TypeOf(KmsFilterSpec{}), fieldNameMap, validators)
}

