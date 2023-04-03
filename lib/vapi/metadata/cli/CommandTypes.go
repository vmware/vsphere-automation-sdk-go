// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Command.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cli

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// The ``FormatterType`` enumeration class defines supported CLI output formatter types. See CommandInfo#formatter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CommandFormatterTypeEnum string

const (
	// Displays command output as it is.
	CommandFormatterType_SIMPLE CommandFormatterTypeEnum = "SIMPLE"
	// Displays command output in table format.
	CommandFormatterType_TABLE CommandFormatterTypeEnum = "TABLE"
	// Displays command output in JSON format.
	CommandFormatterType_JSON CommandFormatterTypeEnum = "JSON"
	// Displays command output in XML format.
	CommandFormatterType_XML CommandFormatterTypeEnum = "XML"
	// Displays command output in CSV format.
	CommandFormatterType_CSV CommandFormatterTypeEnum = "CSV"
	// Displays command output in HTML format.
	CommandFormatterType_HTML CommandFormatterTypeEnum = "HTML"
)

func (f CommandFormatterTypeEnum) CommandFormatterTypeEnum() bool {
	switch f {
	case CommandFormatterType_SIMPLE:
		return true
	case CommandFormatterType_TABLE:
		return true
	case CommandFormatterType_JSON:
		return true
	case CommandFormatterType_XML:
		return true
	case CommandFormatterType_CSV:
		return true
	case CommandFormatterType_HTML:
		return true
	default:
		return false
	}
}

// The ``GenericType`` enumeration class defines generic types supported by ``Command`` interface. See CommandOptionInfo#generic.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CommandGenericTypeEnum string

const (
	// Default case.
	CommandGenericType_NONE CommandGenericTypeEnum = "NONE"
	// Input parameter is an optional.
	CommandGenericType_OPTIONAL CommandGenericTypeEnum = "OPTIONAL"
	// Input parameter is a list.
	CommandGenericType_LIST CommandGenericTypeEnum = "LIST"
	// Input parameter is an optional of type list.
	CommandGenericType_OPTIONAL_LIST CommandGenericTypeEnum = "OPTIONAL_LIST"
	// Input parameter is a list of optionals.
	CommandGenericType_LIST_OPTIONAL CommandGenericTypeEnum = "LIST_OPTIONAL"
)

func (g CommandGenericTypeEnum) CommandGenericTypeEnum() bool {
	switch g {
	case CommandGenericType_NONE:
		return true
	case CommandGenericType_OPTIONAL:
		return true
	case CommandGenericType_LIST:
		return true
	case CommandGenericType_OPTIONAL_LIST:
		return true
	case CommandGenericType_LIST_OPTIONAL:
		return true
	default:
		return false
	}
}

// The ``OutputFieldInfo`` class describes the name used by the CLI to display a single property of a class element in the interface definition language.
type CommandOutputFieldInfo struct {
	// Name of the property.
	FieldName string
	// Name used by the CLI to display the property.
	DisplayName string
}

func (s *CommandOutputFieldInfo) GetType__() vapiBindings_.BindingType {
	return CommandOutputFieldInfoBindingType()
}

