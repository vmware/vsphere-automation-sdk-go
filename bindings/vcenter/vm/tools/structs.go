/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tools.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tools

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// Current run state of VMware Tools in the guest operating system.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type RunState string

const (
    // VMware Tools is not running.
     RunState_NOT_RUNNING RunState = "NOT_RUNNING"
    // VMware Tools is running.
     RunState_RUNNING RunState = "RUNNING"
    // VMware Tools is running scripts as part of a state transition.
     RunState_EXECUTING_SCRIPTS RunState = "EXECUTING_SCRIPTS"
)

func (r RunState) RunState() bool {
    switch r {
        case RunState_NOT_RUNNING:
            return true
        case RunState_RUNNING:
            return true
        case RunState_EXECUTING_SCRIPTS:
            return true
        default:
            return false
    }
}




// The ``UpgradePolicy`` enumeration class defines when Tools are auto-upgraded for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UpgradePolicy string

const (
    // No auto-upgrades for Tools will be performed for this virtual machine. Users must manually invoke the Tools#upgrade method to update Tools.
     UpgradePolicy_MANUAL UpgradePolicy = "MANUAL"
    // When the virtual machine is power-cycled, the system checks for a newer version of Tools when the virtual machine is powered on. If it is available, a Tools upgrade is automatically performed on the virtual machine and it is rebooted if necessary.
     UpgradePolicy_UPGRADE_AT_POWER_CYCLE UpgradePolicy = "UPGRADE_AT_POWER_CYCLE"
)

func (u UpgradePolicy) UpgradePolicy() bool {
    switch u {
        case UpgradePolicy_MANUAL:
            return true
        case UpgradePolicy_UPGRADE_AT_POWER_CYCLE:
            return true
        default:
            return false
    }
}




// The ``VersionStatus`` enumeration class defines the version status types of VMware Tools installed in the guest operating system.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type VersionStatus string

const (
    // VMware Tools has never been installed.
     VersionStatus_NOT_INSTALLED VersionStatus = "NOT_INSTALLED"
    // VMware Tools is installed, and the version is current.
     VersionStatus_CURRENT VersionStatus = "CURRENT"
    // VMware Tools is installed, but it is not managed by VMware. This includes open-vm-tools or OSPs which should be managed by the guest operating system.
     VersionStatus_UNMANAGED VersionStatus = "UNMANAGED"
    // VMware Tools is installed, but the version is too old.
     VersionStatus_TOO_OLD_UNSUPPORTED VersionStatus = "TOO_OLD_UNSUPPORTED"
    // VMware Tools is installed, supported, but a newer version is available.
     VersionStatus_SUPPORTED_OLD VersionStatus = "SUPPORTED_OLD"
    // VMware Tools is installed, supported, and newer than the version available on the host.
     VersionStatus_SUPPORTED_NEW VersionStatus = "SUPPORTED_NEW"
    // VMware Tools is installed, and the version is known to be too new to work correctly with this virtual machine.
     VersionStatus_TOO_NEW VersionStatus = "TOO_NEW"
    // VMware Tools is installed, but the installed version is known to have a grave bug and should be immediately upgraded.
     VersionStatus_BLACKLISTED VersionStatus = "BLACKLISTED"
)

func (v VersionStatus) VersionStatus() bool {
    switch v {
        case VersionStatus_NOT_INSTALLED:
            return true
        case VersionStatus_CURRENT:
            return true
        case VersionStatus_UNMANAGED:
            return true
        case VersionStatus_TOO_OLD_UNSUPPORTED:
            return true
        case VersionStatus_SUPPORTED_OLD:
            return true
        case VersionStatus_SUPPORTED_NEW:
            return true
        case VersionStatus_TOO_NEW:
            return true
        case VersionStatus_BLACKLISTED:
            return true
        default:
            return false
    }
}




// The ``ToolsInstallType`` enumeration class defines the installation type of the Tools in the guest operating system.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ToolsInstallType string

const (
    // Installation type is not known. Most likely tools have been installed by OSPs or open-vm-tools, but a version that does not report its install type or an install type that we do not recognize.
     ToolsInstallType_UNKNOWN ToolsInstallType = "UNKNOWN"
    // MSI is the installation type used for VMware Tools on Windows.
     ToolsInstallType_MSI ToolsInstallType = "MSI"
    // Tools have been installed by the tar installer.
     ToolsInstallType_TAR ToolsInstallType = "TAR"
    // OSPs are RPM or Debian packages tailored for the OS in the VM. See http://packages.vmware.com
     ToolsInstallType_OSP ToolsInstallType = "OSP"
    // open-vm-tools are the open-source version of VMware Tools, may have been packaged by the OS vendor.
     ToolsInstallType_OPEN_VM_TOOLS ToolsInstallType = "OPEN_VM_TOOLS"
)

func (t ToolsInstallType) ToolsInstallType() bool {
    switch t {
        case ToolsInstallType_UNKNOWN:
            return true
        case ToolsInstallType_MSI:
            return true
        case ToolsInstallType_TAR:
            return true
        case ToolsInstallType_OSP:
            return true
        case ToolsInstallType_OPEN_VM_TOOLS:
            return true
        default:
            return false
    }
}





// The ``Info`` class describes the VMWare Tools properties of a virtual machine.
 type Info struct {
    AutoUpdateSupported bool
    InstallAttemptCount *int64
    Error *data.ErrorValue
    VersionNumber *int64
    Version *string
    UpgradePolicy UpgradePolicy
    VersionStatus *VersionStatus
    InstallType *ToolsInstallType
    RunState RunState
}






// The (\\\\@name UpdateSpec} class describes the VMware Tools properties of a virtual machine that can be updated.
 type UpdateSpec struct {
    UpgradePolicy *UpgradePolicy
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
       map[string]int{"Error": 500,"NotFound": 404})
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
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404})
}


func upgradeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["command_line_options"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["command_line_options"] = "CommandLineOptions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func upgradeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"Error": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auto_update_supported"] = bindings.NewBooleanType()
    fieldNameMap["auto_update_supported"] = "AutoUpdateSupported"
    fields["install_attempt_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["install_attempt_count"] = "InstallAttemptCount"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["version_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["version_number"] = "VersionNumber"
    fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["version"] = "Version"
    fields["upgrade_policy"] = bindings.NewEnumType("com.vmware.vcenter.vm.tools.upgrade_policy", reflect.TypeOf(UpgradePolicy(UpgradePolicy_MANUAL)))
    fieldNameMap["upgrade_policy"] = "UpgradePolicy"
    fields["version_status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.tools.version_status", reflect.TypeOf(VersionStatus(VersionStatus_NOT_INSTALLED))))
    fieldNameMap["version_status"] = "VersionStatus"
    fields["install_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.tools.tools_install_type", reflect.TypeOf(ToolsInstallType(ToolsInstallType_UNKNOWN))))
    fieldNameMap["install_type"] = "InstallType"
    fields["run_state"] = bindings.NewEnumType("com.vmware.vcenter.vm.tools.run_state", reflect.TypeOf(RunState(RunState_NOT_RUNNING)))
    fieldNameMap["run_state"] = "RunState"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.tools.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["upgrade_policy"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.tools.upgrade_policy", reflect.TypeOf(UpgradePolicy(UpgradePolicy_MANUAL))))
    fieldNameMap["upgrade_policy"] = "UpgradePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.tools.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


