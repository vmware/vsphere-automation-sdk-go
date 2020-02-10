/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Tls
 * Used by client-side stubs.
 */

package vcenter


// The ``Tls`` interface provides methods to replace Tls certificate. This interface was added in vSphere API 6.7.2.
type TlsClient interface {

    // Replaces the rhttpproxy TLS certificate with the specified certificate. This method can be used in three scenarios : 
    //
    // #. When the CSR is created and the private key is already stored, this method can replace the certificate. The certificate but not the private key and root certificate must be provided as input.
    // #. When the certificate is signed by a third party certificate authority/VMCA and the root certificate of the third party certificate authority/VMCA is already one of the trusted roots in the trust store, this method can replace the certificate and private key. The certificate and the private key but not the root certificate must be provided as input.
    // #. When the certificate is signed by a third party certificate authority and the root certificate of the third party certificate authority is not one of the trusted roots in the trust store, this method can replace the certificate, private key and root CA certificate. The certificate, private key and root certificate must be provided as input.
    //
    //  After this method completes, the services using the certificate will be restarted for the new certificate to take effect. 
    //
    // The above three scenarios are only supported from vsphere 7.0 onwards.. This method was added in vSphere API 6.7.2.
    //
    // @param specParam The information needed to replace the TLS certificate.
    // @throws NotFound If the private key is not present in the VECS store.
    // @throws AlreadyExists If the specified certificate thumbprint is the same as the existing TLS certificate thumbprint.
    // @throws Error If the system failed to replace the TLS certificate.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``CertificateManagement.Administer``.
	Set(specParam TlsSpec) error

    // Returns the rhttpproxy TLS certificate. This method was added in vSphere API 6.7.2.
    // @return TLS certificate.
    // @throws NotFound if the rhttpproxy certificate is not present in VECS store.
    // @throws Error if failed due to generic exception.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	Get() (TlsInfo, error)

    // Renews the TLS certificate for the given duration period. 
    //
    // After this method completes, the services using the certificate will be restarted for the new certificate to take effect.. This method was added in vSphere API 6.7.2.
    //
    // @param durationParam The duration (in days) of the new TLS certificate. The duration should be less than or equal to 730 days.
    // If null, the duration will be 730 days (two years).
    // @throws Unsupported If the TLS certificate is not VMCA generated.
    // @throws InvalidArgument If the duration period specified is invalid.
    // @throws Error If the system failed to renew the TLS certificate.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``CertificateManagement.Administer``.
	Renew(durationParam *int64) error

    // Replace MACHINE SSL with VMCA signed one with the given Spec.The system will go for restart. 
    //
    //  After this method completes, the services using the certificate will be restarted for the new certificate to take effect.. This method was added in vSphere API 6.9.1.
    //
    // @param specParam The information needed to generate VMCA signed Machine SSL
    // @throws InvalidArgument If the Spec given is not complete or invalid
    // @throws Error If the system failed to replace the machine ssl certificate
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``CertificateManagement.Administer``.
	ReplaceVmcaSigned(specParam TlsReplaceSpec) error
}
