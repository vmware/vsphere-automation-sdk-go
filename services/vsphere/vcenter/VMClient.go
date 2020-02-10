/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: VM
 * Used by client-side stubs.
 */

package vcenter

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/VM"
)

// The ``VM`` interface provides methods for managing the lifecycle of a virtual machine.
type VMClient interface {

    // Creates a virtual machine.
    //
    // @param specParam Virtual machine specification.
    // @return ID of newly-created virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists if a virtual machine with the specified name already exists.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec could not be found
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ResourceInUse if any of the specified storage addresses (eg. IDE, SATA, SCSI, NVMe) result in a storage address conflict.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to create the virtual machine could not be allocated.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if ``guestOS`` is not supported for the requested virtual hardware version and spec includes null properties that default to guest-specific values.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``Folder`` referenced by the property VMInventoryPlacementSpec#folder requires ``VirtualMachine.Inventory.Create``.
    // * The resource ``ResourcePool`` referenced by the property VMComputePlacementSpec#resourcePool requires ``Resource.AssignVMToPool``.
    // * The resource ``Datastore`` referenced by the property VMStoragePlacementSpec#datastore requires ``Datastore.AllocateSpace``.
    // * The resource ``Network`` referenced by the property hardware.EthernetBackingSpec#network requires ``Network.Assign``.
	Create(specParam VM.VMCreateSpec) (string, error)

    // Creates a virtual machine from an existing virtual machine. 
    //
    // . This method was added in vSphere API 7.0.0.
    //
    // @param specParam Virtual machine clone specification.
    // @return ID of newly-created virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists if a virtual machine with the specified name already exists.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec could not be found
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to clone the virtual machine could not be allocated.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``Datastore`` referenced by the property VM.VMDiskCloneSpec#datastore requires ``Datastore.AllocateSpace``.
    // * The resource ``Datastore`` referenced by the property VMClonePlacementSpec#datastore requires ``Datastore.AllocateSpace``.
    // * The resource ``Folder`` referenced by the property VMClonePlacementSpec#folder requires ``VirtualMachine.Inventory.CreateFromExisting``.
    // * The resource ``ResourcePool`` referenced by the property VMClonePlacementSpec#resourcePool requires ``Resource.AssignVMToPool``.
    // * The resource ``VirtualMachine`` referenced by the property VM.VMCloneSpec#source requires ``VirtualMachine.Provisioning.Clone``.
	Clone(specParam VM.VMCloneSpec) (string, error)

    // Relocates a virtual machine based on the specification. The parts of the virtual machine that can move are: FOLDER, RESOURCE_POOL, HOST, CLUSTER and DATASTORE of home of the virtual machine and disks. 
    //
    // . This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Existing Virtual machine to relocate.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Relocate specification.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec or the given "vm" could not be found
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``Resource.ColdMigrate``.
    // * The resource ``ResourcePool`` referenced by the property VMRelocatePlacementSpec#resourcePool requires ``Resource.AssignVMToPool``.
	Relocate(vmParam string, specParam VM.VMRelocateSpec) error

    // Create an instant clone of an existing virtual machine. This method was added in vSphere API 6.7.1.
    //
    // @param specParam Virtual machine InstantCloneSpec.
    // @return ID of newly-created virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists if a virtual machine with the specified name already exists.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec could not be found
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to create an instant clone could not be allocated.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the property VMInstantCloneSpec#source requires ``VirtualMachine.Provisioning.Clone`` and ``VirtualMachine.Inventory.CreateFromExisting``.
    // * The resource ``Folder`` referenced by the property VMInstantClonePlacementSpec#folder requires ``VirtualMachine.Interact.PowerOn``.
    // * The resource ``ResourcePool`` referenced by the property VMInstantClonePlacementSpec#resourcePool requires ``Resource.AssignVMToPool``.
    // * The resource ``Datastore`` referenced by the property VMInstantClonePlacementSpec#datastore requires ``Datastore.AllocateSpace``.
    // * The resource ``Network`` referenced by the property hardware.EthernetBackingSpec#network requires ``Network.Assign``.
	InstantClone(specParam VMInstantCloneSpec) (string, error)

    // Returns information about a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the specified virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``System.Read``.
	Get(vmParam string) (VM.VMInfo, error)

    // Deletes a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the virtual machine is running (powered on).
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Inventory.Delete``.
	Delete(vmParam string) error

    // Returns information about at most 4000 visible (subject to permission checks) virtual machines in vCenter matching the VMFilterSpec.
    //
    // @param filterParam Specification of matching virtual machines for which information should be returned.
    // If null, the behavior is equivalent to a VMFilterSpec with all properties null which means all virtual machines match the filter.
    // @return Commonly used information about the virtual machines matching the VMFilterSpec.
    // @throws InvalidArgument if the VMFilterSpec#powerStates property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 4000 virtual machines match the VMFilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(filterParam *VMFilterSpec) ([]VMSummary, error)

    // Creates a virtual machine from existing virtual machine files on storage. This method was added in vSphere API 6.8.7.
    //
    // @param specParam Specification of the location of the virtual machine files and the placement of the new virtual machine.
    // @return Identifier of the newly-created virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists if a virtual machine with the specified name already exists or if a virtual machine using the specified virtual machine files already exists.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec could not be found.
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to register the virtual machine could not be allocated.
    // @throws Unauthenticated if the user cannot be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``Datastore`` referenced by the property VMRegisterSpec#datastore requires ``System.Read``.
    // * The resource ``Folder`` referenced by the property VMInventoryPlacementSpec#folder requires ``VirtualMachine.Inventory.Register``.
    // * The resource ``ResourcePool`` referenced by the property VMComputePlacementSpec#resourcePool requires ``Resource.AssignVMToPool``.
	Register(specParam VMRegisterSpec) (string, error)

    // Removes the virtual machine corresponding to ``vm`` from the vCenter inventory without removing any of the virtual machine's files from storage. All high-level information stored with the management server (ESXi or vCenter) is removed, including information such as statistics, resource pool association, permissions, and alarms. This method was added in vSphere API 6.8.7.
    //
    // @param vmParam Identifier of the virtual machine to be unregistered.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws NotAllowedInCurrentState if the virtual machine is running (powered on).
    // @throws NotFound if there is no virtual machine associated with ``vm`` in the system.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Inventory.Unregister``.
	Unregister(vmParam string) error
}
