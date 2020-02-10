/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Timezone
 * Used by client-side stubs.
 */

package time


// The ``Timezone`` interface provides methods to get and set the appliance timezone. This interface was added in vSphere API 6.7.
type TimezoneClient interface {

    // Set time zone. This method was added in vSphere API 6.7.
    //
    // @param nameParam Time zone name.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws Error if any error occurs during the execution of the operation.
	Set(nameParam string) error

    // Get time zone. This method was added in vSphere API 6.7.
    // @return Time zone name.
    // @throws Error if timezone cannot be read.
	Get() (string, error)
}
