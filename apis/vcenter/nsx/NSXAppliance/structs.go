/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: NSXAppliance.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package NSXAppliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/nsx"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Operation`` enumeration class describes the operation in progress. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Operation string

const (
    // Install the NSX appliance on the vCenter. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Operation_ENABLE Operation = "ENABLE"
    // Delete the NSX appliance on the vCenter. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Operation_DISABLE Operation = "DISABLE"
    // No ongoing operations. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Operation_NONE Operation = "NONE"
)

func (o Operation) Operation() bool {
    switch o {
        case Operation_ENABLE:
            return true
        case Operation_DISABLE:
            return true
        case Operation_NONE:
            return true
        default:
            return false
    }
}




// The ``Phase`` enumeration class represents the current phase of the object with respect to realizing the deployment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Phase string

const (
    // Precheck phase validates if the NSX deployment can complete successfully. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Phase_PRECHECK Phase = "PRECHECK"
    // NSX ovf deployment phase. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Phase_DEPLOY Phase = "DEPLOY"
    // Configuration phase of the deployed NSX appliance. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Phase_CONFIGURE Phase = "CONFIGURE"
)

func (p Phase) Phase() bool {
    switch p {
        case Phase_PRECHECK:
            return true
        case Phase_DEPLOY:
            return true
        case Phase_CONFIGURE:
            return true
        default:
            return false
    }
}




// The ``ConfigStatus`` enumeration class describes the status of NSX deployment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConfigStatus string

const (
    // NSX is not installed on the vCenter Server. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_UNCONFIGURED ConfigStatus = "UNCONFIGURED"
    // NSX is successfully configured on the vCenter Server. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_CONFIGURED ConfigStatus = "CONFIGURED"
    // NSX deployment is in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_INPROGRESS ConfigStatus = "INPROGRESS"
    // NSX deployment is failed. Retry the deployment after resolving the errors. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_ERROR ConfigStatus = "ERROR"
)

func (c ConfigStatus) ConfigStatus() bool {
    switch c {
        case ConfigStatus_UNCONFIGURED:
            return true
        case ConfigStatus_CONFIGURED:
            return true
        case ConfigStatus_INPROGRESS:
            return true
        case ConfigStatus_ERROR:
            return true
        default:
            return false
    }
}




// The ``ProductType`` enumeration class describes the type of NSX deployment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ProductType string

const (
    // Integrated NSX shipped with vSphere. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ProductType_NSXI ProductType = "NSXI"
    // Licensed version of NSX. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ProductType_NSXT ProductType = "NSXT"
)

func (p ProductType) ProductType() bool {
    switch p {
        case ProductType_NSXI:
            return true
        case ProductType_NSXT:
            return true
        default:
            return false
    }
}





// The ``Message`` class contains the information about the ongoing deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Message struct {
    Severity Message_Severity
    Details *std.LocalizableMessage
}




    
    // The ``Severity`` enumeration class represents the severity of the message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Message_Severity string

    const (
        // Informational message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_INFO Message_Severity = "INFO"
        // Warning message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_WARNING Message_Severity = "WARNING"
        // Error message. This is accompanied by vCenter event and/or alarm. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_ERROR Message_Severity = "ERROR"
    )

    func (s Message_Severity) Message_Severity() bool {
        switch s {
            case Message_Severity_INFO:
                return true
            case Message_Severity_WARNING:
                return true
            case Message_Severity_ERROR:
                return true
            default:
                return false
        }
    }



// The ``Info`` class contains detailed information about the status of NSX deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    ConfigStatus ConfigStatus
    Operation Operation
    Phase *Phase
    ProductType *ProductType
    Messages []Message
}






// The ``InstallSpec`` contains the inputs related to appliance deployment operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstallSpec struct {
    DestinationLocation nsx.PlacementDetails
    DestinationAppliance nsx.ApplianceConfig
    ManagementVcenter *nsx.Connection
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
      "GET",
      "/vcenter/nsx/nsx-appliance",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(InstallSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(InstallSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/nsx/nsx-appliance",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"AlreadyExists": 400})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
      "DELETE",
      "/vcenter/nsx/nsx-appliance",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}



func MessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.message.severity", reflect.TypeOf(Message_Severity(Message_Severity_INFO)))
    fieldNameMap["severity"] = "Severity"
    fields["details"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.NSX_appliance.message",fields, reflect.TypeOf(Message{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.config_status", reflect.TypeOf(ConfigStatus(ConfigStatus_UNCONFIGURED)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["operation"] = bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.operation", reflect.TypeOf(Operation(Operation_ENABLE)))
    fieldNameMap["operation"] = "Operation"
    fields["phase"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.phase", reflect.TypeOf(Phase(Phase_PRECHECK))))
    fieldNameMap["phase"] = "Phase"
    fields["product_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.product_type", reflect.TypeOf(ProductType(ProductType_NSXI))))
    fieldNameMap["product_type"] = "ProductType"
    fields["messages"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{})))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.NSX_appliance.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func InstallSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination_location"] = bindings.NewReferenceType(nsx.PlacementDetailsBindingType)
    fieldNameMap["destination_location"] = "DestinationLocation"
    fields["destination_appliance"] = bindings.NewReferenceType(nsx.ApplianceConfigBindingType)
    fieldNameMap["destination_appliance"] = "DestinationAppliance"
    fields["management_vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(nsx.ConnectionBindingType))
    fieldNameMap["management_vcenter"] = "ManagementVcenter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.NSX_appliance.install_spec",fields, reflect.TypeOf(InstallSpec{}), fieldNameMap, validators)
}


