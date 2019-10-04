/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VcenterAppliance.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenterAppliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/nsx"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``ApplianceInfo`` contains information related to the inventory location and configuration for vCenter appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ApplianceInfo struct {
    Configuration nsx.ApplianceConfig
    Location nsx.PlacementDetails
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["management_vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(nsx.ConnectionBindingType))
    fieldNameMap["management_vcenter"] = "ManagementVcenter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ApplianceInfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["management_vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(nsx.ConnectionBindingType))
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "management_vcenter",
      "POST",
      "/vcenter/nsx/vcenter-appliance",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func ApplianceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["configuration"] = bindings.NewReferenceType(nsx.ApplianceConfigBindingType)
    fieldNameMap["configuration"] = "Configuration"
    fields["location"] = bindings.NewReferenceType(nsx.PlacementDetailsBindingType)
    fieldNameMap["location"] = "Location"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.vcenter_appliance.appliance_info",fields, reflect.TypeOf(ApplianceInfo{}), fieldNameMap, validators)
}


