/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Scsi.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package scsi

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual SCSI adapter device.
const Scsi_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.ScsiAdapter"


// The ``Type`` enumeration class defines the valid emulation types for a virtual SCSI adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Type string

const (
    // BusLogic host bus adapter.
     Type_BUSLOGIC Type = "BUSLOGIC"
    // LSI Logic host bus adapter.
     Type_LSILOGIC Type = "LSILOGIC"
    // LSI Logic SAS 1068 host bus adapter.
     Type_LSILOGICSAS Type = "LSILOGICSAS"
    // Paravirtualized host bus adapter.
     Type_PVSCSI Type = "PVSCSI"
)

func (t Type) Type() bool {
    switch t {
        case Type_BUSLOGIC:
            return true
        case Type_LSILOGIC:
            return true
        case Type_LSILOGICSAS:
            return true
        case Type_PVSCSI:
            return true
        default:
            return false
    }
}




// The ``Sharing`` enumeration class defines the valid bus sharing modes for a virtual SCSI adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Sharing string

const (
    // The virtual SCSI bus is not shared.
     Sharing_NONE Sharing = "NONE"
    // The virtual SCSI bus is shared between two or more virtual machines. In this case, no physical machine is involved.
     Sharing_VIRTUAL Sharing = "VIRTUAL"
    // The virtual SCSI bus is shared between two or more virtual machines residing on different physical hosts.
     Sharing_PHYSICAL Sharing = "PHYSICAL"
)

func (s Sharing) Sharing() bool {
    switch s {
        case Sharing_NONE:
            return true
        case Sharing_VIRTUAL:
            return true
        case Sharing_PHYSICAL:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about a virtual SCSI adapter.
 type Info struct {
    Label string
    Type_ Type
    Scsi hardware.ScsiAddressInfo
    PciSlotNumber *int64
    Sharing Sharing
}






// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual SCSI adapter.
 type CreateSpec struct {
    Type_ *Type
    Bus *int64
    PciSlotNumber *int64
    Sharing *Sharing
}






// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual SCSI adapter.
 type UpdateSpec struct {
    Sharing *Sharing
}






// The ``Summary`` class contains commonly used information about a Virtual SCSI adapter.
 type Summary struct {
    Adapter string
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
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["adapter"] = "Adapter"
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
    return bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
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
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"UnableToAllocateResource": 500,"ResourceInUse": 400,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["adapter"] = "Adapter"
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
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["adapter"] = "Adapter"
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
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.type", reflect.TypeOf(Type(Type_BUSLOGIC)))
    fieldNameMap["type"] = "Type_"
    fields["scsi"] = bindings.NewReferenceType(hardware.ScsiAddressInfoBindingType)
    fieldNameMap["scsi"] = "Scsi"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    fields["sharing"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.sharing", reflect.TypeOf(Sharing(Sharing_NONE)))
    fieldNameMap["sharing"] = "Sharing"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.scsi.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.type", reflect.TypeOf(Type(Type_BUSLOGIC))))
    fieldNameMap["type"] = "Type_"
    fields["bus"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["bus"] = "Bus"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    fields["sharing"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.sharing", reflect.TypeOf(Sharing(Sharing_NONE))))
    fieldNameMap["sharing"] = "Sharing"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.scsi.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sharing"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.sharing", reflect.TypeOf(Sharing(Sharing_NONE))))
    fieldNameMap["sharing"] = "Sharing"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.scsi.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
    fieldNameMap["adapter"] = "Adapter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.scsi.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


