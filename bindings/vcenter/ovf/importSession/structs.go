/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ImportSession.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package importSession

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/ovf"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
)



// The ``State`` enumeration class defines the different states of a transfer. 
//
//  For pull imports files are pulled by the server using HTTP GET, for push imports the client uses HTTP PUT to upload the data. For a content library import, the server retrieves the files from the content library. 
//
//  The transition diagram is as follows for a pull or content library import: 
//
// * State#State_IMPORT_OVF_TRANSFER: Initial state, where server is retrieving OVF descriptor, message bundles, certificates, and manifests.
// * State#State_IMPORT_SELECTING_OVF_PARAMS: Waiting for user to call Instantiate to begin the instantiation.
// * State#State_IMPORT_FILE_TRANSFER: Transfering files from the source.
// * State#State_IMPORT_INSTANTIATING: Instantiating the virtual machines/virtual appliance.
// * State#State_IMPORT_COMPLETED: Import completed successfully.
//
//  
//
//  The transition diagram is as follows for a push import: 
//
// * State#State_IMPORT_OVF_TRANSFER: Initial state where server is waiting for the OVF descriptor to be uploaded.
// * State#State_IMPORT_MSG_BUNDLES_TRANSFER: Server is waiting for message bundles to be uploaded.
// * State#State_IMPORT_SELECTING_OVF_PARAMS: Waiting for user to begin the instantiation.
// * State#State_IMPORT_FILE_TRANSFER: Server is waiting for all files to be uploaded.
// * State#State_IMPORT_INSTANTIATING: Instantiating the virtual machines/virtual appliance.
// * State#State_IMPORT_COMPLETED: Import completed successfully.
//
//  
//
//  The client is allowed to upload files as soon as the appear in the import session object. Thus upload can begin before reaching the State#State_IMPORT_FILE_TRANSFER state.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type State string

const (
    // State of an import transfer that does not have any files available. The transfer needs the OVF descriptor to continue. If this is a push transfer, the client must upload the OVF descriptor, and the transfer file list has one file info entry with a URL to which the client must upload the OVF descriptor using HTTP PUT. For pull transfers (including content library), the server is in the process of retrieving the OVF descriptor. 
    //
    //  Transition to the next state is done when the server has retrieved the complete OVF content and parsed it.
     State_IMPORT_OVF_TRANSFER State = "IMPORT_OVF_TRANSFER"
    // The file list contains a number of message bundles that need to be transferred to the server. If this is a push transfer, the client must PUT the requested files to the server. 
    //
    //  In case the OVF descriptor does not specify any bundles this state is skipped. 
    //
    //  Transition to next state is done when the complete content of all message bundles has been retrieved by the server.
     State_IMPORT_MSG_BUNDLES_TRANSFER State = "IMPORT_MSG_BUNDLES_TRANSFER"
    // The server can be queried for OVF parameters, and the client can specify instantiation parameters. 
    //
    //  Specifying an OVF instantiation parameter might affect other OVF instantiation parameters and change the set of files that needs to get transferred. 
    //
    //  During this state the client is allowed to push other files to the server with HTTP PUT. 
    //
    //  Transition to the next state is done by calling the instantiate method. If all needed files have been transferred to the server a transition is made to State#State_IMPORT_INSTANTIATING, otherwise a transition is made to State#State_IMPORT_FILE_TRANSFER.
     State_IMPORT_SELECTING_OVF_PARAMS State = "IMPORT_SELECTING_OVF_PARAMS"
    // State for transferring remaining files that have not been transferred during an earlier state. 
    //
    //  When the content of all files has been transferred the state is changed to State#State_IMPORT_INSTANTIATING: 
    //
    // * Pull transfer: All needed files have been completely transferred, and the manifest and certificate have been transferred if they are available.
    // * Push transfer: All files marked as required (attribute optional == false) have been transferred and all initiated transfers of optional files are complete. In particular the manifest and certificate are optional so the upload of these has to at least be initiated before other files are completely uploaded.
     State_IMPORT_FILE_TRANSFER State = "IMPORT_FILE_TRANSFER"
    // The virtual machine or virtual appliance is being instantiated.
     State_IMPORT_INSTANTIATING State = "IMPORT_INSTANTIATING"
    // The virtual machine or virtual appliance is instantiated, and the upload transfer is done.
     State_IMPORT_COMPLETED State = "IMPORT_COMPLETED"
    // The transfer failed.
     State_IMPORT_ERROR State = "IMPORT_ERROR"
)

func (s State) State() bool {
    switch s {
        case State_IMPORT_OVF_TRANSFER:
            return true
        case State_IMPORT_MSG_BUNDLES_TRANSFER:
            return true
        case State_IMPORT_SELECTING_OVF_PARAMS:
            return true
        case State_IMPORT_FILE_TRANSFER:
            return true
        case State_IMPORT_INSTANTIATING:
            return true
        case State_IMPORT_COMPLETED:
            return true
        case State_IMPORT_ERROR:
            return true
        default:
            return false
    }
}




