/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Restore
 * Used by client-side stubs.
 */

package restore

import (
)

// ``Restore`` interface provides methods Performs restore operations
type RestoreClient interface {


    // Get metadata before restore
    //
    // @param pieceParam RestoreRequest Structure
    // @return Metadata Structure
    // @throws Error Generic error
    Validate(pieceParam RestoreRequest) (Metadata, error) 

}
