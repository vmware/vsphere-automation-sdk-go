/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Floppy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package floppy

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual floppy drive device.
const Floppy_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.Floppy"


// The ``BackingType`` enumeration class defines the valid backing types for a virtual floppy drive.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type BackingType string

const (
    // Virtual floppy drive is backed by an image file.
     BackingType_IMAGE_FILE BackingType = "IMAGE_FILE"
    // Virtual floppy drive is backed by a device on the host where the virtual machine is running.
     BackingType_HOST_DEVICE BackingType = "HOST_DEVICE"
    // Virtual floppy drive is backed by a device on the client that is connected to the virtual machine console.
     BackingType_CLIENT_DEVICE BackingType = "CLIENT_DEVICE"
)

func (b BackingType) BackingType() bool {
    switch b {
        case BackingType_IMAGE_FILE:
            return true
        case BackingType_HOST_DEVICE:
            return true
        case BackingType_CLIENT_DEVICE:
            return true
        default:
            return false
    }
}





// The ``BackingInfo`` class contains information about the physical resource backing a virtual floppy drive.
 type BackingInfo struct {
    Type_ BackingType
    ImageFile *string
    HostDevice *string
    AutoDetect *bool
}






// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual floppy drive.
 type BackingSpec struct {
    Type_ BackingType
    ImageFile *string
    HostDevice *string
}






// The ``Info`` class contains information about a virtual floppy drive.
 type Info struct {
    Label string
    Backing BackingInfo
    State hardware.ConnectionState
    StartConnected bool
    AllowGuestControl bool
}






// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual floppy drive.
 type CreateSpec struct {
    Backing *BackingSpec
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual floppy drive.
 type UpdateSpec struct {
    Backing *BackingSpec
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``Summary`` class contains commonly used information about a virtual floppy drive.
 type Summary struct {
    Floppy string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["floppy"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["floppy"] = "Floppy"
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


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"UnableToAllocateResource": 500,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["floppy"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["floppy"] = "Floppy"
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["floppy"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["floppy"] = "Floppy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func connectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["floppy"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["floppy"] = "Floppy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func connectOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func connectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func disconnectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["floppy"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["floppy"] = "Floppy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func disconnectOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func disconnectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func BackingInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.floppy.backing_type", reflect.TypeOf(BackingType(BackingType_IMAGE_FILE)))
    fieldNameMap["type"] = "Type_"
    fields["image_file"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["image_file"] = "ImageFile"
    fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_device"] = "HostDevice"
    fields["auto_detect"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_detect"] = "AutoDetect"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "IMAGE_FILE": []bindings.FieldData {
                 bindings.NewFieldData("image_file", true),
            },
            "HOST_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("host_device", false),
                 bindings.NewFieldData("auto_detect", true),
            },
            "CLIENT_DEVICE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.backing_info",fields, reflect.TypeOf(BackingInfo{}), fieldNameMap, validators)
}

func BackingSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.floppy.backing_type", reflect.TypeOf(BackingType(BackingType_IMAGE_FILE)))
    fieldNameMap["type"] = "Type_"
    fields["image_file"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["image_file"] = "ImageFile"
    fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_device"] = "HostDevice"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "IMAGE_FILE": []bindings.FieldData {
                 bindings.NewFieldData("image_file", true),
            },
            "HOST_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("host_device", false),
            },
            "CLIENT_DEVICE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.backing_spec",fields, reflect.TypeOf(BackingSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["backing"] = bindings.NewReferenceType(BackingInfoBindingType)
    fieldNameMap["backing"] = "Backing"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(hardware.ConnectionState(hardware.ConnectionState_CONNECTED)))
    fieldNameMap["state"] = "State"
    fields["start_connected"] = bindings.NewBooleanType()
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewBooleanType()
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["floppy"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, "")
    fieldNameMap["floppy"] = "Floppy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


