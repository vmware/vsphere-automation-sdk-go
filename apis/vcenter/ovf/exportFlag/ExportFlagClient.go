/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ExportFlag
 * Used by client-side stubs.
 */

package exportFlag

import (
)

// The ``ExportFlag`` interface provides methods for retrieving information about the export flags supported by the server. Export flags can be specified in a libraryItem.CreateSpec to customize an OVF export.
type ExportFlagClient interface {


    // Returns information about the supported export flags by the server. 
    //
    //  The supported flags are: 
    //
    // PRESERVE_MAC
    //     Include MAC addresses for network adapters.
    // EXTRA_CONFIG
    //     Include extra configuration in OVF export.
    //  
    //
    //  Future server versions might support additional flags.
    // @return A array of supported export flags.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    List() ([]Info, error) 

}
