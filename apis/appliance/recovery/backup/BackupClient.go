/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Backup
 * Used by client-side stubs.
 */

package backup

import (
)

// ``Backup`` interface provides methods Performs backup restore operations
type BackupClient interface {


    // Check for backup errors without starting backup.
    //
    // @param pieceParam BackupRequest Structure
    // @return ReturnResult Structure
    // @throws FeatureInUse A backup or restore is already in progress
    // @throws Error Generic error
    Validate(pieceParam BackupRequest) (ReturnResult, error) 

}
