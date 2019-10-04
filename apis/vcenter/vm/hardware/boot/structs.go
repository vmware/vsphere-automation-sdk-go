/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Boot.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package boot

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Type`` enumeration class defines the valid firmware types for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Type string

const (
    // Basic Input/Output System (BIOS) firmware.
     Type_BIOS Type = "BIOS"
    // Extensible Firmware Interface (EFI) firmware.
     Type_EFI Type = "EFI"
)

func (t Type) Type() bool {
    switch t {
        case Type_BIOS:
            return true
        case Type_EFI:
            return true
        default:
            return false
    }
}




// The ``NetworkProtocol`` enumeration class defines the valid network boot protocols supported when booting a virtual machine with Type#Type_EFI firmware over the network.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type NetworkProtocol string

const (
    // PXE or Apple NetBoot over IPv4.
     NetworkProtocol_IPV4 NetworkProtocol = "IPV4"
    // PXE over IPv6.
     NetworkProtocol_IPV6 NetworkProtocol = "IPV6"
)

func (n NetworkProtocol) NetworkProtocol() bool {
    switch n {
        case NetworkProtocol_IPV4:
            return true
        case NetworkProtocol_IPV6:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about the virtual machine boot process.
 type Info struct {
    Type_ Type
    EfiLegacyBoot *bool
    NetworkProtocol *NetworkProtocol
    Delay int64
    Retry bool
    RetryDelay int64
    EnterSetupMode bool
}






// The ``CreateSpec`` class describes settings used when booting a virtual machine.
 type CreateSpec struct {
    Type_ *Type
    EfiLegacyBoot *bool
    NetworkProtocol *NetworkProtocol
    Delay *int64
    Retry *bool
    RetryDelay *int64
    EnterSetupMode *bool
}






// The ``UpdateSpec`` class describes the updates to the settings used when booting a virtual machine.
 type UpdateSpec struct {
    Type_ *Type
    EfiLegacyBoot *bool
    NetworkProtocol *NetworkProtocol
    Delay *int64
    Retry *bool
    RetryDelay *int64
    EnterSetupMode *bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.type", reflect.TypeOf(Type(Type_BIOS)))
    fieldNameMap["type"] = "Type_"
    fields["efi_legacy_boot"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["efi_legacy_boot"] = "EfiLegacyBoot"
    fields["network_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.network_protocol", reflect.TypeOf(NetworkProtocol(NetworkProtocol_IPV4))))
    fieldNameMap["network_protocol"] = "NetworkProtocol"
    fields["delay"] = bindings.NewIntegerType()
    fieldNameMap["delay"] = "Delay"
    fields["retry"] = bindings.NewBooleanType()
    fieldNameMap["retry"] = "Retry"
    fields["retry_delay"] = bindings.NewIntegerType()
    fieldNameMap["retry_delay"] = "RetryDelay"
    fields["enter_setup_mode"] = bindings.NewBooleanType()
    fieldNameMap["enter_setup_mode"] = "EnterSetupMode"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "EFI": []bindings.FieldData {
                 bindings.NewFieldData("efi_legacy_boot", true),
                 bindings.NewFieldData("network_protocol", true),
            },
            "BIOS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.boot.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.type", reflect.TypeOf(Type(Type_BIOS))))
    fieldNameMap["type"] = "Type_"
    fields["efi_legacy_boot"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["efi_legacy_boot"] = "EfiLegacyBoot"
    fields["network_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.network_protocol", reflect.TypeOf(NetworkProtocol(NetworkProtocol_IPV4))))
    fieldNameMap["network_protocol"] = "NetworkProtocol"
    fields["delay"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["delay"] = "Delay"
    fields["retry"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["retry"] = "Retry"
    fields["retry_delay"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["retry_delay"] = "RetryDelay"
    fields["enter_setup_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["enter_setup_mode"] = "EnterSetupMode"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "EFI": []bindings.FieldData {
                 bindings.NewFieldData("efi_legacy_boot", false),
                 bindings.NewFieldData("network_protocol", false),
            },
            "BIOS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.boot.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.type", reflect.TypeOf(Type(Type_BIOS))))
    fieldNameMap["type"] = "Type_"
    fields["efi_legacy_boot"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["efi_legacy_boot"] = "EfiLegacyBoot"
    fields["network_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.network_protocol", reflect.TypeOf(NetworkProtocol(NetworkProtocol_IPV4))))
    fieldNameMap["network_protocol"] = "NetworkProtocol"
    fields["delay"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["delay"] = "Delay"
    fields["retry"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["retry"] = "Retry"
    fields["retry_delay"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["retry_delay"] = "RetryDelay"
    fields["enter_setup_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["enter_setup_mode"] = "EnterSetupMode"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "EFI": []bindings.FieldData {
                 bindings.NewFieldData("efi_legacy_boot", false),
                 bindings.NewFieldData("network_protocol", false),
            },
            "BIOS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.boot.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


