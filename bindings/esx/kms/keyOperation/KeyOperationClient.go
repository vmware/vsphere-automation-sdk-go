/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: KeyOperation
 * Used by client-side stubs.
 */

package keyOperation

import (
)

// The ``KeyOperation`` interface provides methods to encrypt and decrypt data using a key that is provisioned for a Key Provider.
type KeyOperationClient interface {


    // Generate a new data encryption key.
    //
    // @param providerParam Identifier of the Key Provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param numOfBytesParam Key length.
    // @return A new data encryption key.
    // @throws InvalidArgument if the arguments are invalid.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``\\\\@ATTEST``.
    GenerateKey(providerParam string, numOfBytesParam int64) (GeneratedKey, error) 


    // Encrypt plaintext.
    //
    // @param providerParam Identifier of the Key Provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param plaintextParam Plaintext to encrypt.
    // @return Encrypted plaintext.
    // @throws InvalidArgument if the arguments are invalid.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``\\\\@ATTEST``.
    Encrypt(providerParam string, plaintextParam string) (EncryptResult, error) 


    // Decrypt ciphertext.
    //
    // @param providerParam Identifier of the Key Provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param ciphertextParam Ciphertext to decrypt.
    // @return Decrypted ciphertext.
    // @throws InvalidArgument if the arguments are invalid.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``\\\\@ATTEST``.
    Decrypt(providerParam string, ciphertextParam string) (DecryptResult, error) 

}
