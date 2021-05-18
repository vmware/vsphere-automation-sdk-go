// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.metadata.routing.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package routing

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// Routing information of the vAPI component along with its checksum
type ComponentData struct {
	// Routing information of the vAPI component
	Info ComponentInfo
	// Fingerprint of metadata of a vAPI component
	Fingerprint string
}

func (s *ComponentData) GetType__() bindings.BindingType {
	return ComponentDataBindingType()
}

func (s *ComponentData) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ComponentData._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information about a vAPI component that contains routing information For an explanation of routing information within components, see Component
type ComponentInfo struct {
	// Routing information of all the vAPI packages. The key in the map is the ID of the package and the value in the map is the routing information for the package For an explanation of routing information within packages, see Package
	Packages map[string]PackageInfo
}

func (s *ComponentInfo) GetType__() bindings.BindingType {
	return ComponentInfoBindingType()
}

func (s *ComponentInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ComponentInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information about a vAPI operation that contains routing information. For an explanation of containment within operations, see com.vmware.vapi.metadata.routing.Operation
type OperationInfo struct {
	// The routing information assigned for this operation. For an explanation of routing information, see RoutingInfo
	RoutingInfo RoutingInfo
}

func (s *OperationInfo) GetType__() bindings.BindingType {
	return OperationInfoBindingType()
}

func (s *OperationInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OperationInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information about a vAPI package containing routing information.
//
//  For an explanation of routing information within packages, see Package
type PackageInfo struct {
	// The routing information to be used for all the operations present in this package. If a particular operation has no explicit routing information defined in the routing definition file, this routing info will be used for enforcing routing.
	RoutingInfo RoutingInfo
	// Information about all services in this package that contain routing information. The key in the map is the ID of the service and the value in the map is the routing information for the service For an explanation of routing information within service, see Service
	Services map[string]ServiceInfo
}

func (s *PackageInfo) GetType__() bindings.BindingType {
	return PackageInfoBindingType()
}

func (s *PackageInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PackageInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Routing information
type RoutingInfo struct {
	// The initial version of the routing info allowed routing by single parameter. Clients requested allowing them to place more than one parameters so that the routing is performed by the first non-null argument in the list. To achieve that we have added method com.vmware.vapi.metadata.RoutingInfo#getRoutingPaths() which should be preferred over com.vmware.vapi.metadata.RoutingInfo#getRoutingPath() which is deprecated. The deprecated method will return string representation of the comma-separated list of ids, while the com.vmware.vapi.metadata.RoutingInfo#getRoutingPaths() will return instance of ``java.util.List<String>`` containing the ids.
	RoutingPath string
	// The routingStrategy is the actual strategy, based on which will be performed the routing. If the routingStrategy is IDROUTE, in RoutingInfo#routingPath must be assigned the id for the routing. There are also default strategies like IDFIRSTROUTE, LOCAL for which there is no need to specify routingPath. The name of these strategies is clear about where we should look for an ID to route, or if we need ID at all.
	RoutingStrategy string
	// This is comma-separated list of hints from the input ini file. Here the user must mention the type of the invoked method, e.g. HINTS(create) or HINTS(delete). In the future we expect this field to contain other hints also e.g. HINTS(create,lazy).
	OperationHints []string
	// This is map of specifically predefined resource types in the routing metadata. For example id types that do not require storage in the Inventory Service. Those type of objects are called 'positioned' - it is well known in advance where those objects will be routed, because their ids contain VC server guid. Example: Content Library Sessions are considered transient objects that do not need to be persisted in the IS. Routing ini file must contain section: [types] com.vmware.content.DownloadSession=positioned The map therefore will contain: {{"com.vmware.content.DownloadSession", "positioned"}} Note: This should not be final solution. To avoid duplication, currently this map will be stored only in one RoutingInfo object across the whole ProductModel. In the future, it might be moved to a common place as ComponentInfo, for example.
	IdTypes map[string]string
}

func (s *RoutingInfo) GetType__() bindings.BindingType {
	return RoutingInfoBindingType()
}

func (s *RoutingInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for RoutingInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information about a vAPI service that has routing information A service is said to contain routing information if any of its operations have routing information
type ServiceInfo struct {
	// The routing information to be used for all the operations present in this service. If a particular operation has no explicit routing information defined in the routing definition file, this routing info will be used for enforcing routing.
	RoutingInfo RoutingInfo
	// Information about all operations in this service that contain routing Information. The key in the map is the ID of the operation and the value in the map is the routing information for this operation.
	//
	//  For an explanation of routing information within operations, see com.vmware.vapi.metadata.routing.Operation
	Operations map[string]OperationInfo
}

func (s *ServiceInfo) GetType__() bindings.BindingType {
	return ServiceInfoBindingType()
}

func (s *ServiceInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ServiceInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ComponentDataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["info"] = bindings.NewReferenceType(ComponentInfoBindingType)
	fieldNameMap["info"] = "Info"
	fields["fingerprint"] = bindings.NewStringType()
	fieldNameMap["fingerprint"] = "Fingerprint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.routing.component_data", fields, reflect.TypeOf(ComponentData{}), fieldNameMap, validators)
}

func ComponentInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["packages"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.package"}, ""), bindings.NewReferenceType(PackageInfoBindingType), reflect.TypeOf(map[string]PackageInfo{}))
	fieldNameMap["packages"] = "Packages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.routing.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func OperationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["routing_info"] = bindings.NewReferenceType(RoutingInfoBindingType)
	fieldNameMap["routing_info"] = "RoutingInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.routing.operation_info", fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}

func PackageInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["routing_info"] = bindings.NewReferenceType(RoutingInfoBindingType)
	fieldNameMap["routing_info"] = "RoutingInfo"
	fields["services"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.service"}, ""), bindings.NewReferenceType(ServiceInfoBindingType), reflect.TypeOf(map[string]ServiceInfo{}))
	fieldNameMap["services"] = "Services"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.routing.package_info", fields, reflect.TypeOf(PackageInfo{}), fieldNameMap, validators)
}

func RoutingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["routing_path"] = bindings.NewStringType()
	fieldNameMap["routing_path"] = "RoutingPath"
	fields["routing_strategy"] = bindings.NewStringType()
	fieldNameMap["routing_strategy"] = "RoutingStrategy"
	fields["operation_hints"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["operation_hints"] = "OperationHints"
	fields["id_types"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(), reflect.TypeOf(map[string]string{}))
	fieldNameMap["id_types"] = "IdTypes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.routing.routing_info", fields, reflect.TypeOf(RoutingInfo{}), fieldNameMap, validators)
}

func ServiceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["routing_info"] = bindings.NewReferenceType(RoutingInfoBindingType)
	fieldNameMap["routing_info"] = "RoutingInfo"
	fields["operations"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), bindings.NewReferenceType(OperationInfoBindingType), reflect.TypeOf(map[string]OperationInfo{}))
	fieldNameMap["operations"] = "Operations"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.routing.service_info", fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}