func (s *CommandOutputFieldInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CommandOutputFieldInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``OutputInfo`` class describes the names used by the CLI to display the properties of a class element in the interface definition language as well as the order in which the properties will be displayed.
type CommandOutputInfo struct {
	// Name of the class.
	StructureId string
	// The order in which the properties of the class will be displayed by the CLI as well as the names used to display the properties.
	OutputFields []CommandOutputFieldInfo
}

func (s *CommandOutputInfo) GetType__() vapiBindings_.BindingType {
	return CommandOutputInfoBindingType()
}

func (s *CommandOutputInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CommandOutputInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``OptionInfo`` class describes information about a specific input option of a command.
type CommandOptionInfo struct {
	// The long option name of the parameter as used by the user.
	LongOption string
	// The single character value option name.
	ShortOption *string
	// The fully qualified name of the option referred to by the operation element in CommandInfo#operationId.
	FieldName string
	// The description of the option to be displayed to the user when they request usage information for a CLI command.
	Description string
	// The type of option. This is used to display information about what kind of data is expected (string, number, boolean, etc.) for the option when they request usage information for a CLI command. For enumeration class this stores the fully qualified enumeration class id.
	Type_ string
	// This is used to tell the user whether the option is required or optional, or whether they can specify the option multiple times.
	Generic CommandGenericTypeEnum
}

func (s *CommandOptionInfo) GetType__() vapiBindings_.BindingType {
	return CommandOptionInfoBindingType()
}

func (s *CommandOptionInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CommandOptionInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Identity`` class uniquely identifies a command in the CLI commands tree.
type CommandIdentity struct {
	// The dot-separated path of the namespace containing the command in the CLI command tree.
	Path string
	// Name of the command.
	Name string
}

func (s *CommandIdentity) GetType__() vapiBindings_.BindingType {
	return CommandIdentityBindingType()
}

func (s *CommandIdentity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CommandIdentity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Info`` class contains information about a command. It includes the identity of the command, a description, information about the interface and method that implement the command, and CLI-specific information for the command.
type CommandInfo struct {
	// Basic command identity.
	Identity CommandIdentity
	// The text description displayed to the user in help output.
	Description string
	// The service identifier that contains the operations for this CLI command.
	ServiceId string
	// The operation identifier corresponding to this CLI command.
	OperationId string
	// The input for this command.
	Options []CommandOptionInfo
	// The formatter to use when displaying the output of this command.
	Formatter *CommandFormatterTypeEnum
	// List of output structure name and output field info.
	OutputFieldList []CommandOutputInfo
}

func (s *CommandInfo) GetType__() vapiBindings_.BindingType {
	return CommandInfoBindingType()
}

func (s *CommandInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CommandInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func commandListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommandListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CommandIdentityBindingType), reflect.TypeOf([]CommandIdentity{}))
}

func commandListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	paramsTypeMap["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	queryParams["path"] = "path"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/vapi/metadata/cli/command",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func commandGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity"] = vapiBindings_.NewReferenceType(CommandIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommandGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(CommandInfoBindingType)
}

func commandGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity"] = vapiBindings_.NewReferenceType(CommandIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	bodyFieldsMap["identity"] = "identity"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=get",
		"",
		"POST",
		"/vapi/metadata/cli/command",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func commandFingerprintInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommandFingerprintOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewStringType()
}

func commandFingerprintRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/vapi/metadata/cli/command/fingerprint",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func CommandOutputFieldInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["field_name"] = vapiBindings_.NewStringType()
	fieldNameMap["field_name"] = "FieldName"
	fields["display_name"] = vapiBindings_.NewStringType()
	fieldNameMap["display_name"] = "DisplayName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.cli.command.output_field_info", fields, reflect.TypeOf(CommandOutputFieldInfo{}), fieldNameMap, validators)
}

func CommandOutputInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["structure_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.structure"}, "")
	fieldNameMap["structure_id"] = "StructureId"
	fields["output_fields"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CommandOutputFieldInfoBindingType), reflect.TypeOf([]CommandOutputFieldInfo{}))
	fieldNameMap["output_fields"] = "OutputFields"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.cli.command.output_info", fields, reflect.TypeOf(CommandOutputInfo{}), fieldNameMap, validators)
}

func CommandOptionInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["long_option"] = vapiBindings_.NewStringType()
	fieldNameMap["long_option"] = "LongOption"
	fields["short_option"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["short_option"] = "ShortOption"
	fields["field_name"] = vapiBindings_.NewStringType()
	fieldNameMap["field_name"] = "FieldName"
	fields["description"] = vapiBindings_.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = vapiBindings_.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["generic"] = vapiBindings_.NewEnumType("com.vmware.vapi.metadata.cli.command.generic_type", reflect.TypeOf(CommandGenericTypeEnum(CommandGenericType_NONE)))
	fieldNameMap["generic"] = "Generic"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.cli.command.option_info", fields, reflect.TypeOf(CommandOptionInfo{}), fieldNameMap, validators)
}

func CommandIdentityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["path"] = vapiBindings_.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.cli.command.identity", fields, reflect.TypeOf(CommandIdentity{}), fieldNameMap, validators)
}

func CommandInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity"] = vapiBindings_.NewReferenceType(CommandIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	fields["description"] = vapiBindings_.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fields["operation_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["operation_id"] = "OperationId"
	fields["options"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CommandOptionInfoBindingType), reflect.TypeOf([]CommandOptionInfo{}))
	fieldNameMap["options"] = "Options"
	fields["formatter"] = vapiBindings_.NewOptionalType(vapiBindings_.NewEnumType("com.vmware.vapi.metadata.cli.command.formatter_type", reflect.TypeOf(CommandFormatterTypeEnum(CommandFormatterType_SIMPLE))))
	fieldNameMap["formatter"] = "Formatter"
	fields["output_field_list"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CommandOutputInfoBindingType), reflect.TypeOf([]CommandOutputInfo{}))
	fieldNameMap["output_field_list"] = "OutputFieldList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.cli.command.info", fields, reflect.TypeOf(CommandInfo{}), fieldNameMap, validators)
}