// The ``SourceType`` enumeration class defines the type of source.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SourceType string

const (
    // The client is uploading content using HTTP(S) PUT requests.
     SourceType_PUSH_SOURCE SourceType = "PUSH_SOURCE"
    // The server is pulling content from a URL.
     SourceType_PULL_SOURCE SourceType = "PULL_SOURCE"
    // The server is pulling content from a library item.
     SourceType_CONTENT_LIBRARY_SOURCE SourceType = "CONTENT_LIBRARY_SOURCE"
)

func (s SourceType) SourceType() bool {
    switch s {
        case SourceType_PUSH_SOURCE:
            return true
        case SourceType_PULL_SOURCE:
            return true
        case SourceType_CONTENT_LIBRARY_SOURCE:
            return true
        default:
            return false
    }
}




// The ``PushSourceContentType`` enumeration class defines the content type of source by push.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type PushSourceContentType string

const (
    // The client is uploading an OVF template.
     PushSourceContentType_OVF_SOURCE PushSourceContentType = "OVF_SOURCE"
    // The client is uploading an OVA template.
     PushSourceContentType_OVA_SOURCE PushSourceContentType = "OVA_SOURCE"
)

func (p PushSourceContentType) PushSourceContentType() bool {
    switch p {
        case PushSourceContentType_OVF_SOURCE:
            return true
        case PushSourceContentType_OVA_SOURCE:
            return true
        default:
            return false
    }
}




// The ``PushSourceOvfOption`` class specifies the optional information that the OVF template provides. This only applies to PushSourceContentType#PushSourceContentType_OVF_SOURCE.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type PushSourceOvfOption string

const (
    // Indicates that manifest file will be provided when pushing the OVF template.
     PushSourceOvfOption_MANIFEST PushSourceOvfOption = "MANIFEST"
    // Indicates that manifest and certificate file will be provided when pushing the OVF template.
     PushSourceOvfOption_MANIFEST_CERTIFICATE PushSourceOvfOption = "MANIFEST_CERTIFICATE"
    // Indicates that neither manifest nor certificate file will be provided when pushing the OVF template.
     PushSourceOvfOption_NONE PushSourceOvfOption = "NONE"
)

func (p PushSourceOvfOption) PushSourceOvfOption() bool {
    switch p {
        case PushSourceOvfOption_MANIFEST:
            return true
        case PushSourceOvfOption_MANIFEST_CERTIFICATE:
            return true
        case PushSourceOvfOption_NONE:
            return true
        default:
            return false
    }
}




// The ``FileTransferMode`` class specifies the file transfer mode for disk files during the OVF template deployment.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type FileTransferMode string

const (
    // Indicates that the server will directly handle all file transfers. This is the default transfer mode.
     FileTransferMode_SERVER_TRANSFER_MODE FileTransferMode = "SERVER_TRANSFER_MODE"
    // Indicates that the server delegates files transfer to external services. In this mode, if the CreateSpec#sourceType is {SourceType#SourceType_PUSH_SOURCE, client should call ImportSession#progress to update file transfer progress, besides pushing files to ovf.OvfFileInfo#fileUrl.
     FileTransferMode_EXTERNAL_TRANSFER_MODE FileTransferMode = "EXTERNAL_TRANSFER_MODE"
)

func (f FileTransferMode) FileTransferMode() bool {
    switch f {
        case FileTransferMode_SERVER_TRANSFER_MODE:
            return true
        case FileTransferMode_EXTERNAL_TRANSFER_MODE:
            return true
        default:
            return false
    }
}





// The ``CreateSpec`` class contains information about the import parameters.
 type CreateSpec struct {
    Locale *string
    ImportFlags []string
    SourceType SourceType
    PullSource *url.URL
    SslCertificateThumbprint *string
    PushSourceContentType *PushSourceContentType
    PushSourceOvfOption *PushSourceOvfOption
    ContentLibraryItem *string
}






// The ``OvfValidationResult`` class contains information about the result of calling tryInstantiate.
 type OvfValidationResult struct {
    Parameters []*data.StructValue
    Errors []ovf.OvfError
    Warnings []ovf.OvfWarning
    Information []ovf.OvfInfo
}






// The ``Info`` class represents an import session.
 type Info struct {
    State State
    Progress int64
    FileTransferMode *FileTransferMode
    Files []ovf.OvfFileInfo
    Errors []ovf.OvfError
    Warnings []ovf.OvfWarning
    Information []ovf.OvfInfo
    VappId *string
    VmId *string
    CustomizationResults map[string]bool
}






