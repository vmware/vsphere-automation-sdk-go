package task

import (
	"fmt"
	"log"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
)

func (t *taskImpl) Execute(sessionManager session.SessionManager) error {
	if sessionManager == nil {
		return fmt.Errorf("Task Execution Error: Invalid Session")
	}
	err := sessionManager.Login()
	if err != nil {
		return err
	}

	defer sessionManager.Logout()

	err = t.task(sessionManager)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
