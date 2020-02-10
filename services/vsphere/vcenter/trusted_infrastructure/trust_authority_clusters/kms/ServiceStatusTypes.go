/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ServiceStatus.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package kms

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Health`` enumeration class defines the possible service health states. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ServiceStatusHealth string

const (
    // No status available. This constant field was added in vSphere API 7.0.0.
	ServiceStatusHealth_NONE ServiceStatusHealth = "NONE"
    // Service is functioning normally. This constant field was added in vSphere API 7.0.0.
	ServiceStatusHealth_OK ServiceStatusHealth = "OK"
    // Service is functioning, however there is an issue that requires attention. This constant field was added in vSphere API 7.0.0.
	ServiceStatusHealth_WARNING ServiceStatusHealth = "WARNING"
    // Service is not functioning. This constant field was added in vSphere API 7.0.0.
	ServiceStatusHealth_ERROR ServiceStatusHealth = "ERROR"
)

func (h ServiceStatusHealth) ServiceStatusHealth() bool {
	switch h {
	case ServiceStatusHealth_NONE:
		return true
	case ServiceStatusHealth_OK:
		return true
	case ServiceStatusHealth_WARNING:
		return true
	case ServiceStatusHealth_ERROR:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information that describes the status of the service. This class was added in vSphere API 7.0.0.
type ServiceStatusInfo struct {
    // The service health status. This property was added in vSphere API 7.0.0.
	Health ServiceStatusHealth
    // Details regarding the health of the service. 
    //
    //  When the service ``Health`` is not ServiceStatusHealth#ServiceStatusHealth_OK or ServiceStatusHealth#ServiceStatusHealth_NONE, this member will provide an actionable description of the issues present.. This property was added in vSphere API 7.0.0.
	Details []std.LocalizableMessage
}



func serviceStatusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceStatusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ServiceStatusInfoBindingType)
}

func serviceStatusGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
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
		"",
		"GET",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/service-status",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}


func ServiceStatusInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.service_status.health", reflect.TypeOf(ServiceStatusHealth(ServiceStatusHealth_NONE)))
	fieldNameMap["health"] = "Health"
	fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["details"] = "Details"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.service_status.info", fields, reflect.TypeOf(ServiceStatusInfo{}), fieldNameMap, validators)
}

