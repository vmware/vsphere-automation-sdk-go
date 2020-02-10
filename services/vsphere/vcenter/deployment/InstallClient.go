/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Install
 * Used by client-side stubs.
 */

package deployment


// The ``Install`` interface provides methods to configure the installation of the appliance. This interface was added in vSphere API 6.7.
type InstallClient interface {

    // Get the parameters used to configure the ongoing appliance installation. This method was added in vSphere API 6.7.
    // @return InstallSpec parameters being used to configure appliance install.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if appliance is not in INSTALL_PROGRESS state.
	Get() (InstallInstallSpec, error)

    // Run sanity checks using the InstallSpec parameters passed. This method was added in vSphere API 6.7.
    //
    // @param specParam InstallSpec parameters to run sanity check with.
    // @return CheckInfo containing the check results.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if external PSC credentials are not valid when configuring PSC to replicate with an external existing PSC.
    // @throws Unauthenticated if external PSC credentials are not valid when configuring a VCSA_EXTERNAL appliance.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
	Check(specParam InstallInstallSpec) (CheckInfo, error)

    // Start the appliance installation. This method was added in vSphere API 6.7.
    //
    // @param specParam InstallSpec parameters to configure the appliance install.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if the partner PSC credentials are not valid when configuring PSC to replicate with partner PSC.
    // @throws Unauthenticated if external PSC credentials are not valid when configuring a VCSA_EXTERNAL appliance.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
	Start(specParam InstallInstallSpec) error

    // Cancel the appliance installation that is in progress. This method was added in vSphere API 6.7.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if the appliance is not in CONFIG_IN_PROGRESS state and if the operation is not INSTALL.
	Cancel() error
}
