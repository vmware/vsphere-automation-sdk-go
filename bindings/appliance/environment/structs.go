/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Environment.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package environment

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``ManagedBy`` enumeration class defines the possible entities managing the appliance.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ManagedBy string

const (
    // Managed by VMware
     ManagedBy_VMWARE ManagedBy = "VMWARE"
    // Managed by the user
     ManagedBy_USER ManagedBy = "USER"
)

func (m ManagedBy) ManagedBy() bool {
    switch m {
        case ManagedBy_VMWARE:
            return true
        case ManagedBy_USER:
            return true
        default:
            return false
    }
}




// The ``DeployedBy`` enumeration class defines the possible entities deploying the appliance.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type DeployedBy string

const (
    // Deployed by VMware
     DeployedBy_VMWARE DeployedBy = "VMWARE"
    // Deployed by the user
     DeployedBy_USER DeployedBy = "USER"
)

func (d DeployedBy) DeployedBy() bool {
    switch d {
        case DeployedBy_VMWARE:
            return true
        case DeployedBy_USER:
            return true
        default:
            return false
    }
}




// The ``Provider`` enumeration class defines the possible providers. Now it has only AWS, in the future we foresee soft-layer, google cloud, azure, etc.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Provider string

const (
    // The appliance is located in AWS-backed VMC
     Provider_AWS Provider = "AWS"
    // The appliance location is UNKNOWN
     Provider_UNKNOWN Provider = "UNKNOWN"
)

func (p Provider) Provider() bool {
    switch p {
        case Provider_AWS:
            return true
        case Provider_UNKNOWN:
            return true
        default:
            return false
    }
}





// The ``Display`` class contains information what needs to be displayed in the user interface.
 type Display struct {
    DatabaseMonitoring bool
    SsoStatus bool
}






// The ``Info`` class contains the information about the appliance environment.
 type Info struct {
    ManagedBy ManagedBy
    DeployedBy DeployedBy
    Provider Provider
    Display Display
}






// The ``Config`` class describes the configurable settings for the appliance environment.
 type Config struct {
    ManagedBy ManagedBy
    DeployedBy DeployedBy
    Provider Provider
    Display Display
}









func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(ConfigBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
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
       map[string]int{"Error": 500})
}



func DisplayBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["database_monitoring"] = bindings.NewBooleanType()
    fieldNameMap["database_monitoring"] = "DatabaseMonitoring"
    fields["sso_status"] = bindings.NewBooleanType()
    fieldNameMap["sso_status"] = "SsoStatus"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.environment.display",fields, reflect.TypeOf(Display{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["managed_by"] = bindings.NewEnumType("com.vmware.appliance.environment.managed_by", reflect.TypeOf(ManagedBy(ManagedBy_VMWARE)))
    fieldNameMap["managed_by"] = "ManagedBy"
    fields["deployed_by"] = bindings.NewEnumType("com.vmware.appliance.environment.deployed_by", reflect.TypeOf(DeployedBy(DeployedBy_VMWARE)))
    fieldNameMap["deployed_by"] = "DeployedBy"
    fields["provider"] = bindings.NewEnumType("com.vmware.appliance.environment.provider", reflect.TypeOf(Provider(Provider_AWS)))
    fieldNameMap["provider"] = "Provider"
    fields["display"] = bindings.NewReferenceType(DisplayBindingType)
    fieldNameMap["display"] = "Display"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.environment.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["managed_by"] = bindings.NewEnumType("com.vmware.appliance.environment.managed_by", reflect.TypeOf(ManagedBy(ManagedBy_VMWARE)))
    fieldNameMap["managed_by"] = "ManagedBy"
    fields["deployed_by"] = bindings.NewEnumType("com.vmware.appliance.environment.deployed_by", reflect.TypeOf(DeployedBy(DeployedBy_VMWARE)))
    fieldNameMap["deployed_by"] = "DeployedBy"
    fields["provider"] = bindings.NewEnumType("com.vmware.appliance.environment.provider", reflect.TypeOf(Provider(Provider_AWS)))
    fieldNameMap["provider"] = "Provider"
    fields["display"] = bindings.NewReferenceType(DisplayBindingType)
    fieldNameMap["display"] = "Display"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.environment.config",fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}


