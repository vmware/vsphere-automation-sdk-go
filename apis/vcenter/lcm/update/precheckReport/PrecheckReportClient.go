/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: PrecheckReport
 * Used by client-side stubs.
 */

package precheckReport

import (
)

// The ``PrecheckReport`` interface generates precheck report for a vCenter Server instance against a target update version.
type PrecheckReportClient interface {


    // Creates a vCenter Server pre-update compatibility check report for the pending update version. The report can be exported and downloaded in CSV format. 
    //
    //  The result of this operation can be queried by calling the null method where ``task`` is the response of this operation.
    //
    // @param versionParam Pending update version for which pre-update compatibility check will be executed.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.update.pending``.
    // @return The precheck report, which contains a link to download the CSV report as well
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws NotFound if there is no pending update assosiated with the ``version`` in the system.
    // @throws NotAllowedInCurrentState if a precheck is already in progress.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    Create(versionParam string) (Result, error) 

}
