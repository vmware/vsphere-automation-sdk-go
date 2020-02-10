/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Forwarding
 * Used by client-side stubs.
 */

package logging


// The ``Forwarding`` interface provides methods to manage forwarding of log messages to remote logging servers. This interface was added in vSphere API 6.7.
type ForwardingClient interface {

    // Validates the current log forwarding configuration by checking the liveness of the remote machine and optionally sending a test diagnostic log message from the appliance to all configured logging servers to allow manual end-to-end validation. The message that is sent is: "This is a diagnostic log test message from vCenter Server.". This method was added in vSphere API 6.7.
    //
    // @param sendTestMessageParam Flag specifying whether a default test message should be sent to the configured logging servers.
    // If null, no test message will be sent to the configured remote logging servers.
    // @return Information about the status of the connection to each of the remote logging servers.
	Test(sendTestMessageParam *bool) ([]ForwardingConnectionStatus, error)

    // Sets the configuration for forwarding log messages to remote log servers. This method was added in vSphere API 6.7.
    //
    // @param cfgListParam The cfgList is a list of Config structure that contains the log message forwarding rules in terms of the host, port, protocol of the log message.
    // @throws InvalidArgument if an invalid configuration is provided.
    // @throws UnableToAllocateResource if the number of configurations exceeds the maximum number of supported configurations.
    // @throws Error if there is any internal error during the execution of the operation.
	Set(cfgListParam []ForwardingConfig) error

    // Returns the configuration for forwarding log messages to remote logging servers. This method was added in vSphere API 6.7.
    // @return Information about the configuration for forwarding log messages to remote logging servers.
	Get() ([]ForwardingConfig, error)
}