// The ``PreviewFile`` class contains information about a file that is referenced in the OVF descriptor and will be imported according to an import preview.
 type PreviewFile struct {
    Name string
}






// The ``Preview`` class contains information about the result of an OVF import preview based on OVF descriptor, which includes referenced files and OVF validation information.
 type Preview struct {
    Files []PreviewFile
    Errors []ovf.OvfError
    Warnings []ovf.OvfWarning
    Information []ovf.OvfInfo
}






// The ``ProbeResult`` class contains information about the accessibility of a pull source URI.
 type ProbeResult struct {
    Status ProbeResult_Status
    SslThumbprint *string
    ErrorMessages []std.LocalizableMessage
}




    
    // The ``Status`` enumeration class defines the possible status values from an attempt to access a pull source URI.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ProbeResult_Status string

    const (
        // Indicates that the probe was successful.
         ProbeResult_Status_SUCCESS ProbeResult_Status = "SUCCESS"
        // Indicates that the supplied URL was not valid.
         ProbeResult_Status_INVALID_URL ProbeResult_Status = "INVALID_URL"
        // Indicates that the probe timed out while attempting to connect to the URL.
         ProbeResult_Status_TIMED_OUT ProbeResult_Status = "TIMED_OUT"
        // Indicates that the host in the URL could not be found.
         ProbeResult_Status_HOST_NOT_FOUND ProbeResult_Status = "HOST_NOT_FOUND"
        // Indicates that the provided server certificate thumbprint is invalid. In this case, the returned certificate thumbprint should be provided.
         ProbeResult_Status_CERTIFICATE_ERROR ProbeResult_Status = "CERTIFICATE_ERROR"
        // Indicates an unspecified error different from the other error cases defined in ProbeResult_Status.
         ProbeResult_Status_UNKNOWN_ERROR ProbeResult_Status = "UNKNOWN_ERROR"
    )

    func (s ProbeResult_Status) ProbeResult_Status() bool {
        switch s {
            case ProbeResult_Status_SUCCESS:
                return true
            case ProbeResult_Status_INVALID_URL:
                return true
            case ProbeResult_Status_TIMED_OUT:
                return true
            case ProbeResult_Status_HOST_NOT_FOUND:
                return true
            case ProbeResult_Status_CERTIFICATE_ERROR:
                return true
            case ProbeResult_Status_UNKNOWN_ERROR:
                return true
            default:
                return false
        }
    }






func createForResourcePoolInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["resource_pool"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fields["host_system"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fields["create_spec"] = bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(CreateSpecBindingType),}, bindings.JSONRPC)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["resource_pool"] = "ResourcePool"
    fieldNameMap["host_system"] = "HostSystem"
    fieldNameMap["folder"] = "Folder"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createForResourcePoolOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
}

func createForResourcePoolRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
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


func tryInstantiateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
    fields["instantiation_parameters"] = bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ovf.OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{}))
    fieldNameMap["id"] = "Id"
    fieldNameMap["instantiation_parameters"] = "InstantiationParameters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tryInstantiateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(OvfValidationResultBindingType)
}

func tryInstantiateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"Unauthorized": 403})
}


func instantiateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
    fields["instantiation_parameters"] = bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ovf.OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{}))
    fieldNameMap["id"] = "Id"
    fieldNameMap["instantiation_parameters"] = "InstantiationParameters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instantiateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func instantiateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"InvalidArgument": 400,"Unsupported": 400,"Unauthorized": 403})
}


func progressInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
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


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
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


func previewInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ovf_descriptor"] = bindings.NewStringType()
    fieldNameMap["ovf_descriptor"] = "OvfDescriptor"
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
       map[string]int{"InvalidArgument": 400})
}


func probeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["uri"] = bindings.NewUriType()
    fields["ssl_certificate_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["uri"] = "Uri"
    fieldNameMap["ssl_certificate_thumbprint"] = "SslCertificateThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func probeOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProbeResultBindingType)
}

