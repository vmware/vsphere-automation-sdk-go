/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Question
 * Used by client-side stubs.
 */

package question

import (
)

// The ``Question`` interface provides methods to get the question raised during deployment and to answer them.
type QuestionClient interface {


    // Get the question that was raised during the configuration.
    // @return Info structure containing the question.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if the appliance is not in QUESTION_RAISED state.
    // @throws InternalServerError if questions could not be retrieved although the appliance is in QUESTION_RAISED state.
    Get() (Info, error) 


    // Supply answer to the raised question.
    //
    // @param specParam AnswerSpec with the answer to the raised question.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is NOT in QUESTION_RAISED state.
    // @throws InternalServerError if answer file could not be created.
    Answer(specParam AnswerSpec) error 

}
