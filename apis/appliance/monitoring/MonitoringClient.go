/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Monitoring
 * Used by client-side stubs.
 */

package monitoring

import (
)

// ``Monitoring`` interface provides methods Get and list monitoring data for requested item.
type MonitoringClient interface {


    // Get monitoring data.
    //
    // @param itemParam MonitoredItemDataRequest Structure
    // @return list of MonitoredItemData structure
    // @throws Error Generic error
    Query(itemParam MonitoredItemDataRequest) ([]MonitoredItemData, error) 


    // Get monitored items list
    // @return list of names
    // @throws Error Generic error
    List() ([]MonitoredItem, error) 


    // Get monitored item info
    //
    // @param statIdParam statistic item id
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.monitoring``.
    // @return MonitoredItem structure
    // @throws Error Generic error
    Get(statIdParam string) (MonitoredItem, error) 

}
