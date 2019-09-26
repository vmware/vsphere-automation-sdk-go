/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Configuration
 * Used by client-side stubs.
 */

package configuration

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content"
)

// The ``Configuration`` interface provides methods to configure the global settings of the Content Library Service. 
//
//  The configuration settings are used by the Content Library Service to control the behavior of various operations.
type ConfigurationClient interface {


    // Updates the configuration. The update is incremental. Any property in the content.ConfigurationModel class that is null will not be modified. Note that this update method doesn't guarantee an atomic change of all the properties. In the case of a system crash or failure, some of the properties could be left unchanged while others may be updated.
    //
    // @param modelParam  The content.ConfigurationModel specifying the settings to update.
    // @throws InvalidArgument  if one of the configuration properties is not within the proper range.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``ContentLibrary.UpdateConfiguration``.
    Update(modelParam content.ConfigurationModel) error 


    // Retrieves the current configuration values.
    // @return The content.ConfigurationModel instance representing the configuration of the Content Library Service.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``ContentLibrary.GetConfiguration``.
    Get() (content.ConfigurationModel, error) 

}
