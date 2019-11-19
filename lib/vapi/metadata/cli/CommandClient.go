/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Command
 * Used by client-side stubs.
 */

package cli


// The ``Command`` interface provides methods to get information about command line interface (CLI) commands.
type CommandClient interface {

    // Returns the identifiers of all commands, or commands in a specific namespace.
    //
    // @param pathParam The dot-separated path of the namespace for which command identifiers should be returned.
    // If null identifiers of all commands registered with the infrastructure will be returned.
    // @return Identifiers of the requested commands.
    // @throws NotFound if a namespace corresponding to ``path`` doesn't exist.
	List(pathParam *string) ([]CommandIdentity, error)

    // Retrieves information about a command including information about how to execute that command.
    //
    // @param identityParam Identifier of the command for which to retreive information.
    // @return Information about the command including information about how to execute that command.
    // @throws NotFound if a command corresponding to ``identity`` doesn't exist.
	Get(identityParam CommandIdentity) (CommandInfo, error)

    // Returns the aggregate fingerprint of all the command metadata from all the metadata sources. 
    //
    //  The fingerprint provides clients an efficient way to check if the metadata for commands has been modified on the server.
    // @return Fingerprint of all the command metadata present on the server.
	Fingerprint() (string, error)
}
