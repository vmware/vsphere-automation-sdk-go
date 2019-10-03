package utils

import (
	"log"
)

func (t *taskImpl) Execute() error {
	ParseArguments()
	httpClient := GetHttpClientInstance()
	sessionManager := GetSessionManager()
	sessionManager.CreateSession(*httpClient)

	err := sessionManager.Login()
	if err != nil {
		return err
	}

	defer sessionManager.Logout()

	err = t.task(sessionManager.Connector())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
