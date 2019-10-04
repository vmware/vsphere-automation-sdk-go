/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Processes
 * Used by client-side stubs.
 */

package processes

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/guest"
)

// The ``Processes`` interface provides methods to manage processes in the guest operating system. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ProcessesClient interface {


    // Starts a program in the guest operating system. 
    //
    //  A process started this way can have its status queried with Processes#list or Processes#get. When the process completes, its exit code and end time will be available for 5 minutes after completion. 
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine to perform the operation on.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param credentialsParam The guest authentication data. See guest.Credentials. The program will be run as the user associated with this data.
    // @param specParam The arguments describing the program to be started.
    // @return The process id of the program started.
    // @throws NotFound if the virtual machine ``vm`` is not found.
    // @throws NotFound if the program path does not exist.
    // @throws Unsupported if the operation is not supported by the VMware tools in the guest OS.
    // @throws Unsupported if the operation is disabled by the VMware tools in the guest OS.
    // @throws NotAllowedInCurrentState if the virtual machine ``vm`` is not running.
    // @throws UnableToAllocateResource if the program fails to start. XXX may not work today; hostd doesn't map VIX_E_PROGRAM_NOT_STARTED.
    // @throws ResourceBusy if the virtual machine is busy.
    // @throws Unauthenticated if the ``credentials`` was not valid.
    // @throws Unauthorized if the ``path`` property of ``spec`` cannot be accessed.
    // @throws Unauthorized if the ``path`` property of ``spec`` cannot be run because ``credentials`` will not allow the operation.
    // @throws ServiceUnavailable if the VMware tools are not running.
    Create(vmParam string, credentialsParam guest.Credentials, specParam CreateSpec) (int64, error) 


    // Gets the status of a process running in the guest operating system, including those started by Processes#create that may have recently completed. 
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine to perform the operation on.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param credentialsParam The guest authentication data. See guest.Credentials.
    // @param pidParam Specifies the process to query.
    // @return The Processes.Info for the process with id ``pid``.
    // @throws NotFound if the virtual machine ``vm`` is not found.
    // @throws NotFound if the process ``pid`` is not found.
    // @throws Unsupported if the operation is not supported by the VMware tools in the guest OS.
    // @throws Unsupported if the operation is disabled by the VMware tools in the guest OS.
    // @throws NotAllowedInCurrentState if the virtual machine ``vm`` is not running.
    // @throws ResourceBusy if the virtual machine is busy.
    // @throws Unauthenticated if the ``credentials`` is not valid.
    // @throws ServiceUnavailable if the VMware tools are not running.
    Get(vmParam string, credentialsParam guest.Credentials, pidParam int64) (Info, error) 


    // List the processes running in the guest operating system, plus those started by Processes#create that have recently completed. 
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine to perform the operation on.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param credentialsParam The guest authentication data. See guest.Credentials.
    // @param pidsParam Specifies which processes should be listed. If a specified processes does not exist, nothing will be returned for that process.
    // If null information about all processes is returned.
    // @return The list of running processes is returned in an array of Processes.Info classes.
    // @throws NotFound if the virtual machine ``vm`` is not found.
    // @throws Unsupported if the operation is not supported by the VMware tools in the guest OS.
    // @throws Unsupported if the operation is disabled by the VMware tools in the guest OS.
    // @throws NotAllowedInCurrentState if the virtual machine ``vm`` is not running.
    // @throws ResourceBusy if the virtual machine is busy.
    // @throws Unauthenticated if the ``credentials`` is not valid.
    // @throws ServiceUnavailable if the VMware tools are not running.
    List(vmParam string, credentialsParam guest.Credentials, pidsParam []int64) ([]Info, error) 


    // Terminates a process in the guest OS. 
    //
    //  On Posix guests, the process is sent a TERM signal. If that doesn't terminate the process, a KILL signal is sent. A process may still be running if it's stuck. 
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine to perform the operation on.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param credentialsParam The guest authentication data. See guest.Credentials.
    // @param pidParam Process ID of the process to be terminated
    // @throws NotFound if the virtual machine ``vm`` is not found.
    // @throws NotFound if the ``pid`` is not found.
    // @throws Unsupported if the operation is not supported by the VMware tools in the guest OS.
    // @throws Unsupported if the operation is disabled by the VMware tools in the guest OS.
    // @throws NotAllowedInCurrentState if the virtual machine ``vm`` is not running.
    // @throws ResourceBusy if the virtual machine is busy.
    // @throws Unauthenticated if ``credentials`` is not valid.
    // @throws Unauthorized if ``credentials`` does not have permission to terminate the process.
    // @throws ServiceUnavailable if the VMware tools are not running.
    Delete(vmParam string, credentialsParam guest.Credentials, pidParam int64) error 

}
