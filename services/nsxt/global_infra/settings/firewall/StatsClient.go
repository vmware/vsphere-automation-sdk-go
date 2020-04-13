/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Stats
 * Used by client-side stubs.
 */

package firewall


type StatsClient interface {

    // Sets firewall rule statistics counter to zero. This operation is supported for given category, for example: DFW i.e. for all layer3 firewall (transport nodes only) rules or EDGE i.e. for all layer3 edge firewall (edge nodes only) rules. - no enforcement point path specified: On global manager, it is mandatory to give an enforcement point path. On local manager, reset of stats will be executed for each enforcement point. - {enforcement_point_path}: Reset of stats will be executed only for the given enforcement point.
    //
    // @param categoryParam Aggregation statistic category (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Reset(categoryParam string, enforcementPointPathParam *string) error
}
