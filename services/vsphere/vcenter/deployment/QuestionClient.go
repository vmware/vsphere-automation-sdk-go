/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Question
 * Used by client-side stubs.
 */

package deployment


// The ``Question`` interface provides methods to get the question raised during deployment and to answer them. This interface was added in vSphere API 6.7.
type QuestionClient interface {

    // Get the question that was raised during the configuration. This method was added in vSphere API 6.7.
    // @return Info structure containing the question.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if the appliance is not in QUESTION_RAISED state.
    // @throws InternalServerError if questions could not be retrieved although the appliance is in QUESTION_RAISED state.
	Get() (QuestionInfo, error)

    // Supply answer to the raised question. This method was added in vSphere API 6.7.
    //
    // @param specParam AnswerSpec with the answer to the raised question.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is NOT in QUESTION_RAISED state.
    // @throws InternalServerError if answer file could not be created.
	Answer(specParam QuestionAnswerSpec) error
}
