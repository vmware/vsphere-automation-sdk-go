/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Policies.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policies

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/compute/policies"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Info`` class contains information about the compliance of a virtual machine with a compute policy. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Info struct {
    Status policies.ObjectCompliance
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.Resources.COMPUTE_POLICY"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["policy"] = "Policy"
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
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.Resources.COMPUTE_POLICY"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.Resources.COMPUTE_POLICY"}, "")
    pathParams["vm"] = "vm"
    pathParams["policy"] = "policy"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/vm/{vm}/compute/policies/{policy}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.compute.policies.object_compliance", reflect.TypeOf(policies.ObjectCompliance(policies.ObjectCompliance_UNKNOWN)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.compute.policies.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


