/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ResourceAddressSchemas.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vstats

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for resource addressing schemas. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const ResourceAddressSchemas_RESOURCE_TYPE = "com.vmware.vstats.model.RsrcAddrSchema"


// Declares which predicates are supported by the ``ResourceIdDefinition``. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ResourceAddressSchemasQueryCapabilities string

const (
    // Equal. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ResourceAddressSchemasQueryCapabilities_EQUAL ResourceAddressSchemasQueryCapabilities = "EQUAL"
    // Supports both Equality and Aggregation. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ResourceAddressSchemasQueryCapabilities_EQUAL_ALL ResourceAddressSchemasQueryCapabilities = "EQUAL_ALL"
)

func (q ResourceAddressSchemasQueryCapabilities) ResourceAddressSchemasQueryCapabilities() bool {
	switch q {
	case ResourceAddressSchemasQueryCapabilities_EQUAL:
		return true
	case ResourceAddressSchemasQueryCapabilities_EQUAL_ALL:
		return true
	default:
		return false
	}
}


// The ``ResourceIdDefinition`` class describes a single identifier of the Resource Addressing Schema. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ResourceAddressSchemasResourceIdDefinition struct {
    // Designates a semantic key for the resource identifier. This could be "vm" for virtual machine whose CPU usage is metered or "source" to identify the virtual machine that is origin of measured network traffic. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Key string
    // Type of the resource. This represents various entities for which statistical data is gathered such as virtual machines, containers, servers, disks etc. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Type_ string
    // Designates the supported query-options. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	QueryOptions ResourceAddressSchemasQueryCapabilities
}

// The ``Info`` class defines addressing schema for a counter. This is set of named placeholders for different resource types. For example a network link between VMs will take two arguments "source" and "destination" both of type VM. For each argument query capability is defined. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ResourceAddressSchemasInfo struct {
    // Identifier. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Id string
    // List of ResourceAddressSchemasResourceIdDefinitions. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Schema []ResourceAddressSchemasResourceIdDefinition
}



func resourceAddressSchemasGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourceAddressSchemasGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ResourceAddressSchemasInfoBindingType)
}

func resourceAddressSchemasGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
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
		"/stats/rsrc-addr-schemas/{id}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func ResourceAddressSchemasResourceIdDefinitionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key"] = bindings.NewStringType()
	fieldNameMap["key"] = "Key"
	fields["type"] = bindings.NewIdType([]string{"com.vmware.vstats.model.RsrcType"}, "")
	fieldNameMap["type"] = "Type_"
	fields["query_options"] = bindings.NewEnumType("com.vmware.vstats.resource_address_schemas.query_capabilities", reflect.TypeOf(ResourceAddressSchemasQueryCapabilities(ResourceAddressSchemasQueryCapabilities_EQUAL)))
	fieldNameMap["query_options"] = "QueryOptions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.resource_address_schemas.resource_id_definition", fields, reflect.TypeOf(ResourceAddressSchemasResourceIdDefinition{}), fieldNameMap, validators)
}

func ResourceAddressSchemasInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vstats.model.RsrcAddrSchema"}, "")
	fieldNameMap["id"] = "Id"
	fields["schema"] = bindings.NewListType(bindings.NewReferenceType(ResourceAddressSchemasResourceIdDefinitionBindingType), reflect.TypeOf([]ResourceAddressSchemasResourceIdDefinition{}))
	fieldNameMap["schema"] = "Schema"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.resource_address_schemas.info", fields, reflect.TypeOf(ResourceAddressSchemasInfo{}), fieldNameMap, validators)
}

