/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Report
 * Used by client-side stubs.
 */

package report

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

// The ``Report`` interface provides methods to get attestation reports.
type ReportClient interface {


    // Request a report using remote attestation.
    //
    // @param requestParam The report request.
    // The parameter must contain all the properties defined in AttestRequest.
    // @return The report result.
    // The return value will contain all the properties defined in AttestResult.
    // @throws Error If a generic error.
    // @throws InvalidArgument If the request is invalid.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Unauthorized If the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``\\\\@ANONYMOUS``.
    Attest(requestParam *data.StructValue) (*data.StructValue, error) 

}
