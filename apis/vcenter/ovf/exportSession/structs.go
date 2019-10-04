/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ExportSession.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package exportSession

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/ovf"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``State`` enumeration class defines the different states for an export session.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type State string

const (
    // Files are being made available for download. During this state entries will be added to the file list as the files become ready.
     State_EXPORT_PREPARING State = "EXPORT_PREPARING"
    // All files are available for download, and the file list is complete. This state is used when the files are being downloaded by the client using HTTP GETs.
     State_EXPORT_READY State = "EXPORT_READY"
    // Files are in the process of being transferred. This state is used when the target is a content library.
     State_EXPORT_IN_PROGRESS State = "EXPORT_IN_PROGRESS"
    // Files have been transferred. This state is used when the target is a content library.
     State_EXPORT_COMPLETED State = "EXPORT_COMPLETED"
    // The transfer failed.
     State_EXPORT_ERROR State = "EXPORT_ERROR"
)

func (s State) State() bool {
    switch s {
        case State_EXPORT_PREPARING:
            return true
        case State_EXPORT_READY:
            return true
        case State_EXPORT_IN_PROGRESS:
            return true
        case State_EXPORT_COMPLETED:
            return true
        case State_EXPORT_ERROR:
            return true
        default:
            return false
    }
}




// The ``TargetType`` enumeration class defines the target types of export session. If CreateSpec#contentLibrary is provided, the OVF is exported to a local content library.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type TargetType string

const (
    // A set of URLs are published where the files can be downloaded from.
     TargetType_DOWNLOAD_TARGET TargetType = "DOWNLOAD_TARGET"
    // The OVF package is directly transferred to a content library.
     TargetType_CONTENT_LIBRARY_TARGET TargetType = "CONTENT_LIBRARY_TARGET"
)

func (t TargetType) TargetType() bool {
    switch t {
        case TargetType_DOWNLOAD_TARGET:
            return true
        case TargetType_CONTENT_LIBRARY_TARGET:
            return true
        default:
            return false
    }
}




// The ``TargetContentType`` enumeration class defines the target content types of export session. This only applies to TargetType#TargetType_DOWNLOAD_TARGET.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type TargetContentType string

const (
    // Export the OVF package as a set of files.
     TargetContentType_OVF_TARGET TargetContentType = "OVF_TARGET"
    // Export the OVF package as a single OVA file.
     TargetContentType_OVA_TARGET TargetContentType = "OVA_TARGET"
)

func (t TargetContentType) TargetContentType() bool {
    switch t {
        case TargetContentType_OVF_TARGET:
            return true
        case TargetContentType_OVA_TARGET:
            return true
        default:
            return false
    }
}





// The ``SourceInfo`` class specifies either a virtual machine or a virtual appliance.
 type SourceInfo struct {
    VappId *string
    VmId *string
}






// The ``CreateSpec`` class contains export parameters.
 type CreateSpec struct {
    TargetType TargetType
    TargetContentType *TargetContentType
    ExportFlags []string
    Name *string
    Description *string
    ContentLibrary *string
    ContentLibraryItem *string
}






// The ``Info`` class represents an export session.
 type Info struct {
    State State
    Progress *int64
    Files []ovf.OvfFileInfo
    Errors []ovf.OvfError
    Warnings []ovf.OvfWarning
    Information []ovf.OvfInfo
    LibraryItemId *string
}






// The ``PreviewSpec`` class contains information about preview parameters.
 type PreviewSpec struct {
    ExportFlags []string
    TargetType *TargetType
}






// The ``PreviewFile`` class contains information about a file that will be exported according to a preview.
 type PreviewFile struct {
    Name string
}






// The ``Preview`` class contains information about the result of an export preview.
 type Preview struct {
    Files []PreviewFile
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["source"] = bindings.NewReferenceType(SourceInfoBindingType)
    fields["create_spec"] = bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(CreateSpecBindingType),}, bindings.JSONRPC)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["source"] = "Source"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.OvfExportSession"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"ResourceInaccessible": 500,"ResourceBusy": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfExportSession"}, "")
    fieldNameMap["id"] = "Id"
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
       map[string]int{"NotFound": 404})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfExportSession"}, "")
    fieldNameMap["id"] = "Id"
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
       map[string]int{})
}


func progressInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfExportSession"}, "")
    fields["percent"] = bindings.NewIntegerType()
    fieldNameMap["id"] = "Id"
    fieldNameMap["percent"] = "Percent"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func progressOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func progressRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}


func previewInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source"] = bindings.NewReferenceType(SourceInfoBindingType)
    fields["preview_spec"] = bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(PreviewSpecBindingType),}, bindings.JSONRPC)
    fieldNameMap["source"] = "Source"
    fieldNameMap["preview_spec"] = "PreviewSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func previewOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PreviewBindingType)
}

func previewRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceInaccessible": 500})
}



func SourceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vapp_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"VirtualApp"}, ""))
    fieldNameMap["vapp_id"] = "VappId"
    fields["vm_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"VirtualMachine"}, ""))
    fieldNameMap["vm_id"] = "VmId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.source_info",fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["target_type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.export_session.target_type", reflect.TypeOf(TargetType(TargetType_DOWNLOAD_TARGET)))
    fieldNameMap["target_type"] = "TargetType"
    fields["target_content_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.export_session.target_content_type", reflect.TypeOf(TargetContentType(TargetContentType_OVF_TARGET))))
    fieldNameMap["target_content_type"] = "TargetContentType"
    fields["export_flags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["export_flags"] = "ExportFlags"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["content_library"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""))
    fieldNameMap["content_library"] = "ContentLibrary"
    fields["content_library_item"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["content_library_item"] = "ContentLibraryItem"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("target_type",
        map[string][]bindings.FieldData {
            "DOWNLOAD_TARGET": []bindings.FieldData {
                 bindings.NewFieldData("target_content_type", true),
            },
            "CONTENT_LIBRARY_TARGET": []bindings.FieldData {
                 bindings.NewFieldData("content_library", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("target_content_type",
        map[string][]bindings.FieldData {
            "OVF_TARGET": []bindings.FieldData {},
            "OVA_TARGET": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.ovf.export_session.state", reflect.TypeOf(State(State_EXPORT_PREPARING)))
    fieldNameMap["state"] = "State"
    fields["progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["progress"] = "Progress"
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfFileInfoBindingType), reflect.TypeOf([]ovf.OvfFileInfo{}))
    fieldNameMap["files"] = "Files"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfErrorBindingType), reflect.TypeOf([]ovf.OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfWarningBindingType), reflect.TypeOf([]ovf.OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfInfoBindingType), reflect.TypeOf([]ovf.OvfInfo{}))
    fieldNameMap["information"] = "Information"
    fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["library_item_id"] = "LibraryItemId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func PreviewSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["export_flags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["export_flags"] = "ExportFlags"
    fields["target_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.export_session.target_type", reflect.TypeOf(TargetType(TargetType_DOWNLOAD_TARGET))))
    fieldNameMap["target_type"] = "TargetType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.preview_spec",fields, reflect.TypeOf(PreviewSpec{}), fieldNameMap, validators)
}

func PreviewFileBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.preview_file",fields, reflect.TypeOf(PreviewFile{}), fieldNameMap, validators)
}

func PreviewBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(PreviewFileBindingType), reflect.TypeOf([]PreviewFile{}))
    fieldNameMap["files"] = "Files"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.preview",fields, reflect.TypeOf(Preview{}), fieldNameMap, validators)
}