func probeRestMetadata() protocol.OperationRestMetadata {
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



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["locale"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["locale"] = "Locale"
    fields["import_flags"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["import_flags"] = "ImportFlags"
    fields["source_type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.source_type", reflect.TypeOf(SourceType(SourceType_PUSH_SOURCE)))
    fieldNameMap["source_type"] = "SourceType"
    fields["pull_source"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["pull_source"] = "PullSource"
    fields["ssl_certificate_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_certificate_thumbprint"] = "SslCertificateThumbprint"
    fields["push_source_content_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.push_source_content_type", reflect.TypeOf(PushSourceContentType(PushSourceContentType_OVF_SOURCE))))
    fieldNameMap["push_source_content_type"] = "PushSourceContentType"
    fields["push_source_ovf_option"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.push_source_ovf_option", reflect.TypeOf(PushSourceOvfOption(PushSourceOvfOption_MANIFEST))))
    fieldNameMap["push_source_ovf_option"] = "PushSourceOvfOption"
    fields["content_library_item"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["content_library_item"] = "ContentLibraryItem"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("source_type",
        map[string][]bindings.FieldData {
            "PULL_SOURCE": []bindings.FieldData {
                 bindings.NewFieldData("pull_source", true),
                 bindings.NewFieldData("ssl_certificate_thumbprint", false),
            },
            "PUSH_SOURCE": []bindings.FieldData {
                 bindings.NewFieldData("push_source_content_type", false),
            },
            "CONTENT_LIBRARY_SOURCE": []bindings.FieldData {
                 bindings.NewFieldData("content_library_item", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("push_source_content_type",
        map[string][]bindings.FieldData {
            "OVF_SOURCE": []bindings.FieldData {
                 bindings.NewFieldData("push_source_ovf_option", false),
            },
            "OVA_SOURCE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func OvfValidationResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["parameters"] = bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ovf.OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{}))
    fieldNameMap["parameters"] = "Parameters"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfErrorBindingType), reflect.TypeOf([]ovf.OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfWarningBindingType), reflect.TypeOf([]ovf.OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfInfoBindingType), reflect.TypeOf([]ovf.OvfInfo{}))
    fieldNameMap["information"] = "Information"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.ovf_validation_result",fields, reflect.TypeOf(OvfValidationResult{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.state", reflect.TypeOf(State(State_IMPORT_OVF_TRANSFER)))
    fieldNameMap["state"] = "State"
    fields["progress"] = bindings.NewIntegerType()
    fieldNameMap["progress"] = "Progress"
    fields["file_transfer_mode"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.file_transfer_mode", reflect.TypeOf(FileTransferMode(FileTransferMode_SERVER_TRANSFER_MODE))))
    fieldNameMap["file_transfer_mode"] = "FileTransferMode"
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfFileInfoBindingType), reflect.TypeOf([]ovf.OvfFileInfo{}))
    fieldNameMap["files"] = "Files"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfErrorBindingType), reflect.TypeOf([]ovf.OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfWarningBindingType), reflect.TypeOf([]ovf.OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfInfoBindingType), reflect.TypeOf([]ovf.OvfInfo{}))
    fieldNameMap["information"] = "Information"
    fields["vapp_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"VirtualApp"}, ""))
    fieldNameMap["vapp_id"] = "VappId"
    fields["vm_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"VirtualMachine"}, ""))
    fieldNameMap["vm_id"] = "VmId"
    fields["customization_results"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewBooleanType(),reflect.TypeOf(map[string]bool{})))
    fieldNameMap["customization_results"] = "CustomizationResults"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("state",
        map[string][]bindings.FieldData {
            "IMPORT_FILE_TRANSFER": []bindings.FieldData {
                 bindings.NewFieldData("file_transfer_mode", false),
            },
            "IMPORT_INSTANTIATING": []bindings.FieldData {
                 bindings.NewFieldData("file_transfer_mode", false),
            },
            "IMPORT_COMPLETED": []bindings.FieldData {
                 bindings.NewFieldData("file_transfer_mode", false),
            },
            "IMPORT_ERROR": []bindings.FieldData {
                 bindings.NewFieldData("file_transfer_mode", false),
            },
            "IMPORT_OVF_TRANSFER": []bindings.FieldData {},
            "IMPORT_MSG_BUNDLES_TRANSFER": []bindings.FieldData {},
            "IMPORT_SELECTING_OVF_PARAMS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func PreviewFileBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.preview_file",fields, reflect.TypeOf(PreviewFile{}), fieldNameMap, validators)
}

func PreviewBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(PreviewFileBindingType), reflect.TypeOf([]PreviewFile{}))
    fieldNameMap["files"] = "Files"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfErrorBindingType), reflect.TypeOf([]ovf.OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfWarningBindingType), reflect.TypeOf([]ovf.OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfInfoBindingType), reflect.TypeOf([]ovf.OvfInfo{}))
    fieldNameMap["information"] = "Information"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.preview",fields, reflect.TypeOf(Preview{}), fieldNameMap, validators)
}

func ProbeResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.probe_result.status", reflect.TypeOf(ProbeResult_Status(ProbeResult_Status_SUCCESS)))
    fieldNameMap["status"] = "Status"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["error_messages"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["error_messages"] = "ErrorMessages"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "CERTIFICATE_ERROR": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
                 bindings.NewFieldData("error_messages", false),
            },
            "SUCCESS": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
            },
            "UNKNOWN_ERROR": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
                 bindings.NewFieldData("error_messages", false),
            },
            "INVALID_URL": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", false),
            },
            "TIMED_OUT": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", false),
            },
            "HOST_NOT_FOUND": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.probe_result",fields, reflect.TypeOf(ProbeResult{}), fieldNameMap, validators)
}


